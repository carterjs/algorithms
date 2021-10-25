package bfs

import (
	"reflect"
	"testing"
)

func TestExplore(t *testing.T) {

	// Text graph with just one unconnected node
	t.Run("SingleNode", func(t *testing.T) {
		s := &Node{}

		g := &Graph{
			Vertices:      []*Node{s},
			AdjacencyList: map[*Node][]*Node{},
		}

		want := &Graph{
			Vertices:      []*Node{
				{
					Color:       Black,
					Distance:    0,
					Predecessor: nil,
				},
			},
			AdjacencyList: map[*Node][]*Node{},
		}

		if Explore(g, s); !reflect.DeepEqual(g, want) {
			t.Errorf("Explore() => %v, want %v", g, want)
		}
	})

	// Sample directed graph with more than one node
	t.Run("SeveralNodes", func(t *testing.T) {
		n1 := &Node{}
		n2 := &Node{}
		n3 := &Node{}
		n4 := &Node{}
		n5 := &Node{}


		g := &Graph{
			Vertices:      []*Node{
				n1,
				n2,
				n3,
				n4,
				n5,
			},
			AdjacencyList: map[*Node][]*Node{
				n2: {n3},
				n1: {n2},
				n4: {n5},
				n5: {n1},
				n3: {n4},
			},
		}

		want := &Graph{
			Vertices:      []*Node{
				// n1
				{
					Color:       Black,
					Distance:    0,
					Predecessor: nil,
				},
				// n2
				{
					Color:       Black,
					Distance:    1,
					Predecessor: n1,
				},
				// n3
				{
					Color:       Black,
					Distance:    2,
					Predecessor: n2,
				},
				// n4
				{
					Color:       Black,
					Distance:    3,
					Predecessor: n3,
				},
				// n5
				{
					Color: Black,
					Distance: 4,
					Predecessor: n4,
				},
			},
			AdjacencyList: g.AdjacencyList,
		}

		if Explore(g, n1); !reflect.DeepEqual(g, want) {
			t.Errorf("Explore() => %v, want %v", g, want)
		}
	})
}
