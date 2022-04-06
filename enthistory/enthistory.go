package enthistory

import (
	"bytes"
	_ "embed"
	"fmt"
	"text/template"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

// Extension extends ent with tracking history capabilities.
type Extension struct {
	entc.DefaultExtension
}

// NewExtension creates a new Extension.
func NewExtension() *Extension {
	return &Extension{}
}

func isHistoryField(f *gen.Field) bool {
	return f.Annotations != nil && f.Annotations[annotationName] != nil
}

// historyNodes checks if any field of the node type is annotated with an enthistory annotation
func isHistoryNode(t *gen.Type) bool {
	for _, f := range t.Fields {
		if isHistoryField(f) {
			return true
		}
	}
	return false
}

// filterHistoryNodes returns only the nodes that are tracked.
func filterHistoryNodes(nodes []*gen.Type) []*gen.Type {
	var l []*gen.Type
	for _, n := range nodes {
		if isHistoryNode(n) {
			l = append(l, n)
		}
	}
	return l
}

//go:embed templates/convert.tmpl
var fieldConvertTmpl string

func fieldToConverterName(field *gen.Field) string {
	switch {
	case field.Type != nil && field.Type.Type.Numeric():
		return "int"
	case field.Type != nil && field.Type.Type.Float():
		return "float"
	case field.IsTime():
		return "time"
	case field.IsBool():
		return "bool"
	case field.IsString():
		return "string"
	case field.IsJSON():
		return "json"
	default:
		panic(fmt.Sprintf("enthistory: unsupported field type %s", field.Type))
	}
}

// fieldMarshal returns code that marshals the given field into a string.
func fieldMarshal(field *gen.Field, varName, resVarName string) (string, error) {
	var err error
	buf := bytes.NewBuffer(nil)
	converter := fieldToConverterName(field)

	t, err := template.New(fmt.Sprintf("%s/marshal", converter)).Parse(fieldConvertTmpl)
	if err != nil {
		return "", err
	}
	if err = t.Execute(buf, map[string]interface{}{"f": field, "varName": varName, "resVarName": resVarName}); err != nil {
		return "", err
	}
	return buf.String(), nil
}

//go:embed templates/enthistory.tmpl
var tmpl string

// Templates of the extension.
func (*Extension) Templates() []*gen.Template {
	t := gen.NewTemplate("enthistory")
	t.Funcs(template.FuncMap{
		"isHistoryNode":      isHistoryNode,
		"isHistoryField":     isHistoryField,
		"filterHistoryNodes": filterHistoryNodes,
		"fieldMarshal":       fieldMarshal,
	})
	return []*gen.Template{
		gen.MustParse(t.Parse(tmpl)),
	}
}
