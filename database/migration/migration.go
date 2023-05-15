package migration

import (
	"fmt"
	"go-test/database"
	"go-test/models/entity"
	"log"
)

func Migration() {
	err := database.DB.AutoMigrate(&entity.User{}, &entity.Notification{}, &entity.FcmTokens{}, &entity.Topic{})

	if err != nil {
		log.Println(err)
	}

	fmt.Println("Database migrated")
}
