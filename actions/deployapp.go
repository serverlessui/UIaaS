package actions

import (
	"errors"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/serverlessui/UIaaS/commands"
	"github.com/serverlessui/UIaaS/dns"
)

//Bucket is an interface to define creation of Bucket based sites
type Bucket interface {
	DeployCDN(input *commands.BucketInput) (*cloudformation.Stack, error)
}

//DNS is an interface to represent Cloud DNS Services
type DNS interface {
	DeploySite(input *commands.DNSInput) (*dns.Route53Output, error)
}

//Uploader is an interface defined to upload an application
type Uploader interface {
	UploadApplication(bucketName string, bucketPrefix string, dirPath string) error
}

//ServerlessUI struct to implement DeployAction
type ServerlessUI struct {
	Uploader Uploader
	Bucket   Bucket
	DNS      DNS
}

//Deploy method to deploy serverless UI
func (serverless ServerlessUI) Deploy(dnsInput *commands.DNSInput, bucketInput *commands.BucketInput, appDir string) error {
	output, err := serverless.DNS.DeploySite(dnsInput)
	if err != nil {
		log.Println("error creating hosted zone ", err)
		os.Exit(1)
	}
	//grab the arn output so we don't have to have the user provide it
	if output.WebsiteArn == "" {
		return errors.New("Failed to retrieve Certificate")
	}
	if output.WebsiteURL == "" {
		return errors.New("Failed to retrieve WebsiteURL")
	}
	bucketInput.AcmCertificateArn = output.WebsiteArn
	bucketInput.WebsiteURL = output.WebsiteURL
	_, err = serverless.Bucket.DeployCDN(bucketInput)
	if err != nil {
		log.Println("error creating hosted zone ", err)
		os.Exit(1)
	}
	return deploy(serverless.Uploader, bucketInput, appDir)
}

func deploy(uploader Uploader, bucketInput *commands.BucketInput, appDir string) error {
	err := uploader.UploadApplication(bucketInput.FullDomainName, "/", appDir)
	if err != nil {
		log.Println("error creating hosted zone ", err)
		os.Exit(1)
	}
	return nil
}
