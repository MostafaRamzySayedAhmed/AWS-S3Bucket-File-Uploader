package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	// Set up AWS session
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2")},
	)
	if err != nil {
		log.Fatal(err)
	}

	// Create an S3 service client
	svc := s3.New(sess)

	// Open the file
	file, err := os.Open("example.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a buffer to read the file into
	buf := new(bytes.Buffer)
	buf.ReadFrom(file)

	// Upload the file to S3
	_, err = svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String("your-bucket-name"),
		Key:    aws.String("example.txt"),
		Body:   bytes.NewReader(buf.Bytes()),
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("File uploaded successfully!")
}
