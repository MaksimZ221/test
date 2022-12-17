package main

import(
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

)

type Cat struct {
	ID string `json:"id"`
	Name string `json:"name"`
	IsStrip bool `json:"is_strip"`
	Color string `json:"color"`
}

func removeByIndex(array []Cat, index int) []Cat {
	return append(array[:index], array[])
	var cats []Cat
	r.POST("/api/cat/add", func(c*gin.Context){

		var cat Cat

		if err :=c.BindJSON(&cat); err != nil {
			return
		}

		cat.ID = uuid.NewString()

		cats = append(cats, cat)

		c.JSON(200,cat)

	})

	r.GET("/api/cats:ID", func(c*gin.Context) {
		c.JSON(200,cats)
	})



	r.GET("/api/cat/:ID", func(c*gin.Context) {

		id := c.Param("id")

		var cat Cat 

        for _, ct := range cats {
	        if ct.ID == id {
		cat = ct
	}
}

		c.JSON(200,cats)

	})

	r.Run("0.0.0.0:8000")
}