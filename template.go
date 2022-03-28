package pointable

import (
	"fmt"
	"io"
	"text/template"
)

// Type represents type information to pass to templates.
type Type interface {
	Symbol() string
	FuncName() string
	// TODO: add Imports() for some cases like time.Time
}

// Template represent the specific template.
type Template struct {
	*Dist
	*template.Template

	name   string
	types  []Type
	output io.WriteCloser
}

func NewDefaultTemplate(dist *Dist, name string, tmpl *template.Template, types []Type) *Template {
	file, err := dist.createFile(name)
	if err != nil {
		panic(err)
	}
	return NewTemplateWithOutput(dist, name, file, tmpl, types)
}

func NewTemplateWithOutput(dist *Dist, name string, output io.WriteCloser, tmpl *template.Template, types []Type) *Template {
	return &Template{
		Dist:     dist,
		Template: tmpl,
		name:     name,
		types:    types,
		output:   output,
	}
}

func (t Template) Do() error {
	defer t.output.Close()
	vars := TemplateVars{
		Pkg:   t.Dist.name,
		Types: t.types,
	}
	if err := t.Execute(t.output, vars); err != nil {
		return fmt.Errorf("generating template for %s: %v", t.name, err)
	}
	return nil
}

// TemplateVars represent the schema of data that will be passed to the template.
type TemplateVars struct {
	Pkg   string
	Types []Type
}

// CommonTemplate is the common template.
// It is recommended to use this template, but you can insert the template if you need.
var CommonTemplate = template.Must(template.New("CommonTemplate").Parse(`package {{ $.Pkg }}
{{ range $_, $t := $.Types }}
{{ template "func" $t }}
{{- end }}
{{- define "func" }}
// {{ $.FuncName }} returns a pointer to the given {{ $.Symbol }}.
func {{ $.FuncName }}(v {{ $.Symbol }}) *{{ $.Symbol }} {
	return &v
}
{{- end }}
`))
