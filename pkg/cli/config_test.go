package cli_test

import (
	"reflect"
	"testing"

	"github.com/theprimeagen/projector/pkg/cli"
)

func getOpts(args []string) *cli.Opts {
	opts := &cli.Opts{
		Args:   args,
		Config: "",
		Pwd:    "",
	}
	return opts
}

func testConfig(t *testing.T, args []string, expectedArgs []string, operation cli.Operation) {
	opts := getOpts(args)

	config, err := cli.NewConfig(opts)
	if err != nil {
		t.Errorf("expected to get no error %v", err)
	}

	if !reflect.DeepEqual(expectedArgs, config.Args) {
		t.Errorf("expected args to be %+v but got %+v", expectedArgs, config.Args)
	}

	if config.Operation != operation {
		t.Errorf("operation expected was %v but got %v", operation, config.Operation)
	}
}

func TestConfigPrint(t *testing.T) {
	testConfig(t, []string{}, []string{}, cli.Print)
}

func TestConfigPrintKey(t *testing.T) {
	testConfig(t, []string{"foo"}, []string{"foo"}, cli.Print)
}

func TestConfigAddKeyValue(t *testing.T) {
	testConfig(t, []string{"add", "foo", "bar"}, []string{"foo", "bar"}, cli.Add)
}

func TestConfigRemoveKey(t *testing.T) {
	testConfig(t, []string{"remove", "foo"}, []string{"foo"}, cli.Remove)
}
