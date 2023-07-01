package main

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"todo-go/database/models"
	database "todo-go/database"
)

func getTodos(ctx *gin.Context) {
	var todos []models.Todo;
	if err := database.DB.Find(&todos).Error; err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message" : "Aucune todo trouvée.", " Erreur : " : err.Error()});
		return
	}
	ctx.IndentedJSON(http.StatusOK, todos);
}

func sayHello(ctx *gin.Context)  {
	ctx.IndentedJSON(http.StatusOK, gin.H{"message" : "Hello"});
}

func addTodo(ctx *gin.Context) {
	var newTodo models.Todo;

	if err := ctx.BindJSON(&newTodo); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message" : "Données invalides"});
		return
 	}

 	if err := database.DB.Create(&newTodo).Error; err != nil {
 		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message" : "Données invalides"});
 		return
 	}

 	ctx.IndentedJSON(http.StatusOK, newTodo);
 }

 func getTodo(ctx *gin.Context) {
 	id, err := strconv.Atoi(ctx.Param("id"));
 	if err != nil {
 		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message" : "Todo introuvable"})
 		return
 	}

 	var todo models.Todo

 	if err := database.DB.First(&todo, id).Error; err != nil {
 		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo introuvable"})
 		return
 	}

 	ctx.IndentedJSON(http.StatusOK, todo)
 }


 func toggleTodoCompleted(ctx *gin.Context) {
 	id, err := strconv.Atoi(ctx.Param("id"))
 	if err != nil {
 		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo introuvable"})
 		return
 	}

 	var todo models.Todo
 	if err := database.DB.First(&todo, id).Error; err != nil {
 		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo introuvable"})
 		return
 	}

 	todo.Completed = !todo.Completed

 	if err := database.DB.Save(&todo).Error; err != nil {
 		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la mise à jour du todo"})
 		return
 	}

 	ctx.IndentedJSON(http.StatusOK, todo)
 }

 func deleteTodoById(ctx *gin.Context) {
 	id, err := strconv.Atoi(ctx.Param("id"))
 	if err != nil {
 		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo introuvable"})
 		return
 	}

 	if err := database.DB.Delete(&models.Todo{}, id).Error; err != nil {
 		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la suppression du todo"})
 		return
 	}

 	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Todo supprimée"})
 }


const endPoint = "/todos";

func main() {
	database.ConnectToDb();
	router := gin.Default();
	router.GET(endPoint, getTodos);
	router.GET("/hello", sayHello);

	router.GET(endPoint + "/:id", getTodo);
	router.PATCH(endPoint + "/:id", toggleTodoCompleted);
	router.POST(endPoint, addTodo);
	router.DELETE(endPoint + "/:id", deleteTodoById);
	router.Run(":8000");
}
