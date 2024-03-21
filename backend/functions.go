package backend

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func FetchMonsterData(url string) ([]Monster, error) {

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	
	var monsters []Monster

	
	err = json.Unmarshal(body, &monsters)
	if err != nil {
		return nil, err
	}

	return monsters, nil
}

func FetchArmorSets(link string) (ArmorsSets, error) {
	response, err := http.Get(link)
	if err != nil {
		return ArmorsSets{}, fmt.Errorf("Erreur lors de la requête HTTP: %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return ArmorsSets{}, fmt.Errorf("Erreur de requête: %s", response.Status)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return ArmorsSets{}, fmt.Errorf("Erreur lors de la lecture du corps de la réponse: %v", err)
	}

	var armorSets []ArmorSet
	err = json.Unmarshal(body, &armorSets)
	if err != nil {
		return ArmorsSets{}, fmt.Errorf("Erreur lors du décodage JSON: %v", err)
	}

	return ArmorsSets{ArmorSets: armorSets}, nil
}


func FetchWeaponData(url string) ([]Weapon, error) {

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()


	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	
	var weapons []Weapon

	err = json.Unmarshal(body, &weapons)
	if err != nil {
		return nil, err
	}

	return weapons, nil
}


