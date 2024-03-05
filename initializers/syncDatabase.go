package initializers

import "awesomeProject1/models"

func SyncDatabase() {
	err := DB.AutoMigrate(&models.User{}, &models.Client{})
	if err != nil {
		return
	}

}
