package finder

import (
	"github.com/dmitrygulevich2000/deps-dotted/pkg/graph"
	"github.com/dmitrygulevich2000/deps-dotted/pkg/set"
)

type MatchFunc func(string) bool

func dfs(g *graph.Graph, matches MatchFunc, curr string, used map[string]bool, path set.Set, vertices, matched set.Set) {
	if matches(curr) {
		vertices.Append(path)
		matched.Insert(curr)
		used[curr] = false
		return
	}

	next := g.Next(curr)
	for n, _ := range next {
		if !used[n] {
			used[n] = true
			path.Insert(n)

			dfs(g, matches, n, used, path, vertices, matched)

			path.Remove(n)
			used[n] = false
		}
	}
}

func VerticesBetween(graph *graph.Graph, from string, to MatchFunc) (vertices set.Set, matched set.Set) {
	used := set.New(from)
	parents := set.New(from)
	vertices = set.New()
	matched = set.New()

	dfs(graph, to, from, used, parents, vertices, matched)

	return
}
