package routes

import (
	home_controller "github.com/takaaki-mizuno/goji-boilerplate/app/http/controllers/home"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
)

func HomeRoutes() {
	homeMux := web.New()
	goji.Handle("/*", homeMux)
	homeMux.Get("/", home_controller.Index_get_handler)
}
