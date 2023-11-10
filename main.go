package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

// VARIABLES & CONSTANTES
const port = "localhost:8080"

var (
	counter int
	temp    *template.Template
	err     error
)

type DataCounter struct {
	Message string
	Counter int
}

type Promotion struct {
	Nom      string
	Filiere  string
	Niveau   int
	Students []Student
}
type Student struct {
	Nom    string
	Prenom string
	Sexe   string
	Age    int
}
type DataPromotion struct { //informations sur les édonnées envoyés
	Promotion1      Promotion
	NombreEtudiants int
}

func main() {
	temp, err = template.ParseGlob("./templates/*.html")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fileServer := http.FileServer(http.Dir("CSS"))
	http.Handle("/CSS/", http.StripPrefix("/CSS/", fileServer))
	http.HandleFunc("/promo", promoHandler)
	http.HandleFunc("/change", changeHandler)
	http.HandleFunc(" /user/init", UserHandler)
	http.HandleFunc(" /user/treatment", TreatHandler)
	http.ListenAndServe(port, nil)

}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	//data := PageData{}
	//temp.ExecuteTemplate(w, "user", data)
}

func TreatHandler(w http.ResponseWriter, r *http.Request) {
	//	data := PageData{}
	//	temp.ExecuteTemplate(w, "treat", data)
}

func changeHandler(w http.ResponseWriter, r *http.Request) {
	counter++
	message := ""

	if counter%2 == 0 {
		message = "Le chiffre est pair"
	} else {
		message = "Le chiffre est impair"
	}

	data := DataCounter{
		Message: message,
		Counter: counter,
	}
	temp.ExecuteTemplate(w, "compt", data)
}

func promoHandler(w http.ResponseWriter, r *http.Request) {
	DataPage := DataPromotion{
		Promotion1: Promotion{
			Nom:     "Mentor'ac",
			Filiere: "Informatique",
			Niveau:  5,
			Students: []Student{
				{Nom: "RODRIGUES", Prenom: "Cyril", Sexe: "M", Age: 22},
				{Nom: "MEDERREG", Prenom: "Kheir-eddine", Sexe: "F", Age: 22},
				{Nom: "PHILIPIERT", Prenom: "Alan", Sexe: "M", Age: 26}},
		},
		NombreEtudiants: 3,
	}
	temp.ExecuteTemplate(w, "promot", DataPage)
}
