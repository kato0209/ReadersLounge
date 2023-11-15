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

	"H4sIAAAAAAAC/8xWXU/jSg/+K9G872UgXfZmlTtA5wOxWhAs0pEQqqYTtx1IxmHGAfWg/Pcje9KWtAln",
	"dwXSueqH7bGfx+Nn/KIMVjU6cBRU/qKCWUKl5esJ4gN/1h5r8GRB/j27PvlmsAD+TqsaVK4CeesWqk2V",
	"bmiJftA0Q3yY2uKVzTqCBXg22kovhk+0BNW08eWgsfbWDIfVzay0YQnFVBM7zNFX/E0VmuCAbAUqHY8a",
	"RkCWyqFs7eYknN2DIfa9xED73M06Rv/vYa5y9b9sy33WEZ8J622qDDoCR4OVGA+afhLbOMU1BhrtjNfE",
	"boO2JkSm3oJzwz6DFF3B41dcWHeCxWqfKuOhAEdWD3feinFuB1vVxtOv7cI19fsfH4E7XcFo7nAa/Pw7",
	"PoAbSB38fEpr20D4TUdrP2wkHw8Bzm0J0/EGc7kjDd5vTJuqAKbxllbX3MOY/a+D0+ur3w++X5z/9o1/",
	"W6dytQRdgFdpV1vfaXOwru05rLiOBeKihCmLhNzbEp/l8Kga9m9NFt1ppy29P29YANSSqA55lmljsHEU",
	"DuOBhwarDDPkiKPs6SiTBKkKButYPVTa8gHymeii8BCC2lCncjXTwZqk+51YF2fKopPB52b1S4jJugJ0",
	"bYMUEbv6mlRx4zPun+m4wy3cGcQHC1vu7p+puxQDxAXSBKOh0boXxmUwEI4rIBhvawGUqyvpW/iKjVtA",
	"cnx5pjbiNmh7Ah9i5KfDyeGEK8IanK6tytVn+StVtaalkJ3xBd/c7wWIRvFNFj7PCpWr041HqjyEGl2I",
	"jTqaTGREttqn67q0RkKz+4Bu+0b9m/L0plDY6LNwcd677Cq/vUtVaKpK+5XKufCEkSTbpmYl65WMZifv",
	"fVgiZwLpsYFAa+V5JzSv1FLQcBbroVA5+QbaD2Sy0/AfYHBXKG7v2h6pkcA1l9jQm2Sy/dfZ3NW1doSi",
	"d8DEhQoomfcsqkJmdFnOtHkYHYM/xO+CY07XvjxIXldA4IMk75d3zbOebFySOfqEi2PpIjDilUaZeGzA",
	"r/ZVon9t0n3CtmvNbvLj15Kc8BKYeKDGOyiS2SqJcBKGkwTwT/I0DJUi6+PPVHK307nPk6MhVSusB0MJ",
	"YfInVrDfx977c9u9CttX4I4Rd1K72+R1f4TwDuiFCLr0na9wGO3zpVg/cEBl2/zVAU1fto/TLuyISyAG",
	"2abG5zVuWx+nfq+2uR+Sv0/vMdtx0U4cPCfNepHN1mvvYK9FK/+rWvxWq/m9azaHx/mNCtRsNp88y0o0",
	"ulxioPzL5MtEtXftPwEAAP//+Ojgu0IOAAA=",
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
