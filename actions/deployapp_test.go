package actions

import (
	"errors"
	"testing"

	"github.com/serverlessui/UIaaS/commands"
)

type mockUploader struct {
}

func (mock mockUploader) UploadApplication(bucketName string, bucketPrefix string, dirPath string) error {
	return nil
}

type mockBadUploader struct {
}

func (mock mockBadUploader) UploadApplication(bucketName string, bucketPrefix string, dirPath string) error {
	return errors.New("ERROR")
}
func TestServerlessUIDeploy(t *testing.T) {
	ui := ServerlessUI{mockUploader{}}

	err := ui.Deploy(&commands.BucketInput{}, "")

	if err != nil {
		t.Log("error encountered when none expected ", err)
		t.Fail()
	}
}
