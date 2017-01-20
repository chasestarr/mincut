package graph

// Graph maintains slices of graph vertices and edges
type Graph struct {
	Vertices int
	Edges    []edge
}

type edge struct {
	tail int
	head int
}

// type vertex struct {
// 	label int
// }

// New returns pointer to a new graph struct
func New() *Graph {
	return &Graph{}
}

// AddEdge appends new edge struct to a graph's edges slice
func (g *Graph) AddEdge(tail, head int) {
	e := edge{tail, head}
	g.Edges = append(g.Edges, e)
}

// AddVertex increments vertices count of a Graph struct
func (g *Graph) AddVertex(label int) {
	// v := vertex{label}
	// g.vertices = append(g.vertices, v)
	g.Vertices++
}

// ContractEdge will collapse an edge and merge vertex elements
func (g *Graph) ContractEdge(i int) {
	e := g.Edges[i]

	// update edges to remove e
	g.Edges = append(g.Edges[:i], g.Edges[i+1:]...)

	// loop to remove e.head from edges slice
	for j, edge := range g.Edges {
		if edge.head == e.head {
			g.Edges[j].head = e.tail
		} else if edge.tail == e.head {
			g.Edges[j].tail = e.tail
		}
	}

	// remove self loops
	temp := []edge{}
	for _, edge := range g.Edges {
		if edge.head != edge.tail {
			temp = append(temp, edge)
		}
	}
	g.Edges = temp

	g.Vertices--
}
