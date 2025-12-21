package model

import (
	"context"
	"fmt"

	"github.com/uptrace/bun"
)

type BlogPost struct {
	ID           int64  `bun:",pk,autoincrement,notnull"`
	Title        string `bun:",notnull"`
	Introductory string `bun:",notnull"`
	Slug         string `bun:",notnull"`
	Contents     string `bun:",notnull"`
	TimeStamp    int64  `bun:",notnull"`
	Category     string `bun:",notnull"`
}

func ModelBlogPost(db *bun.DB) error {
	ctx := context.Background()
	_, err := db.NewCreateTable().
		Model((*BlogPost)(nil)).
		IfNotExists().
		Exec(ctx)

	return err
}

func GetBlogPost(id int, db *bun.DB) (*BlogPost, error) {
	var blogPostModel BlogPost
	ctx := context.Background()
	err := db.NewSelect().
		Model(&blogPostModel).
		Where("id = ?", id).
		Scan(ctx)

	if err != nil {
		fmt.Println("Error querying blog post:", err)
		return nil, err
	}

	fmt.Printf("Blog Post: %+v\n", blogPostModel)

	return &blogPostModel, nil
}

// func NewBlogPost(db *bun.DB, name string, category string, timestamp int64)
