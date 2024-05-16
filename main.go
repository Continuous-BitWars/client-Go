package main

import (
	"net/http"

	"github.com/Continuous-BitWars/player-Go/logic"
	"github.com/Continuous-BitWars/player-Go/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.SetTrustedProxies(nil)
	r.GET("/", Identify)
	r.POST("/", Index)
	r.Run(":3000")

}

func Identify(c *gin.Context) {
	c.String(http.StatusOK, "Bitwars Player Go")
}

func Index(c *gin.Context) {
	var gameState models.GameState
	if c.Bind(&gameState) != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "StatusInternalServerError"})
		return
	}

	c.JSON(http.StatusOK, logic.Decide(gameState))
}
