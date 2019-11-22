package yaml2json

import (
	"testing"
)

func TestMyFunc(t *testing.T) {
	body, err := YamlToJSON(testData())
	if err != nil || body == nil {
		t.Error("Yaml convertion failed")
	}
}

func testData() []byte {
	data := `value1:
  value1a: aaaa
value2:
  value2a: aaaa
  value2b: bbbb`

	return []byte(data)
}
