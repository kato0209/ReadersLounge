// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.15.0 DO NOT EDIT.
package openapi

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// fetch book data from RakutenAPI
	// (GET /books)
	FetchBookData(ctx echo.Context, params FetchBookDataParams) error
	// Get books genres
	// (GET /books-genres)
	GetBooksGenres(ctx echo.Context) error
	// return users chat room list
	// (GET /chat-rooms)
	GetChatRooms(ctx echo.Context) error
	// create chat room
	// (POST /chat-rooms)
	CreateChatRoom(ctx echo.Context) error
	// WebSocket Connection for chat
	// (GET /chats)
	ChatSocket(ctx echo.Context, params ChatSocketParams) error
	// Create like of Comment
	// (POST /comment-likes)
	CreateCommentLike(ctx echo.Context) error
	// delete like of Comment
	// (DELETE /comment-likes/comment/{CommentId})
	DeleteCommentLike(ctx echo.Context, commentId int) error
	// create comment
	// (POST /comments)
	CreateComment(ctx echo.Context) error
	// get comments by postID
	// (GET /comments/post/{postId})
	GetCommentsByPostID(ctx echo.Context, postId int) error
	// delete comment
	// (DELETE /comments/{commentId})
	DeleteComment(ctx echo.Context, commentId int) error
	// get csrf token
	// (GET /csrftoken)
	Csrftoken(ctx echo.Context) error
	// get followers connections
	// (GET /followers)
	GetFollowerConnections(ctx echo.Context, params GetFollowerConnectionsParams) error
	// get following connections
	// (GET /followings)
	GetFollowingConnections(ctx echo.Context, params GetFollowingConnectionsParams) error
	// create connection of following
	// (POST /follows)
	CreateConnection(ctx echo.Context) error
	// delete connection of following
	// (DELETE /follows/{connectionId})
	DeleteConnection(ctx echo.Context, connectionId int) error
	// Health Check
	// (GET /health)
	HealthCheck(ctx echo.Context) error
	// Get commentID list of User liked
	// (GET /liked-comments)
	GetLikedCommentList(ctx echo.Context) error
	// Get postID list of User liked
	// (GET /liked-posts)
	GetLikedPostList(ctx echo.Context) error
	// login
	// (POST /login)
	Login(ctx echo.Context) error
	// logout
	// (POST /logout)
	Logout(ctx echo.Context) error
	// return messages in a chat room
	// (GET /messages)
	GetMessages(ctx echo.Context, params GetMessagesParams) error
	// Callback for Google OAuth
	// (GET /oauth/google/callback)
	GoogleOauthCallback(ctx echo.Context, params GoogleOauthCallbackParams) error
	// Create like of Post
	// (POST /post-likes)
	CreatePostLike(ctx echo.Context) error
	// delete like of Post
	// (DELETE /post-likes/post/{PostId})
	DeletePostLike(ctx echo.Context, postId int) error
	// get posts
	// (GET /posts)
	GetPosts(ctx echo.Context) error
	// create post
	// (POST /posts)
	CreatePost(ctx echo.Context) error
	// get posts corresponding to userId
	// (GET /posts/user/{userId})
	GetPostsOfUser(ctx echo.Context, userId int) error
	// delete a post
	// (DELETE /posts/{postId})
	DeletePost(ctx echo.Context, postId int) error
	// get post by postID
	// (GET /posts/{postId})
	GetPostByPostID(ctx echo.Context, postId int) error
	// search user by keyword
	// (GET /search-user)
	SearchUser(ctx echo.Context, params SearchUserParams) error
	// create new user
	// (POST /signup)
	Signup(ctx echo.Context) error
	// get user
	// (GET /user)
	GetLoginUser(ctx echo.Context) error
	// update user
	// (PUT /user)
	UpdateUser(ctx echo.Context) error
	// get user by user_id
	// (GET /user/{userId})
	GetUser(ctx echo.Context, userId int) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// FetchBookData converts echo context to params.
