package helpers

import (
	"errors"

	"github.com/rizal282/go-fiber-api/database"
	"github.com/rizal282/go-fiber-api/models"
)

func FindUser(id int, user *models.User) error {
	database.Database.Db.Find(&user, "id = ?", id)

	if user.ID == 0 {
		return errors.New("User tidak ada")
	}

	return nil
}