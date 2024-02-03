package utils

import (
	"bytes"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

func SaveImage(ctx echo.Context, fileName string, source *[]byte) error {
	if os.Getenv("GO_ENV") == "dev" {
		filePath := os.Getenv("UPLOAD_IMAGE_PATH")
		file, err := os.Create(filePath + fileName)
		if err != nil {
			return errors.WithStack(err)
		}
		defer file.Close()

		_, err = file.Write(*source)
		if err != nil {
			return errors.WithStack(err)
		}
		return nil
	} else {
		sess := session.Must(session.NewSession())
		uploader := s3manager.NewUploader(sess)

		bucketName := os.Getenv("S3_BUCKET_NAME")
		objectKey := fileName

		reader := bytes.NewReader(*source)

		_, err := uploader.Upload(&s3manager.UploadInput{
			Bucket: aws.String(bucketName),
			Key:    aws.String(objectKey),
			Body:   reader,
		})
		if err != nil {
			return errors.WithStack(err)
		}

		return nil
	}
}
