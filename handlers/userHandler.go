package handlers

import (
	"html/template"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
	"github.com/gofrs/uuid"
)

type User struct {
	
	// Gio - I don't we need ID within the Golang User struct.
	// The sqlite3.db already has this field and it is set to auto-increment
	// Delete this once confirmed by the team.
	UserID   int
	Username string
	Email    string
	Password string
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Parse form data
		username := r.FormValue("username")
		email := r.FormValue("useremail")
		password := r.FormValue("userpassword")
		// Check if the email is already taken (you need to implement this)
		if IsEmailTaken(email) {
			// Handle the case where the email is already taken
			http.Error(w, "Email is already registered", http.StatusConflict)
			return
		}
		// Hash the password using bcrypt
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			// Handle error
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		// Generate a UUID for the user
		userID := uuid.NewV4().String()
		// Store user information in the database (you need to implement this)
		if err := CreateUser(userID, username, email, string(hashedPassword)); err != nil {
			// Handle error
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		// Redirect the user to the login page or home page
		http.Redirect(w, r, "/index", http.StatusSeeOther)
		return
	}
	// Render the registration form template (you need to create this template)
	tmpl, err := template.ParseFiles("templates/register.html")
	if err != nil {
		// Handle error
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}
/*HTTP handler to process the registration form submission. This handler should:
Parse the form data submitted by the user.
Check if the email is already taken.
Hash the password using bcrypt.
Store the user's information securely in the database.
Generate a UUID for the user.*/
/*In this code, you'll need to implement
IsEmailTaken to check if an email is already registered and
CreateUser to store user information in the database.*/

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Parse form data
		email := r.FormValue("email")
		password := r.FormValue("password")
		// Retrieve the user's information from the database by email (you need to implement this)
		user, err := GetUserByEmail(email)
		if err != nil {
			// Handle error
			http.Error(w, "Invalid email or password", http.StatusUnauthorized)
			return
		}
		// Compare the hashed password with the provided password
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
			// Handle error
			http.Error(w, "Invalid email or password", http.StatusUnauthorized)
			return
		}
		// Create and set a session cookie (you need to implement this)
		sessionID := uuid.NewV4().String()
		SetSessionCookie(w, sessionID)
		// Redirect the user to the home page or their profile page
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	// Render the login form template (you need to create this template)
	tmpl, err := template.ParseFiles("templates/login.html")
	if err != nil {
		// Handle error
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}
/*
HTTP handler to process the login form submission. This handler should:
Parse the form data submitted by the user.
Verify the user's credentials.
Create and set a session cookie upon successful login.
*/
/*you'll need to implement
GetUserByEmail to retrieve user information from the database and
SetSessionCookie to create and set a session cookie.
*/