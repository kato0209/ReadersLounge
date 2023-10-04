// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.15.0 DO NOT EDIT.
package openapi

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
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
	// create new user
	// (POST /signup)
	Signup(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// Csrftoken converts echo context to params.
func (w *ServerInterfaceWrapper) Csrftoken(ctx echo.Context) error {
	var err error

	ctx.Set(X_CSRF_TOKENScopes, []string{})

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

// Signup converts echo context to params.
func (w *ServerInterfaceWrapper) Signup(ctx echo.Context) error {
	var err error

	ctx.Set(X_CSRF_TOKENScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.Signup(ctx)
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
	router.POST(baseURL+"/signup", wrapper.Signup)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/7RUwW7bMAz9FYPb0a297VLotgUbUKRYh6SHAUVQaDbtqLMllaI3GIX/faCctEnhtEOR",
	"nmKFeuR7j6TuoXCtdxYtB1D3EIo1tjp+LvDuwtXGfnFlL2dPziOxwRgtCEu0bHQjJ+49goLAZGwNQwom",
	"BiuDNBEehlSyL01tO3/89Cl0AcnqFg/WDrNA1ZX7jXaidKDqhrexafhz1KX4jSl3wMYy1kiCHlIIWHRk",
	"uF+K1SPk58lsufh2cnU5//pdzsaCgjXqEglSGJXsX0q3ubU3c+xHYsZWTuAlhoKMZ+MkzyLmCReuszUm",
	"n3+cC9hwgwdif5DCiPxwmp/m4qfzaLU3oOBT/CsFr3kdqWdi14NbNbL8iB1ayp+XoGD2cCMFwuCdDaPs",
	"j3keDXeW0Uag9r4xRYRmt8HZx5GUr/eEFSh4lz3ObLYZ2Gyvp9GNfRcu5xDN79pWUw9KqCbCPeENJIWs",
	"kXmPHXVhQkhchyjirsPA2/Yfif/OtkX+UsUQlqCYOhymvXtB5Choq811/Kw4ib9e3WYg3a9bLHgz66+j",
	"LDwi5xDX7DDncQ3friM7a/7/LTnWOD8t/oJvBaFmTCz+TeQFgr23BtT101fmejWs5ALJvsd4R408O8xe",
	"ZVnjCt2sXWB1lp/lMKyGfwEAAP//9v28fysGAAA=",
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
