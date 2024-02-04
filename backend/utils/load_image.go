package utils

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

func LoadImage(ctx echo.Context, fileName string) (string, error) {
	if os.Getenv("GO_ENV") == "dev" {
		filePath := os.Getenv("UPLOAD_IMAGE_PATH")
		file, err := os.Open(filePath + fileName)
		if err != nil {
			return "", errors.WithStack(err)
		}
		defer file.Close()

		img, _, err := image.Decode(file)
		if err != nil {
			return "", errors.WithStack(err)
		}

		split := strings.Split(fileName, ".")
		var buf bytes.Buffer
		switch split[len(split)-1] {
		case "png":
			if err := png.Encode(&buf, img); err != nil {
				return "", errors.WithStack(err)
			}
		case "jpeg":
			if err := jpeg.Encode(&buf, img, &jpeg.Options{Quality: 100}); err != nil {
				return "", errors.WithStack(err)
			}
		default:
		}

		return base64.StdEncoding.EncodeToString(buf.Bytes()), nil
	} else {
		sess := session.Must(session.NewSession())
		bucketName := os.Getenv("S3_BUCKET_NAME")
		objectKey := fileName

		downloader := s3manager.NewDownloader(sess)
		buf := aws.NewWriteAtBuffer([]byte{})
		_, err := downloader.Download(buf, &s3.GetObjectInput{
			Bucket: aws.String(bucketName),
			Key:    aws.String(objectKey),
		})
		if err != nil {
			return "", errors.WithStack(err)
		}

		return base64.StdEncoding.EncodeToString(buf.Bytes()), nil
	}
}
