package actions

import (
	"errors"
	"os"
	"os/exec"
	"testing"

	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/serverlessui/UIaaS/commands"
	"github.com/serverlessui/UIaaS/dns"
)

type mockBucket struct {
}

func (mock mockBucket) DeploySite(input *commands.BucketInput) (*cloudformation.Stack, error) {
	return nil, nil
}

type mockBadBucket struct {
}

func (mock mockBadBucket) DeploySite(input *commands.BucketInput) (*cloudformation.Stack, error) {
	return nil, errors.New("")
}

type mockDNS struct {
}

func (mock mockDNS) DeployHostedZone(input *commands.DNSInput) (*dns.Route53Output, error) {
	return &dns.Route53Output{WebsiteArn: "SOMEARN"}, nil
}

type mockBadDNS struct {
}

func (mock mockBadDNS) DeployHostedZone(input *commands.DNSInput) (*dns.Route53Output, error) {
	return nil, errors.New("")
}

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
	ui := ServerlessUI{mockUploader{}, mockBucket{}, mockDNS{}}

	err := ui.Deploy(&commands.DNSInput{}, &commands.BucketInput{}, "")

	if err != nil {
		t.Log("error encountered when none expected ", err)
		t.Fail()
	}
}

func TestServerlessUIDeployErrorProcessExitWithCode1(t *testing.T) {
	ui := ServerlessUI{mockUploader{}, mockBucket{}, mockBadDNS{}}

	if os.Getenv("BE_CRASHER") == "1" {
		ui.Deploy(&commands.DNSInput{}, &commands.BucketInput{}, "")
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestServerlessUIDeployErrorProcessExitWithCode1")
	cmd.Env = append(os.Environ(), "BE_CRASHER=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("process ran with err %v, want exit status 1", err)

}
func TestServerlessUIDeployBadUploaderErrorProcessExitWithCode1(t *testing.T) {
	ui := ServerlessUI{mockUploader{}, mockBadBucket{}, mockDNS{}}

	if os.Getenv("BE_CRASHER") == "1" {
		ui.Deploy(&commands.DNSInput{}, &commands.BucketInput{}, "")
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestServerlessUIDeployBadUploaderErrorProcessExitWithCode1")
	cmd.Env = append(os.Environ(), "BE_CRASHER=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("process ran with err %v, want exit status 1", err)

}
