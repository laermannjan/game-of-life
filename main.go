package main

import (
	"fmt"
	"math"
	"time"
)

type Pos struct {
	x, y int
}

func (p Pos) Neighbors() []Pos {
	ns := []Pos{}
	for _, dir := range directions {
		ns = append(ns, Pos{p.x + dir.x, p.y + dir.y})
	}

	return ns
}

var directions = []Pos{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
	{1, 1},
	{1, -1},
	{-1, -1},
	{-1, 1},
}

type World struct {
	cells map[Pos]bool
}

func (w *World) String() string {
	s := ""
	min_x, min_y, max_x, max_y := w.getBoundary()
	for i := min_y - 1; i <= max_y+1; i++ {
		for j := min_x - 1; j <= max_x+1; j++ {
			if w.cells[Pos{j, i}] {
				s += "██"
			} else {
				s += "  "
			}
		}
		s += "\n"
	}
	return s
}

func (w *World) getBoundary() (int, int, int, int) {
	min_x, min_y, max_x, max_y := math.MaxInt, math.MaxInt, math.MinInt, math.MinInt
	for cell := range w.cells {
		if cell.x < min_x {
			min_x = cell.x
		}
		if cell.x > max_x {
			max_x = cell.x
		}
		if cell.y < min_y {
			min_y = cell.y
		}
		if cell.y > max_y {
			max_y = cell.y
		}
	}
	return min_x, min_y, max_x, max_y
}

func (w *World) Next() {
	next_dead := []Pos{}
	next_alive := []Pos{}

	stack := []Pos{}
	seen := map[Pos]bool{}

	for cell := range w.cells {
		stack = append(stack, cell)
	}

	for len(stack) > 0 {
		cell := stack[0]
		stack = stack[1:]
		seen[cell] = true

		n_neighbors := 0

		for _, neighbor := range cell.Neighbors() {
			if w.cells[neighbor] {
				n_neighbors++
			}
		}

		if w.cells[cell] {
			if n_neighbors < 2 || n_neighbors > 3 {
				next_dead = append(next_dead, cell)
			}

			for _, neighbor := range cell.Neighbors() {
				if !w.cells[neighbor] && !seen[neighbor] {
					stack = append(stack, neighbor)
				}
			}
		} else {
			if n_neighbors == 3 {
				next_alive = append(next_alive, cell)
			}
		}

	}
	for _, cell := range next_dead {
		delete(w.cells, cell)
	}

	for _, cell := range next_alive {
		w.cells[cell] = true
	}
}

func main() {
	world := World{cells: map[Pos]bool{}}
	world.cells[Pos{0, 0}] = true
	world.cells[Pos{0, 1}] = true
	world.cells[Pos{0, 2}] = true
	world.cells[Pos{-1, 2}] = true
	world.cells[Pos{-2, 1}] = true
	fmt.Print("\033[H\033[2J")
	fmt.Println(world.String())
	for i := 1; i < 10; i++ {
		fmt.Print("\033[H\033[2J")
		fmt.Println("Generation", i)
		world.Next()
		fmt.Println(world.String())
		time.Sleep(500 * time.Millisecond)
	}
}
