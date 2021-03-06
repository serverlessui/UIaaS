package bucket

import (
	"log"
	"strings"

	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/serverlessui/UIaaS/commands"
	"github.com/serverlessui/UIaaS/stacks"
)

const (
	//route53 param values
	domainNameParam     = "HostedZone"
	fullDomainNameParam = "FullDomainName"
	acmCertARNParam     = "AcmCertificateArn"
	ttlCacheValueParam  = "CacheValueTTL"
)

//S3Bucket is a struct to define needed S3 Bucket dependencies
type S3Bucket struct {
	Stack stacks.Stack
}

//DeployCDN is a function to Create an S3 Site with CDN and ACM
func (s3Bucket S3Bucket) DeployCDN(input *commands.BucketInput) (*cloudformation.Stack, error) {
	stackName := getStackName(input)
	stack, err := s3Bucket.Stack.Get(stackName)
	_, resourceNotFound := err.(stacks.NotFoundError)

	if resourceNotFound {
		s3Bucket.Stack.CreateBucket(input, stackName)
		log.Println("Cloudfront distribution creation begun. This may take up to 30 minutes...")
	} else {
		log.Println("DNS Stack already exists")
	}
	return stack, nil

}

//Method to convert DomainName from input to stack name
//route53 does not allow for full stop (.) characters
func getStackName(input *commands.BucketInput) string {
	return strings.Replace(input.FullDomainName, ".", "-", -1) + "-cdn"
}
