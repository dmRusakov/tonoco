package admin_app_web_v1

import (
	"fmt"
	"html/template"
	"net/http"
)

func (c Controller) Render(w http.ResponseWriter, pageTemplate string) {
	// make template
	tmpl := c.makeTemplate(pageTemplate)

	// render page
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	return
}

// make template
func (c Controller) makeTemplate(pageTemplate string) *template.Template {
	// page template
	var templateSlice []string
	templateSlice = append(templateSlice, fmt.Sprintf("%s%s", c.tmlPath, pageTemplate))

	// static templates
	partials := []string{
		"base.layout",
		"element/head_file_imports.partial",
		"element/head.partial",
		"element/menu.partial",
		"element/footer.partial",
		"element/footer_file_imports.partial",
		"grid/products.partial",
		"grid/product.partial",
		"element/footer_logs.partial",
	}
	for _, x := range partials {
		templateSlice = append(templateSlice, fmt.Sprintf("%s%s.gohtml", c.tmlPath, x))
	}

	// parse templates
	tmpl, err := template.ParseFiles(templateSlice...)
	if err != nil {
		return nil
	}

	return tmpl
}
