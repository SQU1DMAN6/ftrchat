package model

import (
	"context"
	"strings"

	"fmt"

	"github.com/uptrace/bun"
)

type BlogCategory struct {
	ID     int64      `bun:",pk,autoincrement,notnull"`
	Name   string     `bun:",notnull"`
	Slug   string     `bun:",notnull"`
	UserID int64      `bun:",notnull"`
	User   *User      `bun:"rel:belongs-to,join:user_id=id"`
	Posts  []BlogPost `bun:"rel:has-many,join:id=user_id"`
}

func ModelBlogCategory(db *bun.DB) error {
	ctx := context.Background()
	_, err := db.NewCreateTable().
		Model((*BlogCategory)(nil)).
		IfNotExists().
		Exec(ctx)

	return err
}

func GetBlogCategory(id int, db *bun.DB) (*BlogCategory, error) {
	var blogCategoryModel BlogCategory
	ctx := context.Background()
	err := db.NewSelect().
		Model(&blogCategoryModel).
		Where("id = ?", id).
		Scan(ctx)

	if err != nil {
		fmt.Println("Error querying blog category:", err)
		return nil, err
	}

	fmt.Printf("Blog Category Name: %+v\n", blogCategoryModel)

	return &blogCategoryModel, nil
}

func GetBlogCategoryByName(name string, db *bun.DB) (*BlogCategory, error) {
	var blogCategory BlogCategory
	ctx := context.Background()

	err := db.NewSelect().Model(&blogCategory).Where("name = ?", name).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return &blogCategory, nil
}

func CreateBlogPostCategory(db *bun.DB, name string) {
	ctx := context.Background()
	slug := strings.ReplaceAll(name, " ", "-")
	blogPostCategory := &BlogCategory{Name: name, Slug: slug}
	db.NewInsert().Model(blogPostCategory).Exec(ctx)
}
