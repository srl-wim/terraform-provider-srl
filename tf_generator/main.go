package main

import (
	"bufio"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"sort"

	"github.com/openconfig/goyang/pkg/yang"
	"github.com/spf13/pflag"
)

// Resources defines the terraform resources we want output files for
var Resources []string

// Resource is a global resource to hold state
var Resource string

// F is a global variable to hold the file
var F *os.File

// Indent is a global varibale to hold the indent state
var Indent string

// Flags
var tpl string
var yangDir string
var resFile string
var paths []string
var outDir string

// Intput struct
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
	pflag.StringSliceVarP(&paths, "path-dir", "p", []string{"./yang/ietf"}, "comma separated list of directories to add to search path DIR[,DIR...]")
	pflag.StringVarP(&yangDir, "yang-dir", "y", "./yang/srl", "yang models directory")
	pflag.StringVarP(&resFile, "res-file", "r", "./res/resourceMapInput.conf", "resource input file")
	pflag.StringVarP(&outDir, "out-dir", "o", "./out", " output directory where resource files are stored")
	pflag.StringVarP(&tpl, "template", "t", "./templates", "templates path")
	pflag.Parse()

	var inputs []Intput
	//
	// TODO: add reading yang models, and build a []Input
	//
	// Check if the resource input file exists
	if !fileExists(resFile) {
		fmt.Printf("Resource Map input file does not exist, please supply a valid file as input. Exiting ...")
		os.Exit(1)
	}
	// Read resource input file
	r, err := os.Open(resFile)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		Resources = append(Resources, scanner.Text())
	}

	// read the path directory
	for _, path := range paths {
		expanded, err := yang.PathsWithModules(path)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		yang.AddPath(expanded...)
	}

	// initialize yang modules
	ms := yang.NewModules()

	// read the yang directory
	files, err := ioutil.ReadDir(yangDir)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		fmt.Println(yangDir + "/" + f.Name())
		if err := ms.Read(yangDir + "/" + f.Name()); err != nil {
			fmt.Printf("ERROR: %v %s\n", os.Stderr, err)
			continue
		}
	}

	// process the yang modules
	errs := ms.Process()
	if len(errs) > 0 {
		for err := range errs {
			fmt.Printf("Error: %v \n", err)
		}
	}

	// Keep track of the top level modules we read in.
	// Those are the only modules we want to process.
	mods := map[string]*yang.Module{}
	var names []string
	for _, m := range ms.Modules {
		if mods[m.Name] == nil {
			mods[m.Name] = m
			names = append(names, m.Name)
		}
	}
	sort.Strings(names)
	entries := make([]*yang.Entry, len(names))
	for x, n := range names {
		entries[x] = yang.ToEntry(mods[n])
	}

	for _, e := range entries {
		Resource = ""
		Indent = ""
		resourceWrite("", e)
		if F != nil {
			F.Close()
			F = nil
		}
	}

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
