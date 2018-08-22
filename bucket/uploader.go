package bucket

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/larse514/serverlessui/serverless-ui/fileutil"
)

//S3Uploader is a struct to upload an application to S3
type S3Uploader struct {
	Client   s3iface.S3API
	FileUtil fileutil.FileUtil
}

//UploadApplication is a method to upload an application to s3 bucket
func (uploader S3Uploader) UploadApplication(bucketName string, bucketPrefix string, dirPath string) error {
	//retrieve list of files to be uploaded
	fileList := uploader.FileUtil.GetFilesInDirectory(dirPath)
	log.Println("file list ", fileList)
	//upload each
	for _, file := range fileList {
		log.Println("about to upload file ", file, " to ", bucketName, " ", bucketPrefix)
		uploadFileToS3(uploader.Client, bucketName, bucketPrefix, file, dirPath)
	}
	return nil
}

func uploadFileToS3(client s3iface.S3API, bucketName string, bucketPrefix string, fileName string, root string) error {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Failed to open file", file, err)
		return errors.New("Error opening file")
	}
	//defer close until after method returns
	defer file.Close()
	key := strings.Replace(fileName, root, "", -1)
	// Create S3 upload parameters
	params := &s3.PutObjectInput{
		Bucket:      aws.String(bucketName), // Required
		Key:         aws.String(key),        // Required
		ContentType: aws.String("text/html"),
		Body:        file,
	}
	//perform the upload
	_, err = client.PutObject(params)
	if err != nil {
		fmt.Printf("Failed to upload data to %s/%s, %s\n",
			bucketName, key, err.Error())
		return errors.New("Error uploading object")
	}
	return nil
}
