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

	ctx.Set(JwtAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetBooksGenres(ctx)
	return err
}

// GetChatRooms converts echo context to params.
func (w *ServerInterfaceWrapper) GetChatRooms(ctx echo.Context) error {
	var err error

	ctx.Set(JwtAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetChatRooms(ctx)
	return err
}

// CreateChatRoom converts echo context to params.
func (w *ServerInterfaceWrapper) CreateChatRoom(ctx echo.Context) error {
	var err error

	ctx.Set(JwtAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.CreateChatRoom(ctx)
	return err
}

// ChatSocket converts echo context to params.
func (w *ServerInterfaceWrapper) ChatSocket(ctx echo.Context) error {
	var err error

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

	ctx.Set(JwtAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.DeleteCommentLike(ctx, commentId)
	return err
}

// CreateComment converts echo context to params.
func (w *ServerInterfaceWrapper) CreateComment(ctx echo.Context) error {
	var err error

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

	ctx.Set(JwtAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.DeleteComment(ctx, commentId)
	return err
}

// GetFollowerConnections converts echo context to params.
func (w *ServerInterfaceWrapper) GetFollowerConnections(ctx echo.Context) error {
	var err error

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

	ctx.Set(JwtAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetLikedCommentList(ctx)
	return err
}

// GetLikedPostList converts echo context to params.
func (w *ServerInterfaceWrapper) GetLikedPostList(ctx echo.Context) error {
	var err error

	ctx.Set(JwtAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetLikedPostList(ctx)
	return err
}

// Login converts echo context to params.
func (w *ServerInterfaceWrapper) Login(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.Login(ctx)
	return err
}

// Logout converts echo context to params.
func (w *ServerInterfaceWrapper) Logout(ctx echo.Context) error {
	var err error

	ctx.Set(JwtAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.Logout(ctx)
	return err
}

// GetMessages converts echo context to params.
func (w *ServerInterfaceWrapper) GetMessages(ctx echo.Context) error {
	var err error

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

	ctx.Set(JwtAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.DeletePostLike(ctx, postId)
	return err
}

// GetPosts converts echo context to params.
func (w *ServerInterfaceWrapper) GetPosts(ctx echo.Context) error {
	var err error

	ctx.Set(JwtAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetPosts(ctx)
	return err
}

// CreatePost converts echo context to params.
func (w *ServerInterfaceWrapper) CreatePost(ctx echo.Context) error {
	var err error

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

	ctx.Set(JwtAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetPostByPostID(ctx, postId)
	return err
}

// SearchUser converts echo context to params.
func (w *ServerInterfaceWrapper) SearchUser(ctx echo.Context) error {
	var err error

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

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.Signup(ctx)
	return err
}

// GetLoginUser converts echo context to params.
func (w *ServerInterfaceWrapper) GetLoginUser(ctx echo.Context) error {
	var err error

	ctx.Set(JwtAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetLoginUser(ctx)
	return err
}

// UpdateUser converts echo context to params.
func (w *ServerInterfaceWrapper) UpdateUser(ctx echo.Context) error {
	var err error

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
