package main

import (
	"log"

	"github.com/ssuareza/yaml2json"
)

func main() {
	input, err := readInput()
	if err != nil {
		log.Fatal(err)
	}

	// get json
	j, err := yaml2json.YamlToJSON(input)
	if err != nil {
		log.Fatal(err)
	}

	// print json
	yaml2json.PrintJSON(j)
}
