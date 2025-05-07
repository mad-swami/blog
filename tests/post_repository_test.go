package tests

import (
	"os"
	"testing"
	
	"github.com/user/blog/database"
	"github.com/user/blog/models"
	"github.com/user/blog/repositories"
)

func TestPostRepository(t *testing.T) {
	// Setup test database
	os.Remove("../test_blog.db")
	db, err := database.InitDB()
	if err != nil {
		t.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()
	
	// Execute schema
	schemaSQL, err := os.ReadFile("../database/schema.sql")
	if err != nil {
		t.Fatalf("Failed to read schema file: %v", err)
	}
	
	_, err = db.Exec(string(schemaSQL))
	if err != nil {
		t.Fatalf("Failed to execute schema: %v", err)
	}
	
	// Create repository
	repo := repositories.NewSQLitePostRepository(db)
	
	// Test post creation
	post := &models.Post{
		Title:   "Test Post",
		Content: "This is a test post content",
	}
	
	err = repo.Create(post)
	if err != nil {
		t.Fatalf("Failed to create post: %v", err)
	}
	
	if post.ID == 0 {
		t.Fatal("Expected post ID to be non-zero after creation")
	}
	
	// Test get post by ID
	fetchedPost, err := repo.GetByID(post.ID)
	if err != nil {
		t.Fatalf("Failed to get post by ID: %v", err)
	}
	
	if fetchedPost.Title != post.Title {
		t.Errorf("Expected title %s, got %s", post.Title, fetchedPost.Title)
	}
	
	if fetchedPost.Content != post.Content {
		t.Errorf("Expected content %s, got %s", post.Content, fetchedPost.Content)
	}
	
	// Test update post
	fetchedPost.Title = "Updated Title"
	fetchedPost.Content = "Updated content for test post"
	
	err = repo.Update(fetchedPost)
	if err != nil {
		t.Fatalf("Failed to update post: %v", err)
	}
	
	updatedPost, err := repo.GetByID(fetchedPost.ID)
	if err != nil {
		t.Fatalf("Failed to get updated post: %v", err)
	}
	
	if updatedPost.Title != "Updated Title" {
		t.Errorf("Expected updated title to be 'Updated Title', got %s", updatedPost.Title)
	}
	
	// Test get all posts
	posts, err := repo.GetAll()
	if err != nil {
		t.Fatalf("Failed to get all posts: %v", err)
	}
	
	if len(posts) != 1 {
		t.Errorf("Expected 1 post, got %d", len(posts))
	}
	
	// Test delete post
	err = repo.Delete(post.ID)
	if err != nil {
		t.Fatalf("Failed to delete post: %v", err)
	}
	
	posts, err = repo.GetAll()
	if err != nil {
		t.Fatalf("Failed to get all posts after deletion: %v", err)
	}
	
	if len(posts) != 0 {
		t.Errorf("Expected 0 posts after deletion, got %d", len(posts))
	}
}
