// game logic goes here
package main

import (
	"fmt"
)

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
		fmt.Println("Invalid input")
		return false
	}
	return true
}

func modifyBoard(input int, turn int, tic_tac []int) {
	if turn%2 == 0 {
		tic_tac[input] = 1
	} else {
		tic_tac[input] = 2
	}
}

func showXO(tic_tac []int) {
	for i := 0; i < len(tic_tac); i++ {
		if tic_tac[i] == 1 {
			fmt.Print("X")
		} else if tic_tac[i] == 2 {
			fmt.Print("O")
		} else {
			fmt.Print("~")
		}
		if (i+1)%3 != 0 {
			fmt.Print("|")
		} else {
			fmt.Print("\n")
			fmt.Print("---------\n")
		}
	}
}

func checkPlayerTurn(turn int, tic_tac []int) {
	var input int = 0
	loopTillCorrect := func(turn int, input int) {
		for !checkValidInput(input) {
			fmt.Println("Please enter a valid input (0-8), Player ", turn, "...")
			fmt.Scanln(&input)
		}
	}
	fmt.Println("Player ", (turn%2)+1, "'s turn")
	fmt.Scanln(&input)
	loopTillCorrect((turn%2)+1, input)
	modifyBoard(input, turn, tic_tac)
}

// How to check if a player has won with an efficient algorithm
func checkPlayerWin(tic_tac []int) bool {
	//check rows
	for i := 0; i < len(tic_tac); i += 3 {
		if tic_tac[i] == tic_tac[i+1] && tic_tac[i] == tic_tac[i+2] {
			return true
		}
	}
	//check columns
	for i := 0; i < len(tic_tac); i++ {
		if tic_tac[i] == tic_tac[i+3] && tic_tac[i] == tic_tac[i+6] {
			return true
		}
	}
	//check diagonals
	if tic_tac[0] == tic_tac[4] && tic_tac[0] == tic_tac[8] {
		return true
	}
	if tic_tac[2] == tic_tac[4] && tic_tac[2] == tic_tac[6] {
		return true
	}
	return false
}

func main() {
	//initialize board
	fmt.Println("Hello world")

	//for battoru
	// var p1_board [49]byte
	// var p2_board [49]byte
	var tic_tac []int = []int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	var turn int = 0
	fmt.Println("Welcome to Tic Tac Toe")
	fmt.Println("Initial board should be like this...")
	showXO(tic_tac)
	fmt.Println("Now let's play!")
	for checkPlayerWin(tic_tac) { //will loop if false
		checkPlayerTurn(turn, tic_tac)
		showXO(tic_tac)
		turn++
	}

}
