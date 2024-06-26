package denalijsontest

import (
	"testing"

	fr "github.com/Dharitri-org/drtg-vm-util/test-util/denali/json/fileresolver"
	mjparse "github.com/Dharitri-org/drtg-vm-util/test-util/denali/json/parse"
	mjwrite "github.com/Dharitri-org/drtg-vm-util/test-util/denali/json/write"
	"github.com/stretchr/testify/require"
)

func TestWriteScenario(t *testing.T) {
	contents, err := loadExampleFile("example.scen.json")
	require.Nil(t, err)

	p := mjparse.NewParser(
		fr.NewDefaultFileResolver().ReplacePath(
			"smart-contract.wasm",
			"exampleFile.txt"))

	scenario, parseErr := p.ParseScenarioFile(contents)
	require.Nil(t, parseErr)

	serialized := mjwrite.ScenarioToJSONString(scenario)

	// good for debugging:
	// _ = ioutil.WriteFile("example_re.scen.json", []byte(serialized), 0644)

	require.Equal(t, contents, []byte(serialized))
}
