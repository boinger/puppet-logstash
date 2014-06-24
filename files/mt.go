// to cross compile from linux:
// GOOS=windows GOARCH=386 go build -o mt.exe mt.go

package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"text/template"
)

func main() {

	// read arguments
	if len(os.Args) < 4 {
		log.Fatal("\n\n./mt.go json_settings_file template_input_file output_file\n\n")
	}

	settings_file := os.Args[1]
	template_file := os.Args[2]
	output_file := os.Args[3]

	log.Println("settings_file =", settings_file, "template_file =", template_file, "output_file =", output_file)

	// read and parse the settings json file
	settings_json, err := ioutil.ReadFile(settings_file)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(settings_json))

	settings := make(map[string]string)

	err = json.Unmarshal(settings_json, &settings)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(settings)

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
