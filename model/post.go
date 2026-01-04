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

func NewBlogPost(db *bun.DB, title string, contents string, category string, timestamp int64, userID int64) (blogID int64, err error) {
	ctx := context.Background()
	var ids []int64
	userModel, err := GetUser(int(userID), db)
	if err != nil {
		return -1, err
	}
	slug := strings.ToLower(strings.ReplaceAll(title, " ", "-"))
	blogPost := &BlogPost{Title: title, Contents: contents, Category: category, TimeStamp: timestamp, UserID: userID, User: userModel, Slug: slug}
	_, err = db.NewInsert().
		Model(blogPost).
		Returning("id").
		Exec(ctx, &ids)
	if err != nil {
		return -1, err
	}
	fmt.Println("=======================we are testng=======================")
	fmt.Println(ids)
	fmt.Println("=======================we are testng=======================")

	return ids[0], nil
}

func ListBlogPosts(db *bun.DB) ([]BlogPost, error) {
	// Select multiple users <----
	ctx := context.Background()
	var posts []BlogPost
	err := db.NewSelect().
		Model(&posts).
		Where("category = ?", "GENERAL").
		Scan(ctx)
	if err != nil {
		fmt.Println("Error querying blog post:", err)
		return nil, err
	}

	return posts, nil
}

func ListBlogPostsWithPagination(db *bun.DB, paginationOffset int) ([]BlogPost, error) {
	ctx := context.Background()
	var posts []BlogPost
	err := db.NewSelect().
		Model(&posts).
		Where("category = ?", "GENERAL").
		Order("time_stamp DESC").
		Limit(3).
		Offset(paginationOffset).
		Scan(ctx)
	if err != nil {
		fmt.Println("Error querying blog post:", err)
		return nil, err
	}

	return posts, nil
}
