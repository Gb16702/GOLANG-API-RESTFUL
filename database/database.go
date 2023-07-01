package database

import (
	"log"
	"todo-go/database/config"
	"todo-go/database/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)


var DB *gorm.DB

func ConnectToDb() {

	connectionString := config.Settings.ConnString

	var err error
    DB, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})

    if err != nil {
        panic(err)
    }

    DB.Logger = logger.Default.LogMode(logger.Info)
    log.Println("Connecté à la base de données")

	// Migrate the schema

	defer func() {
		if r := recover(); r != nil {
			log.Println("Erreur lors de la migration de la base de données : ", r)
		}
	}()

	DB.AutoMigrate(&models.Todo{})
	log.Println("Migrations terminées")

	insertTodo()

	if err != nil {
		log.Fatal("Erreur lors de la migration de la base de données : ", err)
	}

}

var todos = []*models.Todo{
	{Item: "Apprendre Go", Completed: false},
}

func insertTodo() {
	for _, todo := range todos {
		t := *todo
		if err := DB.Where("item = ?", t.Item).First(&t).Error; err == nil {
			continue
		}
		if err := DB.Create(&t).Error; err != nil {
			panic(err)
		}
		log.Println("Todo insérée")
	}
}