package main

import "fmt"

var b = [][]int {
  {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
  {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
  {0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0},
  {0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0},
  {0, 0, 1, 1, 1, 1, 1, 1, 1, 0, 0},
  {0, 0, 1, 1, 1, 2, 1, 1, 1, 0, 0},
  {0, 0, 1, 1, 1, 1, 1, 1, 1, 0, 0},
  {0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0},
  {0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0},
  {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
  {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}}

type pos struct {
  i, j int
}

type move struct {
  from, in, to pos
}

func (m *move) String() string {
  return fmt.Sprintf("%v, %v -> %v, %v", m.from.i-1, m.from.j-1, m.to.i-1, m.to.j-1)
}

func valid() []move {
  var moves []move
  for i := 2; i <= 8; i++ {
    for j := 2; j <= 8; j++ {
      if b[i][j] != 1 {
        continue
      }

      moves = append(moves, try(i, j, 1, 0)...)
      moves = append(moves, try(i, j, 0, 1)...)
      moves = append(moves, try(i, j, -1, 0)...)
      moves = append(moves, try(i, j, 0, -1)...)
    }
  }
  return moves
}

func try(i, j, di, dj int) []move {
  if b[i + di][j + dj] == 1 && b[i + 2*di][j + 2*dj] == 2 {
    return []move{move{pos{i, j}, pos{i+di, j+dj}, pos{i+2*di, j+2*dj}}}
  }
  return nil
}

func run(n int) bool {
  if n == 1 {
    return true
  }
  moves := valid()
  for _, m := range moves {
    apply(m)
    if run(n - 1) {
      fmt.Println(m.String())
      // print()
      return true
    }
    undo(m)
  }
  return false
}

func apply(m move) {
  b[m.from.i][m.from.j] = 2
  b[m.in.i][m.in.j] = 2
  b[m.to.i][m.to.j] = 1
}

func undo(m move) {
  b[m.from.i][m.from.j] = 1
  b[m.in.i][m.in.j] = 1
  b[m.to.i][m.to.j] = 2
}

func print() {
  for i := 2; i <= 8; i++ {
    fmt.Printf("\n")
    for j := 2; j <= 8; j++ {
      switch (b[i][j]) {
        case 0:
          fmt.Printf(" ")
        case 1:
          fmt.Printf("#")
        case 2:
          fmt.Printf(".")
      }
    }
  }
}

func main() {
  run(32)
}