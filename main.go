package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Data struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Gmail       string `json:"gmail"`
	PhoneNumber int    `json:"phoneNumber"`
}

var datas = []Data{
	{ID: "1", Name: "Top", Gmail: "dragontops55", PhoneNumber: 54778644},
	{ID: "2", Name: "Mina", Gmail: "minar120304", PhoneNumber: 56266920},
}

func main() {
	fmt.Printf("http://localhost:8080/person \n\n")
	r := gin.Default()
	r.GET("/person", getDatas)
	r.POST("/person", postDatas)
	r.Run(":8080")
}

func getDatas(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, datas)
}

func postDatas(c *gin.Context) {
	var newDatas Data
	datas = append(datas, newDatas)
	c.IndentedJSON(http.StatusCreated, newDatas)
}