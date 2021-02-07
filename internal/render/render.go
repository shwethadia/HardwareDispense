package render

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/justinas/nosurf"
	"github.com/shwethadia/BOOKINGAPI/internal/config"
	"github.com/shwethadia/BOOKINGAPI/internal/models"
)

var functions = template.FuncMap{}

var app *config.AppConfig

//NewTemplates ...
func NewTemplates(a *config.AppConfig) {

	app = a

}

//AddDefaultData ...
func AddDefaultData(td *models.TemplateData,r *http.Request) *models.TemplateData {

	td.CSRFToken = nosurf.Token(r)
	return td
}

//RenderTemplate ...
func RenderTemplate(w http.ResponseWriter,r *http.Request, tmpl string, td *models.TemplateData) {

	var tc map[string]*template.Template
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td,r)
	_ = t.Execute(buf, td)

	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template to browser", err)

	}
}

//CreateTemplateCache ...
func CreateTemplateCache() (map[string]*template.Template, error) {

	myCache := map[string]*template.Template{}
	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return myCache, err
	}
	for _, page := range pages {

		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		matches, err := filepath.Glob("./templates/*layout.html")
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}
	return myCache, nil
}
