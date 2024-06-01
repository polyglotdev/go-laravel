package render

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

// Render is a struct that contains the configuration for the renderer.
type Render struct {
	Renderer   string
	RootPath   string
	Secure     bool
	Port       string
	ServerName string
}

// TemplateData is a struct that contains the data to be passed to the template.
type TemplateData struct {
	IsAuthenticated bool
	IntMap          map[string]int
	StringMap       map[string]string
	FloatMap        map[string]float32
	Data            map[string]interface{}
	CSRFToken       string
	Port            string
	ServerName      string
	Secure          bool
}

// Page renders a web page using the specified view and data.
// It selects the rendering engine based on the Renderer's value.
//
// Parameters:
//   - w: The HTTP response writer.
//   - r: The HTTP request.
//   - view: The name of the view template to render.
//   - variables: Additional variables for rendering (currently unused).
//   - data: The data to pass to the view template.
//
// Returns:
//   - error: An error if the rendering fails, otherwise nil.
func (c *Render) Page(w http.ResponseWriter, r *http.Request, view string, variables,
	data interface{}) error {
	switch strings.ToLower(c.Renderer) {
	case "go":
		return c.GoPage(w, r, view, data)
	}
	return nil
}

// GoPage is a method on the Render struct that renders a Go template page.
// It takes a http.ResponseWriter, http.Request, a string representing the view, and an interface{} for data.
// The method first attempts to parse the template file corresponding to the view.
// If an error occurs during parsing, it returns the error.
// If the data passed is not nil, it asserts the data to be of type *TemplateData.
// It then executes the template with the TemplateData and writes the output to the http.ResponseWriter.
func (c *Render) GoPage(w http.ResponseWriter, r *http.Request, view string, data interface{}) error {
	tmpl, err := template.ParseFiles(fmt.Sprintf("%s/views/%s.page.tmpl",
		c.RootPath, view))
	if err != nil {
		return err
	}

	td := &TemplateData{}
	if data != nil {
		td = data.(*TemplateData)
	}

	err = tmpl.Execute(w, td)
	if err != nil {
		return err
	}

	return nil
}
