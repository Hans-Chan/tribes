package home

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/tribes/controller/home/authenticate"
)

const page = "Home"

func RequestHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Step 1
	msg := fmt.Sprintf("%s: Step 1 - Render Page", page)
	err := RenderHomePage(w, r)
	if err != nil {
		fmt.Println(fmt.Sprintf("%s Error: %s"), msg, err)
		// RenderErrorPage(w, form, err)
	}
}

func RequestLoginHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Variable declaration
	page := "tribe"

	// Logging in
	authenticate.Login(w, r)

	// Check if authenticated
	errLogin := authenticate.CheckLogin(w, r)
	if errLogin != nil {
		// RenderErrorPage(w, form, err)
	}

	// Step 2
	msg := fmt.Sprintf("%s: Step 2 - Login", page)
	err := RenderHomePage(w, r)
	if err != nil {
		fmt.Println(fmt.Sprintf("%s Error: %s"), msg, err)
		// RenderErrorPage(w, form, err)
	}
}
