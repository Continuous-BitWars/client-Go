package logic

import (
	"reflect"
	"testing"

	"github.com/Continuous-BitWars/player-Go/models"
)

func StrategyTest(t *testing.T) {
	gameState := models.GameState{}
	result := Decide(gameState)
	want := []models.PlayerAction{}

	if reflect.DeepEqual(result, want) {
		t.Fatalf("Result does not equal want.")
	}
}
