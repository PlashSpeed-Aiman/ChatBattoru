// game logic goes here
package main

import (
	"fmt"
)

type GameState struct {
	tic_tac_state []int
	current_turn  int
}

func (gs GameState) UpdateGameState(tic_tac_state []int) float64 {
	panic("UpdateGameState not implemented")
}

/*
	0 | 0 | 0
	---------
	0 | 0 | 0
	---------
	0 | 0 | 0

	default - 0
	user 1 - X
	user 2 - O

	arr(string)[3][3] = {0,0,0
				 		 0,0,0
				 		 0,0,0}

	if 3 connected, user wins
*/

//1D array, length 49
//0,0
//0,1
//1,0

func checkValidInput(input int) bool {
	//check if input is valid
	if input < 0 || input > 8 {
		fmt.Println("\n\nInvalid input")
		return false
	}
	return true
}

func modifyBoard(input int, turn int, tic_tac []int) {
	tic_tac[input] = (turn % 2) + 1
}

func showXO(tic_tac []int) {
	var char_map = map[int]string{
		0: " ~ ",
		1: " X ",
		2: " O ",
	}
	var spaces int = 6
	var space_flag = true
	printSpacesOffset := func(spaces int) {
		for i := 0; i < spaces; i++ {
			fmt.Print(" ")
		}
	}
	fmt.Print("//////////BOARD//////////\n\n")
	for i := 0; i < len(tic_tac); i++ {
		if space_flag {
			printSpacesOffset(spaces)
			space_flag = false
		}
		fmt.Print(char_map[tic_tac[i]])
		if (i+1)%3 != 0 {
			fmt.Print("|")
		} else {
			fmt.Print("\n")
			printSpacesOffset(spaces)
			fmt.Print("-----------\n")
			space_flag = true
		}
	}
	fmt.Print("\n//////////BOARD//////////\n")
}

func checkPlayerTurn(turn int, tic_tac []int) {
	var input int = -1
	getInput := func(input int) int {
		fmt.Scanln(&input)
		return input - 1
	}
	loopTillCorrect := func(turn int, input int) int {
		for !checkValidInput(input) {
			fmt.Println("Please enter a valid input (1-9), Player ", turn, "...")
			fmt.Print("Enter your input: ")
			input = getInput(input)
		}
		return input
	}
	fmt.Println("\n\nPlayer ", (turn%2)+1, "'s turn")
	fmt.Print("Enter your input (1-9): ")
	input = getInput(input)
	input = loopTillCorrect((turn%2)+1, input)
	fmt.Print("\n\n\n")
	modifyBoard(input, turn, tic_tac)
}

// How to check if a player has won with an efficient algorithm
func checkPlayerWin(tic_tac []int) bool {
	var win bool = false
	var char_match int = 0
	//check rows
	for i := 0; i < len(tic_tac); i += 3 {
		if tic_tac[i] == tic_tac[i+1] && tic_tac[i] == tic_tac[i+2] {
			char_match = tic_tac[i]
			break
		}
	}
	//check columns
	for i := 0; i < 3; i++ {
		if tic_tac[i] == tic_tac[i+3] && tic_tac[i] == tic_tac[i+6] {
			char_match = tic_tac[i]
			break
		}
	}
	//check diagonals
	if tic_tac[0] == tic_tac[4] && tic_tac[0] == tic_tac[8] {
		char_match = tic_tac[0]
	}
	if tic_tac[2] == tic_tac[4] && tic_tac[2] == tic_tac[6] {
		char_match = tic_tac[2]
	}
	if char_match != 0 {
		win = true
	}
	return win
}

func main() {
	fmt.Println("Hello world")

	//for battoru
	// var p1_board [49]byte
	// var p2_board [49]byte
	var tic_tac []int = []int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	var turn int = -1
	fmt.Println("Welcome to Tic Tac Toe")
	fmt.Print("Initial board should be like this...\n\n")
	showXO(tic_tac)
	fmt.Print("\nNow let's play!\n")
	for !checkPlayerWin(tic_tac) {
		turn++
		checkPlayerTurn(turn, tic_tac)
		showXO(tic_tac)
	}
	fmt.Println("\n\nPlayer ", (turn%2)+1, " wins!")
}
