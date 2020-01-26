package fizzbuzz

import "testing"

func TestFizzBuzz(t *testing.T) {
	type arguments struct {
		mult1 int
		mult2 int
		limit int
		fizz  string
		buzz  string
	}
	type expected struct {
		index int
		value string
	}

	var tests = []struct {
		name string
		args arguments
		expe []expected
	}{
		{
			name: "basic",
			args: arguments{
				mult1: 3,
				mult2: 5,
				limit: 30,
				fizz:  "fizz",
				buzz:  "buzz",
			},
			expe: []expected{
				expected{0, "1"},
				expected{2, "fizz"},
				expected{4, "buzz"},
				expected{14, "fizzbuzz"},
			},
		},
		{
			name: "different fizz",
			args: arguments{
				mult1: 3,
				mult2: 5,
				limit: 30,
				fizz:  "toto",
				buzz:  "bobo",
			},
			expe: []expected{
				expected{0, "1"},
				expected{2, "toto"},
				expected{4, "bobo"},
				expected{29, "totobobo"},
			},
		},
		{
			name: "non numbers",
			args: arguments{
				mult1: 1,
				mult2: 5,
				limit: 30,
				fizz:  "fizz",
				buzz:  "buzz",
			},
			expe: []expected{
				expected{0, "fizz"},
				expected{1, "fizz"},
				expected{4, "fizzbuzz"},
				expected{29, "fizzbuzz"},
			},
		},
	}

	for _, tst := range tests {
		tf := func(t *testing.T) {
			res, err := FizzBuzz(tst.args.mult1, tst.args.mult2, tst.args.limit, tst.args.fizz, tst.args.buzz)
			if err != nil {
				t.Error(err)
			}
			if len(res) != 30 {
				t.Errorf("expected lenth of result slice 30 found %v", len(res))
			}
			for _, e := range tst.expe {
				if r := res[e.index]; r != e.value {
					t.Errorf("expected %v found %v", e.value, r)
				}
			}
		}
		t.Run(tst.name, tf)
	}

}

func TestFuzzBuzzErrors(t *testing.T) {
	// multiple of 0
	_, err := FizzBuzz(0, 5, 30, "fizz", "bizz")
	if err != ErrInvalidMultiple {
		t.Error("Expected error ErrInvalidMultiple")
	}

	// negative limit
	_, err = FizzBuzz(3, 5, -1, "fizz", "bizz")
	if err != ErrInvalidLimit {
		t.Error("Expected error ErrInvalidLimit")
	}
}

func BenchmarkFizzBuzz(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FizzBuzz(3, 5, 30, "fizz", "bizz")
	}
}

func TestFizzBuzzElement(t *testing.T) {
	if res := fizzBuzzElement(1, 3, 5, "fizz", "buzz"); res != "1" {
		t.Errorf("fizzbuzz of 1 must be 1 and not: %v", res)
	}
	if res := fizzBuzzElement(6, 3, 5, "fizz", "buzz"); res != "fizz" {
		t.Errorf("fizzbuzz of 6 must be fizz and not: %v", res)
	}
	if res := fizzBuzzElement(10, 3, 5, "fizz", "buzz"); res != "buzz" {
		t.Errorf("fizzbuzz of 10 must be buzz and not: %v", res)
	}
	if res := fizzBuzzElement(15, 3, 5, "fizz", "buzz"); res != "fizzbuzz" {
		t.Errorf("fizzbuzz of 15 must be fizzbuzz and not: %v", res)
	}

}

func TestReplaceIfMutilpleOf(t *testing.T) {
	if ok, res := replaceIfMutilpleOf(15, 5, "fizz"); !ok || res != "fizz" {
		t.Error("Must substitute 15 with fizz since it's mutiple of 5")
	}

	if ok, _ := replaceIfMutilpleOf(7, 5, "buzz"); ok {
		t.Error("Must not substitute 7 since it's not multiple of 5")
	}
}
