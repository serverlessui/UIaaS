package actions

import (
	"log"
	"os"

	"github.com/serverlessui/UIaaS/commands"
)

//Uploader is an interface defined to upload an application
type Uploader interface {
	UploadApplication(bucketName string, bucketPrefix string, dirPath string) error
}

//ServerlessUI struct to implement DeployAction
type ServerlessUI struct {
	Uploader Uploader
}

//Deploy method to deploy serverless UI
func (serverless ServerlessUI) Deploy(bucketInput *commands.BucketInput, appDir string) error {

	err := serverless.Uploader.UploadApplication(bucketInput.FullDomainName, "/", appDir)
	if err != nil {
		log.Println("error creating hosted zone ", err)
		os.Exit(1)
	}
	return nil
}
