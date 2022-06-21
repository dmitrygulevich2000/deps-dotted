package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/dmitrygulevich2000/deps-dotted/pkg/set"
	"os"
	"strings"
	"text/template"

	"github.com/Masterminds/semver/v3"

	"github.com/dmitrygulevich2000/deps-dotted/internal/finder"
	"github.com/dmitrygulevich2000/deps-dotted/pkg/graph"
)

var tpl = template.Must(template.New(".dot desc").Parse(
	`strict digraph {
{{range .Edges}}    "{{index . 0}}" -> "{{index . 1}}"
{{end}}{{range $v, $flg := .Targets}}    "{{$v}}" [fontcolor = red]
{{end}}}`,
))

var depFlag = flag.String("dep", "", "dependency name")
var modFlag = flag.String("mod", "", "module name")
var verFlag = flag.String("ver", "*", "version constraint")

func main() {
	flag.Parse()
	if *depFlag == "" || *modFlag == "" {
		fmt.Fprintf(os.Stderr, `"mod" and "dep" arguments required`)
		os.Exit(1)
	}

	sc := bufio.NewScanner(os.Stdin)
	graph := graph.New()
	for sc.Scan() {
		edge := strings.Split(sc.Text(), " ")
		if len(edge) != 2 {
			fmt.Fprintf(os.Stderr, "wrong edge representation: %v\n", edge)
			os.Exit(1)
		}
		graph.AddEdge(edge[0], edge[1])
	}

	verCons, err := semver.NewConstraint(*verFlag)
	if err != nil {
		fmt.Fprintf(os.Stderr, "invalid version constraint: %v\n", err)
		os.Exit(1)
	}
	vert, targets := finder.VerticesBetween(graph, *modFlag, finder.ModuleMatchFunc(*depFlag, verCons))
	subgr := graph.Subgraph(vert)

	err = tpl.Execute(os.Stdout, struct{
		Edges [][]string
		Targets set.Set
	}{subgr.EdgesList(), targets})
	if err != nil {
		fmt.Fprintf(os.Stderr, "template execution error: %v\n", err)
		os.Exit(1)
	}
}
