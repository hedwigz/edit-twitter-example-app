package enthistory

import (
	"context"
	"testing"
	"time"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/ariga/edit-twitter-example-app/enthistory/ent"
	"github.com/ariga/edit-twitter-example-app/enthistory/ent/entity"
	"github.com/ariga/edit-twitter-example-app/enthistory/ent/enttest"
	"github.com/ariga/edit-twitter-example-app/enthistory/ent/history"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCodeGeneration(t *testing.T) {
	err := entc.Generate("./ent/schema", &gen.Config{}, entc.Extensions(NewExtension()))
	if err != nil {
		t.Fatalf("running ent codegen: %v", err)
	}
}

func TestQueryHistoryAsEdge(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()

	e := client.Entity.Create().
		SetData("enthistory").
		SaveX(context.Background())

	client.History.Create().
		SetAction(history.ActionCreate).
		SetEntityName("mock").
		SetRecordID(e.ID).
		SetTimestamp(time.Now()).SaveX(context.Background())

	h := client.Entity.QueryHistory(e).AllX(context.Background())

	assert.Len(t, h, 1)
	assert.EqualValues(t, h[0].RecordID, e.ID)
}

func TestCreateEntityCreatesHistoryRow(t *testing.T) {
	ctx := context.Background()

	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()

	ent.HookHistory(client, nil)

	tx, err := client.Tx(ctx)
	assert.NoError(t, err)

	e := tx.Entity.Create().
		SetData("a8m").
		SetStrings([]string{"a", "b"}).
		SaveX(ctx)
	err = tx.Commit()
	assert.NoError(t, err)

	h := client.Entity.QueryHistory(e).WithChanges().AllX(ctx)

	assert.Len(t, h, 1)
	assert.EqualValues(t, h[0].RecordID, e.ID)
	assert.Equal(t, h[0].Edges.Changes[0].Value, "a8m")
	assert.Equal(t, h[0].Edges.Changes[1].Value, `["a","b"]`)
}

func TestCreateBulkEntityCreatesHistoryRow(t *testing.T) {
	ctx := context.Background()

	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()

	ent.HookHistory(client, nil)

	tx, err := client.Tx(ctx)
	assert.NoError(t, err)

	e := tx.Entity.CreateBulk(tx.Entity.Create().SetData("a8m"), tx.Entity.Create().SetData("a9m")).SaveX(ctx)
	err = tx.Commit()
	assert.NoError(t, err)

	h := client.Entity.QueryHistory(e[0]).WithChanges().AllX(ctx)

	assert.Len(t, h, 1)
	assert.EqualValues(t, h[0].RecordID, e[0].ID)
	assert.Equal(t, h[0].Edges.Changes[0].Value, "a8m")

	h = client.Entity.QueryHistory(e[1]).WithChanges().AllX(ctx)
	assert.Len(t, h, 1)
	assert.EqualValues(t, h[0].RecordID, e[1].ID)
	assert.Equal(t, h[0].Edges.Changes[0].Value, "a9m")
}

func TestUpdateEntityCreatesHistoryRow(t *testing.T) {
	ctx := context.Background()

	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()

	ent.HookHistory(client, nil)

	tx, err := client.Tx(ctx)
	assert.NoError(t, err)

	e := tx.Entity.Create().
		SetData("a8m").
		SetCounter(5).
		SaveX(ctx)
	e = tx.Entity.UpdateOne(e).
		SetData("a9m").
		SetCounter(10).
		SaveX(ctx)

	err = tx.Commit()
	assert.NoError(t, err)

	h := client.Entity.QueryHistory(e).WithChanges().AllX(ctx)

	assert.Len(t, h, 2)
	assert.EqualValues(t, h[0].RecordID, e.ID)
	assert.Equal(t, h[0].Edges.Changes[0].Value, "a8m")
	assert.Equal(t, h[0].Edges.Changes[1].Value, "5")
	assert.Equal(t, h[1].Edges.Changes[0].Previous, "a8m")
	assert.Equal(t, h[1].Edges.Changes[0].Value, "a9m")
	assert.Equal(t, h[1].Edges.Changes[1].Previous, "5")
	assert.Equal(t, h[1].Edges.Changes[1].Value, "10")
}

func TestUpdateManyEntityCreatesHistoryRow(t *testing.T) {
	ctx := context.Background()

	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()

	ent.HookHistory(client, nil)

	tx, err := client.Tx(ctx)
	assert.NoError(t, err)

	e := tx.Entity.Create().
		SetData("a8m").
		SetIsFun(true).
		SaveX(ctx)
	tx.Entity.Create().
		SetData("a9m").
		SetIsFun(true).
		SaveX(ctx)
	// the following updates multiple fields and enthistory hook does not support
	tx.Entity.Update().
		Where(entity.IsFun(true)).
		SetData("a10m").
		SaveX(ctx)

	err = tx.Commit()
	assert.NoError(t, err)

	h := client.Entity.QueryHistory(e).WithChanges().AllX(ctx)

	assert.Len(t, h, 1)
	assert.EqualValues(t, h[0].RecordID, e.ID)
	assert.Equal(t, h[0].Edges.Changes[0].Value, "a8m")
}

func TestNoOp(t *testing.T) {
	ctx := context.Background()

	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()

	ent.HookHistory(client, nil)

	tx, err := client.Tx(ctx)
	assert.NoError(t, err)

	e := tx.Entity.Create().
		SetData("a8m").
		SetIsFun(true).
		SaveX(ctx)
	require.NoError(t, tx.Commit())

	// Perform a no-op
	tx, err = client.Tx(ctx)
	require.NoError(t, err)
	tx.Entity.UpdateOne(e).ExecX(ctx)
	require.NoError(t, tx.Commit())

	// Assert no history was created from the UPDATE.
	c := client.History.
		Query().
		Where(
			history.RecordID(e.ID),
			history.EntityName("Entity"),
			history.ActionEQ(history.ActionUpdate),
		).
		CountX(ctx)
	require.Zero(t, c)
}
