package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type game struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Year string `json:"year"`
}

var games = []game{
	{ID: "1", Name: "Cyberpunk 2077", Year: "2020"},
	{ID: "2", Name: "Devil My Cry 5", Year: "2019"},
	{ID: "3", Name: "Black Desert", Year: "2014"},
}

func getGames(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, games)
}

func postGames(c *gin.Context) {
	var newGame game

	// Call BindJSON to bind the received JSON to
	// newGame.
	if err := c.BindJSON(&newGame); err != nil {
		return
	}

	// Add the new game to the slice.
	games = append(games, newGame)
	c.IndentedJSON(http.StatusCreated, newGame)
}

func getGameByID(c *gin.Context) {
	id := c.Param("id")
	// Loop over the list of games, looking for
	// an game whose ID value matches the parameter.
	for _, a := range games {
		if a.ID == id {			
			c.IndentedJSON(http.StatusOK, a)
			return			
		}

	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "game not found"})
}

func main() {
	router := gin.Default()
	router.GET("/games", getGames)
	router.POST("/games", postGames)
	router.GET("/games/:id", getGameByID)
	router.Run("localhost:80")
}
