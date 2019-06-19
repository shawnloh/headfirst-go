package main

import "github.com/shawnloh/headfirstgo/chapter-11/gadget"

type Player interface {
	Play(string)
	Stop()
}

func playList(device Player, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func main() {
	player := gadget.TapePlayer{}
	mixTape := []string{"Jessie's Girl", "Whip it", "9 to 5"}
	playList(player, mixTape)
	recPlayer := gadget.TapeRecorder{}
	playList(recPlayer, mixTape)
}
