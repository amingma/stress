package gen

import (
	"fmt"
	"math/rand"
	"slices"
	"strings"
)

type Edge struct {
	a int
	b int
}

func getRandomEdge(adj [][]bool) Edge {
	n := len(adj)
	a := rand.Intn(n) + 1
	b := rand.Intn(n) + 1
	for a == b || adj[a][b] {
		a = rand.Intn(n) + 1
		b = rand.Intn(n) + 1
	}
	return Edge{a, b}
}

func GenerateRandomGraph(n int, m int) string {
	var ret strings.Builder
	adj := make([][]bool, n)
	for i := range adj {
		adj[i] = make([]bool, n)
	}
	fmt.Fprintf(&ret, "%d %d", n, m)
	in_tree := []int{1}
	out_tree := []int{}
	for i := 2; i <= n; i++ {
		out_tree = append(out_tree, i)
	}
	for range n - 1 {
		in_tree = shuffleArray(in_tree)
		out_tree = shuffleArray(out_tree)
		fmt.Fprintf(&ret, "%d %d\n", in_tree[len(in_tree)-1], out_tree[len(out_tree)-1])
		adj[in_tree[len(in_tree)-1]][out_tree[len(out_tree)-1]] = true
		adj[out_tree[len(out_tree)-1]][in_tree[len(in_tree)-1]] = true
		in_tree = append(in_tree, out_tree[len(out_tree)-1])
		out_tree = slices.Delete(out_tree, len(out_tree)-1, len(out_tree))
	}
	for range m - n + 1 {
		edge := getRandomEdge(adj)
		adj[edge.a][edge.b] = true
		adj[edge.b][edge.a] = true
		fmt.Fprintf(&ret, "%d %d\n", edge.a, edge.b)
	}
	return ret.String()
}
