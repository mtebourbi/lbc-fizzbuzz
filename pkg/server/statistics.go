package server

import "sync"

type RequestParams struct {
	Mult1 int    `json:"int1"`
	Mult2 int    `json:"int2"`
	Limit int    `json:"limit"`
	Fuzz  string `json:"str1"`
	Buzz  string `json:"str2"`
}

type RequestHit struct {
	RequestParams
	Counter int64 `json:"hits"`
}

type requestRepo struct {
	requestStats map[RequestParams]int64
	topRequest   RequestHit
	sync.RWMutex
}

func NewRequestRepo() *requestRepo {
	return &requestRepo{
		requestStats: make(map[RequestParams]int64),
	}
}

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

func (r *requestRepo) TopRequest() RequestHit {
	r.RLock()
	defer r.RUnlock()
	return r.topRequest
}

var defaultRequestRepo = NewRequestRepo()

func Accept(mult1, mult2, limit int, fuzz, buzz string) {
	defaultRequestRepo.Accept(RequestParams{
		Mult1: mult1,
		Mult2: mult2,
		Limit: limit,
		Fuzz:  fuzz,
		Buzz:  buzz,
	})
}

func TopHitRequest() RequestHit {
	return defaultRequestRepo.TopRequest()
}
