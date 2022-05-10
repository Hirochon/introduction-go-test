package main

import (
	"testing"
	"time"
)

func TestSum(t *testing.T) {
	got := sum(2, 2)
	want := 4
	if got != want {
		t.Errorf("got %d want %d given, %d, %d", got, want, 2, 2)
	}
}

func TestSumSub(t *testing.T) {
	t.Run("２足す２のテスト", func(t *testing.T) {
		got := sum(2, 2)
		want := 4
		if got != want {
			t.Errorf("got %d want %d given, %d, %d", got, want, 2, 2)
		}
	})
}

func TestSumTableDriven(t *testing.T) {
	cases := []struct {
		name             string
		arg1, arg2, want int
	}{
		{"２足す２のテスト", 2, 2, 4},
		//{"３足す４のテスト", 3, 4, 7},
		//{"４足す３のテスト", 4, 3, 7},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := sum(c.arg1, c.arg2)
			if got != c.want {
				t.Errorf("got %d want %d given, %d, %d", got, c.want, c.arg1, c.arg2)
			}
		})
	}
}

func TestSumParallel1(t *testing.T) {
	//t.Parallel()
	time.Sleep(time.Second)
	got := sum(2, 2)
	want := 4
	if got != want {
		t.Errorf("got %d want %d given, %d, %d", got, want, 2, 2)
	}
}

func TestSumParallel2(t *testing.T) {
	//t.Parallel()
	time.Sleep(time.Second)
	got := sum(2, 2)
	want := 4
	if got != want {
		t.Errorf("got %d want %d given, %d, %d", got, want, 2, 2)
	}
}

func TestSumParallel3(t *testing.T) {
	//t.Parallel()
	time.Sleep(time.Second)
	got := sum(2, 2)
	want := 4
	if got != want {
		t.Errorf("got %d want %d given, %d, %d", got, want, 2, 2)
	}
}

func TestSumParallel4(t *testing.T) {
	//t.Parallel()
	time.Sleep(time.Second)
	got := sum(2, 2)
	want := 4
	if got != want {
		t.Errorf("got %d want %d given, %d, %d", got, want, 2, 2)
	}
}

//func TestCoverage100(t *testing.T) {
//	main()
//}
