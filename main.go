package yaml2json

import (
	"fmt"

	"github.com/hokaccha/go-prettyjson"
	"gopkg.in/yaml.v2"
)

// YamlToJSON convert yaml format to json
func YamlToJSON(m []byte) (interface{}, error) {
	var body interface{}
	yaml.Unmarshal([]byte(m), &body)
	body = FormatJSON(body)
	return body, nil
}

// FormatJSON adds json format
func FormatJSON(i interface{}) interface{} {
	switch x := i.(type) {
	case map[interface{}]interface{}:
		m2 := map[string]interface{}{}
		for k, v := range x {
			m2[k.(string)] = FormatJSON(v)
		}
		return m2
	case []interface{}:
		for i, v := range x {
			x[i] = FormatJSON(v)
		}
	}
	return i
}

// PrintYaml prints yaml
func PrintYaml(in map[interface{}]interface{}) {
	output, _ := yaml.Marshal(in)
	fmt.Printf("%s", output)
}

// PrintJSON prints json
func PrintJSON(v interface{}) {
	output, _ := prettyjson.Marshal(v)
	fmt.Println(string(output))
}
