package error_handling

import (
	"testing"
)

func TestPanic(t *testing.T) {
	t.Parallel()
	defer func() {
		err := recover()
		if err != "illegal processing" {
			t.Errorf("got %v want %v", err, "illegal processing")
		}
	}()

	downProcess(true)
}

func TestInitializeSQLServer(t *testing.T) {
	t.Parallel()
	defer func() {
		err := recover()
		if err != "はにゃ？" {
			t.Errorf("sql server down...: %v", err)
		}
	}()

	initializeSQLServer(true)
}

func TestRequestValidationError(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name   string
		age    int
		height int
		weight int
	}{
		{
			name: "invalid1",
			age:  -1,
		},
		{
			name: "invalid2",
			age:  200,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := RequestValidation(c.age)
			t.Logf("age: %d, %v", c.age, err)
			if err == nil {
				t.Errorf("不適切な年齢なのに通ってる！＞＜: %v", err)
			}
		})
	}
}
