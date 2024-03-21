package controllers

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/Mathews9396/go-userMngmnt/pkg/models"
	"github.com/gin-gonic/gin"
)

func GetAllUsers(w *gin.Context){
	allUsers := models.GetAllUsers()
	fmt.Printf("%v",len(allUsers))
	if len(allUsers) == 0 {
		w.JSON(http.StatusOK, gin.H{"message": "No users in the database"})
		return
	}
	w.JSON(http.StatusOK, gin.H{"users": allUsers, "message": "Users retrieved successfully"})
}

func CreateUser(w *gin.Context){
	var user = &models.User{}

    // Bind the JSON request body to the CreateUserRequest struct
    if err := w.ShouldBindJSON(user); err != nil {
		fmt.Printf("Error binding JSON request body: %v\n", err)
        // Respond with an error message
        w.JSON(http.StatusBadRequest, gin.H{"error": "Error binding JSON request body"})
        return
    }

	// Print the user object to the console
    fmt.Printf("User in the request body: %+v\n", user)

	if user.Firstname == "" || user.Lastname == "" || user.Email == ""{
        w.JSON(http.StatusBadRequest, gin.H{"error": "User data is missing values"})
		return

	}
	user.ID = strconv.Itoa(rand.Intn(100000000))

	userCreated := models.CreateUser(user)
    // Respond with a success message
	w.JSON(http.StatusOK, gin.H{"newUser": userCreated, "message": "User created successfully"})
}

func GetUser(w *gin.Context){
	id := w.Param("userId")
	if id == "" {
		w.JSON(http.StatusBadRequest, gin.H{"error": "ID not provided"})
		return
	}
	userFound := models.GetUser(id)
	if userFound == nil {
		w.JSON(http.StatusBadRequest, gin.H{"message": "User not found for the given id"})
		return
	}
	w.JSON(http.StatusOK, gin.H{"user": userFound, "message": "User found successfully"})
}

func UpdateUser(w *gin.Context){
	var user = &models.User{}
	id := w.Param("userId")
	if id == "" {
		w.JSON(http.StatusBadRequest, gin.H{"error": "ID not provided"})
		return
	}
	userFound := models.GetUser(id)
	if userFound == nil {
		w.JSON(http.StatusBadRequest, gin.H{"message": "User not found for the given id"})
		return
	}
    // Bind the JSON request body to the CreateUserRequest struct
    if err := w.ShouldBindJSON(user); err != nil {
		fmt.Printf("Error binding JSON request body: %v\n", err)
        // Respond with an error message
        w.JSON(http.StatusBadRequest, gin.H{"error": "Error binding JSON request body"})
        return
    }
	user.ID = id
	// Print the user object to the console
    fmt.Printf("User in the request body: %+v\n", user)

	if user.Firstname == "" && user.Lastname == "" && user.Email == ""{
        w.JSON(http.StatusBadRequest, gin.H{"error": "User data is missing values"})
		return
	}
	

	userUpdated := models.UpdateUser(user)
	if userUpdated == nil {
		w.JSON(http.StatusInternalServerError, gin.H{"error": "User data updation failed"})
		return
	}
    // Respond with a success message
	w.JSON(http.StatusOK, gin.H{"newUser": userUpdated, "message": "User updated successfully"})
}

func DeleteUser(w *gin.Context){
	id := w.Param("userId")
	if id == "" {
		w.JSON(http.StatusBadRequest, gin.H{"error": "ID not provided"})
		return
	}
	userFound := models.GetUser(id)
	if userFound == nil {
		w.JSON(http.StatusBadRequest, gin.H{"message": "User not found for the given id"})
		return
	}
	userDeleted := models.DeleteUser(id)
	if !userDeleted {
		w.JSON(http.StatusInternalServerError, gin.H{"error": "User deletion failed"})
		return
	}
    // Respond with a success message
	w.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}