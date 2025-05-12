package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/user/blog/database"
	"github.com/user/blog/models"
	"github.com/user/blog/repositories"
)

type PageData struct {
	Title       string
	Posts       []*models.Post
	Post        *models.Post
	Comments    []*models.Comment
	CurrentYear int
}

func main() {
	// Initialize database
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	// Create repositories
	postRepo := repositories.NewSQLitePostRepository(db)
	commentRepo := repositories.NewSQLiteCommentRepository(db)

	// Set up HTTP routes
	// Serve static files
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Route handlers
	http.HandleFunc("/", handleHome(postRepo))
	http.HandleFunc("/post/", handlePostDetail(postRepo, commentRepo))
	http.HandleFunc("/comment", handleAddComment(commentRepo))

	// Start the server
	fmt.Println("Starting server at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleHome(postRepo repositories.PostRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		posts, err := postRepo.GetAll()
		if err != nil {
			http.Error(w, "Failed to fetch posts", http.StatusInternalServerError)
			return
		}

		tmpl, err := template.ParseFiles("templates/home.html", "templates/layout.html")
		if err != nil {
			http.Error(w, "Failed to load template", http.StatusInternalServerError)
			return
		}

		data := PageData{
			Title:       "My Blog",
			Posts:       posts,
			CurrentYear: time.Now().Year(),
		}

		err = tmpl.ExecuteTemplate(w, "layout", data)
		if err != nil {
			http.Error(w, "Failed to render template", http.StatusInternalServerError)
		}
	}
}

func handlePostDetail(postRepo repositories.PostRepository, commentRepo repositories.CommentRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.URL.Path[len("/post/"):]
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			http.Error(w, "Invalid post ID", http.StatusBadRequest)
			return
		}

		post, err := postRepo.GetByID(id)
		if err != nil {
			http.Error(w, "Post not found", http.StatusNotFound)
			return
		}

		comments, err := commentRepo.GetByPostID(id)
		if err != nil {
			http.Error(w, "Failed to load comments", http.StatusInternalServerError)
			return
		}

		tmpl, err := template.ParseFiles("templates/post.html", "templates/layout.html")
		if err != nil {
			http.Error(w, "Failed to load template", http.StatusInternalServerError)
			return
		}

		data := PageData{
			Title:       post.Title,
			Post:        post,
			Comments:    comments,
			CurrentYear: time.Now().Year(),
		}

		err = tmpl.ExecuteTemplate(w, "layout", data)
		if err != nil {
			http.Error(w, "Failed to render template", http.StatusInternalServerError)
		}
	}
}

func handleAddComment(commentRepo repositories.CommentRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Failed to parse form", http.StatusBadRequest)
			return
		}

		postIDStr := r.FormValue("post_id")
		postID, err := strconv.ParseInt(postIDStr, 10, 64)
		if err != nil {
			http.Error(w, "Invalid post ID", http.StatusBadRequest)
			return
		}

		commenterName := r.FormValue("commenter_name")
		content := r.FormValue("content")

		if commenterName == "" || content == "" {
			http.Error(w, "Name and comment are required", http.StatusBadRequest)
			return
		}

		comment := &models.Comment{
			PostID:        postID,
			CommenterName: commenterName,
			Content:       content,
		}

		err = commentRepo.Create(comment)
		if err != nil {
			http.Error(w, "Failed to add comment", http.StatusInternalServerError)
			return
		}

		// Redirect back to the post
		http.Redirect(w, r, "/post/"+postIDStr, http.StatusSeeOther)
	}
}


