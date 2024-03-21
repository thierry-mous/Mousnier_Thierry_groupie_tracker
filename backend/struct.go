package backend

type Location struct {
	ID        int    `json:"id"`
	ZoneCount int    `json:"zoneCount"`
	Name      string `json:"name"`
}

type Item struct {
	ID          int    `json:"id"`
	Rarity      int    `json:"rarity"`
	CarryLimit  int    `json:"carryLimit"`
	Value       int    `json:"value"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Condition struct {
	Type     string `json:"type"`
	Rank     string `json:"rank"`
	Quantity int    `json:"quantity"`
	Chance   int    `json:"chance"`
	Subtype  string `json:"subtype,omitempty"`
}

type Resistance struct {
	Element   string `json:"element"`
	Condition string `json:"condition,omitempty"`
}

type Weakness struct {
	Element   string `json:"element"`
	Stars     int    `json:"stars"`
	Condition string `json:"condition,omitempty"`
}

type Reward struct {
	ID         int         `json:"id"`
	Item       Item        `json:"item"`
	Conditions []Condition `json:"conditions"`
}

type Monster struct {
	ID          int          `json:"id"`
	Type        string       `json:"type"`
	Species     string       `json:"species"`
	Elements    []string     `json:"elements"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Ailments    []Ailment    `json:"ailments"`
	Locations   []Location   `json:"locations"`
	Resistances []Resistance `json:"resistances"`
	Weaknesses  []Weakness   `json:"weaknesses"`
	Rewards     []Reward     `json:"rewards"`
}

type Ailment struct {
	ID          int        `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Recovery    Recovery   `json:"recovery"`
	Protection  Protection `json:"protection"`
}

type Protection struct {
	Skills []SkillMonster `json:"skills"`
	Items  []Item         `json:"items"`
}

type Recovery struct {
	Actions []string `json:"actions"`
	Items   []Item   `json:"items"`
}
type SkillMonster struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// ArmorSet représente un ensemble d'armure
type ArmorSet struct {
	ID     int    `json:"id"`
	Rank   string `json:"rank"`
	Name   string `json:"name"`
	Pieces []struct {
		ID      int    `json:"id"`
		Type    string `json:"type"`
		Rank    string `json:"rank"`
		Rarity  int    `json:"rarity"`
		Assets2 struct {
			ImageMale2 string `json:"imageMale"`
		} `json:"assets"`
	} `json:"pieces"`
	Bonus interface{} `json:"bonus"`
}

type FullAmrorSet struct {
	Armor ArmorSet
	Piece Piece
}

type Piece struct {
	ID      int    `json:"id"`
	Type    string `json:"type"`
	Rank    string `json:"rank"`
	Rarity  int    `json:"rarity"`
	Assets2 struct {
		ImageMale2 string `json:"imageMale"`
	} `json:"assets"`
}

type ArmorsSets struct {
	ArmorSets []ArmorSet
	PrevPage  int
	NextPage  int
}

// weapon//

type Weapon struct {
	ID     int    `json:"id"`
	Type   string `json:"type"`
	Rarity int    `json:"rarity"`
	Attack struct {
		Display int `json:"display"`
		Raw     int `json:"raw"`
	} `json:"attack"`
	Elderseal  interface{} `json:"elderseal"`
	Attributes struct{}    `json:"attributes"`
	DamageType string      `json:"damageType"`
	Name       string      `json:"name"`
	Durability []struct {
		Red    int `json:"red"`
		Orange int `json:"orange"`
		Yellow int `json:"yellow"`
		Green  int `json:"green"`
		Blue   int `json:"blue"`
		White  int `json:"white"`
		Purple int `json:"purple"`
	} `json:"durability"`
	Slots    []interface{} `json:"slots"`
	Elements []interface{} `json:"elements"`
	Crafting struct {
		Craftable         bool  `json:"craftable"`
		Previous          int   `json:"previous"`
		Branches          []int `json:"branches"`
		CraftingMaterials []struct {
			Quantity int `json:"quantity"`
			Item     struct {
				ID          int    `json:"id"`
				Rarity      int    `json:"rarity"`
				CarryLimit  int    `json:"carryLimit"`
				Value       int    `json:"value"`
				Name        string `json:"name"`
				Description string `json:"description"`
			} `json:"item"`
		} `json:"craftingMaterials"`
		UpgradeMaterials []struct {
			Quantity int `json:"quantity"`
			Item     struct {
				ID          int    `json:"id"`
				Rarity      int    `json:"rarity"`
				CarryLimit  int    `json:"carryLimit"`
				Value       int    `json:"value"`
				Name        string `json:"name"`
				Description string `json:"description"`
			} `json:"item"`
		} `json:"upgradeMaterials"`
	} `json:"crafting"`
	Assets struct {
		Icon  string `json:"icon"`
		Image string `json:"image"`
	} `json:"assets"`
}

type Weapons struct {
	Weapons []Weapon
}

