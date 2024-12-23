package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

//+++++++++++++++++++++++++++++++++++++++
// init
//+++++++++++++++++++++++++++++++++++++++

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

func init() {
	sc.Buffer([]byte{}, math.MaxInt64)
	sc.Split(bufio.ScanWords)
}

//+++++++++++++++++++++++++++++++++++++++
// main
//+++++++++++++++++++++++++++++++++++++++

func main() {
	defer func() { wr.Flush() }()

	n, m := in2()
	sx, sy := in2()

	point := make(map[[2]int]bool)
	for i := 0; i < n; i++ {
		x, y := in2()
		point[[2]int{x, y}] = true
	}

	d := make([]string, m)
	c := make([]int, m)
	for i := 0; i < m; i++ {
		d[i] = ins()
		c[i] = in()
	}

	count := 0
	visited := make(map[[2]int]bool)
	visited[[2]int{sx, sy}] = true

	cur := [2]int{sx, sy}
	for i := 0; i < m; i++ {
		prev := cur
		if d[i] == "U" {
			cur[1] += c[i]
			for j := prev[1]; j <= cur[1]; j++ {
				if point[[2]int{cur[0], j}] && !visited[[2]int{cur[0], j}] {
					count++
				}
				visited[[2]int{cur[0], j}] = true
			}
		} else if d[i] == "D" {
			cur[1] -= c[i]
			for j := cur[1]; j <= prev[1]; j++ {
				if point[[2]int{cur[0], j}] && !visited[[2]int{cur[0], j}] {
					count++
				}
				visited[[2]int{cur[0], j}] = true
			}
		} else if d[i] == "L" {
			cur[0] -= c[i]
			for j := cur[0]; j <= prev[0]; j++ {
				if point[[2]int{j, cur[1]}] && !visited[[2]int{j, cur[1]}] {
					count++
				}
				visited[[2]int{j, cur[1]}] = true
			}
		} else if d[i] == "R" {
			cur[0] += c[i]
			for j := prev[0]; j <= cur[0]; j++ {
				if point[[2]int{j, cur[1]}] && !visited[[2]int{j, cur[1]}] {
					count++
				}
				visited[[2]int{j, cur[1]}] = true
			}
		}
	}

	outf("%d %d %d\n", cur[0], cur[1], count)
}

//+++++++++++++++++++++++++++++++++++++++
// i/o
//+++++++++++++++++++++++++++++++++++++++

// 文字列を読み込む関数
func ins() string {
	sc.Scan()
	return sc.Text()
}

// Intを読み込む関数
func in() int {
	return atoi(ins())
}

// Intを読み込む関数
// 2個の変数にいっぺんで読み込むパターン
func in2() (int, int) {
	return atoi(ins()), atoi(ins())
}

// 浮動小数点数を読み込む関数
func infl() float64 {
	return atof(ins())
}

// 改行付き出力
func out(x ...interface{}) {
	fmt.Fprintln(wr, x...)
}

// フォーマット出力
func outf(s string, x ...interface{}) {
	fmt.Fprintf(wr, s, x...)
}

func outSlice[T any](s []T) {
	for i := 0; i < len(s)-1; i++ {
		fmt.Fprint(wr, s[i], " ")
	}
	fmt.Fprintln(wr, s[len(s)-1])
}

//+++++++++++++++++++++++++++++++++++++++
// convert
//+++++++++++++++++++++++++++++++++++++++

func atoi(s string) int {
	i, e := strconv.Atoi(s)
	if e != nil {
		panic(e)
	}
	return i
}

func atof(s string) float64 {
	f, e := strconv.ParseFloat(s, 64)
	if e != nil {
		panic(e)
	}
	return f
}

func itoa(i int) string {
	return strconv.Itoa(i)
}

//+++++++++++++++++++++++++++++++++++++++
// util
//+++++++++++++++++++++++++++++++++++++++

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// min for n entry
func nmin(a ...int) int {
	ret := a[0]
	for _, e := range a {
		ret = min(ret, e)
	}
	return ret
}

// max for n entry
func nmax(a ...int) int {
	ret := a[0]
	for _, e := range a {
		ret = max(ret, e)
	}
	return ret
}

func asub(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

//+++++++++++++++++++++++++++++++++++++++
// UnionFind
//+++++++++++++++++++++++++++++++++++++++

type UnionFind struct {
	n    int // 要素数
	root []int
	l, r []int
}

func newUnionFind(n int) *UnionFind {
	root, l, r := make([]int, n), make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		root[i] = -1
		l[i] = i
		r[i] = i
	}
	uf := &UnionFind{n: n, root: root, l: l, r: r}
	return uf
}

func (uf *UnionFind) find(x int) int {
	if uf.root[x] < 0 {
		return x
	}
	uf.root[x] = uf.find(uf.root[x])
	return uf.root[x]
}

func (uf *UnionFind) unite(x, y int) {
	x = uf.find(x)
	y = uf.find(y)
	if x == y {
		return
	}
	if uf.size(x) < uf.size(y) {
		x, y = y, x
	}
	// 要素の大きい方へマージ
	uf.l[x] = min(uf.l[x], uf.l[y])
	uf.r[x] = max(uf.r[x], uf.r[y])

	uf.root[x] += uf.root[y]
	uf.root[y] = x
}

func (uf *UnionFind) isSame(x, y int) bool {
	return uf.find(x) == uf.find(y)
}

func (uf *UnionFind) size(x int) int {
	return -uf.root[uf.find(x)]
}

//+++++++++++++++++++++++++++++++++++++++
// PriorityQueue
//+++++++++++++++++++++++++++++++++++++++

type pqi struct {
	val      [2]int
	priority int // The priority of the item in the queue.
}

type PriorityQueue []pqi

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// default: ascending order
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	*pq = append(*pq, x.(pqi))
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}
