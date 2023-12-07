package utils

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

func LoadImage(ctx echo.Context, fileName string) (string, error) {
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
}
