package main

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/dialect"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	app "github.com/ariga/edit-twitter-example-app"
	"github.com/ariga/edit-twitter-example-app/ent"
	"github.com/ariga/edit-twitter-example-app/ent/migrate"
	_ "github.com/mattn/go-sqlite3"
)

func seed(client *ent.Client) {
	c := client.User.Query().CountX(context.Background())
	// no need to seed
	if c != 0 {
		return
	}
	client.User.Create().
		SetName("Hedwigz").
		SetAge(26).
		SetEmail("amit@ariga.io").
		SetRole(1).
		SetImageURL("https://avatars.githubusercontent.com/u/8277210?v=4").
		SaveX(context.Background())
	a8m := client.User.Create().
		SetName("a8m").
		SetAge(32).
		SetEmail("a@ariga.io").
		SetRole(1).
		SetImageURL("https://avatars.githubusercontent.com/u/7413593?v=4").
		SaveX(context.Background())
	client.Tweet.Create().
		SetAuthor(a8m).
		SetContent("hello world").
		SaveX(context.Background())
}

func main() {
	// Create ent.Client and run the schema migration.
	client, err := ent.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")

	if err != nil {
		log.Fatal("opening ent client", err)
	}
	if err := client.Schema.Create(
		context.Background(),
		migrate.WithGlobalUniqueID(true),
	); err != nil {
		log.Fatal("opening ent client", err)
	}

	seed(client)
	ent.HookHistory(client, func(ctx context.Context, hc *ent.HistoryCreate) *ent.HistoryCreate {
		return hc
	})

	// Configure the server and start listening on :8081.
	srv := handler.NewDefaultServer(app.NewSchema(client))

	srv.Use(entgql.Transactioner{TxOpener: client})
	http.Handle("/",
		playground.Handler("Todo", "/query"),
	)
	http.Handle("/query", Authenticator(srv))
	log.Println("listening on :8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal("http server terminated", err)
	}
}

func Authenticator(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		uid := r.URL.Query().Get("userid")
		userID, err := strconv.Atoi(uid)
		if err != nil {
			next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "userid", -1)))
			return
		}
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "userid", userID)))
	}
	return http.HandlerFunc(fn)
}
