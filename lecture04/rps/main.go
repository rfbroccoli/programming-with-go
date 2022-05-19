package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"os"
)

var reader = bufio.NewReader(os.Stdin)

const rock = "rock"
const paper = "paper"
const scissors = "scissors"

func scan() string {
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", 1)
	return text
}

func getGameResult(userMove string, computerMove string) string {
	// user r p s
	// com r p s

	// new
	if userMove == computerMove {
		// draw
		return "DRAW"
	} else if (userMove == rock && computerMove == paper) || (userMove == paper && computerMove == scissors) || (userMove == scissors && computerMove == rock) {
		// lose
		return "YOU LOST."
	} else {
		// win
		return "YOU WON!"
	}

	// old
	// if (userMove == rock && computerMove == scissors) || (userMove == paper && computerMove == rock) || (userMove == scissors && computerMove == paper) {
	// 	// win
	// 	return "YOU WON!"
	// } else if (userMove == rock && computerMove == paper) || (userMove == paper && computerMove == scissors) || (userMove == scissors && computerMove == rock) {
	// 	// lose
	// 	return "YOU LOST."
	// } else {
	// 	// draw
	// 	return "DRAW"
	// }
}

func main() {
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)

	computerMoveNum := random.Intn(3)
	moves := []string{rock, paper, scissors}
	computerMove := moves[computerMoveNum]

	fmt.Println("choose a move: (r)ock (p)aper (s)cissors")
	// userMove = strings.Replace(userMove, "\n", "", 1)
	var userMove string

	userMoveChar := scan()

	switch userMoveChar {
	case "r":
		userMove = moves[0]
	case "p":
		userMove = moves[1]
	case "s":
		userMove = moves[2]
	default:
		fmt.Println("You choose an invalid move.")
		return
	}
	// fmt.Printf("computer chose %s\n", computerMove)
	// fmt.Printf("you chose %s\n", userMove)
	fmt.Printf("you %s vs computer %s\n", userMove, computerMove)
	result := getGameResult(userMove, computerMove)
	fmt.Println(result)

}
