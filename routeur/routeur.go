package routeur

import (
	"fmt"
	"groupie/controller"
	"log"
	"net/http"
	"os"
)

func InitServ() {
	rootDoc, _ := os.Getwd()
	fileserver := http.FileServer(http.Dir(rootDoc + "/asset"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))

	http.HandleFunc("/accueil", controller.AccueilPage)
	http.HandleFunc("/category", controller.CategoryPage)
	http.HandleFunc("/monster", controller.MonsterPage)
	http.HandleFunc("/armor", controller.ArmorPage)
	http.HandleFunc("/weapon", controller.WeaponPage)
	http.HandleFunc("/add", controller.AddFav)
	http.HandleFunc("/search", controller.SearchFunc)

	//Init serv
	log.Println(" Serveur lancé !")
	fmt.Println("http://localhost:8080/accueil")
	http.ListenAndServe("localhost:8080", nil)
}
