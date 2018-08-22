package flags

import (
	"testing"
)

func TestDeploy(t *testing.T) {
	flags := Deploy()

	if len(flags) != 7 {
		t.Log("Flags returned more than expected")
		t.Fail()
	}
}
