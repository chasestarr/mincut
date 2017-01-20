package mincut

import (
	"math/rand"
	"time"

	"github.com/chasestarr/mincut/graph"
)

// Karger takes a pointer to Graph struct as argument
// returns min cut of graph based on Kerger alg, nondeterministic
// https://en.wikipedia.org/wiki/Karger's_algorithm
// set listedTwice to true if eges are represented twice in graph
func Karger(g *graph.Graph, listedTwice bool) int {
	min := len(g.Edges)
	for i := 0; i < 1000; i++ {
		if m := iterate(g, listedTwice); m < min {
			min = m
		}
	}

	return min
}

func iterate(g *graph.Graph, listedTwice bool) int {
	for g.Vertices > 2 {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		i := r1.Intn(len(g.Edges))

		g.ContractEdge(i)
	}

	if listedTwice {
		return len(g.Edges) / 2
	}
	return len(g.Edges)
}
