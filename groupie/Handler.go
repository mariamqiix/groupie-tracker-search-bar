package groupie

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"encoding/json"
)



func StartTheSite() {
	artists, TheLocations, TheDates, _, flag = UnmarshalData()
}

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/style.css" {
		http.ServeFile(w, r, "./template/style.css")
		return

	} else if r.URL.Path == "/home.html" || r.URL.Path == "/" {
		http.ServeFile(w, r, "./template/home.html")
		return

	} else if r.URL.Path == "/aboutus.html" {
		http.ServeFile(w, r, "./template/aboutus.html")
		return

	} else if r.URL.Path == "/index.html" {
		indexTemplate, err := template.ParseFiles("./template/index.html")

		if !flag || err != nil{
			w.WriteHeader(http.StatusInternalServerError)
			indexTemplate, _:= template.ParseFiles("./template/500.html")
			indexTemplate.Execute(w, "")
			return

		} 
			pageData := PageData{
				All: artists,
			}
	
			err = indexTemplate.Execute(w, pageData)
			if err != nil {
				fmt.Print(err)
				w.WriteHeader(http.StatusInternalServerError)
				http.ServeFile(w, r, "./template/500.html")
			}

		return

	} else if r.URL.Path == "/submit" {
		if !flag {
			w.WriteHeader(http.StatusInternalServerError)
			http.ServeFile(w, r, "./template/500.html")
			return
		} 

		valueStr := r.URL.Query().Get("value")
		value, _ := strconv.Atoi(valueStr)
		indexTemplate, err := template.ParseFiles("./template/artist.html")
		if err != nil {
			fmt.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			http.ServeFile(w, r, "./template/500.html")
			return
		}

		if valueStr == "" || value > 52 || value < 1 {
			w.WriteHeader(http.StatusBadRequest)
			http.ServeFile(w, r, "./template/400.html")
			return
		}

		pageDataArtice := PageDataArtice{
			All:                    artists[value-1],
			MergeDatesAndLocations: MergeDatesAndLocations(value - 1),
		}

		err = indexTemplate.Execute(w, pageDataArtice)
		if err != nil {
			fmt.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			http.ServeFile(w, r, "./template/500.html")
		}

		return
	} else if r.URL.Path == "/search" {

		if r.Method != "POST" {
			w.WriteHeader(http.StatusBadRequest)
			http.ServeFile(w, r, "./template/400.html")
			return
		}

		indexTemplate, err := template.ParseFiles("./template/index.html")
		search := r.FormValue("theArtist")
		if err != nil {
			fmt.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			http.ServeFile(w, r, "./template/500.html")
			return
		}
		artists, _ := Search(search)

		if  len(search) > 50 || !CheckLetter(search) {
			w.WriteHeader(http.StatusBadRequest)
			http.ServeFile(w, r, "./template/400.html")
			return
		} else if len(artists) == 0 {
			http.ServeFile(w, r, "./template/Noresult.html")
			return
		}

		pageData := PageData{
			All:       artists,
		}
		err = indexTemplate.Execute(w, pageData)
		if err != nil {
			fmt.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			http.ServeFile(w, r, "./template/500.html")
		}

		return

	} else if r.URL.Path == "/suggestions" {
		query := r.URL.Query().Get("query")
		suggested := suggestions(query)

		// Convert the suggestions slice to JSON
		suggestionsJSON, err := json.Marshal(suggested)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Set the appropriate content type header
		w.Header().Set("Content-Type", "application/json")

		// Write the JSON response
		w.Write(suggestionsJSON)
	} else {
		w.WriteHeader(http.StatusNotFound)
		http.ServeFile(w, r, "./template/404.html")
		return
	}

}

func CheckLetter(s string) bool {
	for g := 0; g < len(s); g++ {
		if s[g] > 126 || s[g] < 32 {
			return false
		}
	}
	return true
}