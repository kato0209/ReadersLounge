package utils

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

func SaveImage(ctx echo.Context, fileName string, source *[]byte) error {
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
}
