package server

import "sync"

// RequestParams contain the parameters used to call a FizzBuzz.
type RequestParams struct {
	Mult1 int    `json:"int1"`
	Mult2 int    `json:"int2"`
	Limit int    `json:"limit"`
	Fuzz  string `json:"str1"`
	Buzz  string `json:"str2"`
}

// RequestHit contain the number of hits for the request params.
type RequestHit struct {
	RequestParams
	Counter int64 `json:"hits"`
}

type requestRepo struct {
	requestStats map[RequestParams]int64
	topRequest   RequestHit
	sync.RWMutex
}

// NewRequestRepo create a request repo.
// The created repo is tread safe.
func NewRequestRepo() *requestRepo {
	return &requestRepo{
		requestStats: make(map[RequestParams]int64),
	}
}

// Accept process a RequestParams incrementing a counter of calls
// using the same values.
func (r *requestRepo) Accept(rp RequestParams) {
	r.Lock()
	defer r.Unlock()
	r.requestStats[rp]++

	if r.requestStats[rp] > r.topRequest.Counter {
		r.topRequest = RequestHit{
			RequestParams: rp,
			Counter:       r.requestStats[rp],
		}
	}

}

// TopRequest return the RequestHit which has the grater counter in this repo.
func (r *requestRepo) TopRequest() RequestHit {
	r.RLock()
	defer r.RUnlock()
	return r.topRequest
}

var DefaultRequestRepo = NewRequestRepo()

// Accept process a RequestParams incrementing a counter of calls
// using the same values.
// This version use the default repo.
func Accept(mult1, mult2, limit int, fuzz, buzz string) {
	DefaultRequestRepo.Accept(RequestParams{
		Mult1: mult1,
		Mult2: mult2,
		Limit: limit,
		Fuzz:  fuzz,
		Buzz:  buzz,
	})
}

// TopHitRequest return the RequestHit which has the grater counter in this repo.
// This version use the default repo.
func TopHitRequest() RequestHit {
	return DefaultRequestRepo.TopRequest()
}
