package main

import (
	"fmt"
	"html/template"
	"log"
	"os"

	"github.com/spf13/pflag"
)

var tpl string
var yangDir string

type Intput struct {
	PkgName      string
	ResourceName string
	Resources    []struct {
		Name     string
		Path     string
		Type     string
		Optional string
	}
}

func main() {
	pflag.StringVarP(&yangDir, "yang-dir", "y", "./yang", "yang models directory")
	pflag.StringVarP(&tpl, "template", "t", "./templates", "templates path")
	pflag.Parse()

	var inputs []Intput
	//
	// TODO: add reading yang models, and build a []Input
	//

	funcMap := template.FuncMap{
		"ToSnakeCase":  ToSnakeCase,
		"YangToTFType": YangToTFType,
	}

	templates := template.Must(template.ParseGlob(tpl + "/*")).Funcs(funcMap)
	for _, in := range inputs {
		filename := fmt.Sprintf("resource_%s.go", ToSnakeCase(in.ResourceName))
		f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Printf("failed to create file %s: %v", filename, err)
			return
		}
		err = templates.ExecuteTemplate(f, "base", in)
		if err != nil {
			log.Printf("failed to execute template with input %v: %v", in, err)
			return
		}
	}
}
