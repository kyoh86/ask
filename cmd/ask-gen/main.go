// !build gen

package main

import (
	"html/template"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

func main() {
	genTypes()
	genPrototype()
}

func generate(definition interface{}, templateText, filename string) (retErr error) {
	buf, readErr := ioutil.ReadFile(filename + ".yaml")
	if readErr != nil {
		return readErr
	}
	if err := yaml.Unmarshal(buf, definition); err != nil {
		return err
	}

	out, openErr := os.OpenFile(filename+"_gen.go", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if openErr != nil {
		return openErr
	}
	defer func() {
		if err := out.Close(); err != nil && retErr == nil {
			retErr = err
		}
	}()

	tmp, err := template.New("definition").
		Funcs(template.FuncMap{"title": strings.Title}).
		Parse(templateText)
	if err != nil {
		return err
	}
	if err := tmp.Execute(out, map[string]interface{}{"definition": definition}); err != nil {
		return err
	}
	cmd := exec.Command("goimports", "-w", filename+"_gen.go")
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func genTypes() {
	var types []struct {
		Type string `yaml:"type"`
		Name string `yaml:"name"`
		Conv string `yaml:"conv"`
	}

	var typesTemplate = `// Code generated by ask-gen. DO NOT EDIT!!

package ask

{{range .definition}}
// {{.Name|title}} takes {{.Type}} value from user input
func (s Service) {{.Name|title}}() (*{{.Type}}, error) {
	var v {{.Type}}
	if err := s.{{.Name|title}}Var(&v).Do(); err != nil {
		return nil, err
	}
	return &v, nil
}

// {{.Name|title}} takes {{.Type}} value from user input
func {{.Name|title}}() (*{{.Type}}, error) {
  return static.{{.Name|title}}()
}

// {{.Name|title}}Var sets a {{.Type}} variable, "v" to accept user input
func (s Service) {{.Name|title}}Var(v *{{.Type}}) Doer {
	return doFunc(func() error {
		return s.AskFunc(func(input string) error {
			{{if .Conv -}}
			p, err := {{.Conv}}
			if err != nil {
				return err
			}
			*v = {{.Type}}(p)
			{{- else -}}
			*v = input
			{{- end}}
			return nil
		})
	})
}

// {{.Name|title}}Var sets a {{.Type}} variable, "v" to accept user input
func {{.Name|title}}Var(v *{{.Type}}) Doer {
  return static.{{.Name|title}}Var(v)
}
{{end}}
`

	if err := generate(&types, typesTemplate, "types"); err != nil {
		panic(err)
	}
}

func genPrototype() {
	var options []struct {
		Name string `yaml:"name"`
		Type string `yaml:"type"`
	}

	var prototypeTemplate = `// Code generated by ask-gen. DO NOT EDIT!!

package ask

// Prototype gives handler some options
type Prototype struct {
{{- range .definition}}
	{{.Name|title}} {{.Type}}
{{- end}}
}
{{range .definition}}
// {{.Name|title}} makes new Service with new {{.Name|title}} option
func (s Service) {{.Name|title}}(v {{.Type}}) *Service {
	n := s.Prototype
	n.{{.Name|title}} = v
	return &Service{Prototype: n}
}

// {{.Name|title}} makes new Service with new {{.Name|title}} option
func {{.Name|title}}(v {{.Type}}) *Service {
	return static.{{.Name|title}}(v)
}
{{end}}
`

	if err := generate(&options, prototypeTemplate, "prototype"); err != nil {
		panic(err)
	}
}
