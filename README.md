## [Karger min cut](https://en.wikipedia.org/wiki/Karger's_algorithm)

[![GoDoc](https://godoc.org/github.com/chasestarr/mincut?status.svg)](https://godoc.org/github.com/chasestarr/mincut)

```go
import (
  "fmt"

  "github.com/chasestarr/mincut"
  "github.com/chasestarr/mincut/graph"
)

func main() {
  g := graph.New()
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(0, 3)
	g.AddEdge(1, 3)
	g.AddEdge(1, 0)
	g.AddEdge(2, 0)
	g.AddEdge(2, 3)
	g.AddEdge(3, 0)
	g.AddEdge(3, 1)
	g.AddEdge(3, 2)
	g.Vertices = 4
	fmt.Println(mincut.Karger(g, true)) // 2 or 3
}