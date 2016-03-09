package user

import (
	"fmt"
	"net/http"

	"github.com/zenazn/goji/web"
)

func Index_get_handler(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", "User Page")
}
