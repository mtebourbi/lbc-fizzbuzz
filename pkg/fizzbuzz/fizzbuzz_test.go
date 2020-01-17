package fizzbuzz

import "testing"

func TestFizzBuzz(t *testing.T) {
	res, err := FizzBuzz(3, 5, 30, "fizz", "bizz")
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
	if res[0] != "1" {
		t.Error("Error")
	}

	// multiple of 0
	FizzBuzz(0, 5, 30, "fizz", "bizz")

	// negative limit
	FizzBuzz(3, 5, -1, "fizz", "bizz")

}

func TestFizzBuzzElement(t *testing.T) {
	if res := fizzBuzzElement(1, 3, 5, "fizz", "buzz"); res != "1" {
		// TODO: Check string format (verify all)
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
