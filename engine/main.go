package main

import (
	"os"
	"os/exec"
	"io"
	"fmt"
//	"strings"
	"encoding/json"
	"bytes"
	"time"
	"database/sql"
	_ "github.com/ziutek/mymysql/godrv"
)

/*type coord string

func (c chord) setCoord(x, y int) {
	c = strconv.Itoa(x), "x"
}*/

type (
	GenericType struct {
		X, Y int
		T string
	}
	
	Point struct {
		X, Y int
	}
	
	BigData struct {
		Structures, Units []GenericType
		Obstacles []Point
	}
	
	Map struct {
		W, H int
		Tiles [][]byte
	}
)

func MapGenerator(w, h int) (m Map) {
	m.W, m.H = w, h
	
	m.Tiles = make([][]byte, MAP_WIDTH)
	
	for i:=0;i<w;i++ {
		m.Tiles[i] = make([]byte, MAP_HEIGHT)
		for j:=0;j<h;j++ {
			m.Tiles[i][j] = 'P'
		}
	}

	return
}

const MAP_WIDTH int = 64
const MAP_HEIGHT int = 64


func fetchAndStartBattle() (c bool) {
	c = false;
	con, err := sql.Open("mymysql", "nullCase/nullCase/7dVCSzvVcvvzmnJG")
	
	if err != nil {
		fmt.Println(err)
		return
	}
	
	battle := con.QueryRow("SELECT `id` FROM `battles` WHERE `played` = 0")
	
	var id int;
	err = battle.Scan(&id)
	
	if err != nil {
		// No battles in queue
		fmt.Println(err)
		return
	}
	
	// Get players from new table: battle_players
	// Should probably not define this anymore, but just a regular variable
		const PLAYERS int = 2
	
	// Probably also get the path to some storage area where the code that
	// was just uploaded are... This can be defined by a const.. maybe..
	
	con.Close()
	
	// Generate Map
//	FullMap := MapGenerator(MAP_WIDTH,MAP_HEIGHT)
	
	/*/ Maybe this data is not required, START {
	PlayerMap := make([]Map, PLAYERS)
	
	for i:=0;i<PLAYERS;i++ {
		PlayerMap[i] = Map{
				W: MAP_WIDTH,
				H: MAP_HEIGHT,
				Tiles: make([][]byte, MAP_WIDTH),
			}
		
		for j:=0;j<MAP_WIDTH;j++ {
			PlayerMap[i].Tiles[j] = make([]byte, MAP_HEIGHT)
		}
	}
	// } END
	
	// Just done this to use the two variables.. LAME? YES!
	fmt.Println(FullMap, PlayerMap)*/
	
	// At this point we should spin up a new thread to handle the battle
	// It should ...
	
	go doBattle()
	
	return true
}

func doBattle() {
	// Do some initial setup
	
	// create a loop for each round
	cmd := exec.Command("php", "-f", "package/php/test.php")
	
	stdOut, errOut := cmd.StdoutPipe()
	if errOut != nil {
		fmt.Println(errOut)
		return
	}
	
	stdIn, errIn := cmd.StdinPipe()
	if errIn != nil {
		fmt.Println(errIn)
		return
	}
	
	if err := cmd.Start(); err != nil {
		fmt.Println(err)
		return
	}
	
	data := BigData{
		Structures: []GenericType{
				GenericType{X:1,Y:2,T:"Q"},
			},
		Units: []GenericType{
				GenericType{X:2,Y:2,T:"B"},
				GenericType{X:4,Y:3,T:"W"},
			},
		Obstacles: []Point{
				Point{X:9,Y:2},
			},
	}
	
	b, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	
	jsonR := bytes.NewReader(b)
	
	fmt.Println("In")
	io.Copy(stdIn, jsonR)
	
	fmt.Print("Out: ")
    io.Copy(os.Stdout, stdOut)
	fmt.Println("")
	
	fmt.Println("Wait")
	if err := cmd.Wait(); err != nil {
		fmt.Println(err)
		return
	}
	
	fmt.Println("Closed")
}

func battleRound() {
	
}

func main() {
	sleep := 1000
	for {
		if !fetchAndStartBattle() {
			sleep += 500
		}
		
		if sleep > 60000 {
			// Don't sleep more than a minute
			sleep = 60000
		}
		
		time.Sleep(time.Duration(sleep) * time.Millisecond)
	}

	fmt.Println("\n\n## CLOSING DOWN\n")
}
