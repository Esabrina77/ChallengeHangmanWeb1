package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"time"
)

// VARIABLES & CONSTANTES
const port = "localhost:8080"

var (
	counter int
	temp    *template.Template
	err     error
	user    DataUser
)

type DataUser struct {
	Nom           string
	Prenom        string
	DateNaissance string
	Sexe          string
}

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
	http.HandleFunc("/user/init", InitHandler)
	http.HandleFunc("/user/treatment", TreatHandler)
	http.HandleFunc("/user/display", DisplayHandler)
	http.ListenAndServe(port, nil)

}

func InitHandler(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "init", nil)
}

func TreatHandler(w http.ResponseWriter, r *http.Request) {
	user = DataUser{
		Nom:           r.FormValue("name"),
		Prenom:        r.FormValue("surname"),
		DateNaissance: r.FormValue("Date"),
		Sexe:          r.FormValue("gender"),
	}
	if user.Nom == "" || user.Prenom == "" || user.DateNaissance == "" || user.Sexe == "" {
		errorMessage := "VEILLEZ REMPLIR TOUS  LES CHAMPS DU FORMULAIRE"
		http.Redirect(w, r, "/user/init?error="+errorMessage, http.StatusSeeOther)
		return
	}

	dateStr := r.FormValue("Date")
	layout := "02/01/2006" // Format attendu : jour/mois/année (02/01/2006)
	date, err := time.Parse(layout, dateStr)

	if err != nil {
		errorMessage := "Format de date incorrect. Veuillez entrer la date au format jj/mm/aaaa| Fevrier:Respectez la norme 'année bisextile/année ordinaire'"
		http.Error(w, errorMessage, http.StatusBadRequest)
		return
	}
	if date.Year() < 1900 || date.Year() > 2023 {
		errorMessage := "Format de date incorrect.problème au niveau de 'aa' .La chronologie n'est pas respectée"
		http.Error(w, errorMessage, http.StatusBadRequest)
		return
	}
	if date.Month() < 0 || date.Month() > 12 {
		errorMessage := "Format de date incorrect. mm est compris entre 1 & 12"
		http.Error(w, errorMessage, http.StatusBadRequest)
		return
	}

	if date.Day() < 0 || date.Day() > 31 {
		errorMessage := "Format de jj incorrect pour cette date."
		http.Error(w, errorMessage, http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "/user/display", http.StatusSeeOther)
}

func DisplayHandler(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "display", user)
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
