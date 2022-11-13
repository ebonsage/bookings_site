package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/ebonsage/learngo/pkg/config"
	"github.com/ebonsage/learngo/pkg/handlers"
	"github.com/ebonsage/learngo/pkg/render"
	"log"
	"net/http"
	"time"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	// change this to true in prod
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatalf("Cannont create template cache \n")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	// fmt.Println("Hello, World!")
	//http.HandleFunc("/", handlers.Repo.Home)
	//http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Starting the application on port: %s", portNumber))

	//err = http.ListenAndServe(portNumber, nil)
	//if err != nil {
	//	fmt.Println(err)
	//}
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}
