package transport

import (
	"github.com/gorilla/mux"
	"html/template"
	"l0/internal/models"
	"log"
	"net/http"
	"os"
)

const Port = ":8181"

func MainPage(resp http.ResponseWriter, req *http.Request) {
	Title := "Main page"

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	tmpl, err := template.ParseFiles(wd+"/templates/index.html",
		wd+"/templates/header.html", wd+"/templates/footer.html")
	if err != nil {
		log.Fatalf("Error in parsing templates: %v\n", err)
	}

	var ordersList []models.Order
	for _, v := range models.Cache {
		ordersList = append(ordersList, v)
	}

	err = tmpl.ExecuteTemplate(resp, "index", struct {
		Title string
		List  []models.Order
	}{Title: Title, List: ordersList})
	if err != nil {
		log.Println("Error: can't execute template.")
	}
}

func MessageHandler(resp http.ResponseWriter, req *http.Request) {
	Title := "Message page"
	vars := mux.Vars(req)
	uid := vars["uid"]

	a, ok := models.Cache[uid]

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	if ok == false {
		tmpl, err := template.ParseFiles(wd+"/templates/404.html",
			wd+"/templates/header.html", wd+"/templates/footer.html")
		if err != nil {
			log.Fatalf("Error in parsing templates: %v\n", err)
		}

		Title = "Not found"

		err = tmpl.ExecuteTemplate(resp, "404", struct {
			Title string
		}{Title: Title})
		if err != nil {
			log.Println("Error: can't execute template.")
		}
	} else {

		tmpl, err := template.ParseFiles(wd+"/templates/message.html",
			wd+"/templates/header.html", wd+"/templates/footer.html")
		if err != nil {
			log.Fatalf("Error in parsing templates: %v\n", err)
		}

		err = tmpl.ExecuteTemplate(resp, "message", struct {
			Title string
			Order models.Order
		}{Title: Title, Order: a})
		if err != nil {
			log.Println("Error: can't execute template.")
		}
	}
}
