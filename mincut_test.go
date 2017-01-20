package mincut

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"testing"

	"github.com/chasestarr/mincut/graph"
)

func graphFromTSV(name string) *graph.Graph {
	input, err := ioutil.ReadFile(name)
	if err != nil {
		log.Fatal("error reading test_input.txt")
	}

	// initialize new graph struct
	g := graph.New()

	// split by line
	s := strings.Split(string(input), "\n")
	for _, line := range s {
		items := strings.Split(line, "\t")
		tail, _ := strconv.Atoi(items[0])
		for i := 1; i < len(items); i++ {
			head, _ := strconv.Atoi(items[i])
			g.AddEdge(tail, head)
		}
		g.Vertices++
	}

	return g
}

func TestMinCutShort(t *testing.T) {
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
	t.Log(Karger(g, true))
}

func TestMinCutLong(t *testing.T) {
	g := graphFromTSV("test_inputLong.txt")
	t.Log(Karger(g, true))
}
