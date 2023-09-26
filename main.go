package main

import (
	"bufio"
	"bytes"
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// 0:未入力
// 1-9:入力
type Board [9][9]int

// BoardをCLIに表示する
func pretty(b Board) string {
	var buf bytes.Buffer
	for i := 0; i < 9; i++ {
		if i%3 == 0 {
			buf.WriteString("+---+---+---+\n")
		}
		for j := 0; j < 9; j++ {
			if j%3 == 0 {
				buf.WriteString("|")
			}
			buf.WriteString(strconv.Itoa(b[i][j]))
		}
		buf.WriteString("|\n")
	}
	buf.WriteString("+---+---+---+\n")
	return buf.String()
}

// 1-9で重複した数値があるかを判定する
// indexを数独の数値にしていて、cには行or列or矩形のいずれかの配列が入る
func duplicated(c [10]int) bool {
	for k, v := range c {
		if k == 0 {
			continue
		}
		if v >= 2 {
			return true
		}
	}
	return false
}

// 行・列・矩形で重複した数値がないかを判定する
func verify(b Board) bool {
	// 行チェック
	for i := 0; i < 9; i++ {
		var c [10]int
		for j := 0; j < 9; j++ {
			num := b[i][j]
			c[num] += 1
		}
		if duplicated(c) {
			return false
		}
	}
	// 列チェック
	for i := 0; i < 9; i++ {
		var c [10]int
		for j := 0; j < 9; j++ {
			num := b[j][i]
			c[num] += 1
		}
		if duplicated(c) {
			return false
		}
	}
	// 矩形チェック
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			var c [10]int
			for row := i; row < i+3; row++ {
				for col := j; col < j+3; col++ {
					num := b[row][col]
					c[num] += 1
				}
			}
			if duplicated(c) {
				return false
			}
		}
	}
	return true
}

// 全てのマスが回答済みかを判定する
func solved(b Board) bool {
	if !verify(b) {
		return false
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if b[i][j] == 0 {
				return false
			}
		}
	}
	return true
}

func backtrack(b *Board, showProgress bool) bool {
	if showProgress {
		time.Sleep(time.Second * 1)
		fmt.Printf("%v\n", pretty(*b))
	}

	if solved(*b) {
		return true
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			// 未入力のマスに9から順に数値を入れてみる
			if b[i][j] != 0 {
				continue
			}
			for c := 9; c >= 1; c-- {
				b[i][j] = c
				// 重複がなければ再帰的に探索する
				if verify(*b) {
					if backtrack(b, showProgress) {
						return true
					}
				}
				// 再帰的な探索の中で9-1を全て試行して失敗した際に、呼び出し元から別の数値で試行するために0に戻しておく
				b[i][j] = 0
			}
			return false
		}
	}
	return false
}

// question : .5..83.17...1..4..3.4..56.8....3...9.9.8245....6....7...9....5...729..861.36.72.4
func convertToBoard(question string) (*Board, error) {
	if len(question) != 81 {
		return nil, errors.New("questionの長さは81(9×9)にしてください。")
	}
	s := bufio.NewScanner(strings.NewReader(question))
	s.Split(bufio.ScanRunes)

	var b Board
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if !s.Scan() {
				break
			}
			token := s.Text()
			if token == "." {
				b[i][j] = 0
				continue
			}
			n, err := strconv.Atoi(token)
			if err != nil {
				return nil, err
			}
			b[i][j] = n
		}
	}
	return &b, nil
}

var defaultQuestion = strings.Repeat(".", 81)

func main() {
	// コマンドラインの引数を解析する
	showProgress := flag.Bool("showProgress", false, "回答の探索状況を出力する場合は指定する")
	question := flag.String("question", defaultQuestion, "数独の問題")
	flag.Parse()
	fmt.Printf("question = %+v\n", *question)

	// 問題をBoardに置き換える
	b, err := convertToBoard(*question)
	fmt.Printf("board = %+v\n", *b)
	if err != nil {
		panic(err)
	}

	// 1マスずつ探索して回答を行う
	if backtrack(b, *showProgress) {
		fmt.Println(pretty(*b))
	} else {
		fmt.Fprint(os.Stderr, "cannot solve")
	}
}
