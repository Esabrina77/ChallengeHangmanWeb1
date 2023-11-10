package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

const port = "localhost:8080"

type Promotion struct { //informations sur la promotion
	Nom      string
	Filiere  string
	Niveau   int
	Students []Student
}
type Student struct { //informations sur les étudiants
	Nom    string
	Prenom string
	Sexe   string
	Age    int
}

type DataPromotion struct { //informations sur les édonnées envoyés
	Promotion1      Promotion
	NombreEtudiants int
}
type CSS string

func main() {
	temp, err := template.ParseGlob("./templates/*.html")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	http.HandleFunc("/promo", func(w http.ResponseWriter, r *http.Request) {

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
	})

	fileServer := http.FileServer(http.Dir("CSS"))
	http.Handle("/CSS/", http.StripPrefix("/CSS/", fileServer))
	http.ListenAndServe(port, nil)
}
