package main

import "github.com/gin-gonic/gin"

type fruit struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func getFruits(c *gin.Context) {
	fruits := []fruit{
		{ID: "1", Name: "Maçã", Price: 2.50},
		{ID: "2", Name: "Banana", Price: 1.75},
		{ID: "3", Name: "Laranja", Price: 2.00},
		{ID: "4", Name: "Manga", Price: 3.00},
		{ID: "5", Name: "Uva", Price: 2.20},
		{ID: "6", Name: "Morango", Price: 4.50},
		{ID: "7", Name: "Abacaxi", Price: 3.80},
		{ID: "8", Name: "Melancia", Price: 5.00},
		{ID: "9", Name: "Pera", Price: 2.30},
		{ID: "10", Name: "Kiwi", Price: 3.50},
	}

	c.JSON(200, fruits)
}

func main() {
	r := gin.Default()

	r.GET("/fruits", getFruits)

	r.Run(":8000")
}
