package encodingtools

/*
 Base package for tools used to manipulate data in different format : json, yaml, binary (from protobuf def)
*/
import (
	"encoding/json"

	"gopkg.in/yaml.v3"
)

//YAML to JSON Converter

// y2jConvert : internal function used by yaml2Json
func y2jConvert(i interface{}) interface{} {
	switch x := i.(type) {
	case map[interface{}]interface{}:
		m2 := map[string]interface{}{}
		for k, v := range x {
			m2[k.(string)] = y2jConvert(v)
		}
		return m2
	case []interface{}:
		for i, v := range x {
			x[i] = y2jConvert(v)
		}
	}
	return i
}

// Yaml2Json : credits : stackoverflow ;)
func Yaml2JSON(buf []byte) ([]byte, error) {
	var body interface{}
	if err := yaml.Unmarshal(buf, &body); err != nil {
		return nil, err
	}

	body = y2jConvert(body)

	b, err := json.Marshal(body)

	return b, err
}
