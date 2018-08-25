package config

import (
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/serverlessui/UIaaS/actions"
	"github.com/serverlessui/UIaaS/bucket"
	"github.com/serverlessui/UIaaS/commands"
	"github.com/serverlessui/UIaaS/dns"
	"github.com/serverlessui/UIaaS/fileutil"
	"github.com/serverlessui/UIaaS/stacks"
	"github.com/urfave/cli"
)

const (
	url             = "https://6456l3jgrl.execute-api.us-east-1.amazonaws.com/dev/stacks/:name"
	createBucketURL = "https://6456l3jgrl.execute-api.us-east-1.amazonaws.com/dev/sites/:name/cdns"
	createDNSURL    = "https://6456l3jgrl.execute-api.us-east-1.amazonaws.com/dev/sites/:name"
)

//CreateApp method to create initial app
func CreateApp() *cli.App {
	// Initialize a session in us-west-2 that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)

	if err != nil {
		log.Fatal("error creating session")
		os.Exit(1)
	}

	s3 := s3.New(sess)
	uploader := bucket.S3Uploader{Client: s3, FileUtil: fileutil.FileUtility{}}
	httpClient := http.Client{}

	stack := stacks.CloudformationStack{Client: &httpClient, URL: url, CreateBucketURL: createBucketURL, CreateDNSURL: createDNSURL}

	dns := dns.Route53{Stack: stack}
	bucket := bucket.S3Bucket{Stack: stack}
	deployAction := actions.ServerlessUI{Uploader: uploader, DNS: dns, Bucket: bucket}
	app := cli.NewApp()

	app.Name = "serverless-ui"
	app.Usage = "Command line interface for serverless ui deployment"
	app.Version = "0.0.1"
	app.Author = "VSS"
	app.Commands = []cli.Command{
		commands.Deploy(deployAction),
	}

	return app
}
