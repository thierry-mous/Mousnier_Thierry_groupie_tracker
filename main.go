package main

import (
	"groupie/routeur"
	"groupie/templates"
)

func main() {
	templates.InitTemplate()
	routeur.InitServ()
}
