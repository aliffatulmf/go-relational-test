package main

import (
	"flag"

	"test/content"
	"test/database"
	"test/finder"
	"test/models"
)

func main() {
	db := database.New()
	defer db.Close()
	belongsTo := flag.Bool("belongsTo", false, `auto migrate "belongs to" model`)
	hasMany := flag.Bool("hasMany", false, `auto migrate "has many" model`)

	findBelongsTo := flag.Bool("findBelongsTo", false, `find "belongs to"`)
	findHasMany := flag.Bool("findHasMany", false, `find "has many"`)
	query := flag.String("query", "", "")
	flag.Parse()

	if *belongsTo {
		db.DB.Migrator().DropTable(&models.Employee{}, &models.Company{})
		db.DB.AutoMigrate(&models.Employee{}, &models.Company{})
		content.InjectBelongsTo(db.DB)
	}

	if *hasMany {
		db.DB.Migrator().DropTable(&models.CreditCard{}, &models.Customer{})
		db.DB.AutoMigrate(&models.CreditCard{}, &models.Customer{})
		content.InjectHasMany(db.DB)
	}

	if *findBelongsTo {
		finder.FindBelongsTo(db.DB, *query)
	}

	if *findHasMany {
		finder.FindHasMany(db.DB, *query)
	}
}
