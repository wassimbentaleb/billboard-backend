package initializers

import (
	"awesomeProject1/models"
	"fmt"
)

func SyncDatabase() {
	err := DB.AutoMigrate(&models.User{}, &models.Client{}, &models.Plans{})
	if err != nil {
		fmt.Println(err)
		return
	}

}
