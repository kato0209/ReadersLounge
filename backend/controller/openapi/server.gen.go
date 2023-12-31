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
	// (DELETE /post-likes/{PostId})
	DeletePostLike(ctx echo.Context, postId int) error
	// get posts
	// (GET /posts)
	GetPosts(ctx echo.Context) error
	// create post
	// (POST /posts)
	CreatePost(ctx echo.Context) error
	// delete a post
	// (DELETE /posts/{postId})
	DeletePost(ctx echo.Context, postId int) error
	// get posts corresponding to userId
	// (GET /posts/{userId})
	GetPostsOfUser(ctx echo.Context, userId int) error
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
	router.GET(baseURL+"/csrftoken", wrapper.Csrftoken)
	router.GET(baseURL+"/followers", wrapper.GetFollowerConnections)
	router.GET(baseURL+"/followings", wrapper.GetFollowingConnections)
	router.POST(baseURL+"/follows", wrapper.CreateConnection)
	router.DELETE(baseURL+"/follows/:connectionId", wrapper.DeleteConnection)
	router.GET(baseURL+"/liked-posts", wrapper.GetLikedPostList)
	router.POST(baseURL+"/login", wrapper.Login)
	router.POST(baseURL+"/logout", wrapper.Logout)
	router.GET(baseURL+"/messages", wrapper.GetMessages)
	router.GET(baseURL+"/oauth/google/callback", wrapper.GoogleOauthCallback)
	router.POST(baseURL+"/post-likes", wrapper.CreatePostLike)
	router.DELETE(baseURL+"/post-likes/:PostId", wrapper.DeletePostLike)
	router.GET(baseURL+"/posts", wrapper.GetPosts)
	router.POST(baseURL+"/posts", wrapper.CreatePost)
	router.DELETE(baseURL+"/posts/:postId", wrapper.DeletePost)
	router.GET(baseURL+"/posts/:userId", wrapper.GetPostsOfUser)
	router.GET(baseURL+"/search-user", wrapper.SearchUser)
	router.POST(baseURL+"/signup", wrapper.Signup)
	router.GET(baseURL+"/user", wrapper.GetLoginUser)
	router.PUT(baseURL+"/user", wrapper.UpdateUser)
	router.GET(baseURL+"/user/:userId", wrapper.GetUser)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"AAA=",
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
