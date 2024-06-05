package controllers

import (
	"aplikasi-layanan-umum/database"
	"aplikasi-layanan-umum/models"
)

func GetAllUser() ([]models.User, error) {
	db := database.GetConnection()

	script := "SELECT id, username, password, role FROM users;"
	rows, err := db.Query(script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		user := models.User{}
		if err = rows.Scan(&user.Id, &user.Username, &user.Password, &user.Role); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
