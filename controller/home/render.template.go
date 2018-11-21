package home

import (
	"net/http"

	"github.com/tribes/controller/home/authenticate"
	"github.com/tribes/view/render"
)

func RenderHomePage(w http.ResponseWriter, r *http.Request) error {
	// Get user details from cookie
	templateParameter, err := authenticate.GetUserData(w, r)

	if err != nil {
		err.Error()
	}

	render.RenderPage(w, "home", templateParameter)
	return nil
}
