// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.15.0 DO NOT EDIT.
package openapi

import (
	openapi_types "github.com/oapi-codegen/runtime/types"
)

const (
	X_CSRF_TOKENScopes = "X_CSRF_TOKEN.Scopes"
	Google_authScopes  = "google_auth.Scopes"
	JwtAuthScopes      = "jwtAuth.Scopes"
	StateScopes        = "state.Scopes"
)

// Book defines model for Book.
type Book struct {
	ISBNcode    string `json:"ISBNcode"`
	Author      string `json:"author"`
	BookId      int    `json:"book_id"`
	Image       string `json:"image"`
	ItemUrl     string `json:"item_url"`
	Price       int    `json:"price"`
	PublishedAt string `json:"published_at"`
	Publisher   string `json:"publisher"`
	Title       string `json:"title"`
}

// BookGenreNode defines model for BookGenreNode.
type BookGenreNode struct {
	BooksGenreId   string          `json:"books_genre_id"`
	BooksGenreName string          `json:"books_genre_name"`
	Children       []BookGenreNode `json:"children"`
	GenreLevel     int             `json:"genre_level"`
	Id             int             `json:"id"`
	ParentGenreId  string          `json:"parent_genre_id"`
}

// ChatRoom defines model for ChatRoom.
type ChatRoom struct {
	LastMessage            string `json:"last_message"`
	LastMessageSentAt      string `json:"last_message_sent_at"`
	RoomId                 int    `json:"room_id"`
	TargetUserId           int    `json:"target_user_id"`
	TargetUserName         string `json:"target_user_name"`
	TargetUserProfileImage string `json:"target_user_profile_image"`
}

// Message defines model for Message.
type Message struct {
	Content   string `json:"content"`
	MessageId int    `json:"message_id"`
	SentAt    string `json:"sent_at"`
	UserId    int    `json:"user_id"`
}

// Post defines model for Post.
type Post struct {
	Book      Book    `json:"book"`
	Content   string  `json:"content"`
	CreatedAt string  `json:"created_at"`
	Image     *string `json:"image,omitempty"`
	PostId    int     `json:"post_id"`
	Rating    int     `json:"rating"`
	User      User    `json:"user"`
}

// ReqCreatePostBody defines model for ReqCreatePostBody.
type ReqCreatePostBody struct {
	ISBNcode string              `json:"ISBNcode"`
	Content  string              `json:"content"`
	Image    *openapi_types.File `json:"image,omitempty"`
	Rating   int                 `json:"rating"`
}

// ReqLoginBody defines model for ReqLoginBody.
type ReqLoginBody struct {
	Credential *string `json:"credential,omitempty"`
	Identifier *string `json:"identifier,omitempty"`
}

// ReqSignupBody defines model for ReqSignupBody.
type ReqSignupBody struct {
	Credential *string `json:"credential,omitempty"`
	Identifier *string `json:"identifier,omitempty"`
	Username   *string `json:"username,omitempty"`
}

// ResCsrfToken defines model for ResCsrfToken.
type ResCsrfToken struct {
	CsrfToken *string `json:"csrf_token,omitempty"`
}

// SendMessageReqBody defines model for SendMessageReqBody.
type SendMessageReqBody struct {
	Content string `json:"content"`
	RoomId  int    `json:"room_id"`
}

// UpdateUserReqBody defines model for UpdateUserReqBody.
type UpdateUserReqBody struct {
	Name         *string `json:"name,omitempty"`
	ProfileImage *string `json:"profile_image,omitempty"`
	ProfileText  *string `json:"profile_text,omitempty"`
}

// User defines model for User.
type User struct {
	Name         string `json:"name"`
	ProfileImage string `json:"profile_image"`
	UserId       int    `json:"user_id"`
}

// FetchBookDataParams defines parameters for FetchBookData.
type FetchBookDataParams struct {
	// BooksGenreId ID to specify the genre in Rakuten Books
	BooksGenreId *string `form:"booksGenreId,omitempty" json:"booksGenreId,omitempty"`

	// Keyword keyword to search books
	Keyword *string `form:"keyword,omitempty" json:"keyword,omitempty"`
}

// ChatSocketParams defines parameters for ChatSocket.
type ChatSocketParams struct {
	// RoomId ID to specify the chat room
	RoomId int `form:"room_id" json:"room_id"`
}

// LogoutJSONBody defines parameters for Logout.
type LogoutJSONBody = map[string]interface{}

// GetMessagesParams defines parameters for GetMessages.
type GetMessagesParams struct {
	// RoomId ID to specify the chat room
	RoomId int `form:"room_id" json:"room_id"`
}

// GoogleOauthCallbackParams defines parameters for GoogleOauthCallback.
type GoogleOauthCallbackParams struct {
	// State State parameter for CSRF protection
	State string `form:"state" json:"state"`

	// Code Authorization code returned by Google auth server
	Code string `form:"code" json:"code"`
}

// LoginJSONRequestBody defines body for Login for application/json ContentType.
type LoginJSONRequestBody = ReqLoginBody

// LogoutJSONRequestBody defines body for Logout for application/json ContentType.
type LogoutJSONRequestBody = LogoutJSONBody

// CreatePostMultipartRequestBody defines body for CreatePost for multipart/form-data ContentType.
type CreatePostMultipartRequestBody = ReqCreatePostBody

// SignupJSONRequestBody defines body for Signup for application/json ContentType.
type SignupJSONRequestBody = ReqSignupBody

// UpdateUserMultipartRequestBody defines body for UpdateUser for multipart/form-data ContentType.
type UpdateUserMultipartRequestBody = UpdateUserReqBody
