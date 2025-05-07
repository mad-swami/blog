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
	
	// Setup the database
	db, err := database.SetupDB()
	if err != nil {
		log.Fatalf("Failed to setup database: %v", err)
	}
	defer db.Close()
	
	// Create repositories
	adminRepo := repositories.NewSQLiteAdminRepository(db)
	postRepo := repositories.NewSQLitePostRepository(db)
	commentRepo := repositories.NewSQLiteCommentRepository(db)
	imageRepo := repositories.NewSQLiteImageRepository(db)
	
	// Create an admin
	admin := &models.Admin{
		Username:     "admin",
		PasswordHash: "hashed_password", // In a real app, this would be properly hashed
		DisplayName:  "Site Admin",
	}
	
	if err := adminRepo.Create(admin); err != nil {
		log.Fatalf("Failed to create admin: %v", err)
	}
	fmt.Printf("Created admin with ID: %d\n", admin.ID)
	
	// Create a post
	post := &models.Post{
		Title:   "My First Blog Post",
		Content: "This is the content of my first blog post. Welcome to my blog!",
	}
	
	if err := postRepo.Create(post); err != nil {
		log.Fatalf("Failed to create post: %v", err)
	}
	fmt.Printf("Created post with ID: %d\n", post.ID)
	
	// Add a comment to the post
	comment := &models.Comment{
		PostID:        post.ID,
		CommenterName: "John Doe",
		Content:       "Great first post! Looking forward to more.",
	}
	
	if err := commentRepo.Create(comment); err != nil {
		log.Fatalf("Failed to create comment: %v", err)
	}
	fmt.Printf("Created comment with ID: %d\n", comment.ID)
	
	// Add an image to the post
	image := &models.Image{
		PostID:   post.ID,
		Filename: "example.jpg",
		FilePath: "/images/example.jpg",
	}
	
	if err := imageRepo.Create(image); err != nil {
		log.Fatalf("Failed to create image: %v", err)
	}
	fmt.Printf("Created image with ID: %d\n", image.ID)
	
	// Retrieve post with comments
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
	
	// Retrieve images for the post
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
