package server

import (
	"fmt"
	"go-dart/common"
	"testing"
)

func TestGamex01End(t *testing.T) {
	fmt.Println()
	fmt.Println("TestGamex01End")

	game := NewGamex01(1)
	game.AddPlayer("Alice")
	game.AddPlayer("Bob")
	game.Start()
	state := game.HandleDart(common.Sector{Val: 5, Pos: 1})

	if !state.Ongoing {
		t.Error("Game should not be ended")
	}

	alice := state.Scores[0]

	if alice.Score != 1 {
		t.Error("Alice should have the same score : 1")
	}

	if state.CurrentPlayer != 1 || state.CurrentDart != 0 {
		t.Errorf("Should be bob's turn, first Dart (%d, %d)", state.CurrentPlayer, state.CurrentDart)
	}

	state = game.HandleDart(common.Sector{Val: 1, Pos: 1})

	if state.Ongoing {
		t.Error("Game should be ended")
	}

	bob := state.Scores[0]

	if bob.Player != "Bob" {
		t.Error("Bob should be first")
	}

}

func TestGame301(t *testing.T) {
	fmt.Println()
	fmt.Println("TestGame301")

	game := NewGamex01(301)
	game.AddPlayer("Alice")
	game.AddPlayer("Bob")
	game.AddPlayer("Charly")
	game.AddPlayer("Dan")

	game.Start()

	// Visit 1, Player 0
	state := game.HandleDart(common.Sector{Val: 20, Pos: 3})
	verifyCurrents(state, 0, 1, t)
	state = game.HandleDart(common.Sector{Val: 20, Pos: 3})
	verifyCurrents(state, 0, 2, t)
	state = game.HandleDart(common.Sector{Val: 20, Pos: 3})
	verifyCurrents(state, 1, 0, t)
	verifyScore(state, 121, 0, t)

	// Visit 1, Player 1
	state = game.HandleDart(common.Sector{Val: 25, Pos: 2})
	verifyCurrents(state, 1, 1, t)
	state = game.HandleDart(common.Sector{Val: 25, Pos: 2})
	verifyCurrents(state, 1, 2, t)
	state = game.HandleDart(common.Sector{Val: 25, Pos: 2})
	verifyCurrents(state, 2, 0, t)
	verifyScore(state, 151, 1, t)

	// Visit 1, Player 2
	state = game.HandleDart(common.Sector{Val: 19, Pos: 2})
	verifyCurrents(state, 2, 1, t)
	state = game.HandleDart(common.Sector{Val: 0, Pos: 0})
	verifyCurrents(state, 2, 2, t)
	state = game.HandleDart(common.Sector{Val: 25, Pos: 2})
	verifyCurrents(state, 3, 0, t)
	verifyScore(state, 213, 2, t)

	// Visit 1, Player 3
	state = game.HandleDart(common.Sector{Val: 0, Pos: 0})
	verifyCurrents(state, 3, 1, t)
	state = game.HandleDart(common.Sector{Val: 0, Pos: 0})
	verifyCurrents(state, 3, 2, t)
	state = game.HandleDart(common.Sector{Val: 0, Pos: 0})
	verifyCurrents(state, 0, 0, t)
	verifyScore(state, 301, 3, t)

	// Visit 2, Player 0
	state = game.HandleDart(common.Sector{Val: 20, Pos: 3})
	verifyCurrents(state, 0, 1, t)
	state = game.HandleDart(common.Sector{Val: 7, Pos: 3})
	verifyCurrents(state, 0, 2, t)
	state = game.HandleDart(common.Sector{Val: 20, Pos: 2})
	verifyCurrents(state, 1, 0, t)
	verifyScore(state, 0, 0, t)
	verifyRank(state, 1, 0, t)

	// Visit 2, Player 1
	state = game.HandleDart(common.Sector{Val: 20, Pos: 3})
	verifyCurrents(state, 1, 1, t)
	verifyScore(state, 91, 1, t)
	state = game.HandleDart(common.Sector{Val: 20, Pos: 3})
	verifyCurrents(state, 1, 2, t)
	verifyScore(state, 31, 1, t)
	state = game.HandleDart(common.Sector{Val: 20, Pos: 2})
	verifyCurrents(state, 2, 0, t)
	verifyScore(state, 151, 1, t)

	// Visit 2, Player 2
	state = game.HandleDart(common.Sector{Val: 20, Pos: 3})
	verifyCurrents(state, 2, 1, t)
	state = game.HandleDart(common.Sector{Val: 20, Pos: 3})
	verifyCurrents(state, 2, 2, t)
	state = game.HandleDart(common.Sector{Val: 20, Pos: 3})
	verifyCurrents(state, 3, 0, t)
	verifyScore(state, 33, 2, t)

	// Visit 2, Player 3
	state = game.HandleDart(common.Sector{Val: 0, Pos: 0})
	verifyCurrents(state, 3, 1, t)
	state = game.HandleDart(common.Sector{Val: 0, Pos: 0})
	verifyCurrents(state, 3, 2, t)
	state = game.HandleDart(common.Sector{Val: 0, Pos: 0})
	verifyCurrents(state, 1, 0, t)
	verifyScore(state, 301, 3, t)

	// Visit 3, Player 1
	state = game.HandleDart(common.Sector{Val: 20, Pos: 3})
	verifyCurrents(state, 1, 1, t)
	state = game.HandleDart(common.Sector{Val: 20, Pos: 3})
	verifyCurrents(state, 1, 2, t)
	state = game.HandleDart(common.Sector{Val: 20, Pos: 1})
	verifyCurrents(state, 2, 0, t)
	verifyScore(state, 11, 1, t)

	// Visit 3, Player 2
	state = game.HandleDart(common.Sector{Val: 10, Pos: 3})
	verifyCurrents(state, 2, 1, t)
	state = game.HandleDart(common.Sector{Val: 1, Pos: 1})
	verifyCurrents(state, 2, 2, t)
	state = game.HandleDart(common.Sector{Val: 1, Pos: 2})
	verifyCurrents(state, 3, 0, t)
	verifyScore(state, 0, 2, t)
	verifyRank(state, 2, 2, t)

	// Visit 3, Player 3
	state = game.HandleDart(common.Sector{Val: 0, Pos: 0})
	verifyCurrents(state, 3, 1, t)
	state = game.HandleDart(common.Sector{Val: 0, Pos: 0})
	verifyCurrents(state, 3, 2, t)
	state = game.HandleDart(common.Sector{Val: 0, Pos: 0})
	verifyCurrents(state, 1, 0, t)
	verifyScore(state, 301, 3, t)

	// Visit 4, Player 1
	state = game.HandleDart(common.Sector{Val: 1, Pos: 1})
	verifyCurrents(state, 1, 1, t)
	state = game.HandleDart(common.Sector{Val: 5, Pos: 2})

	if state.Ongoing {
		t.Error("Game should be ended")
	}

	verifyScore(state, 0, 0, t)
	verifyScore(state, 0, 1, t)
	verifyScore(state, 0, 2, t)
	verifyScore(state, 301, 3, t)

	verifyRank(state, 1, 0, t)
	verifyRank(state, 2, 1, t)
	verifyRank(state, 3, 2, t)
	verifyRank(state, 4, 3, t)

	verifyPlayer(state, "Alice", 0, t)
	verifyPlayer(state, "Charly", 1, t)
	verifyPlayer(state, "Bob", 2, t)
	verifyPlayer(state, "Dan", 3, t)

}

func verifyCurrents(state *common.GameState, p, d int, t *testing.T) {
	if state.CurrentPlayer != p || state.CurrentDart != d {
		t.Errorf("Player should be %d and Dart %d, but was %d and %d", p, d, state.CurrentPlayer, state.CurrentDart)
	}
}

func verifyScore(state *common.GameState, score, player int, t *testing.T) {
	if state.Scores[player].Score != score {
		t.Errorf("Score should be %d but was %d for Player %d", score, state.Scores[player].Score, player)
	}
}

func verifyRank(state *common.GameState, rank, player int, t *testing.T) {
	if state.Scores[player].Rank != rank {
		t.Errorf("Rank should be %d but was %d for Player %d", rank, state.Scores[player].Rank, player)
	}
}

func verifyPlayer(state *common.GameState, name string, player int, t *testing.T) {
	if state.Scores[player].Player != name {
		t.Errorf("Name should be %s but was %s for Player %d", name, state.Scores[player].Player, player)

	}
}