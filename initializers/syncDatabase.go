package initializers

import "awesomeProject1/models"

func SyncDatabase() {
	err := DB.AutoMigrate(&models.User{}, &models.Post{})
	if err != nil {
		return
	}

}
