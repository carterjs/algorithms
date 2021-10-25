// Package bfs implements breadth first search on a directed graph
package bfs

import "fmt"

type nodeColor string
const (
	// White = not visited (default)
	White = "White"
	// Grey = discovered
	Grey = "Grey"
	// Black = visited
	Black = "Black"
)

type Node struct {
	Color nodeColor
	Distance int
	Predecessor  *Node
}
// String converts a node to a readable string
func (n *Node) String() string {
	return fmt.Sprintf("\n(Color: %s, Distance %d, Predecessor: %p)", n.Color, n.Distance, n.Predecessor)
}

type Graph struct {
	Vertices []*Node
	AdjacencyList map[*Node][]*Node
}
// String converts a graph to a readable string (omitting the AdjacencyList)
func (g *Graph) String() string {
	return fmt.Sprintf("%v\n", g.Vertices)
}

// Explore visits each reachable node starting from s
func Explore(g *Graph, s *Node) {
	// Initialize nodes
	for _, n := range g.Vertices {
		if n != s {
			n.Color = White
			n.Distance = -1
			n.Predecessor = nil
		} else {
			// Start node is already discovered
			s.Color = Grey
			s.Distance = 0
			s.Predecessor = nil
		}
	}

	// Buffered channels can be used like queues
	// TODO: ensure this buffered channel is large enough
	q := make(chan *Node, len(g.Vertices))

	// Add start node
	q <- s

	for {
		select {
		// Dequeue node n
		case n := <-q:
			// Set to visited
			n.Color = Black

			// Visit each adjacent node
			for _, a := range g.AdjacencyList[n] {
				if a.Color == White {
					// Set discovered and set distance
					a.Color = Grey
					a.Distance = n.Distance + 1
					a.Predecessor = n

					// Enqueue a
					q <- a
				}
			}
		// Otherwise, no nodes to dequeue so we can exit
		default:
			return
		}
	}
}