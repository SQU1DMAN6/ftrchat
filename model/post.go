package model

import (
	"context"
	"fmt"
	"strings"

	"github.com/uptrace/bun"
)

// type User struct {
//     ID      int64   `bun:",pk,autoincrement"`
//     Name    string  `bun:",notnull"`
//     Posts   []Post  `bun:"rel:has-many,join:id=user_id"`//
// }

// type Post struct {
//     ID     int64 `bun:",pk,autoincrement"`
//     Title  string
//     UserID int64
//     User   *User `bun:"rel:belongs-to,join:user_id=id"`
// }

type BlogPost struct {
	ID        int64  `bun:",pk,autoincrement,notnull"`
	Title     string `bun:",notnull"`
	Slug      string `bun:",notnull"`
	Contents  string `bun:",notnull"`
	TimeStamp int64  `bun:",notnull"`
	Category  string `bun:",notnull"`
	UserID    int64  `bun:",notnull"`
	User      *User  `bun:"rel:belongs-to,join:user_id=id"`
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

		//SELECT * from BlogPost where id = id
	if err != nil {
		fmt.Println("Error querying blog post:", err)
		return nil, err
	}

	fmt.Printf("Blog Post: %+v\n", blogPostModel)

	return &blogPostModel, nil
}

func NewBlogPost(db *bun.DB, title string, category string, timestamp int64, userID int64) error {
	ctx := context.Background()
	userModel, err := GetUser(int(userID), db)
	if err != nil {
		return err
	}
	slug := strings.ToLower(strings.ReplaceAll(title, " ", "-"))
	blogPost := &BlogPost{Title: title, Category: category, TimeStamp: timestamp, UserID: userID, User: userModel, Slug: slug}
	db.NewInsert().Model(blogPost).Exec(ctx)

	return nil
}

func ListBlogPosts(db *bun.DB) ([]BlogPost, error) {
	// var blogPostModel BlogPost
	// ctx := context.Background()
	// err := db.NewSelect().
	// 	Model(&blogPostModel).
	// 	Limit(10).
	// 	Scan(ctx)

	// Select multiple users <----
	ctx := context.Background()
	var posts []BlogPost
	err := db.NewSelect().
		Model(&posts).
		Where("category = ?", "GENERAL").
		Limit(10).
		Scan(ctx)

		//SELECT * from BlogPost limit 10 orderby id asc
	if err != nil {
		fmt.Println("Error querying blog post:", err)
		return nil, err
	}

	return posts, nil
}
