package graphql

import (
    "server/database"
    "server/models"
    
    "github.com/graphql-go/graphql"
)

func GetUsersResolver(params graphql.ResolveParams) (interface{}, error) {
    var users []models.User
    result := database.GetDB().Find(&users)
    return users, result.Error
}

func GetUserResolver(params graphql.ResolveParams) (interface{}, error) {
    id := params.Args["id"].(int)
    var user models.User
    result := database.GetDB().First(&user, id)
    if result.Error != nil {
        return nil, result.Error
    }
    return user, nil
}

func CreateUserResolver(params graphql.ResolveParams) (interface{}, error) {
    input := params.Args["input"].(map[string]interface{})
    
    user := models.User{
        Name:  input["name"].(string),
        Email: input["email"].(string),
    }
    
    if age, ok := input["age"].(int); ok {
        user.Age = age
    }
    
    result := database.GetDB().Create(&user)
    return user, result.Error
}

func UpdateUserResolver(params graphql.ResolveParams) (interface{}, error) {
    id := params.Args["id"].(int)
    input := params.Args["input"].(map[string]interface{})
    
    var user models.User
    if err := database.GetDB().First(&user, id).Error; err != nil {
        return nil, err
    }
    
    user.Name = input["name"].(string)
    user.Email = input["email"].(string)
    if age, ok := input["age"].(int); ok {
        user.Age = age
    }
    
    result := database.GetDB().Save(&user)
    return user, result.Error
}

func DeleteUserResolver(params graphql.ResolveParams) (interface{}, error) {
    id := params.Args["id"].(int)
    
    result := database.GetDB().Delete(&models.User{}, id)
    if result.Error != nil {
        return false, result.Error
    }
    
    return result.RowsAffected > 0, nil
}
