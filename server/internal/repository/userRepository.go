package repository

import (
	"healtech-backend/server/internal/models"

	"golang.org/x/crypto/bcrypt"
)

func GetUserByUsername(username string) (*models.LoginMaster, error) {
	// fmt.Println("inside repo")
	query := `SELECT id, register_id, username, password_hash, role_id FROM loginMaster WHERE username = ?`

	row := DB.QueryRow(query, username)
	// fmt.Println("-->", row)
	var user models.LoginMaster
	err := row.Scan(&user.ID, &user.RegisterID, &user.Username, &user.PasswordHash, &user.RoleID)
	if err != nil {
		return nil, err
	}
	// fmt.Println("-->", &user)
	return &user, nil
}

func RegisterUser(username, password string, registerId, roleId int) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	query := `INSERT INTO loginMaster (register_id, username, password_hash, role_id) 
	 VALUES (?, ?, ?, ?)`

	_, err = DB.Exec(query, registerId, username, string(hash), roleId)
	if err != nil {
		return "", err
	}
	return "SUCCESS", nil
}
