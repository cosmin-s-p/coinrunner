package tests

import (
	"coinrunner/pkg/coinrunner"
	"testing"
)

func TestMoveForwardActionFromStartToPrologue(t *testing.T) {
	token := coinrunner.InitializeRandomToken()

	startState := coinrunner.StartPage
	endState := coinrunner.ProloguePage

	worldData := coinrunner.InitWorld()
	gameData := coinrunner.GameData{
		Token:          &token,
		CurrentState:   startState,
		IsIdle:         true,
		CanMoveForward: true,
	}

	g, _ := coinrunner.HandleAction(gameData, worldData, coinrunner.MoveForwardAction)

	if g.CurrentState != endState {
		t.Error("Moving forward from " + startState.String() + " failed, got: " + g.CurrentState.String() + " expected:" + endState.String())
	}
}

func TestMoveForwardAllCases(t *testing.T) {

	token := coinrunner.InitializeRandomToken()
	worldData := coinrunner.InitWorld()
	gameData := coinrunner.GameData{
		Token:          &token,
		IsIdle:         true,
		CanMoveForward: true,
	}

	startStates := coinrunner.GameStates

	for state := range startStates {
		gameData.CurrentState = state

		g, _ := coinrunner.HandleAction(gameData, worldData, coinrunner.MoveForwardAction)

		nextState := state + 1
		if state == coinrunner.GameOver {
			nextState = coinrunner.StartPage
		}

		if g.CurrentState != nextState {
			t.Error("Moving forward from " + state.String() + " failed, got: " + g.CurrentState.String() + " expected:" + nextState.String())
		}
	}
}
