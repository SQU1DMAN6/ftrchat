package model_user

import (
	"context"
	"fmt"

	"github.com/uptrace/bun"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int64  `bun:",pk,autoincrement,notnull"`
	Name     string `bun:",notnull"`
	Email    string `bun:",notnull"`
	Password string `bun:",notnull"`
}

func ModelUser(db *bun.DB) {
	ctx := context.Background()
	db.NewCreateTable().Model((*User)(nil)).Exec(ctx)
	// Create table
	_, err := db.NewCreateTable().Model((*User)(nil)).IfNotExists().Exec(ctx)
	if err != nil {
		fmt.Println("Error creating table:", err)
		return
	}
}

func GetUser(id int, db *bun.DB) (err error, user *User) {
	ctx := context.Background()
	err = db.NewSelect().
		Model(&user).
		Where("id = ?", id).
		Scan(ctx)

	if err != nil {
		fmt.Println("Error querying user:", err)
		return err, nil
	}

	fmt.Printf("User: %+v\n", user)

	return nil, user
}

func CreateUser(db *bun.DB, name string, email string, password string) {
	ctx := context.Background()
	hashedPassword, _ := HashPassword(password)
	user := &User{Name: name, Email: email, Password: hashedPassword}
	db.NewInsert().Model(user).Exec(ctx)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
