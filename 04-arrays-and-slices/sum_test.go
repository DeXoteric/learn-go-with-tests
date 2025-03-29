package main

import (
	"fmt"
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3}

	got := Sum(numbers)
	want := 6

	if got != want {
		t.Errorf("got %d want %d given %v", got, want, numbers)
	}
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !slices.Equal(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}
	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSums(t, got, want)
	})

}

func BenchmarkSum(b *testing.B) {
	for b.Loop() {
		Sum([]int{1, 3, 4})
	}
}

func BenchmarkSumAll(b *testing.B) {
	for b.Loop() {
		SumAll([]int{1, 3, 4}, []int{2, 3, 1})
	}
}

func BenchmarkSumAllTails(b *testing.B) {
	for b.Loop() {
		SumAllTails([]int{1, 3, 4}, []int{2, 3, 1})
	}
}

func ExampleSum() {
	result := Sum([]int{1, 3, 4})
	fmt.Println(result)
	// Output: 8
}

func ExampleSumAll() {
	result := SumAll([]int{1, 3, 4}, []int{2, 3, 1})
	fmt.Println(result)
	// Output: [8 6]
}

func ExampleSumAllTails() {
	result := SumAllTails([]int{1, 3, 4}, []int{2, 3, 1})
	fmt.Println(result)
	// Output: [7 4]
}
