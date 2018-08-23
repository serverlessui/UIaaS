package stacks

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/serverlessui/UIaaS/commands"
)

const (
	createComplete = "CREATE_COMPLETE"
)

//Stack interface to define interacting with infra stacks
type Stack interface {
	Get(stackName string) (*cloudformation.Stack, error)
	CreateBucket(bucketInput *commands.BucketInput) error
	CreateDNS(dnsInput *commands.DNSInput) error
	WaitForStackCreation(stackName string) (*cloudformation.Stack, error)
}

//CloudformationStack is an implementation of the Stack interface
type CloudformationStack struct {
	Client          *http.Client
	URL             string
	CreateBucketURL string
	CreateDNSURL    string
}

//NotFoundError is an error corresponding to a resource not found
type NotFoundError struct {
	Resource string
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("Resource %s not found", e.Resource)
}

//CreateDNS method to create DNS record
func (cf CloudformationStack) CreateDNS(dnsInput *commands.DNSInput) error {
	log.Println("DEBUG: created url ", cf.CreateDNSURL)
	out, err := json.Marshal(dnsInput)
	if err != nil {
		return err
	}
	return post(out, cf.CreateDNSURL, cf.Client)

}

//WaitForStackCreation for stack creation completion
func (cf CloudformationStack) WaitForStackCreation(stackName string) (*cloudformation.Stack, error) {
	for {
		stack, err := cf.Get(stackName)
		if err != nil {
			return nil, err
		}
		if *stack.StackStatus == createComplete {
			return stack, nil
		}
		log.Println("Stack status ", stack.StackStatus, "...")
	}
}

//CreateBucket method to create Bucket
func (cf CloudformationStack) CreateBucket(bucketInput *commands.BucketInput) error {
	log.Println("DEBUG: created url ", cf.CreateBucketURL)
	out, err := json.Marshal(bucketInput)
	if err != nil {
		return err
	}
	return post(out, cf.CreateBucketURL, cf.Client)

}

//Get method to retrieve stack information
func (cf CloudformationStack) Get(stackName string) (*cloudformation.Stack, error) {
	url := strings.Replace(cf.URL, ":name", stackName, -1)

	log.Println("DEBUG: about to send request to url ", cf.URL)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("Content-Type", "application/json")

	resp, err := cf.Client.Do(req)
	if err != nil {
		log.Println("ERROR: error making http call to source with error", err)
		return &cloudformation.Stack{}, errors.New("error making http call")
	}
	if resp.StatusCode == http.StatusNotFound {
		return &cloudformation.Stack{}, NotFoundError{Resource: stackName}
	}
	defer resp.Body.Close()
	s := cloudformation.Stack{}
	err = json.NewDecoder(resp.Body).Decode(&s)

	if err != nil {
		log.Println("ERROR: error unmarshalling response ", err)
		return nil, errors.New("Unmarshal error")
	}

	return &s, nil
}

func is2xx(status *int) bool {
	switch *status {
	case 200:
		return true
	case 201:
		return true
	case 202:
		return true
	default:
		return false
	}
}

func post(out []byte, url string, client *http.Client) error {
	jsonString := string(out)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(jsonString)))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if !is2xx(&resp.StatusCode) {
		log.Println("client returned error ", resp)
		os.Exit(127)
	}
	return nil
}
