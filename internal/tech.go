package internal

import (
	"log"

	"github.com/meastblue/teckwatch/database"
	"github.com/meastblue/teckwatch/model"
)

func GetTechList() []model.Tech {
	techs := []model.Tech{}
	db, err := database.New()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	stmt, err := db.Preparex(`select * from techs`)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		tech := model.Tech{}
		err := rows.Scan(&tech.ID, &tech.Label, &tech.Description, &tech.CreatedAt, &tech.UpdatedAt)
		if err != nil {
			log.Fatal(err)
		}
		techs = append(techs, tech)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return techs
}

func CreateTech(t model.Tech) string {
	db, err := database.New()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := db.Preparex(`insert into techs(label,description) values($1, $2)`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(t.Label, t.Description)
	if err != nil {
		log.Fatal(err)
	}

	return "OK"
}
