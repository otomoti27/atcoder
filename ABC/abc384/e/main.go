package main

import (
	"bufio"
	"container/heap"
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

	H, W := in2()
	X := in()
	// P, Q 初期位置
	P := in() - 1
	Q := in() - 1

	S := make([][]int, H)
	for i := 0; i < H; i++ {
		S[i] = make([]int, W)
		for j := 0; j < W; j++ {
			S[i][j] = in()
		}
	}

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	dir := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for _, d := range dir {
		np := [2]int{P + d[0], Q + d[1]}
		if np[0] < 0 || np[0] >= H || np[1] < 0 || np[1] >= W {
			continue
		}
		heap.Push(&pq, pqi{val: np, priority: S[np[0]][np[1]]})
	}

	merged := make(map[[2]int]bool)
	merged[[2]int{P, Q}] = true

	ans := S[P][Q]
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(pqi)
		if merged[item.val] {
			continue
		}
		if float64(item.priority) < float64(ans)/float64(X) {
			ans += item.priority
			merged[item.val] = true
			for _, d := range dir {
				np := [2]int{item.val[0] + d[0], item.val[1] + d[1]}
				if np[0] < 0 || np[0] >= H || np[1] < 0 || np[1] >= W {
					continue
				}
				if merged[np] {
					continue
				}
				heap.Push(&pq, pqi{val: np, priority: S[np[0]][np[1]]})
			}
		} else {
			break
		}
	}

	out(ans)
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
