package encodingtools

import (
	"os"
	"testing"
)

func TestConvertion(t *testing.T) {
	json := `{"A":1,"F":1.2,"O":{"foo":"bar"},"S":"foo","T":[1,2,3]}`

	// Nominal JSON2Yaml
	data, err := JSON2Yaml([]byte(json))

	if err != nil {
		t.Errorf("JSON2Yaml: error unexpected %v", err)
	}


	// Nominal Yaml2JSON
	data, err = Yaml2JSON(data)

	if err != nil {
		t.Errorf("Yaml2JSON: error unexpected %v", err)
	}
	if string(data) != json {
		t.Errorf("Yaml2JSON: bad yaml format \n%s vs \n%s", string(data), json)
	}
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
