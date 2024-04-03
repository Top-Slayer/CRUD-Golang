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
	fmt.Printf("http://localhost:8080/ \n\n")
	r := setupRouter()
	r.Run(":8080")
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", getDatas)
	r.GET("/:id", getDataById)
	r.POST("/", postDatas)
	r.DELETE("/:id", deleteDataById)

	return r
}

func getDatas(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, datas)
}

func getDataById(c *gin.Context) {
	id := c.Param("id")
	// fmt.Fprintf(c.Writer, `<script>alert('%s');</script>`, id)
	for _, e := range datas {
		if e.ID == id {
			c.IndentedJSON(http.StatusOK, e)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "datas not found"})
}

func postDatas(c *gin.Context) {
	var newDatas Data
	err := c.BindJSON(&newDatas)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	datas = append(datas, newDatas)
	c.IndentedJSON(http.StatusCreated, newDatas)
}

func deleteDataById(c *gin.Context) {
	id := c.Param("id")
  for i, e := range datas {
    if e.ID == id {
      datas = append(datas[:i], datas[i+1:]...)
      c.IndentedJSON(http.StatusOK, e)
      return
    }
  }
	c.IndentedJSON(http.StatusOK, id)
}
