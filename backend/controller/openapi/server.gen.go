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
	// Get the child genres of booksGenreId
	// (GET /books-genres)
	GetBooksGenres(ctx echo.Context, params GetBooksGenresParams) error
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

// GetBooksGenres converts echo context to params.
func (w *ServerInterfaceWrapper) GetBooksGenres(ctx echo.Context) error {
	var err error

	ctx.Set(X_CSRF_TOKENScopes, []string{})

	ctx.Set(JwtAuthScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetBooksGenresParams
	// ------------- Required query parameter "booksGenreId" -------------

	err = runtime.BindQueryParameter("form", true, true, "booksGenreId", ctx.QueryParams(), &params.BooksGenreId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter booksGenreId: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetBooksGenres(ctx, params)
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
	router.GET(baseURL+"/books-genres", wrapper.GetBooksGenres)
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

	"H4sIAAAAAAAC/8xYUU8bORD+K5bvHgOb0pdq30p67SGqgqBIJ6EocryTxGTXXmwvKFflv59mvLvJJt6U",
	"chTxRIjH9sw338w3zg8uTVEaDdo7nv7gTi6gEPTx1Jgl/i2tKcF6BfTt2fXpN2kywM9+VQJPufNW6Tlf",
	"D7io/MLY6NLUmOVEZVtrSnuYg8VFVYh5/ETloZhUNo8ullZJiJ9YVtNcuQVkE+Hje2uDuLde+Tzm0HrA",
	"LdxXykLG09s2qMEGlmZvC0bj5vaVO/5thdlgMR40V5vpHUiPTmFCvoC2sJ8VdMRN5rjYBbmbgMZEiyIO",
	"d1jO4QHynkz1ZLAUFrQ/5MEOdITajtsRJ7su7d8Tw+nSOB+HCP/+aWHGU/5HsiF+UrM+IcqvB1wa7UHH",
	"iSMtCN/Pq34ql8b53gqwwqNZdK1ygaWHPL9Bm12MmxvrIzZxtfd1wgnwRyG9gvsRGSK2pyZb/WJjOARo",
	"i9jM2AJR5VOlhV3xwb5tP0w7oUcibR0ch4C+mrnS8VikhQy0VyLedxQtzlS0eazD6ddqrqvy5Y8Puewp",
	"4HC3Gzk7+26WoCNXOzub+GYtsv2mJlt3W2/DKK2ZqRwm/bRHd3tov5OzxnLA69LvHr7Py/WAO5CVVX51",
	"jXUQfP3naHR99fno+8X5X9/wf6V5yhcgMiqBEEnXqD1YlOocVtQIjZnnMMEeTtTMzSMdHpq6+ld4ZfSo",
	"pnvnyxsUK77wvnRpkggpTaW9Ow4HHktTJCYxuOMkeThJ6IIBd9KUwXsohMID6C8TWWbBuQ0WWB3CKcnq",
	"/5nSoWyU0SRcmNquC+Gy2gFRKkdOBA4ghA2oZIZn3D36j3XchJ00Zqlgg93do68pFAHOeeGhd2tY3duG",
	"bmAguC8DJ60qKaCUX1He3FdT6Tmwj5dnrcDG1x7AurDz3fHweIgemRK0KBVP+Xv6CkXELwjshPQGP82B",
	"mhNynrA8y3jKP4OXC9SET8KLID6iAA/W8fR219OzT8wb5kqQarZifgGMRIopza7EsvKg2SndNgjQ3FdA",
	"Pa5GhjwhdT/LiBDY16MyunvxElaPxmZ0OwgrF2x64J7a+uAVY6xLVxrtAidPhkPqHZsmLsoyV5KQSu6c",
	"0ZvRkVLvoXBP1dqGCtaKmgnd8C7OO4VOyHdL/HaMoLSsvR1jAK4qChSRlM8wi4QJy4QXbGZN0aQESYOn",
	"ByIcUcb6+fAF/GmbJve6hNg0Sm8reBPZC7Poa6TwC3gCUC5UngUYHTMz1sGI8oj61spbNImj1uJ/4nQI",
	"no4IPwmRTrxz8AwjYZsuneQ4rpAy1+NtNyyaZmqewGZGe6FotoaldVe2kY3r34hkPdg+j1MdUAOADZam",
	"8gfBxPXno7k7qKx7IHqBmNBRCooEPAkyn0iR51Mhl/29jOwucM+osf1JQ7tG8WatCZsZy9A5nEU8SLKK",
	"97JG9p/exPZU7uP2jMVwjmcWfGU1ZGy6YiEchuEwB/aBZr2YK/Uj/fnt9P3wJDamZMqC9Njx/zYF7Oex",
	"M1De1mPeZqyj/lfPTrtJbvJDgNeBXlCjpLwjhQ9q1iUZvIYq0Mv7NQQBG2SIu35YR/p8+149WMdFlXtV",
	"CusTnKOPcEL4pca48yx+Und898zSP4RIeMcTKIEWjh6g/T0uPFB/n2JsPYCfDsobk4waVA2PrGp+YEma",
	"n2Oi5XYTfmh5m3L4s4qq2sNDCw0iULWvyTRJciNFvjDOpx+GH4Z8PV7/FwAA///p3cU2QhYAAA==",
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
