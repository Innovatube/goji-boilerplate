package admin

import (
	"fmt"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/zenazn/goji/web"
)

func Index_get_handler(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", "Admin Page")
	log.Debug("Request admin controller")
}
