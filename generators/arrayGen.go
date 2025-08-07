package gen

import (
	"fmt"
	"math/rand"
	"strings"
)

func GenerateRandomArray(n int, a int, b int) string {
	var ret strings.Builder
	fmt.Fprintf(&ret, "%d\n", n)
	for i := range n {
		x := rand.Intn(b+1-a) + a
		fmt.Fprintf(&ret, "%d", x)
		if i != n-1 {
			ret.WriteString(" ")
		}
	}
	return ret.String()
}
