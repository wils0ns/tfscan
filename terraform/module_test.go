package terraform_test

import (
	"testing"

	"github.com/wils0ns/tfscan/terraform"
	"github.com/wils0ns/tfscan/testdata"
)

type visitor struct {
	ResCount int
}

func (v *visitor) Visit(m, p *terraform.Module) {
	v.ResCount += len(m.Resources)
}

func TestVisitModules(t *testing.T) {
	state, err := terraform.NewState(testdata.StateReader)
	if err != nil {
		t.Fatal(err)
	}

	v := &visitor{}
	state.Values["root_module"].VisitModules(v, nil)

	expectedResCount := 21
	if v.ResCount != expectedResCount {
		t.Errorf("Expected %v resources, got %v", expectedResCount, v.ResCount)
	}
}
