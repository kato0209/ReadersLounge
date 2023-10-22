package validator

import (
	"backend/models"
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type IUserValidator interface {
	SignupValidator(user models.User) error
	LoginValidator(user models.User) error
}

type userValidator struct{}

func NewUserValidator() IUserValidator {
	return &userValidator{}
}

// Email Password方式のバリデーション
func (uv *userValidator) SignupValidator(user models.User) error {
	return validation.ValidateStruct(&user,
		validation.Field(
			&user.Identifier,
			validation.Required.Error("email is required"),
			validation.RuneLength(1, 100).Error("limited max 100 char"),
			is.Email.Error("is not valid email format"),
		),
		validation.Field(
			&user.Credential,
			validation.Required.Error("password is required"),
			validation.By(passwordValidator),
		),
		validation.Field(
			&user.Name,
			validation.Required.Error("username is required"),
			validation.RuneLength(1, 20).Error("limited max 20 char"),
		),
	)
}

func (uv *userValidator) LoginValidator(user models.User) error {
	return validation.ValidateStruct(&user,
		validation.Field(
			&user.Identifier,
			validation.Required.Error("email is required"),
			validation.RuneLength(1, 100).Error("limited max 100 char"),
			is.Email.Error("is not valid email format"),
		),
		validation.Field(
			&user.Credential,
			validation.Required.Error("password is required"),
			validation.By(passwordValidator),
		),
	)
}

func passwordValidator(value interface{}) error {
	password, _ := value.(string)

	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasDigit := regexp.MustCompile(`\d`).MatchString(password)

	if !hasLower || !hasUpper || !hasDigit || len(password) < 8 || len(password) > 100 {
		return validation.NewError("invalid_password", "Please set a password of at least 8 characters and no more than 100 characters, including at least one each of lower and upper case alphabetical and numeric characters.")
	}

	return nil
}
