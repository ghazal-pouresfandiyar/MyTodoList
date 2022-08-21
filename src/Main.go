package src

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type Status int

const (
	Undone Status = iota
	Doing
	Done
)

type Todo struct {
	ID           string `json:"id"`
	Item         Item   `json:"item"`
	Status       string `json:"status"`
	DueDate      string `json:"due_date"`
	DurationTime int    `json:"duration_time"`
}

func makeTodoObj(ID string, Item Item, Status Status, DueDate string, DurationTime int) Todo {
	str := ""
	switch Status {
	case Done:
		str = "Done"
	case Doing:
		str = "Doing"
	case Undone:
		str = "Undone"
	default:
		str = "Undefined"
	}
	return Todo{ID, Item, str, DueDate, DurationTime}
}

type Item struct {
	Name    string
	SubItem []string
}

var todos []Todo

func getTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func getTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range todos {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Todo{})
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var todo Todo
	_ = json.NewDecoder(r.Body).Decode(&todo)
	todo.ID = strconv.Itoa(rand.Int()) //create random id and convert it to string (not secure)
	todos = append(todos, todo)
	json.NewEncoder(w).Encode(todo)
}

func updateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for idx, item := range todos {
		if item.ID == params["id"] {
			todos = append(todos[:idx], todos[idx+1:]...)
			var todo Todo
			_ = json.NewDecoder(r.Body).Decode(&todo)
			todo.ID = strconv.Itoa(rand.Int()) //create random id and convert it to string (not secure)
			todos = append(todos, todo)
			json.NewEncoder(w).Encode(todo)
			return
		}
	}
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range todos {
		if item.ID == params["id"] {
			fmt.Println("here")
			todos = append(todos[:index], todos[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(todos)
}

func RootEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("anything e.g hello :)"))
}

func main() {
	//Init Router
	router := mux.NewRouter()

	//CORS handler
	headers := handlers.AllowedHeaders([]string{"X-Requested-with", "content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})
	router.HandleFunc("/", RootEndpoint).Methods("GET")

	//our data
	then, _ := time.Parse("2 Jan 06 03:04PM", "28 Jul 22 11:59PM")
	item1 := Item{"Study English", []string{"Listening", "Reading", "Writing", "Speaking"}}
	todos = append(todos, makeTodoObj("1", item1, Done, then.Format(time.UnixDate), 5))
	item2 := Item{"Programming", []string{"Syntax", "Installing", "Coding"}}
	todos = append(todos, makeTodoObj("2", item2, Doing, then.Add(time.Hour*24).Format(time.UnixDate), 2))
	item3 := Item{"Cooking", []string{}}
	todos = append(todos, makeTodoObj("3", item3, Undone, then.Add(time.Hour*48).Format(time.UnixDate), 0))

	//Route Handlers / Endpoints
	router.HandleFunc("/api/todo", getTodos).Methods("GET")
	router.HandleFunc("/api/todo/{id}", getTodo).Methods("GET")
	router.HandleFunc("/api/todo", createTodo).Methods("POST")
	router.HandleFunc("/api/todo/{id}", updateTodo).Methods("PUT")
	router.HandleFunc("/api/todo/{id}", deleteTodo).Methods("DELETE")

	//start server
	http.ListenAndServe(":8000", handlers.CORS(headers, methods, origins)(router))
}
