package controllers

import (
	"net/http"
	"web-server/models/member"
	memberService "web-server/services/member"

	"github.com/gin-gonic/gin"
)

func GetMembers(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"message": "Get all users",
	})
}

func CreateMember(c *gin.Context){
	var createMemberRequest member.CreateMemberRequest
	
	if err := c.ShouldBindBodyWithJSON(&createMemberRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}

	member := memberService.CreateMember(&createMemberRequest)

	c.JSON(http.StatusOK, member)
}