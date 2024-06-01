package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/CloudyKit/jet/v6"
)

// Render is a struct that contains the configuration for the renderer.
type Render struct {
	Renderer   string
	RootPath   string
	Secure     bool
	Port       string
	ServerName string
	JetViews   *jet.Set
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
	case "jet":
		return c.JetPage(w, r, view, variables, data)
	}
	return nil
}

// JetPage is a method on the Render struct that renders a Jet template page.
// It takes a http.ResponseWriter, http.Request, a string representing the view, and an interface{} for data.
// The method first attempts to parse the template file corresponding to the view.
// If an error occurs during parsing, it returns the error.
// If the data passed is not nil, it asserts the data to be of type *TemplateData.
// It then executes the template with the TemplateData and writes the output to the http.ResponseWriter.
func (c *Render) JetPage(w http.ResponseWriter, r *http.Request, templateName string, variables interface{}, data interface{}) error {
	var vars jet.VarMap
	if variables != nil {
		var ok bool
		vars, ok = variables.(jet.VarMap)
		if !ok {
			return fmt.Errorf("variables is not of type jet.VarMap")
		}
	} else {
		vars = make(jet.VarMap) // Initialize vars to an empty jet.VarMap if variables is nil
	}

	td := &TemplateData{}
	if data != nil {
		var ok bool
		td, ok = data.(*TemplateData)
		if !ok {
			return fmt.Errorf("data is not of type *TemplateData")
		}
	}

	t, err := c.JetViews.GetTemplate(fmt.Sprintf("%s.jet", templateName))
	if err != nil {
		log.Println("Error getting template:", err)
		return err
	}

	if err := t.Execute(w, vars, td); err != nil {
		log.Println("Error executing template "+templateName+":", err)
		return err
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
