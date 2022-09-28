package api

import (
	"fmt"
	"github.com/gorilla/mux"
	jsoniter "github.com/json-iterator/go"
	"net/http"
	database "url-shortener/internal/db"
	"url-shortener/internal/shortlinkgen"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func handleCreate(db database.DB, apiUrl string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := CreateRequest{}
		raw := make([]byte, 1024)
		r.Body.Read(raw)
		json.Unmarshal(raw, &req)
		short, err := shortlinkgen.GenerateShortLink(req.Url)
		if err != nil {
			return
		}
		err = db.AddLink(req.Url, short)
		if err != nil {
			fmt.Fprint(w, err)
			return
		}
		resp := CreateResponse{ShortUrl: apiUrl + "/" + short}
		res, err := json.Marshal(resp)
		if err != nil {
			fmt.Fprint(w, err)
			return
		}
		w.Write(res)
	}
}

func handleRedirect(db database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		url, err := db.GetInitialLink(id)
		if err != nil {
			return
		}
		http.Redirect(w, r, url, http.StatusSeeOther)
	}
}

func Init(r *mux.Router, db database.DB, apiUrl string) {
	r.HandleFunc("/create", handleCreate(db, apiUrl)).Methods(http.MethodPost)
	r.HandleFunc("/{id:.{10}}", handleRedirect(db)).Methods(http.MethodGet)
}
