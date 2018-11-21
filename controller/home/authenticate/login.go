package authenticate

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

func Login(w http.ResponseWriter, r *http.Request) {
	useremail := r.FormValue("useremail")
	// password := r.FormValue("password")

	session, _ := store.Get(r, "user-cookie")

	// Authentication goes here
	// Get user data from database

	// Set user as authenticated
	session.Values["user-name"] = "Hans-San"
	session.Values["user-email"] = useremail
	session.Values["authenticated"] = true
	session.Save(r, w)
	return
}

func Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "user-cookie")

	// Revoke users authentication
	session.Values["authenticated"] = false
	session.Save(r, w)
}

func CheckLogin(w http.ResponseWriter, r *http.Request) error {
	session, _ := store.Get(r, "user-cookie")

	// Check if user is authenticated
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		return fmt.Errorf("Forbidden", http.StatusForbidden)
	}

	// Print secret message
	return nil
}

func GetUserData(w http.ResponseWriter, r *http.Request) (map[string]interface{}, error) {
	session, _ := store.Get(r, "user-cookie")
	param := map[string]interface{}{
		"Authenticated": session.Values["authenticated"],
	}

	// Check if user is authenticated
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		return map[string]interface{}{}, nil
	}

	// Set parameter value
	param["UserName"] = session.Values["user-name"]

	return param, nil
}
