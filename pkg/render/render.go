package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"

	"path/filepath"

	"github.com/AbahEkang/bookings/pkg/Config"
	"github.com/AbahEkang/bookings/pkg/models"
)

// var functions = template.FuncMap{

// }

var app *Config.AppConfig

// NewTemplates sets the config for the template package
func NewTemplates(a *Config.AppConfig){
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData{
	return td
}
// More complex caching method
func RenderTemplates(w http.ResponseWriter, tmpl string, td *models.TemplateData){

	var tc map[string]*template.Template
	if app.UseCache{
		// get the template cache from the app config
		tc = app.TemplateCache
	}else{
		tc, _ = CreateTemplateCache()
	}
	
	// create a template cache


	// tc, err := CreateTemplateCache()
	// if err != nil{
	// 	log.Fatal(err)
	// }

	// get requested template from cache
	t, ok := tc[tmpl]

	if !ok{
		log.Fatal("could not get template from template cache")
	}

	theBuffer := new(bytes.Buffer)

	td = AddDefaultData(td)	
	_ = t.Execute(theBuffer, td)

	// render the template

	// parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl, "./templates/base.layout.tmpl.html")

	// err := parsedTemplate.Execute(w, nil)

	// if err != nil{
	// 	fmt.Println("error parsing template:", err)
	// }
	
	_, err := theBuffer.WriteTo(w)
	if err != nil{
		log.Println(err)
	}

}

func CreateTemplateCache()(map[string]*template.Template, error){

	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl.html")

	if err != nil{
		return myCache, err
	}

	// range through all files ending with *.page.tmpl.html
	for _, page := range pages{
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)

		if err != nil{
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl.html")

		
		if err != nil{
			return myCache, err
		}

		if len(matches)>0{

			ts, err = ts.ParseGlob("./templates/*.layout.tmpl.html")

			if err != nil{
				return myCache, err
			}

		}
		myCache[name] = ts
		
	}

	return myCache, nil
}


// LESS COMPLEX CACHING
// func RenderTemplate(w http.ResponseWriter, tmpl string) {

// 	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl, "./templates/base.layout.tmpl.html")

// 	err := parsedTemplate.Execute(w, nil)
// 	if err != nil {
// 		fmt.Println("error parsing template", err)
// 		return
// 	}

// }

// var tc = make(map[string]*template.Template)

// func RenderTemplates(w http.ResponseWriter, t string){

// 	var tmpl *template.Template

// 	var err error

// 	// check to see if we already have the template in our cache

// 	_, inMap := tc[t]

// 	if !inMap{

// 		//need to create the template
// 		log.Println("creating template and adding to cache")
// 		err = createTemplateCache(t)
// 		if err != nil{
// 			log.Println(err)
// 		}

// 	}else{

// 		// we have the template in the cache
// 		log.Println("using the cached template")
// 	}

// 	tmpl = tc[t]

// 	err = tmpl.Execute(w, nil)

// 	if err != nil{
// 		log.Println(err)
// 	}



// }

// func createTemplateCache(t string) error{

// 	templates := []string{

// 		fmt.Sprintf("./templates/%s", t), "./templates/base.layout.tmpl.html",

// 	}

// 	// parse the templates

// 	tmpl, err := template.ParseFiles(templates...)

// 	if err != nil{
// 		return err
// 	}

// 	// add template to cache (map)
// 	tc[t] = tmpl

// 	return nil

// }