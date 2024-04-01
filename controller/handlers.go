package controller

import (
	"encoding/json"
	"fmt"
	"groupie/backend"
	"groupie/templates"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func AccueilPage(w http.ResponseWriter, r *http.Request) {
	templates.Temp.ExecuteTemplate(w, "accueil", nil)
}

func CategoryPage(w http.ResponseWriter, r *http.Request) {
	cat := r.URL.Query().Get("type")
	currentPage, errPage := strconv.Atoi(r.URL.Query().Get("page"))
	if errPage != nil || currentPage < 1 {
		currentPage = 1
	}
	var firstElement int = 4 * (currentPage - 1)
	var lastElement int = firstElement + 4

	switch cat {
	case "armor":
		link := "https://mhw-db.com/armor/sets"
		armorSets, err := backend.FetchArmorSets(link)
		if err != nil {
			fmt.Println("Erreur lors de la récupération des données:", err)
			return
		}

		if lastElement > len(armorSets.ArmorSets) {
			lastElement = len(armorSets.ArmorSets)
		}
		templates.Temp.ExecuteTemplate(w, "categoryarmor", backend.ArmorsSets{ArmorSets: armorSets.ArmorSets[firstElement:lastElement], PrevPage: currentPage - 1, NextPage: currentPage + 1})
	case "weapon":
		weapons, err := backend.FetchWeaponData("https://mhw-db.com/weapons")
		if err != nil {
			fmt.Println("Erreur lors de la récupération des données d'armes:", err)
			http.Error(w, "Erreur lors de la récupération des données d'armes", http.StatusInternalServerError)
			return
		}
		fmt.Println(weapons)

		templates.Temp.ExecuteTemplate(w, "categoryweapon", weapons)
	case "monster":
		monsters, err := backend.FetchMonsterData("https://mhw-db.com/monsters")
		if err != nil {
			fmt.Println("Erreur lors de la récupération des données de monstres:", err)
			http.Error(w, "Erreur lors de la récupération des données de monstres", http.StatusInternalServerError)
			return
		}
		templates.Temp.ExecuteTemplate(w, "categorymonster", monsters)
	default:
		http.Error(w, "Catégorie non valide", http.StatusBadRequest)
	}
}

func MonsterPage(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	fmt.Println(id)

	monsters, err := backend.FetchMonsterData("https://mhw-db.com/monsters")
	if err != nil {
		fmt.Println("Erreur lors de la récupération des données de monstres:", err)
		http.Error(w, "Erreur lors de la récupération des données de monstres", http.StatusInternalServerError)
		return
	}

	var monsterIdData backend.Monster

	for _, monster := range monsters {
		if fmt.Sprint(monster.ID) == id {
			monsterIdData = monster
			fmt.Println("Trouvé")
			break
		}
	}

	templates.Temp.ExecuteTemplate(w, "resourcepagemonster", monsterIdData)
}

func ArmorPage(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	fmt.Println(id)

	link := "https://mhw-db.com/armor/sets"
	armorSets, err := backend.FetchArmorSets(link)
	if err != nil {
		fmt.Println("Erreur lors de la récupération des données:", err)
		return
	}

	var pieceData backend.Piece
	var armorSetData backend.ArmorSet
	var fullData backend.FullAmrorSet

	for _, armor := range armorSets.ArmorSets {
		for _, piece := range armor.Pieces {
			if fmt.Sprint(piece.ID) == id {
				pieceData = piece
				armorSetData = armor
				fmt.Println("Trouvé")
				break
			}
		}
	}

	fullData.Armor = armorSetData
	fullData.Piece = pieceData

	templates.Temp.ExecuteTemplate(w, "resourcepagearmor", fullData)

}

func WeaponPage(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	fmt.Println(id)

	weapons, err := backend.FetchWeaponData("https://mhw-db.com/weapons")
	if err != nil {
		fmt.Println("Erreur lors de la récupération des données d'armes:", err)
		http.Error(w, "Erreur lors de la récupération des données d'armes", http.StatusInternalServerError)
		return
	}

	var weaponData backend.Weapon

	for _, weapon := range weapons {
		if fmt.Sprint(weapon.ID) == id {
			weaponData = weapon
			fmt.Println("Trouvé")
			break
		}
	}

	fmt.Println(weaponData)
	fmt.Println("Execute")
	templates.Temp.ExecuteTemplate(w, "ressource-weapon", weaponData)
}

type ArmorFav struct {
	IdArmor int `json:"id_armor"`
	IdPiece int `json:"id_piece"`
}
type Favorites struct {
	ListArmor   []ArmorFav `json:"ListArmor"`
	ListWeapon  []int      `json:"ListWeapon"`
	ListMonster []int      `json:"ListMonster"`
}

// Declarer struct pour le fichier JSON
func AddFav(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	type_fav := r.URL.Query().Get("type")
	id_piece := r.URL.Query().Get("id_piece")
	idFav, _ := strconv.Atoi(id)
	idPiece, _ := strconv.Atoi(id_piece)
	fmt.Println(id, type_fav, id_piece)
	// Afficher les paramètres pour vérification

	// Lire le fichier JSON de favoris
	file, err := os.OpenFile("favorite.json", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Décode les données JSON dans une structure de favoris
	var favorites Favorites
	err = json.NewDecoder(file).Decode(&favorites)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	switch type_fav {
	case "armor":
		favorites.ListArmor = append(favorites.ListArmor, ArmorFav{IdArmor: idFav, IdPiece: idPiece})
	case "monster":
		favorites.ListMonster = append(favorites.ListMonster, idFav)
	case "weapon":
		favorites.ListWeapon = append(favorites.ListWeapon, idFav)
	}
	data, errData := json.Marshal(favorites)
	if errData != nil {
		fmt.Println("Erreur lors de la conversion en JSON:", errData)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	os.WriteFile("favorite.json", data, 0644)
	// Rediriger vers la page des favoris
	http.Redirect(w, r, "/templates/favorite-page.html", http.StatusSeeOther)
}

func SearchFunc(w http.ResponseWriter, r *http.Request) {
	search := r.URL.Query().Get("search")
	fmt.Println(search)
	dataWeapon, errWeapon := backend.FetchWeaponData("https://mhw-db.com/weapons")
	if errWeapon != nil {
		fmt.Println("Erreur lors de la récupération des données d'armes:", errWeapon)
	}

	var ListSearchWeapon []backend.Weapon
	for _, weapon := range dataWeapon {
		if strings.Contains(weapon.Name, search) {
			ListSearchWeapon = append(ListSearchWeapon, weapon)
		}
	}

	templates.Temp.ExecuteTemplate(w, "search", ListSearchWeapon)

}

type ResultFilter struct {
	Monster []backend.Monster
	Armor   []backend.ArmorSet
	Weapon  []backend.Weapon
}

func CollectionPage(w http.ResponseWriter, r *http.Request) {
	typeCollection := r.URL.Query().Get("collection")
	var data ResultFilter

	switch typeCollection {
	case "monster":
		monsters, err := backend.FetchMonsterData("https://mhw-db.com/monsters")
		if err != nil {
			fmt.Println("Erreur lors de la récupération des données de monstres:", err)
			http.Error(w, "Erreur lors de la récupération des données de monstres", http.StatusInternalServerError)
			return
		}
		data.Monster = monsters
		fmt.Println("iciciciciciciccicicicici", data.Monster)
		filterMonster := r.FormValue("monster")
		fmt.Println("iciciciciciciccicicicici", filterMonster)

		if filterMonster != "" {
			var ListFilterMonster []backend.Monster
			for _, monster := range data.Monster {
				for _, element := range monster.Elements {
					if strings.Contains(element, filterMonster) {
						ListFilterMonster = append(ListFilterMonster, monster)
						break
					}
				}
			}
			data.Monster = ListFilterMonster
		}
		break
	case "armor":
		armorSets, err := backend.FetchArmorSets("https://mhw-db.com/armor/sets")
		if err != nil {
			fmt.Println("Erreur lors de la récupération des données:", err)
			return
		}
		data.Armor = armorSets.ArmorSets
		filterArmor := r.FormValue("armor")
		if filterArmor != "" {
			var ListFilterArmor []backend.ArmorSet
			for _, armorSets := range data.Armor {
				for _, piece := range armorSets.Pieces {
					if strings.Contains(piece.Type, filterArmor) {
						ListFilterArmor = append(ListFilterArmor, armorSets)
						break
					}
				}
			}
			data.Armor = ListFilterArmor
		}
		break
	case "weapon":
		weapons, err := backend.FetchWeaponData("https://mhw-db.com/weapons")
		if err != nil {
			fmt.Println("Erreur lors de la récupération des données d'armes:", err)
			http.Error(w, "Erreur lors de la récupération des données d'armes", http.StatusInternalServerError)
			return
		}
		data.Weapon = weapons
		filterWeapon := r.FormValue("weapon")
		if filterWeapon != "" {
			var ListFilterWeapon []backend.Weapon
			for _, weapon := range data.Weapon {
				if strings.Contains(weapon.Type, filterWeapon) {
					ListFilterWeapon = append(ListFilterWeapon, weapon)
				}
			}
			data.Weapon = ListFilterWeapon
		}
		break
	default:
		data.Monster, _ = backend.FetchMonsterData("https://mhw-db.com/monsters")
		armors, _ := backend.FetchArmorSets("https://mhw-db.com/armor/sets")
		data.Weapon, _ = backend.FetchWeaponData("https://mhw-db.com/weapons")
		data.Armor = armors.ArmorSets
	}

	templates.Temp.ExecuteTemplate(w, "collection", data)
}
