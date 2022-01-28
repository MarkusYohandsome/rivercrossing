package main

import (
	"strings"
	"testing"
)

func TestPutBoatWest(t *testing.T) {

	PutBoatWest("Mann")
	PutBoatWest("Rev")
	PrintWorld()

	if strings.Contains(ViewWest(), "Kylling") && strings.Contains(ViewWest(), "Korn") {
		t.Errorf("Illegal state, Korn blir spist opp av Kylling")
	}
	if strings.Contains(ViewWest(), "Kylling") && strings.Contains(ViewWest(), "Rev") {
		t.Errorf("Illegal state, Kylling blir spist opp av Rev")
	}
}

func TestViewBoat(t *testing.T) {

	if strings.Contains(ViewBoat(), "Kylling") && strings.Contains(ViewBoat(), "Mann") && strings.Contains(ViewBoat(), "Rev") {
		t.Errorf("Illegal state, For full båt")
	}
	if strings.Contains(ViewBoat(), "Korn") && strings.Contains(ViewBoat(), "Mann") && strings.Contains(ViewBoat(), "Rev") {
		t.Errorf("Illegal state, For full båt")
	}
	if strings.Contains(ViewBoat(), "Korn") && strings.Contains(ViewBoat(), "Mann") && strings.Contains(ViewBoat(), "Kylling") {
		t.Errorf("Illegal state, For full båt")
	}
}
