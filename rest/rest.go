package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"encoding/json"
  
	"github.com/julienschmidt/httprouter"
)

// Todoのstruct
type Todo struct {
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}
  
type Todos []Todo

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Welcmoe!")
}
  
func TodoIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	todos := Todos{
		Todo{Name: "Write presentation"},
		Todo{Name: "Host meetup"},
	}
  
	json.NewEncoder(w).Encode(todos)
}
  
func TodoShow(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "Todo show: %s", ps.ByName("todoId"))
}
  
func main() {
	// routerの作成
	router := httprouter.New()

	// GEtメソッドの受付
	router.GET("/", Index)
	router.GET("/todos", TodoIndex)
	router.GET("/todos/:todoId", TodoShow)
  
	log.Fatal(http.ListenAndServe(":8080", router))
}

// http://sgykfjsm.github.io/blog/2016/03/13/golang-json-api-tutorial/