package main

import (
	_ "fmt"

	"github.com/gin-gonic/gin"
)

// delete user
func DeleteUserHandler(id uint) gin.HandlerFunc {
	return func(c *gin.Context) {
		result:=DB.Model(&User{}).Where("id=?", id).Update("status", false)
		if result.Error!=nil{

		}
		if result.RowsAffected==0{

		}

		// deleteUserCache(c.Request.Context(), userID uint)
	}
}
