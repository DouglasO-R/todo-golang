package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var DB ITodoDatabase

type ITodoDatabase interface {
	RetrieveTodoList () []Todo
	CreateTodo (todo Todo) 
}

type InMemoryTodoDatabase struct{
	todoList map[uuid.UUID]Todo

}

func NewInMemoryTodoDatabase() InMemoryTodoDatabase {
	result := InMemoryTodoDatabase{}
	result.todoList = make(map[uuid.UUID]Todo)
	return result
}


func (inMemoryTodoDatabase *InMemoryTodoDatabase) RetrieveTodoList() []Todo {
	todos := make([]Todo, 0)
	for _,todo := range inMemoryTodoDatabase.todoList{
		todos = append(todos, todo)
	}

	return todos;
}

func (inMemoryTodoDatabase *InMemoryTodoDatabase) CreateTodo(todo Todo){
	uniqueId := todo.UUID
	inMemoryTodoDatabase.todoList[uniqueId]=todo
}	

type Todo struct{
	UUID uuid.UUID `json:"id"`
	Title string `json:"title"`
}

type TodosResponse struct {
	Items []Todo `json:"items"`
}

func main() {
	router := gin.Default()
	DB := NewInMemoryTodoDatabase();

	router.GET("/todos", func(response *gin.Context) {
		res := TodosResponse{
			Items: DB.RetrieveTodoList(),
		} // Sua lista de itens aqui


		response.JSON(http.StatusOK,res)
	})

	router.POST("/todos", func(response *gin.Context) {
		todo := Todo{UUID: uuid.New(),Title: "test"}
		DB.CreateTodo(todo)

		response.JSON(http.StatusOK,gin.H{})
	})


	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

