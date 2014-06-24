// to cross compile from linux:
// GOOS=windows GOARCH=386 go build -o mt.exe mt.go

package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

func main() {

	// read arguments
	if len(os.Args) < 3 {
		log.Fatal("\n\n./mt.go template_input_file output_file\n\n")
	}

	template_file := os.Args[1]
	output_file := os.Args[2]

	log.Println("template_file =", template_file, "output_file =", output_file)

	// pull settings from the environment
	settings := make(map[string]string)
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		settings[pair[0]] = pair[1]
		log.Println(pair[0], "=", pair[1])
	}

	// read template file
	t, err := template.ParseFiles(template_file)
	if err != nil {
		log.Fatal(err)
	}

	// open output file
	out, err := os.Create(output_file)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := out.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	// render template, write it to file
	t.Execute(out, settings)
}
