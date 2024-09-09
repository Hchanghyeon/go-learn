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

	c.ShouldBindBodyWithJSON(createMemberRequest)
	member := memberService.CreateMember(createMemberRequest)

	c.JSON(http.StatusOK, member)
}