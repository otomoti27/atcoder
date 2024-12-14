package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
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

	// 31人
	a, b, c, d, e := in(), in(), in(), in(), in()

	ans := make([]chunk, 0)
	// 1文字: 5通り
	// A,B,C,D,E
	// 2文字: 10通り
	// AB,AC,AD,AE,BC,BD,BE,CD,CE,DE
	// 3文字: 10通り
	// ABC,ABD,ABE,ACD,ACE,ADE,BCD,BCE,BDE,CDE
	// 4文字: 5通り
	// ABCD,ABCE,ABDE,ACDE,BCDE
	// 5文字: 1通り
	// ABCDE

	ans = append(ans, chunk{name: "ABCDE", score: a + b + c + d + e})
	ans = append(ans, chunk{name: "ABCD", score: a + b + c + d})
	ans = append(ans, chunk{name: "ABCE", score: a + b + c + e})
	ans = append(ans, chunk{name: "ABDE", score: a + b + d + e})
	ans = append(ans, chunk{name: "ACDE", score: a + c + d + e})
	ans = append(ans, chunk{name: "BCDE", score: b + c + d + e})
	ans = append(ans, chunk{name: "ABC", score: a + b + c})
	ans = append(ans, chunk{name: "ABD", score: a + b + d})
	ans = append(ans, chunk{name: "ABE", score: a + b + e})
	ans = append(ans, chunk{name: "ACD", score: a + c + d})
	ans = append(ans, chunk{name: "ACE", score: a + c + e})
	ans = append(ans, chunk{name: "ADE", score: a + d + e})
	ans = append(ans, chunk{name: "BCD", score: b + c + d})
	ans = append(ans, chunk{name: "BCE", score: b + c + e})
	ans = append(ans, chunk{name: "BDE", score: b + d + e})
	ans = append(ans, chunk{name: "CDE", score: c + d + e})
	ans = append(ans, chunk{name: "AB", score: a + b})
	ans = append(ans, chunk{name: "AC", score: a + c})
	ans = append(ans, chunk{name: "AD", score: a + d})
	ans = append(ans, chunk{name: "AE", score: a + e})
	ans = append(ans, chunk{name: "BC", score: b + c})
	ans = append(ans, chunk{name: "BD", score: b + d})
	ans = append(ans, chunk{name: "BE", score: b + e})
	ans = append(ans, chunk{name: "CD", score: c + d})
	ans = append(ans, chunk{name: "CE", score: c + e})
	ans = append(ans, chunk{name: "DE", score: d + e})
	ans = append(ans, chunk{name: "A", score: a})
	ans = append(ans, chunk{name: "B", score: b})
	ans = append(ans, chunk{name: "C", score: c})
	ans = append(ans, chunk{name: "D", score: d})
	ans = append(ans, chunk{name: "E", score: e})

	sort.Slice(ans, func(i, j int) bool {
		if ans[i].score == ans[j].score {
			return ans[i].name < ans[j].name
		}
		return ans[i].score > ans[j].score
	})

	for a := range ans {
		out(ans[a].name)
	}
}

type chunk struct {
	name  string
	score int
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
