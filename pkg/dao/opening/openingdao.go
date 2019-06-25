package openingdao

import (
	"github.com/hunhunyosshy/black-and-white/pkg/db"
)

type Opening struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func FetchIndex() []Opening {
	db := db.Connect()
	defer db.Close()

	//rowを取得
	rows, err := db.Query("SELECT * FROM opening")
	if err != nil {
		panic(err.Error())
	}
	openingArgs := make([]Opening, 0)
	for rows.Next() {
		var opening Opening
		err = rows.Scan(&opening.ID, &opening.Name)
		if err != nil {
			panic(err.Error())
		}
		openingArgs = append(openingArgs, opening)
	}
	return openingArgs
}

func FetchByKey(id string) []Opening {
	db := db.Connect()
	defer db.Close()

	//rowを取得
	rows, err := db.Query("SELECT * FROM opening WHERE opening_id = ?", id)
	if err != nil {
		panic(err.Error())
	}
	openingArgs := make([]Opening, 0)
	for rows.Next() {
		var opening Opening
		err = rows.Scan(&opening.ID, &opening.Name)
		if err != nil {
			panic(err.Error())
		}
		openingArgs = append(openingArgs, opening)
	}
	return openingArgs
}
