package render

import (
	"bytes"
	"github.com/ebonsage/learngo/pkg/config"
	"github.com/ebonsage/learngo/pkg/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// var functions = template.FuncMap{}

var app *config.AppConfig

// NewTemplates Set the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

// RenderingTemplate RenderingTemplates used to
func RenderingTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}
	// create a template cache
	// after... get the template cache from he appconfig
	// tc := app.TemplateCache

	// get the requested template from cache

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	err := t.Execute(buf, td)
	if err != nil {
		log.Println(err)
	}

	// ren2der the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}

	/*parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template", err)
		return
	}*/

}

func CreateTemplateCache() (map[string]*template.Template, error) {
	// make a map 1/2 -> myCache := make(map[string]*template.Template)
	myCache := map[string]*template.Template{}

	// get all the files named *.tmpl from the ./templates folder
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	// range through all files ending with *.page.tmpl
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}

		}

		myCache[name] = ts
	}
	return myCache, err

}

/*

var tc = make(map[string]*template.Template)

func RenderingTemplate(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	// check to see if we already have the template in our cache
	_, inMap := tc[t]
	if !inMap {
		// need to create the template
		err = createTemplateCache(t)
		if err != nil {
			log.Printf("Error adding template to cache. -> %s\n", err)
		}
	} else {
		// we have the template in cache
		log.Println("Using cache template")
	}

	tmpl = tc[t]

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Printf("Error with templeting. -> %s\n", err)
	}
}

func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.tmpl",
	}

	// parse the template
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}

	// add template to cache (which is a map)
	tc[t] = tmpl

	return nil
}
*/
