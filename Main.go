package main

import (
	"fmt"
	"strings"
)

var west = "[ Kylling Mann Korn Rev -V_____"
var boat = "\\___/"
var east = "_________E--]"

func main() {

	PrintWorld()
	PutBoatWest("Kylling")
	PutBoatWest("Mann")
	PrintWorld()
	CrossRiver("Kylling")
	PrintWorld()

}
func ViewWest() string {
	return west
}
func ViewEast() string {
	return east
}
func ViewBoat() string {
	return boat
}

func PutBoatWest(item string) {
	switch item {
	case "Kylling":
		west = strings.ReplaceAll(west, "Kylling", "")
		boat = boat[:len(boat)-2] + " Kylling " + boat[len(boat)-2:]
		break
	case "Korn":
		west = strings.ReplaceAll(west, " Korn", "")
		boat = boat[:len(boat)-2] + " Korn " + boat[len(boat)-2:]
		break
	case "Mann":
		west = strings.ReplaceAll(west, "Mann", "")
		boat = boat[:len(boat)-2] + " Mann " + boat[len(boat)-2:]
		break
	case "Rev":
		west = strings.ReplaceAll(west, "Rev", "")
		boat = boat[:len(boat)-2] + " Rev " + boat[len(boat)-2:]
		break
	}
}
func PrintWorld() {
	fmt.Println(west + boat + east)
}
func CrossRiver(item string) {
	switch item {
	case "Mann":
		boat = strings.ReplaceAll(boat, "Mann", "")
		east = east[:len(east)-2] + " Mann " + east[len(east)-2:]
		break
	case "Korn":
		boat = strings.ReplaceAll(boat, "Korn", "")
		east = east[:len(east)-2] + " Korn " + east[len(east)-2:]
		break
	case "Rev":
		boat = strings.ReplaceAll(boat, "Rev", "")
		east = east[:len(east)-2] + " Rev " + east[len(east)-2:]
		break
	case "Kylling":
		boat = strings.ReplaceAll(boat, "Kylling", "")
		east = east[:len(east)-2] + " Kylling " + east[len(east)-2:]
		break

	}

}
