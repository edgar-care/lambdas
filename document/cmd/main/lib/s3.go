package lib

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func UploadToS3(file io.Reader, filename string) (string, error) {
	awsAccessKeyID := os.Getenv("EDGAR_ACCESS_KEY_ID")
	awsSecretAccessKey := os.Getenv("EDGAR_SECRET_ACCESS_KEY")
	region := os.Getenv("EDGAR_REGION")
	bucketName := "document-patient"

	// Validate the filename
	if filename == "" {
		return "", fmt.Errorf("empty filename")
	}

	// Check for an empty file
	var buf bytes.Buffer
	_, err := io.Copy(&buf, file)
	if err != nil {
		return "", fmt.Errorf("error reading file: %w", err)
	}

	if buf.Len() == 0 {
		return "", fmt.Errorf("empty file content")
	}

	// Set the content type based on the file extension
	contentType := "application/octet-stream"
	if ext := filepath.Ext(filename); ext != "" {
		contentType = mime.TypeByExtension(ext)
	}

	cfg := aws.NewConfig().WithRegion(region)
	if awsAccessKeyID != "" && awsSecretAccessKey != "" {
		cfg = cfg.WithCredentials(credentials.NewStaticCredentials(awsAccessKeyID, awsSecretAccessKey, ""))
	}

	sess, err := session.NewSession(cfg)
	if err != nil {
		return "", fmt.Errorf("error creating session: %w", err)
	}

	s3Client := s3.New(sess)

	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(bucketName),
		Key:         aws.String(filename),
		Body:        bytes.NewReader(buf.Bytes()),
		ContentType: aws.String(contentType),
	})
	if err != nil {
		log.Printf("Error uploading file to S3: %v", err)
		return "", fmt.Errorf("error uploading file to S3: %w", err)
	}

	downloadURL := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", bucketName, filename)
	return downloadURL, nil
}
