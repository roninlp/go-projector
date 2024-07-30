package cli_test

import (
	"testing"

	"github.com/theprimeagen/projector/pkg/cli"
)

func getData() *cli.Data {
	return &cli.Data{
		Projector: map[string]map[string]string{
			"/": {
				"foo": "bar1",
				"fem": "is_great",
			},
			"/foo": {
				"foo": "bar2",
			},
			"/foo/bar": {
				"foo": "bar3",
			},
		},
	}
}

func getProjector(pwd string, data *cli.Data) *cli.Projector {
	return cli.CreateProjector(
		&cli.Config{
			Args:      []string{},
			Operation: cli.Print,
			Pwd:       pwd,
			Config:    "Hello, Frontend Masters",
		},
		data,
	)
}

func test(t *testing.T, proj *cli.Projector, key, value string) {
	v, ok := proj.GetValue(key)
	if !ok {
		t.Errorf("expected to find value %v", key)
	}

	if value != v {
		t.Errorf("expected value %v but got %v", value, v)
	}
}

func TestGetValue(t *testing.T) {
	pwd := "/foo/bar"
	data := getData()
	projector := getProjector(pwd, data)

	test(t, projector, "foo", "bar3")
	test(t, projector, "fem", "is_great")
}

func TestSetValue(t *testing.T) {
	data := getData()
	proj := getProjector("/foo/bar", data)

	test(t, proj, "foo", "bar3")
	proj.SetValue("foo", "bar4")
	test(t, proj, "foo", "bar4")

	proj.SetValue("fem", "is_super_great")
	test(t, proj, "fem", "is_super_great")

	proj = getProjector("/", data)
	test(t, proj, "fem", "is_great")
}

func TestRemoveValue(t *testing.T) {
	data := getData()
	proj := getProjector("/foo/bar", data)

	test(t, proj, "foo", "bar3")
	proj.RemoveValue("foo")
	test(t, proj, "foo", "bar2")
	proj.RemoveValue("fem")
	test(t, proj, "fem", "is_great")
}
