package todoById

import (
	"app/db/todo"

	"github.com/gin-gonic/gin"
)

func GetTodoById(c *gin.Context) {
	id := c.Param("id")
	todo, err := todo.FetchTodo(id)
	if(err != nil){
		c.JSON(400, gin.H{
			"error": err,
		})
	}
	c.JSON(200, gin.H{
		"todo": todo,
	})
}