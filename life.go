
package main

import (
	"math/rand"
	"time"
	"strconv"
)

type Game struct {
	current [][]bool
	next    [][]bool
}

func main() {

	game := newGame(20, 20)
	time.Sleep(1000 * time.Millisecond)
	
	iterations := 0

	for {
		print("\033[H\033[2J")
		iterations++
		game.step()
		print(boardToString(game.current, iterations))
		time.Sleep(100 * time.Millisecond)
	}
}

func (game *Game) step() {
	game.current = game.next
	
	for i := 1; i < (len(game.current)-1); i++ {
		for j := 1; j < (len(game.current[i])-1); j++ {
				game.change(i,j)
		}	
	}
}

func (game *Game) change(x, y int) {
	
	count := 0

	if game.current[x-1][y-1] == true {
		count++
	}
	if game.current[x-1][y] == true {
		count++	
	}
	if game.current[x-1][y+1] == true {
		count++	
	}
	if game.current[x][y+1] == true {
		count++	
	}
	if game.current[x+1][y+1] == true {
		count++	
	}	
	if game.current[x+1][y] == true {
		count++	
	}	
	if game.current[x+1][y-1] == true {
		count++	
	}	
	if game.current[x][y-1] == true {
		count++	
	}	
	
	
	if count == 3 {
		game.next[x][y] = true
	}
	if count < 2 {
		game.next[x][y] = false
	}
	if count > 3 {
		game.next[x][y] = false
	}
}

func boardToString(board [][]bool, iterations int) string {

	text := "" 
	
	count := 0

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == true {
				text = text + "*"
				count ++
			} else {
				text = text + " "
			}
		}
		text = text + "\n"
	}
	
	text = "Conway's Game of Life - Iteration: " + strconv.Itoa(iterations) + " - Number of living cells: " + strconv.Itoa(count) + "\n" + text
	return text

}

func newGame(h, w int) Game {

	board := make([][]bool, h+2)

	for i := range board {
		board[i] = make([]bool, w+2)
	}
	
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	times := r.Intn(w*6)
		
	for t := 0; t < times; t++ {
		i, j := rand.Intn(h), rand.Intn(w)
		board[i][j] = true
	}
	
	game := Game{next: board}
	return game
}

