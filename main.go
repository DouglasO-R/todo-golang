package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ITodoDatabase interface {
	RetrieveTodoList () []Todo
}

type InMemoryTodoDatabase struct{

}

func (inMemoryTodoDatabase *InMemoryTodoDatabase) RetrieveTodoList() []Todo {
	return  []Todo{
		{UUID: uuid.New(),Title: "test"},
		{UUID: uuid.New(),Title: "test2"},
		{UUID: uuid.New(),Title: "test3"},
	}
}

func NewInMemoryTodoDatabase() InMemoryTodoDatabase {
	return InMemoryTodoDatabase{}
}

type Todo struct{
	UUID uuid.UUID `json:"id"`
	Title string `json:"title"`
}

type TodosResponse struct {
	Items []Todo `json:"items"`
}

func main() {
	r := gin.Default()

	r.GET("/todos", func(c *gin.Context) {
		db := NewInMemoryTodoDatabase();
		response := TodosResponse{
			Items: db.RetrieveTodoList(),
		} // Sua lista de itens aqui


		c.JSON(http.StatusOK,response)
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

