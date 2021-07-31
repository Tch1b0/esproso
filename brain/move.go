package brain

import (
	"fmt"

	GameData "github.com/Tch1b0/Esproso/data"
)


func Move(data GameData.Data) string {
	position := data.You.Head
	possibleMoves := []string{
		"up",
		"down",
		"left",
		"right",
	}
	avoidMoves := avoid(data)

	closeFood := findFood(data)

	// Get the direction the player has to go to get to the food
	foodDirection := getDirection(position, closeFood)

	// Go up/down if y-value isn't equal
	if position.Y != closeFood.Y && 
	    !containsString(avoidMoves, foodDirection) {
			fmt.Println("FOOD Y")
		return foodDirection
	} else if position.X != closeFood.X && 
		!containsString(avoidMoves, foodDirection) {
		fmt.Println("FOOD X")
		return foodDirection		
	}

	for _, move := range possibleMoves {
		if !containsString(avoidMoves, move) {
			fmt.Println("FREE MOVE")
			return move
		}
	}

	return "up"
}