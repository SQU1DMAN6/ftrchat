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
	Category  int64  `bun:"rel:belongs-to,notnull"`
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
		Relation("User").
		Where("?TableAlias.id = ?", id).
		// Where("id = ?", id).
		Scan(ctx)

		//SELECT * from BlogPost where id = id
	if err != nil {
		fmt.Println("Error querying blog post:", err)
		return nil, err
	}

	fmt.Printf("Blog Post: %+v\n", blogPostModel)

	return &blogPostModel, nil
}

func NewBlogPost(db *bun.DB, title string, contents string, category int64, timestamp int64, userID int64) (blogID int64, err error) {
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
		Relation("User").
		Where("category = ?", "GENERAL").
		Scan(ctx)
	if err != nil {
		fmt.Println("Error querying blog post:", err)
		return nil, err
	}

	return posts, nil
}

// Get the blog post with pagination
type BlogPostPagination struct {
	Posts        []BlogPost
	TotalPost    int
	PreviousPage int
	NextPage     int
}

func ListBlogPostsWithPagination(db *bun.DB, page int, limit int) (*BlogPostPagination, error) {
	if page < 1 {
		page = 1
	}

	ctx := context.Background()
	offset := (page - 1) * limit

	var posts []BlogPost

	err := db.NewSelect().
		Model(&posts).
		Relation("User").
		Order("time_stamp DESC").
		Limit(limit).
		Offset(offset).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	total, err := db.NewSelect().
		Model((*BlogPost)(nil)).
		Where("category = ?", "GENERAL").
		Count(ctx)
	if err != nil {
		return nil, err
	}

	var prevPage, nextPage int

	if page > 1 {
		prevPage = page - 1
	}

	if offset+len(posts) < total {
		nextPage = page + 1
	}

	return &BlogPostPagination{
		Posts:        posts,
		TotalPost:    total,
		PreviousPage: prevPage,
		NextPage:     nextPage,
	}, nil
}
