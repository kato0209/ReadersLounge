package response

import "backend/models"

type UserWithProfileImage struct {
	models.User
	ProfileImageFileName string `db:"profile_image"`
}
