package services

import (
	"net/http"

	"github.com/ugorji/go/codec"
	"github.com/zenazn/goji/web"
)

type api struct {
}

func Api() *api {
	return &api{}
}

func (api *api) GenerateResponse(c web.C, r *http.Request, w http.ResponseWriter, data interface{}) {
	if r.Header.Get("Accept") == "application/x-msgpack" {
		w.Header().Set("Content-Type", "application/x-msgpack")
		h := new(codec.MsgpackHandle)
		encoder := codec.NewEncoder(w, h)
		encoder.Encode(data)
	} else {
		w.Header().Set("Content-Type", "application/json")
		h := new(codec.JsonHandle)
		encoder := codec.NewEncoder(w, h)
		encoder.Encode(data)
	}

}
