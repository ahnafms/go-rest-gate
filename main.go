package main

import (
	"github.com/gin-gonic/gin"
  Gate "github.com/ahnafms/go-rest-gate/controllers"
  DB "github.com/ahnafms/go-rest-gate/models"
)

func main(){
  DB.ConnectDB()
  router := gin.Default()
  gate := router.Group("/api")
  { 
    gate.POST("/in", Gate.InGate) 
    gate.POST("/out", Gate.OutGate) 
  } 
  router.Run(":8080") }

