package main

import (
	"context"
	"sync"
	"testing"
	"time"
)

// 与えられた複数の整数スライス（例: [][]int{{1,2,3},{4,5,6},{7,8,9}}）を、それぞれ別々のgoroutineで合計値を計算し、計算結果をひとつのスライスとして返す関数を作れ。以下の条件をすべて満たすこと。
//
// 各スライスの合計値は 同じ順番 で返すこと
// 合計値の計算は、スライスごとに必ず 別々のgoroutine で行うこと
// goroutineの終了を待ってから関数をreturnすること
func Test_sumSlicesConcurrently(t *testing.T) {
	input := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
		{10, 11, 12},
	}
	want := []int{6, 15, 24, 33}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	before := time.Now()
	got, err := sumSlicesConcurrently(ctx, input)
	after := time.Now()
	if err != nil {
		t.Fatalf("エラーが発生: %v", err)
	}
	if len(got) != len(want) {
		t.Fatalf("要素数が合わない: got=%v, want=%v", got, want)
	}
	for i := range got {
		if got[i] != want[i] {
			t.Errorf("結果が違う: index=%d, got=%d, want=%d", i, got[i], want[i])
		}
	}

	elapsed := after.Sub(before)
	if elapsed > 3*time.Second {
		t.Fatalf("所要時間が長い: %v", elapsed)
	}
}

func sumSlicesConcurrently(ctx context.Context, input [][]int) ([]int, error) {
	results := make([]int, len(input))

	var wg sync.WaitGroup
	wg.Add(len(input))
	for i, slice := range input {
		go func(i int, slice []int) {
			defer wg.Done()
			time.Sleep(1 * time.Second)
			for _, val := range slice {
				results[i] += val
			}
		}(i, slice)
	}

	wg.Wait()

	return results, nil
}
