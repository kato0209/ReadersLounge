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
	// fetch book genre from RakutenAPI
	// (GET /books-genres)
	FetchBookGenres(ctx echo.Context) error
	// get csrf token
	// (GET /csrftoken)
	Csrftoken(ctx echo.Context) error
	// login
	// (POST /login)
	Login(ctx echo.Context) error
	// logout
	// (POST /logout)
	Logout(ctx echo.Context) error
	// Callback for Google OAuth
	// (GET /oauth/google/callback)
	GoogleOauthCallback(ctx echo.Context, params GoogleOauthCallbackParams) error
	// get posts
	// (GET /posts)
	GetPosts(ctx echo.Context) error
	// create post
	// (POST /posts)
	CreatePost(ctx echo.Context) error
	// create new user
	// (POST /signup)
	Signup(ctx echo.Context) error
	// get user
	// (GET /user)
	User(ctx echo.Context) error
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

// FetchBookGenres converts echo context to params.
func (w *ServerInterfaceWrapper) FetchBookGenres(ctx echo.Context) error {
	var err error

	ctx.Set(X_CSRF_TOKENScopes, []string{})

	ctx.Set(JwtAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.FetchBookGenres(ctx)
	return err
}

// Csrftoken converts echo context to params.
func (w *ServerInterfaceWrapper) Csrftoken(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.Csrftoken(ctx)
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

// Signup converts echo context to params.
func (w *ServerInterfaceWrapper) Signup(ctx echo.Context) error {
	var err error

	ctx.Set(X_CSRF_TOKENScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.Signup(ctx)
	return err
}

// User converts echo context to params.
func (w *ServerInterfaceWrapper) User(ctx echo.Context) error {
	var err error

	ctx.Set(X_CSRF_TOKENScopes, []string{})

	ctx.Set(JwtAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.User(ctx)
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
	router.GET(baseURL+"/books-genres", wrapper.FetchBookGenres)
	router.GET(baseURL+"/csrftoken", wrapper.Csrftoken)
	router.POST(baseURL+"/login", wrapper.Login)
	router.POST(baseURL+"/logout", wrapper.Logout)
	router.GET(baseURL+"/oauth/google/callback", wrapper.GoogleOauthCallback)
	router.GET(baseURL+"/posts", wrapper.GetPosts)
	router.POST(baseURL+"/posts", wrapper.CreatePost)
	router.POST(baseURL+"/signup", wrapper.Signup)
	router.GET(baseURL+"/user", wrapper.User)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/8xYX2/bOAz/KoLuHtM6614Gv63drVdsWId2BQ4oikCRmUSNLXkS3SI35LsfSDlOncpZ",
	"u9uKPcWxKIr88c+P8jepXVU7CxaDzL/JoBdQKX48dm5Jv7V3NXg0wG/PLo8/aVcAPeOqBpnLgN7YuVyP",
	"pGpw4XxyaerccmKKB2vGIszB06Kp1Dyt0SBUk8aXycXaGw1pjXUzLU1YQDFRmN7bCqStRYNlyiBaim/c",
	"9BY0kizhdArWw2OwyOkwmdNi3/c+LhsRq6o0CnG5hDsoBwAcALZWHizusyDl0WcXMO0M/f7pYSZz+Ue2",
	"zZysTZuMc2Y9ktpZBJtGXntQOByY4VyoXcDBFPIKSSy51oQY5n2WX5EMoeHha2M8FDK/7k5sVWz96s7r",
	"uROjKW8SkF7A1xMWJGyPXbF6ZmXtA7RDbOZ8RajKqbHKr+TosewwTDuuJzztDLyJDn10c2PTvmgPBVg0",
	"Kl24hhdnJll966j90sxtU/989TGWA6UWzw4nwc++uCXYxNHBzya4WUtsv2qTrb9tsLRr72amhMlw2pO5",
	"A2m/E7ON5Cget6v8cV6uRzKAbrzB1SXVQbT1n4OTy4v3B1/OP/z1if4bK3O5AFVwCURP+kKdYlWbD7Di",
	"luXcvIQJMQKnZunuWXmkCPOvQuPsSZvuvZdX1O3lArEOeZYprV1jMRxGhYfaVZnLHO04yu6OMj5gJIN2",
	"dbQeKmVIAf8KVRQeQthiQdWhgtGi/S+MjWVjnOXOT6HtmxAPaw1QtQlsRMyBh/2TxUjH7T2+bf1m7LRz",
	"SwNb7G7vsU2hBHABFcLg1rj6aBuZQY7QvgKC9qZmh3J5wXELH11j5yDefj6THbsl1+7Ah7jz1eH4cEwW",
	"uRqsqo3M5Wt+RbSCCwY7Y/qipzlwc6KcZyzPCpnL94B6QZzwTqHifV5VgOCDzK93LT17J9CJUIM2s5XA",
	"BQimLWGsuFDLBsGKYz5tFKH52gD3uBYZtoR5+KzghKC+niS83YOXsLp3vuDTQXm9ENM957TSe4+4oboM",
	"tbMh5uTReMy9Y9vEVV2XRjNS2W1wdjt7cegRqvBUrt2kgveqzYS+e+cfeoXOyPdL/PqGQOmy9vqGHAhN",
	"VRGJ5HJGUWRMRKFQiZl31SYklDSkPSbCAUfsCflwGuVeCqY4nr0wVjF7k2ARiXQckkTqpJP4nxjtg6bH",
	"dE9Co+frHFCQJ2LbCrOSZgKmv3aG7LvFI4OMrAXbQegnefNgIln3uRF9A+tfiGQ7Pf5YPvVAjQBusHQN",
	"7gWT1n8czd1pYD0A0U/wiQxlp5gls8ilmVZlOVV6OVgGpyx3TntONrLfoZFLYkjRiYiZ84KMI8JH0CyV",
	"7uwbbu2nzbOo5O3DQUbQsCw8YOMtFGK6EtEdQe6IAP6OB6qUKTxmP8eSXcZ5PT5KzQKF8aCRmO5vV8Hj",
	"OPamtut2ltrOTtz72gFlN8ib+DDgraPn3CQ57pTCw8RwCviZBV6CEfh6+xJkQA0y+t3eXhN9vrsU7q3j",
	"qinR1MpjRsPqAdHwsxrjzt3zSd3x1Q+W/j5E4mWZQYlpEfiWN9zj4i3w1zHGg1vm00H5zSijBdXCvWg2",
	"XzGyzTePZLldxa8Zvycdfq+imk55bKGRBJruypZnWem0KhcuYP5m/GYs1zfr/wIAAP//zR+jKOgUAAA=",
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
