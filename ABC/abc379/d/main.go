package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

//+++++++++++++++++++++++++++++++++++++++
// 準備用の処理
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
	var q int
	fmt.Scanf("%d", &q)

	sc := bufio.NewScanner(os.Stdin)

	remain := 10 ^ 100
	plot := make([]int, 0)
	for i := 0; i < q; i++ {
		sc.Scan()
		query := strings.Split(sc.Text(), " ")
		q1, _ := strconv.Atoi(query[0])

		if q1 == 1 {
			plot = append(plot, 0)
			remain--
			continue
		}

		q2, _ := strconv.Atoi(query[1])
		if q1 == 2 {
			for i := 0; i < len(plot); i++ {
				plot[i] += q2
			}
			continue
		}
		if q1 == 3 {
			count := 0
			for _, v := range plot {
				if v >= q2 {
					count++
					plot = plot[1:]
				} else {
					// 小さい順にならんでるため、これ以上大きいのはない
					break
				}
			}
			fmt.Println(count)
		}
	}
}

//+++++++++++++++++++++++++++++++++++++++
// 入力用の関数
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

//+++++++++++++++++++++++++++++++++++++++
// 出力用の関数
//+++++++++++++++++++++++++++++++++++++++

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
// 変換用の関数
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
// 基本関数
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
