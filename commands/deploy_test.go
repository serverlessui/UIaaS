package commands

import (
	"reflect"
	"testing"

	"github.com/serverlessui/UIaaS/flags"
)

type mockDeployAction struct {
}

func (mock mockDeployAction) Deploy(bucketInput *BucketInput, appDir string) error {
	return nil
}

func TestDeployName(t *testing.T) {
	mock := mockDeployAction{}
	expected := "deploy"
	actual := Deploy(mock).Name

	if actual != expected {
		t.Log("Incorrect name expected ", expected, " got ", actual)
	}
}

func TestDeployAliases(t *testing.T) {
	mock := mockDeployAction{}
	expected := "d"
	actual := Deploy(mock).Aliases[0]

	if actual != expected {
		t.Log("Incorrect name expected ", expected, " got ", actual)
	}
}

func TestDeployUsage(t *testing.T) {
	mock := mockDeployAction{}
	expected := flags.Deploy()
	actual := Deploy(mock).Flags

	if reflect.DeepEqual(actual, expected) {
		t.Log("Incorrect name expected ", expected, " got ", actual)
	}
}
