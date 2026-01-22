package model

import (
	"context"
	"errors"
	"fmt"

	"github.com/uptrace/bun"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID         int64          `bun:",pk,autoincrement,notnull"`
	Name       string         `bun:",notnull"`
	Email      string         `bun:",notnull"`
	Password   string         `bun:",notnull"`
	Posts      []BlogPost     `bun:"rel:has-many,join:id=user_id"`
	Categories []BlogCategory `bun:"rel:has-many,join:id=user_id"`
}

func ModelUser(db *bun.DB) error {
	ctx := context.Background()
	_, err := db.NewCreateTable().
		Model((*User)(nil)).
		IfNotExists().
		Exec(ctx)

	return err
}

func GetUser(id int, db *bun.DB) (*User, error) {
	var userModel User
	ctx := context.Background()
	err := db.NewSelect().
		Model(&userModel).
		Where("id = ?", id).
		Scan(ctx)

	if err != nil {
		fmt.Println("Error querying user:", err)
		return nil, err
	}

	fmt.Printf("User: %+v\n", userModel)

	return &userModel, nil
}

func GetUserByEmail(email string, db *bun.DB) (*User, error) {
	var user User
	ctx := context.Background()

	err := db.NewSelect().
		Model(&user).
		Where("email = ?", email).
		Scan(ctx)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUserByName(username string, db *bun.DB) (*User, error) {
	var user User

	ctx := context.Background()

	err := db.NewSelect().
		Model(&user).
		Where("name = ?", username).
		Scan(ctx)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func CreateUser(db *bun.DB, name string, email string, password string) {
	ctx := context.Background()
	hashedPassword, _ := HashPassword(password)
	user := &User{Name: name, Email: email, Password: hashedPassword}
	db.NewInsert().Model(user).Exec(ctx)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func CheckPassword(db *bun.DB, email, plainTextPassword string) (*User, error) {
	usr, err := GetUserByEmail(email, db)
	if err != nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword(
		[]byte(usr.Password),
		[]byte(plainTextPassword),
	); err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return nil, err
		}
		return nil, err
	}
	return usr, nil

}
