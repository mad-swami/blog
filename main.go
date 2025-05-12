package main

import (
	"fmt"
	"log"
	
	"github.com/user/blog/database"
	"github.com/user/blog/models"
	"github.com/user/blog/repositories"
)

func main() {
	fmt.Println("Setting up the blog database...")
	
	// setup
	db, err := database.SetupDB()
	if err != nil {
		log.Fatalf("Failed to setup database: %v", err)
	}
	defer db.Close()
	
	adminRepo := repositories.NewSQLiteAdminRepository(db)
	postRepo := repositories.NewSQLitePostRepository(db)
	commentRepo := repositories.NewSQLiteCommentRepository(db)
	imageRepo := repositories.NewSQLiteImageRepository(db)
	
	// make admin
	admin := &models.Admin{
		Username:     "admin",
		PasswordHash: "hashed_password", // In real app, this would be properly hashed
		DisplayName:  "Site Admin",
	}
	
	if err := adminRepo.Create(admin); err != nil {
		log.Fatalf("Failed to create admin: %v", err)
	}
	fmt.Printf("Created admin with ID: %d\n", admin.ID)
	
	// make mock post
	post := &models.Post{
		Title:   "My First Blog Post",
		Content: "Hello databases class. This is my first blog post. I hope you all enjoy my presentation!",
	}
	
	if err := postRepo.Create(post); err != nil {
		log.Fatalf("Failed to create post: %v", err)
	}
	fmt.Printf("Created post with ID: %d\n", post.ID)
	
	// add comment to post
	comment := &models.Comment{
		PostID:        post.ID,
		CommenterName: "John Doe",
		Content:       "Awesome blog bro. Keep up the cool posts",
	}
	
	if err := commentRepo.Create(comment); err != nil {
		log.Fatalf("Failed to create comment: %v", err)
	}
	fmt.Printf("Created comment with ID: %d\n", comment.ID)
	
	// add image to post
	image := &models.Image{
		PostID:   post.ID,
		Filename: "example.jpg",
		FilePath: "/images/example.jpg",
	}
	
	if err := imageRepo.Create(image); err != nil {
		log.Fatalf("Failed to create image: %v", err)
	}
	fmt.Printf("Created image with ID: %d\n", image.ID)
	
	// get post with comments
	comments, err := commentRepo.GetByPostID(post.ID)
	if err != nil {
		log.Fatalf("Failed to retrieve comments: %v", err)
	}
	
	fmt.Printf("\nPost: %s\n", post.Title)
	fmt.Printf("Content: %s\n", post.Content)
	fmt.Printf("Comments (%d):\n", len(comments))
	
	for _, c := range comments {
		fmt.Printf("- %s: %s\n", c.CommenterName, c.Content)
	}
	
	// get images for post
	images, err := imageRepo.GetByPostID(post.ID)
	if err != nil {
		log.Fatalf("Failed to retrieve images: %v", err)
	}
	
	fmt.Printf("Images (%d):\n", len(images))
	for _, img := range images {
		fmt.Printf("- %s (%s)\n", img.Filename, img.FilePath)
	}
	
	fmt.Println("\nBlog database demo completed successfully!")
}
