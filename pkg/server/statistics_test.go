package server

import "testing"

func TestRequestRepoAccept(t *testing.T) {
	rp := RequestParams{
		Mult1: 3,
		Mult2: 5,
		Limit: 100,
		Fuzz:  "str1",
		Buzz:  "str2",
	}
	rr := NewRequestRepo()
	rr.Accept(rp)

	if cnt := rr.requestStats[rp]; cnt != 1 {
		t.Errorf("Expected counter = 1 found : %v", cnt)
	}

	rr.Accept(rp)

	if cnt := rr.requestStats[rp]; cnt != 2 {
		t.Errorf("Expected counter = 2 found : %v", cnt)
	}
}

func TestRequestRepoTopRequest(t *testing.T) {
	rr := NewRequestRepo()

	x := RequestParams{
		Mult1: 3,
		Mult2: 5,
		Limit: 100,
		Fuzz:  "str1",
		Buzz:  "str2",
	}

	rr.Accept(x)

	rr.Accept(RequestParams{
		Mult1: 3,
		Mult2: 5,
		Limit: 100,
		Fuzz:  "str1",
		Buzz:  "str2",
	})

	rr.Accept(RequestParams{
		Mult1: 6,
		Mult2: 9,
		Limit: 100,
		Fuzz:  "str1",
		Buzz:  "str2",
	})

	res := rr.TopRequest()

	if res.RequestParams != x {
		t.Error("Invalid top request")
	}
}
