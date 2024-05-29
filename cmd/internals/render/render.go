package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/justinas/nosurf"
	"github.com/sanjay-xdr/cmd/internals/config"
	"github.com/sanjay-xdr/cmd/internals/models"
)

var app *config.AppConfig

var pathToTemplates="./templates"

func NewTemplate(a *config.AppConfig) {
	app = a //this app is pointing to the same object as in main

}

func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {

	td.Flash = app.Session.PopString(r.Context(), "flash")
	td.Error = app.Session.PopString(r.Context(), "error")
	td.Warning = app.Session.PopString(r.Context(), "warning")

	td.CSRFToken = nosurf.Token(r)

	// data := nosurf.Token(r)

	// fmt.Print("Value of CSRF token ", data)
	return td
}

// This requires to read from disk again and again lets cache it
func RenderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, td *models.TemplateData) {

	var tc map[string]*template.Template

	if app.UseCache {
		tc = app.TemplateCache
	} else {
		res, err := CreateTemplateCache()

		if err != nil {
			log.Fatal(err)
		}

		tc = res

	}

	t, ok := tc[tmpl]
	var err error
	if !ok {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)
	td = AddDefaultData(td, r)
	err = t.Execute(buf, td)

	if err != nil {
		log.Println("SOmething went wrong", err)
	}

	_, err = buf.WriteTo(w)

	if err != nil {
		log.Println("Something went wrong here as well", err)
	}

}

func CreateTemplateCache() (map[string]*template.Template, error) {

	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob(fmt.Sprintf("%s/*.page.html",pathToTemplates))

	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		// fmt.Println(page)
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		// fmt.Println(ts)

		matches, err :=  filepath.Glob(fmt.Sprintf("%s/*.page.html",pathToTemplates)) //checking for layout
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob(fmt.Sprintf("%s/*.layout.html",pathToTemplates))
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts

	}

	return myCache, nil

}
