package handlers

import (
    "net/http"

    "server/database"
    "server/models"

    "github.com/gin-gonic/gin"
)

// GetUsers - GET /api/users
func GetUsers(c *gin.Context) {
    var users []models.User
    result := database.GetDB().Find(&users)
    
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }
    
    c.JSON(http.StatusOK, gin.H{"data": users})
}

// GetUser - GET /api/users/:id
func GetUser(c *gin.Context) {
    id := c.Param("id")
    var user models.User
    
    result := database.GetDB().First(&user, id)
    if result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }
    
    c.JSON(http.StatusOK, gin.H{"data": user})
}

// CreateUser - POST /api/users
func CreateUser(c *gin.Context) {
    var user models.User
    
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    result := database.GetDB().Create(&user)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }
    
    c.JSON(http.StatusCreated, gin.H{"data": user})
}

// UpdateUser - PUT /api/users/:id
func UpdateUser(c *gin.Context) {
    id := c.Param("id")
    var user models.User
    
    if err := database.GetDB().First(&user, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }
    
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    database.GetDB().Save(&user)
    c.JSON(http.StatusOK, gin.H{"data": user})
}

// DeleteUser - DELETE /api/users/:id
func DeleteUser(c *gin.Context) {
    id := c.Param("id")
    
    result := database.GetDB().Delete(&models.User{}, id)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }
    
    if result.RowsAffected == 0 {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }
    
    c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
