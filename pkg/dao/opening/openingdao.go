package openingdao

import (
	"log"

	"github.com/hunhunyosshy/black-and-white/pkg/db"
)

// Opening is struct which stores how to fight about black-and-white.
type Opening struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// FetchIndex gets list of opening.
func FetchIndex() []Opening {
	db := db.Connect()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM opening")
	if err != nil {
		log.Fatal(err.Error())
	}
	openingArgs := make([]Opening, 0)
	for rows.Next() {
		var opening Opening
		err = rows.Scan(&opening.ID, &opening.Name)
		if err != nil {
			log.Fatal(err.Error())
		}
		openingArgs = append(openingArgs, opening)
	}
	return openingArgs
}

// FetchByKey gets opening by id.
func FetchByKey(id string) []Opening {
	db := db.Connect()
	defer db.Close()

	//rowを取得
	rows, err := db.Query("SELECT * FROM opening WHERE opening_id = ?", id)
	if err != nil {
		log.Fatal(err.Error())
	}
	openingArgs := make([]Opening, 0)
	for rows.Next() {
		var opening Opening
		err = rows.Scan(&opening.ID, &opening.Name)
		if err != nil {
			log.Fatal(err.Error())
		}
		openingArgs = append(openingArgs, opening)
	}
	return openingArgs
}
