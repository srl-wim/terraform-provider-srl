package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/openconfig/goyang/pkg/yang"
)

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func fileName(s string) string {
	r := []rune(s)
	r = r[1:len(r)]
	s = string(r)
	s = strings.ReplaceAll(s, "/", "_")
	s = strings.ReplaceAll(s, "-", "_")
	s = s + "_resource.txt"
	//fmt.Printf("Filename: %s\n", s)
	return s
}

func findResources(res string) {
	// When the global Resource was initialized and the string no longer contains the subsctring,
	// it means we have to close the file and reinitilaize the global variabled
	if Resource != "" && !strings.Contains(res, Resource) {
		//fmt.Printf("CLOSE FILE \n")
		F.Close()
		F = nil
		Resource = ""
		Indent = ""
	}

	// check if the resource is part of the global resouce files
	var err error
	for _, r := range Resources {
		if strings.Contains(res, r) {
			// if Resource was not found, open the file; otherwise it should have been open already
			if Resource == "" {
				f := fileName(r)
				//fmt.Printf("OPEN FILE: %s \n", f)
				F, err = os.Create(filepath.Join(outDir, filepath.Base(f)))
				if err != nil {
					log.Fatal(err)
				}

			}
			Resource = r
		}
	}
}

func getTypeName(e *yang.Entry) string {
	if e == nil || e.Type == nil {
		return ""
	}
	// Return our root's type name.
	// This is should be the builtin type-name
	// for this entry.
	return e.Type.Root.Name
}

// function that goes through the yang tree and writes the elements to the file
func resourceWrite(path string, e *yang.Entry) {
	path += "/" + e.Name
	//fmt.Printf("Resource: %s \n", res)
	//fmt.Printf("Indent Length: %d \n", len(Indent))

	// check if the resource is part of the global resource definitions
	findResources(path)
	//fmt.Printf("FOUND-RESOURCE %s \n", Resource)

	// if resource was found, process the rest of the yang tree
	if Resource != "" {
		// check the Indent length for the element in the yang tree
		i := strings.Count(path, "/") - strings.Count(Resource, "/")
		Indent = strings.Repeat("  ", i)
		// check if the elemnt is RW or RO or RPC
		switch {
		case e.RPC != nil:
			fmt.Fprintf(F, Indent+"RPC: ")
		case e.ReadOnly():
			fmt.Fprintf(F, Indent+"RO: ")
		default:
			fmt.Fprintf(F, Indent+"rw: ")
		}
		// check the type
		if e.Type != nil {
			fmt.Fprintf(F, "%s ", getTypeName(e))
		}
		// check if the element has a mandatory flag
		switch e.Mandatory {
		case 0: // TSUnset
			//fmt.Fprintf(w, "M: Unknown ")
		case 1: // TSTrue
			fmt.Fprintf(F, "M: true ")
		case 2: // TSTrue
			fmt.Fprintf(F, "M: false ")
		}
		// check if the element has a default indication
		if e.Default != "" {
			fmt.Fprintf(F, "D: %s ", e.Default)
		}
		name := e.Name
		// check the type of element
		switch {
		case e.Dir == nil && e.ListAttr != nil:
			fmt.Fprintf(F, "[]%s\n", name)
			return
		case e.Dir == nil:
			fmt.Fprintf(F, "%s\n", name)
			return
		case e.ListAttr != nil:
			fmt.Fprintf(F, "[%s]%s {\n", e.Key, name) //}
		default:
			fmt.Fprintf(F, "%s {\n", name) //}
		}
	}

	var names []string
	for k := range e.Dir {
		names = append(names, k)
	}
	sort.Strings(names)

	for _, k := range names {

		resourceWrite(path, e.Dir[k])
	}

	// check the Indent length for the element in the yang tree
	i := strings.Count(path, "/") - strings.Count(Resource, "/")
	if i >= 0 {
		Indent = strings.Repeat("  ", i)
		fmt.Fprintln(F, Indent+"}")
	}

}
