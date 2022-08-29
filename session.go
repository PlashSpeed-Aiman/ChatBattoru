// session logic goes here
package main

type Session struct {
	CurrentGameState GameState
	SessionId        string
}

//I'm not sure of the program flow yet. This is just for reference
func StartNewSession() Session {
	var s Session
	s.CurrentGameState.tic_tac_state = make([]int, 9)
	s.CurrentGameState.current_turn = 0
	//s.SessionId = panic("Some Random Number")
	return s
}
