package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDuplicated(t *testing.T) {
	// 全て未入力の場合、重複していないと判定する
	{
		t.Run(fmt.Sprintf("#%d 全て未入力", 1), func(t *testing.T) {
			c := [10]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
			if duplicated(c) {
				t.Fatal("expected: false, actual: true")
			}
		})
	}

	// 2未満の場合、重複していないと判定する
	{
		t.Run(fmt.Sprintf("#%d 重複なし", 2), func(t *testing.T) {
			c := [10]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
			if duplicated(c) {
				t.Fatal("expected: false, actual: true")
			}
		})
	}

	// 2以上の場合、重複していると判定する
	{
		t.Run(fmt.Sprintf("#%d 重複あり", 3), func(t *testing.T) {
			c := [10]int{1, 1, 1, 1, 0, 1, 0, 2, 1, 0}
			if !duplicated(c) {
				t.Fatal("expected: true, actual: false")
			}
		})
	}
}

func TestVerify(t *testing.T) {
	cases := []struct {
		msg      string
		b        Board
		expected bool
	}{
		{
			msg: "全て未入力",
			b: Board{
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
			expected: true,
		},
		{
			msg: "行チェック",
			b: Board{
				{0, 5, 0, 0, 5, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
			expected: false,
		},
		{
			msg: "列チェック",
			b: Board{
				{0, 0, 0, 0, 5, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 5, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
			expected: false,
		},
		{
			msg: "矩形チェック",
			b: Board{
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 5, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 5, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
			expected: false,
		},
	}

	for k, v := range cases {
		t.Run(fmt.Sprintf("#%d %s", k+1, v.msg), func(t *testing.T) {
			result := verify(v.b)
			if result != v.expected {
				t.Errorf("expexted: %v, actual: %v", v.expected, result)
			}
		})
	}
}

func TestSolve(t *testing.T) {
	// test case from https://github.com/t-dillon/tdoku provided under BSD-2 LICENSE.
	// question: .5..83.17...1..4..3.4..56.8....3...9.9.8245....6....7...9....5...729..861.36.72.4
	b := Board{
		{6, 5, 2, 4, 8, 3, 9, 1, 7},
		{0, 0, 0, 1, 0, 0, 4, 0, 0},
		{3, 0, 4, 0, 0, 5, 6, 0, 8},
		{0, 0, 0, 0, 3, 0, 0, 0, 9},
		{0, 9, 0, 8, 2, 4, 5, 0, 0},
		{0, 0, 6, 0, 0, 0, 0, 7, 0},
		{0, 0, 9, 0, 0, 0, 0, 5, 0},
		{0, 0, 7, 2, 9, 0, 0, 8, 6},
		{1, 0, 3, 6, 0, 7, 2, 0, 4},
	}

	// answer: 652483917978162435314975628825736149791824563436519872269348751547291386183657294
	expected := Board{
		{6, 5, 2, 4, 8, 3, 9, 1, 7},
		{9, 7, 8, 1, 6, 2, 4, 3, 5},
		{3, 1, 4, 9, 7, 5, 6, 2, 8},
		{8, 2, 5, 7, 3, 6, 1, 4, 9},
		{7, 9, 1, 8, 2, 4, 5, 6, 3},
		{4, 3, 6, 5, 1, 9, 8, 7, 2},
		{2, 6, 9, 3, 4, 8, 7, 5, 1},
		{5, 4, 7, 2, 9, 1, 3, 8, 6},
		{1, 8, 3, 6, 5, 7, 2, 9, 4},
	}

	// 回答できなかった場合、失敗とする
	if !backtrack(&b, false) {
		t.Fatal("should solve but cannot")
	}

	// Boardが等価ではない場合、失敗とする
	if !reflect.DeepEqual(expected, b) {
		t.Fatalf("expected: %v, actual: %v", expected, b)
	}
}

func TestConvertToBoard(t *testing.T) {
	result, err := convertToBoard(".5..83.17...1..4..3.4..56.8....3...9.9.8245....6....7...9....5...729..861.36.72.4")
	if err != nil {
		t.Fatalf("parse failed: %s", err)
	}

	expected := Board{
		{0, 5, 0, 0, 8, 3, 0, 1, 7},
		{0, 0, 0, 1, 0, 0, 4, 0, 0},
		{3, 0, 4, 0, 0, 5, 6, 0, 8},
		{0, 0, 0, 0, 3, 0, 0, 0, 9},
		{0, 9, 0, 8, 2, 4, 5, 0, 0},
		{0, 0, 6, 0, 0, 0, 0, 7, 0},
		{0, 0, 9, 0, 0, 0, 0, 5, 0},
		{0, 0, 7, 2, 9, 0, 0, 8, 6},
		{1, 0, 3, 6, 0, 7, 2, 0, 4},
	}

	// Boardが等価ではない場合、失敗とする
	if !reflect.DeepEqual(expected, *result) {
		t.Fatalf("expected: %v, actual: %v", expected, *result)
	}

	// 問題の長さが81より短い場合、失敗とする
	{
		b, err := convertToBoard(".")
		if b != nil {
			t.Fatalf("board should be nil: %s", err)
		}
	}
}
