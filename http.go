package main

import (
	"bytes"
	"fmt"
	"html"
	"html/template"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/shurcooL/httpfs/filter"
	"github.com/shurcooL/httpfs/html/vfstemplate"
	"go.tmthrgd.dev/kim.tmthrgd.dev/internal/assets"
	"go.tmthrgd.dev/kim.tmthrgd.dev/internal/templates"
	"go.tmthrgd.dev/vfshash"
)

const assetsPath = "/assets"

var assetNames = vfshash.NewAssetNames(assets.FileSystem)

var errorTmpl = newTemplate("error.tmpl")

// notFoundHandler returns a handler that serves a 404 Not Found error.
func notFoundHandler() http.HandlerFunc {
	return errorHandler(func(http.ResponseWriter, *http.Request) error {
		return os.ErrNotExist
	})
}

// methodNotAllowedHandler returns a handler that serves a 405 Method Not
// Allowed error.
func methodNotAllowedHandler() http.HandlerFunc {
	var err error = methodNotAllowedError("GET, HEAD")
	return errorHandler(func(http.ResponseWriter, *http.Request) error { return err })
}

// assetNamesHandler returns a handler that serves files from the assets
// directory without name mangling.
func assetNamesHandler() http.HandlerFunc {
	return http.FileServer(assetNames).ServeHTTP
}

// assetsHandler returns a handler that serves site assets.
func assetsHandler() http.Handler {
	return http.StripPrefix(assetsPath, http.FileServer(filter.Skip(assets.FileSystem, excludeAssets)))
}

// excludeAssets returns true if the file should be excluded from the assets handler.
func excludeAssets(path string, info os.FileInfo) bool {
	return info.IsDir() || strings.HasPrefix(info.Name(), ".")
}

// httpsOnly is a http.Handler middleware that redirects http requests to https when
// behind Google Frontend.
func httpsOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-Forwarded-Proto") == "http" {
			u := *r.URL
			u.Scheme = "https"
			u.Host = "kim.tmthrgd.dev"

			http.Redirect(w, r, u.String(), http.StatusPermanentRedirect)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// errorHandler converts a handler with an error return to a http.HandlerFunc,
// sending a HTTP error code to the client appropriate for a given error.
func errorHandler(handler func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := handler(w, r)
		if err == nil {
			return
		}

		hdr := w.Header()
		statusCode := http.StatusInternalServerError
		errorMsg := html.EscapeString(err.Error())
		switch err := err.(type) {
		case methodNotAllowedError:
			hdr.Set("Allow", string(err))

			statusCode = http.StatusMethodNotAllowed
			errorMsg = fmt.Sprintf("The request method <code>%s</code> is inappropriate for the URL <code>%s</code>.",
				html.EscapeString(r.Method), html.EscapeString(r.URL.Path))
		case *tumblrAPIError:
			statusCode = http.StatusBadGateway
		default:
			if os.IsNotExist(err) {
				statusCode = http.StatusNotFound
				errorMsg = "The requested file was not found."
			}
		}

		hdr.Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(statusCode)

		if templateExecute(w, errorTmpl, &struct {
			StatusCode int
			Message    template.HTML
		}{
			statusCode,
			template.HTML(errorMsg),
		}) != nil {
			fmt.Fprintf(w, "%d %s: %s", statusCode,
				http.StatusText(statusCode), errorMsg)
		}
	}
}

// templateExecute calls Execute on the given template, only writing out the result if
// execution was successful.
func templateExecute(w http.ResponseWriter, tmpl *template.Template, data interface{}) error {
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return err
	}

	if n, err := buf.WriteTo(w); n == 0 {
		// Only forward the error if nothing was written.
		return err
	}

	return nil
}

// newTemplate parses the template at templates/name and returns a new
// html/template.Template. It panics if the source is invalid or the template
// doesn't exist.
func newTemplate(name string) *template.Template {
	t := template.New(name).Funcs(templateFuncs)
	return template.Must(vfstemplate.ParseFiles(templates.FileSystem, t, name))
}

var templateFuncs = template.FuncMap{
	"assetPath":      assetPath,
	"httpStatusText": http.StatusText,
}

// assetPath returns the path to a named asset file.
func assetPath(name string) string {
	return path.Join(assetsPath, assetNames.Lookup(name))
}

type methodNotAllowedError string

func (methodNotAllowedError) Error() string {
	return "the request method is inappropriate for the requested URL"
}

type tumblrAPIError struct{ error }

func (te *tumblrAPIError) Error() string {
	return "upstream HTTP error: " + te.error.Error()
}
