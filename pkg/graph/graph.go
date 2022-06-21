package graph

import "github.com/dmitrygulevich2000/deps-dotted/pkg/set"

type Graph struct {
	edges map[string]set.Set
}

func New() *Graph {
	return &Graph{
		edges: map[string]set.Set{},
	}
}

func (g *Graph) VertCount() int {
	return len(g.edges)
}

func (g *Graph) AddEdge(from, to string) {
	if toSet, exists := g.edges[from]; exists {
		toSet.Insert(to)
	} else {
		g.edges[from] = set.Set{to: true}
	}
}

func (g* Graph) Next(n string) set.Set {
	toSet, ok := g.edges[n]
	if !ok {
		return set.Set{}
	}
	return toSet
}

func (g *Graph) EdgesList() [][]string {
	result := make([][]string, 0)
	for from, toSet := range g.edges {
		for to, _ := range toSet {
			result = append(result, []string{from, to})
		}
	}

	return result
}

func (g *Graph) VerticesList() []string {
	result := make([]string, 0, g.VertCount())
	for v, _ := range g.edges {
		result = append(result, v)
	}
	return result
}

func (g *Graph) Subgraph(vertices set.Set) *Graph {
	newGr := New()
	for u, toSet := range g.edges {
		newGr.edges[u] = set.Intersection(toSet, vertices)
	}
	return newGr
}
