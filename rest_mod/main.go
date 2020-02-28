package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"rest_mod/model"
	"strings"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "<h1>Hello</h1>")
}

type contact struct {
	Name  string `json: "name"`
	Phone string `json: "phone"`
}

type server struct{}

func getName(req *http.Request) string {
	path := strings.Split(req.URL.Path, "/")
	return path[len(path)-1]
}

func (s *server) serveGET(w http.ResponseWriter, req *http.Request) {
	name := getName(req)
	if name == "" {
		contacts := model.ReadAll()
		res := make([]contact, len(contacts))
		for i, c := range contacts {
			res[i] = contact{c[0], c[1]}
		}

		json_res, err := json.Marshal(res)
		if err != nil {
			log.Fatal(err)
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, string(json_res))
	} else {
		phone := model.Read(name)
		if phone != "" {
			json_res, err := json.Marshal(&contact{name, phone})
			if err != nil {
				log.Fatal(err)
			}

			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, string(json_res))
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}
}

func (s *server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch req.Method {
	case "GET":
		s.serveGET(w, req)
	case "POST":
		var c contact
		json.NewDecoder(req.Body).Decode(&c)
		id := model.Create(c.Name, c.Phone)
		w.WriteHeader(http.StatusCreated)
		fmt.Println(c.Name, c.Phone, id)
		fmt.Fprintf(w, `{"id": "%d"}`, id)

	case "PUT":
		name := getName(req)
		if name == "" {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			model.Delete(name)
			w.WriteHeader(http.StatusNoContent)
			var c contact
			json.NewDecoder(req.Body).Decode(&c)
			model.Update(name, c.Name, c.Phone)
			json_res, err := json.Marshal(&contact{c.Name, c.Phone})
			if err != nil {
				log.Fatal(err)
			}

			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, string(json_res))
		}

	case "DELETE":
		name := getName(req)
		if name == "" {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			model.Delete(name)
			w.WriteHeader(http.StatusNoContent)
		}
		fmt.Fprintf(w, `{"message": "Delete"}`)

	}
}

func main() {

	defer model.CloseDB()
	http.Handle("/phones/", &server{})
	http.ListenAndServe(":8000", nil)

}
