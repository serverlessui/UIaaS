package dns

import (
	"log"
	"strings"

	"github.com/aws/aws-sdk-go/service/cloudformation"

	"github.com/serverlessui/UIaaS/commands"
	"github.com/serverlessui/UIaaS/stacks"
)

const (
	websiteArnOutput = "WebsiteCertArn"
)

//Route53 is an implementation of the DNS interface
type Route53 struct {
	Stack stacks.Stack
}

//Route53Output struct containing output from Route53
type Route53Output struct {
	WebsiteArn string
}

//DeployHostedZone Method to create Route53 hosted zone
func (route53 Route53) DeployHostedZone(input *commands.DNSInput) (*Route53Output, error) {
	//replace domain name
	stackName := getStackName(input)
	stack, err := route53.Stack.Get(stackName)
	_, resourceNotFound := err.(stacks.NotFoundError)
	websiteOutputValue := input.Environment + "-" + websiteArnOutput

	if resourceNotFound {
		route53.Stack.CreateDNS(input)
	} else {
		log.Println("DNS Stack already exists")
		return &Route53Output{WebsiteArn: getOutputValue(stack, websiteOutputValue)}, nil
	}
	stack, err = route53.Stack.WaitForStackCreation(stackName)

	return &Route53Output{WebsiteArn: getOutputValue(stack, websiteOutputValue)}, nil
}

//Method to convert DomainName from input to stack name
//route53 does not allow for full stop (.) characters
func getStackName(input *commands.DNSInput) string {
	return strings.Replace(input.HostedZone, ".", "-", -1)
}

//getOutputValue method will retrieve an output value from Output array
func getOutputValue(stack *cloudformation.Stack, key string) string {
	for i := range stack.Outputs {
		if *stack.Outputs[i].ExportName == key {
			// Found!
			return *stack.Outputs[i].OutputValue
		}
	}
	return ""
}
