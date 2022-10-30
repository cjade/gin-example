package migrations

import (
	"gin-example/app/models"
	"gin-example/utils/db"
	"log"
)

func Setup() {
	m := db.Mysql.Set("gorm:table_options", "ENGINE=InnoDB").Migrator()
	book := &models.Books{}
	err := m.CreateTable(book)
	if err == nil {
		log.Fatalln(err.Error())
	}
}