func (w *ServerInterfaceWrapper) FetchBookData(ctx echo.Context) error {
	var err error

	ctx.Set(X_CSRF_TOKENScopes, []string{})

	ctx.Set(JwtAuthScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params FetchBookDataParams
	// ------------- Optional query parameter "booksGenreId" -------------

	err = runtime.BindQueryParameter("form", true, false, "booksGenreId", ctx.QueryParams(), &params.BooksGenreId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter booksGenreId: %s", err))
	}

	// ------------- Optional query parameter "keyword" -------------

	err = runtime.BindQueryParameter("form", true, false, "keyword", ctx.QueryParams(), &params.Keyword)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter keyword: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.FetchBookData(ctx, params)
	return err
}

// GetBooksGenres converts echo context to params.
func (w *ServerInterfaceWrapper) GetBooksGenres(ctx echo.Context) error {
	var err error

	ctx.Set(X_CSRF_TOKENScopes, []string{})

	ctx.Set(JwtAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetBooksGenres(ctx)
	return err
}

// GetChatRooms converts echo context to params.
func (w *ServerInterfaceWrapper) GetChatRooms(ctx echo.Context) error {
	var err error

	ctx.Set(X_CSRF_TOKENScopes, []string{})

	ctx.Set(JwtAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetChatRooms(ctx)
	return err
}

// CreateChatRoom converts echo context to params.
func (w *ServerInterfaceWrapper) CreateChatRoom(ctx echo.Context) error {
	var err error

	ctx.Set(X_CSRF_TOKENScopes, []string{})

	ctx.Set(JwtAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.CreateChatRoom(ctx)
	return err
}

// ChatSocket converts echo context to params.
func (w *ServerInterfaceWrapper) ChatSocket(ctx echo.Context) error {
	var err error

	ctx.Set(X_CSRF_TOKENScopes, []string{})

	ctx.Set(JwtAuthScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params ChatSocketParams
	// ------------- Required query parameter "room_id" -------------

	err = runtime.BindQueryParameter("form", true, true, "room_id", ctx.QueryParams(), &params.RoomId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter room_id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.ChatSocket(ctx, params)
	return err
}

// CreateCommentLike converts echo context to params.
func (w *ServerInterfaceWrapper) CreateCommentLike(ctx echo.Context) error {
	var err error

	ctx.Set(X_CSRF_TOKENScopes, []string{})

	ctx.Set(JwtAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.CreateCommentLike(ctx)
	return err
}

// DeleteCommentLike converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteCommentLike(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "CommentId" -------------
	var commentId int

	err = runtime.BindStyledParameterWithLocation("simple", false, "CommentId", runtime.ParamLocationPath, ctx.Param("CommentId"), &commentId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter CommentId: %s", err))
	}

	ctx.Set(X_CSRF_TOKENScopes, []string{})

	ctx.Set(JwtAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.DeleteCommentLike(ctx, commentId)
	return err
}

// CreateComment converts echo context to params.
func (w *ServerInterfaceWrapper) CreateComment(ctx echo.Context) error {
	var err error

	ctx.Set(X_CSRF_TOKENScopes, []string{})

	ctx.Set(JwtAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.CreateComment(ctx)
	return err
}

// GetCommentsByPostID converts echo context to params.
func (w *ServerInterfaceWrapper) GetCommentsByPostID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "postId" -------------
	var postId int

	err = runtime.BindStyledParameterWithLocation("simple", false, "postId", runtime.ParamLocationPath, ctx.Param("postId"), &postId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter postId: %s", err))
	}

	ctx.Set(X_CSRF_TOKENScopes, []string{})

	ctx.Set(JwtAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetCommentsByPostID(ctx, postId)
	return err
}

// DeleteComment converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteComment(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "commentId" -------------
	var commentId int

	err = runtime.BindStyledParameterWithLocation("simple", false, "commentId", runtime.ParamLocationPath, ctx.Param("commentId"), &commentId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter commentId: %s", err))
	}

	ctx.Set(X_CSRF_TOKENScopes, []string{})

	ctx.Set(JwtAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.DeleteComment(ctx, commentId)
	return err
}

// Csrftoken converts echo context to params.
func (w *ServerInterfaceWrapper) Csrftoken(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.Csrftoken(ctx)
	return err
}

// GetFollowerConnections converts echo context to params.
func (w *ServerInterfaceWrapper) GetFollowerConnections(ctx echo.Context) error {
	var err error

	ctx.Set(X_CSRF_TOKENScopes, []string{})

	ctx.Set(JwtAuthScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetFollowerConnectionsParams
	// ------------- Required query parameter "user_id" -------------

	err = runtime.BindQueryParameter("form", true, true, "user_id", ctx.QueryParams(), &params.UserId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter user_id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetFollowerConnections(ctx, params)
	return err
}

// GetFollowingConnections converts echo context to params.
func (w *ServerInterfaceWrapper) GetFollowingConnections(ctx echo.Context) error {
	var err error

	ctx.Set(X_CSRF_TOKENScopes, []string{})

	ctx.Set(JwtAuthScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetFollowingConnectionsParams
	// ------------- Required query parameter "user_id" -------------

	err = runtime.BindQueryParameter("form", true, true, "user_id", ctx.QueryParams(), &params.UserId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter user_id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetFollowingConnections(ctx, params)
	return err
}

// CreateConnection converts echo context to params.
func (w *ServerInterfaceWrapper) CreateConnection(ctx echo.Context) error {
	var err error

	ctx.Set(X_CSRF_TOKENScopes, []string{})

	ctx.Set(JwtAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.CreateConnection(ctx)
	return err
}

// DeleteConnection converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteConnection(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "connectionId" -------------
	var connectionId int

	err = runtime.BindStyledParameterWithLocation("simple", false, "connectionId", runtime.ParamLocationPath, ctx.Param("connectionId"), &connectionId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter connectionId: %s", err))
	}

	ctx.Set(X_CSRF_TOKENScopes, []string{})

	ctx.Set(JwtAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.DeleteConnection(ctx, connectionId)
	return err
}

// HealthCheck converts echo context to params.
func (w *ServerInterfaceWrapper) HealthCheck(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.HealthCheck(ctx)
	return err
}

// GetLikedCommentList converts echo context to params.
func (w *ServerInterfaceWrapper) GetLikedCommentList(ctx echo.Context) error {
	var err error

	ctx.Set(X_CSRF_TOKENScopes, []string{})

	ctx.Set(JwtAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetLikedCommentList(ctx)
	return err
}

// GetLikedPostList converts echo context to params.
func (w *ServerInterfaceWrapper) GetLikedPostList(ctx echo.Context) error {
	var err error

	ctx.Set(X_CSRF_TOKENScopes, []string{})

	ctx.Set(JwtAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetLikedPostList(ctx)
	return err
}

// Login converts echo context to params.
func (w *ServerInterfaceWrapper) Login(ctx echo.Context) error {
	var err error

	ctx.Set(X_CSRF_TOKENScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.Login(ctx)
	return err
}

// Logout converts echo context to params.
func (w *ServerInterfaceWrapper) Logout(ctx echo.Context) error {
	var err error

	ctx.Set(X_CSRF_TOKENScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.Logout(ctx)
	return err
}

// GetMessages converts echo context to params.
func (w *ServerInterfaceWrapper) GetMessages(ctx echo.Context) error {
	var err error

	ctx.Set(X_CSRF_TOKENScopes, []string{})

	ctx.Set(JwtAuthScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetMessagesParams
	// ------------- Required query parameter "room_id" -------------

	err = runtime.BindQueryParameter("form", true, true, "room_id", ctx.QueryParams(), &params.RoomId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter room_id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetMessages(ctx, params)
	return err
}

// GoogleOauthCallback converts echo context to params.
func (w *ServerInterfaceWrapper) GoogleOauthCallback(ctx echo.Context) error {
	var err error

	ctx.Set(Google_authScopes, []string{"email", "profile"})

	ctx.Set(StateScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params GoogleOauthCallbackParams
	// ------------- Required query parameter "state" -------------

	err = runtime.BindQueryParameter("form", true, true, "state", ctx.QueryParams(), &params.State)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter state: %s", err))
	}

	// ------------- Required query parameter "code" -------------

	err = runtime.BindQueryParameter("form", true, true, "code", ctx.QueryParams(), &params.Code)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter code: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GoogleOauthCallback(ctx, params)
	return err
}

// CreatePostLike converts echo context to params.
func (w *ServerInterfaceWrapper) CreatePostLike(ctx echo.Context) error {
	var err error

	ctx.Set(X_CSRF_TOKENScopes, []string{})

	ctx.Set(JwtAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.CreatePostLike(ctx)
	return err
}

// DeletePostLike converts echo context to params.
func (w *ServerInterfaceWrapper) DeletePostLike(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "PostId" -------------
	var postId int

	err = runtime.BindStyledParameterWithLocation("simple", false, "PostId", runtime.ParamLocationPath, ctx.Param("PostId"), &postId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter PostId: %s", err))
	}

	ctx.Set(X_CSRF_TOKENScopes, []string{})

	ctx.Set(JwtAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.DeletePostLike(ctx, postId)
	return err
}

// GetPosts converts echo context to params.
func (w *ServerInterfaceWrapper) GetPosts(ctx echo.Context) error {
	var err error

	ctx.Set(X_CSRF_TOKENScopes, []string{})

	ctx.Set(JwtAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetPosts(ctx)
	return err
}

// CreatePost converts echo context to params.
func (w *ServerInterfaceWrapper) CreatePost(ctx echo.Context) error {
	var err error

	ctx.Set(X_CSRF_TOKENScopes, []string{})

	ctx.Set(JwtAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.CreatePost(ctx)
	return err
}

// GetPostsOfUser converts echo context to params.
func (w *ServerInterfaceWrapper) GetPostsOfUser(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "userId" -------------
	var userId int

	err = runtime.BindStyledParameterWithLocation("simple", false, "userId", runtime.ParamLocationPath, ctx.Param("userId"), &userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter userId: %s", err))
	}

	ctx.Set(X_CSRF_TOKENScopes, []string{})

	ctx.Set(JwtAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetPostsOfUser(ctx, userId)
	return err
}

// DeletePost converts echo context to params.
func (w *ServerInterfaceWrapper) DeletePost(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "postId" -------------
	var postId int

	err = runtime.BindStyledParameterWithLocation("simple", false, "postId", runtime.ParamLocationPath, ctx.Param("postId"), &postId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter postId: %s", err))
	}

	ctx.Set(X_CSRF_TOKENScopes, []string{})

	ctx.Set(JwtAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.DeletePost(ctx, postId)
	return err
}

// GetPostByPostID converts echo context to params.
func (w *ServerInterfaceWrapper) GetPostByPostID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "postId" -------------
	var postId int

	err = runtime.BindStyledParameterWithLocation("simple", false, "postId", runtime.ParamLocationPath, ctx.Param("postId"), &postId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter postId: %s", err))
	}

	ctx.Set(X_CSRF_TOKENScopes, []string{})

	ctx.Set(JwtAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetPostByPostID(ctx, postId)
	return err
}

// SearchUser converts echo context to params.
func (w *ServerInterfaceWrapper) SearchUser(ctx echo.Context) error {
	var err error

	ctx.Set(X_CSRF_TOKENScopes, []string{})

	ctx.Set(JwtAuthScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params SearchUserParams
	// ------------- Required query parameter "keyword" -------------

	err = runtime.BindQueryParameter("form", true, true, "keyword", ctx.QueryParams(), &params.Keyword)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter keyword: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.SearchUser(ctx, params)
	return err
}

// Signup converts echo context to params.
func (w *ServerInterfaceWrapper) Signup(ctx echo.Context) error {
	var err error

	ctx.Set(X_CSRF_TOKENScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.Signup(ctx)
	return err
}

// GetLoginUser converts echo context to params.
func (w *ServerInterfaceWrapper) GetLoginUser(ctx echo.Context) error {
	var err error

	ctx.Set(X_CSRF_TOKENScopes, []string{})

	ctx.Set(JwtAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetLoginUser(ctx)
	return err
}

// UpdateUser converts echo context to params.
func (w *ServerInterfaceWrapper) UpdateUser(ctx echo.Context) error {
	var err error

	ctx.Set(X_CSRF_TOKENScopes, []string{})

	ctx.Set(JwtAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.UpdateUser(ctx)
	return err
}

// GetUser converts echo context to params.
func (w *ServerInterfaceWrapper) GetUser(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "userId" -------------
	var userId int

	err = runtime.BindStyledParameterWithLocation("simple", false, "userId", runtime.ParamLocationPath, ctx.Param("userId"), &userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter userId: %s", err))
	}

	ctx.Set(X_CSRF_TOKENScopes, []string{})

	ctx.Set(JwtAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetUser(ctx, userId)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/books", wrapper.FetchBookData)
	router.GET(baseURL+"/books-genres", wrapper.GetBooksGenres)
	router.GET(baseURL+"/chat-rooms", wrapper.GetChatRooms)
	router.POST(baseURL+"/chat-rooms", wrapper.CreateChatRoom)
	router.GET(baseURL+"/chats", wrapper.ChatSocket)
	router.POST(baseURL+"/comment-likes", wrapper.CreateCommentLike)
	router.DELETE(baseURL+"/comment-likes/comment/:CommentId", wrapper.DeleteCommentLike)
	router.POST(baseURL+"/comments", wrapper.CreateComment)
	router.GET(baseURL+"/comments/post/:postId", wrapper.GetCommentsByPostID)
	router.DELETE(baseURL+"/comments/:commentId", wrapper.DeleteComment)
	router.GET(baseURL+"/csrftoken", wrapper.Csrftoken)
	router.GET(baseURL+"/followers", wrapper.GetFollowerConnections)
	router.GET(baseURL+"/followings", wrapper.GetFollowingConnections)
	router.POST(baseURL+"/follows", wrapper.CreateConnection)
	router.DELETE(baseURL+"/follows/:connectionId", wrapper.DeleteConnection)
	router.GET(baseURL+"/health", wrapper.HealthCheck)
	router.GET(baseURL+"/liked-comments", wrapper.GetLikedCommentList)
	router.GET(baseURL+"/liked-posts", wrapper.GetLikedPostList)
	router.POST(baseURL+"/login", wrapper.Login)
	router.POST(baseURL+"/logout", wrapper.Logout)
	router.GET(baseURL+"/messages", wrapper.GetMessages)
	router.GET(baseURL+"/oauth/google/callback", wrapper.GoogleOauthCallback)
	router.POST(baseURL+"/post-likes", wrapper.CreatePostLike)
	router.DELETE(baseURL+"/post-likes/post/:PostId", wrapper.DeletePostLike)
	router.GET(baseURL+"/posts", wrapper.GetPosts)
	router.POST(baseURL+"/posts", wrapper.CreatePost)
	router.GET(baseURL+"/posts/user/:userId", wrapper.GetPostsOfUser)
	router.DELETE(baseURL+"/posts/:postId", wrapper.DeletePost)
	router.GET(baseURL+"/posts/:postId", wrapper.GetPostByPostID)
	router.GET(baseURL+"/search-user", wrapper.SearchUser)
	router.POST(baseURL+"/signup", wrapper.Signup)
	router.GET(baseURL+"/user", wrapper.GetLoginUser)
	router.PUT(baseURL+"/user", wrapper.UpdateUser)
	router.GET(baseURL+"/user/:userId", wrapper.GetUser)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"AAD//45WDtdeQwAA",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
