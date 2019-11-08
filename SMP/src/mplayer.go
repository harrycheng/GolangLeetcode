package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"hc.com/SMP/pkg/mplayer/library"
	"hc.com/SMP/pkg/mplayer/mp"
)

var lib *library.MusicManager
var id int = 1
var ctrl, signal chan int

func handleLibCommans(tokens []string) {
	switch tokens[1] {
	case "list":
		for i := 0; i < lib.Len(); i++ {
			e, _ := lib.Get(i)
			fmt.Println(i+1, ":", e.Name)
		}
	case "add":
		if len(tokens) == 6 {
			id++
			lib.Add(&library.MusicEntry{strconv.Itoa(id), tokens[2], tokens[3], tokens[4], tokens[5]})
		} else {
			fmt.Println("use age : lib add name artist source type")
		}
	case "remove":
		if len(tokens) == 3 {
			lib.RemoveByName(tokens[2])
		} else {
			fmt.Println("useage : lib remove name")
		}

	default:
		fmt.Println("Unrecognized lib command ", tokens[1])
	}
}

func handlePlayCommand(tokens []string) {
	if len(tokens) != 2 {
		fmt.Println(" usage: play <name>")
		return
	}

	_, e := lib.Find(tokens[1])
	if e == nil {
		fmt.Println("The music ", tokens[1], "does not exist")
		return
	}

	mp.Play(e.Source, e.Type)

}

func main() {
	fmt.Print("Simple media player system...")
	lib = library.NewMusicManager()
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Enter command->")
		rawLine, _, _ := r.ReadLine()
		line := string(rawLine)

		if line == "q" || line == "e" {
			break
		}

		tokens := strings.Split(line, " ")
		fmt.Println(tokens)
		if tokens[0] == "lib" {
			handleLibCommans(tokens)
		} else if tokens[0] == "play" {
			handlePlayCommand(tokens)
		} else {
			fmt.Println("unsupport command : ", tokens[0])
		}
	}
}
