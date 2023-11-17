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
	// posts
	// (GET /posts)
	Posts(ctx echo.Context) error
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

// Posts converts echo context to params.
func (w *ServerInterfaceWrapper) Posts(ctx echo.Context) error {
	var err error

	ctx.Set(X_CSRF_TOKENScopes, []string{})

	ctx.Set(JwtAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.Posts(ctx)
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

	router.GET(baseURL+"/csrftoken", wrapper.Csrftoken)
	router.POST(baseURL+"/login", wrapper.Login)
	router.POST(baseURL+"/logout", wrapper.Logout)
	router.GET(baseURL+"/oauth/google/callback", wrapper.GoogleOauthCallback)
	router.GET(baseURL+"/posts", wrapper.Posts)
	router.POST(baseURL+"/signup", wrapper.Signup)
	router.GET(baseURL+"/user", wrapper.User)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/8xW30/jSAz+V6K5e8ySLvuyyhug+4FYLQgW6SSEqunEbQeScZhxQD2U//1kT9qSNuF2",
	"VyDtU3/YHvv7PP7Gz8pgVaMDR0HlzyqYJVRavh4j3vNn7bEGTxbk39Or468GC+DvtKpB5SqQt26h2lTp",
	"hpboB00zxPupLV7YrCNYgGejrfRi+ERLUE0bXw4aa2/NcFjdzEobllBMNbHDHH3F31ShCT6QrUCl41HD",
	"CMhSOZSt3ZyEszswxL4XGGifu1nH6O8e5ipXv2Vb7rOO+ExYb1Nl0BE4GqzEeND0g9jGKa4x0GhnvCZ2",
	"G7Q1ITL1Gpxr9hmk6BIevuDCumMsVvtUGQ8FOLJ6uPNWjHM72Ko2nn5lF66p3/74CNzpCkZzh5Pg59/w",
	"HtxA6uDnU1rbBsKvO1r7YSP5eAhwbkuYjjeYyx1p8H5j2lQFMI23tLriHsbs/3w4ubr888O387M/vvJv",
	"61SulqAL8Crtaus7bQ7WtT2DFdexQFyUMGWRkHtb4pMcHlXD/qvJojvptKX35zULgFoS1SHPMm0MNo7C",
	"QTzwwGCVYYYccZg9HmaSIFXBYB2rh0pbPkA+E10UHkJQG+pUrmY6WJN0vxPr4kxZdDL43Kx+CTFZV4Cu",
	"bZAiYldfkipufMbdEx11uIU7g3hvYcvd3RN1l2KAuECaYDQ0WvfCuAwGwnEFBONtLYBydSl9C1+wcQtI",
	"ji5O1UbcBm2P4EOM/HgwOZhwRViD07VVufokf6Wq1rQUsjO+4Jv7vQDRKL7JwudpoXJ1svFIlYdQowux",
	"UYeTiYzIVvt0XZfWSGh2F9Bt36j/U57eFAobfRbOz3qXXeU3t6kKTVVpv1I5F54wkmTb1KxkvZLR7OS9",
	"D0vkTCA9NBBorTxvhOaFWgoazmI9FCon30D7jkx2Gv4dDO4Kxc1t2yM1ErjmEht6lUy2/zybu7rWjlD0",
	"Bpi4UAEl855FVciMLsuZNvejY/CX+J1zzMnalwfJ6woIfJDk/fKueNaTjUsyR59wcSxdBEa80igTDw34",
	"1b5K9K9Nuk/Ydq3ZTX70UpITXgITD9R4B0UyWyURTsJwkgD+UZ6GoVJkffyRSm53OvdpcjikaoX1YCgh",
	"TP7GCvb72Ht/brpXYfsK3DLiTmp3m7zujxDeAT0XQZe+8xUOo32+EOs7Dqhsmz87oOnz9nHahR1xCcQg",
	"29T4vMZt6/3U78U2913y9/EXlL+4tycOnpJmvRdn6y168OpI7l9V2l+7Ofx8NpvDoxxEQWs2i1SeZSUa",
	"XS4xUP558nmi2tv2vwAAAP//foKuV5EOAAA=",
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
