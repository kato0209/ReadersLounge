// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.15.0 DO NOT EDIT.
package openapi

import (
	"time"
)

const (
	X_CSRF_TOKENScopes = "X_CSRF_TOKEN.Scopes"
	Google_authScopes  = "google_auth.Scopes"
	JwtAuthScopes      = "jwtAuth.Scopes"
	StateScopes        = "state.Scopes"
)

// Book defines model for Book.
type Book struct {
	ISBNcode    *string    `json:"ISBNcode,omitempty"`
	Author      *string    `json:"author,omitempty"`
	BookId      *int       `json:"book_id,omitempty"`
	Image       *string    `json:"image,omitempty"`
	ItemUrl     *string    `json:"item_url,omitempty"`
	Price       *string    `json:"price,omitempty"`
	PublishedAt *time.Time `json:"published_at,omitempty"`
	Publisher   *string    `json:"publisher,omitempty"`
	Title       *string    `json:"title,omitempty"`
}

// Post defines model for Post.
type Post struct {
	Book      *Book      `json:"book,omitempty"`
	Content   *string    `json:"content,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	Image     *string    `json:"image,omitempty"`
	PostId    *int       `json:"post_id,omitempty"`
	Rating    *int       `json:"rating,omitempty"`
	User      *User      `json:"user,omitempty"`
}

// ReqCreatePostBody defines model for ReqCreatePostBody.
type ReqCreatePostBody struct {
	ISBNcode string  `json:"ISBNcode"`
	Content  string  `json:"content"`
	Image    *string `json:"image,omitempty"`
	Rating   int     `json:"rating"`
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

// User defines model for User.
type User struct {
	Name         *string `json:"name,omitempty"`
	ProfileImage *string `json:"profile_image,omitempty"`
	UserId       *int    `json:"user_id,omitempty"`
}

// LogoutJSONBody defines parameters for Logout.
type LogoutJSONBody = map[string]interface{}

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

// CreatePostJSONRequestBody defines body for CreatePost for application/json ContentType.
type CreatePostJSONRequestBody = ReqCreatePostBody

// SignupJSONRequestBody defines body for Signup for application/json ContentType.
type SignupJSONRequestBody = ReqSignupBody
