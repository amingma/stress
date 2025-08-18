package gen

import (
	"fmt"
	"math/rand"
	"slices"
	"strings"
	"time"
)

func shuffleArray(arr []int) []int {
	ret := make([]int, len(arr))
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i, j := range r.Perm(len(arr)) {
		ret[i] = arr[j]
	}
	return ret
}

func GenerateRandomTree(n int) string {
	var ret strings.Builder
	fmt.Fprintf(&ret, "%d\n", n-1)
	in_tree := []int{0}
	out_tree := []int{}
	for i := 1; i < n; i++ {
		out_tree = append(out_tree, i)
	}
	for range n - 1 {
		in_tree = shuffleArray(in_tree)
		out_tree = shuffleArray(out_tree)
		fmt.Fprintf(&ret, "%d %d\n", in_tree[len(in_tree)-1], out_tree[len(out_tree)-1])
		in_tree = append(in_tree, out_tree[len(out_tree)-1])
		out_tree = slices.Delete(out_tree, len(out_tree)-1, len(out_tree))
	}
	return ret.String()
}
