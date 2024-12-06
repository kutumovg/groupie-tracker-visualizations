package server

import (
	"fmt"
	"html/template"
	"net/http"
	"path"

	"groupie-tracker-visualizations/parsers"
)

type ErrorPage struct {
	StatusCode string
	StatusText string
}

func Server() {
	fmt.Println("Server started on address http://localhost:4000")
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/artist/", ArtistPage)
	http.Handle("/ui/", http.StripPrefix("/ui/", http.FileServer(http.Dir("./ui"))))

	err := http.ListenAndServe(":4000", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func ErrorHandler(w http.ResponseWriter, r *http.Request, statusCode int, statusText string) {
	data := ErrorPage{
		StatusCode: fmt.Sprint(statusCode),
		StatusText: statusText,
	}

	ts, err := template.ParseFiles("./templates/wentwrong.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		return
	}

	// w.WriteHeader(statusCode)
	err = ts.Execute(w, data)
	if err != nil {
		http.Error(w, "Error when executing", http.StatusInternalServerError)
		return
	}
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	if r.URL.Path != "/" {
		ErrorHandler(w, r, http.StatusNotFound, http.StatusText(http.StatusNotFound))
		return
	}
	if r.Method != http.MethodGet {
		ErrorHandler(w, r, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
		return
	}

	artists := parsers.GetArtists()
	data := struct {
		Artists []parsers.Artists
	}{
		Artists: artists,
	}

	err = ts.Execute(w, data)
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
}

func ArtistPage(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("./templates/artist.html")
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	if r.URL.Path != "/artist/"+path.Base(r.URL.Path) {
		ErrorHandler(w, r, http.StatusNotFound, http.StatusText(http.StatusNotFound))
		return
	}
	if r.Method != http.MethodGet {
		ErrorHandler(w, r, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
		return
	}

	id := path.Base(r.URL.Path)
	if !parsers.CheckId(id) {
		ErrorHandler(w, r, http.StatusNotFound, http.StatusText(http.StatusNotFound))
		return
	}
	artist := parsers.GetArtist(id)
	relations := parsers.GetRelations(id)

	artist.Relations = relations.DatesLocations

	err = ts.Execute(w, artist)
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
}
