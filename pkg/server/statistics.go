package server

import "sync"

type requestRepo struct {
	requestStats map[RequestParams]int64
	sync.RWMutex
}

func NewRequestRepo() requestRepo {
	return requestRepo{
		requestStats: make(map[RequestParams]int64),
	}
}

func (r *requestRepo) Accept(rp RequestParams) {
	r.Lock()
	defer r.Unlock()
	r.requestStats[rp]++
}

func (r *requestRepo) TopRequest() RequestHit {
	var req RequestParams
	var hits int64
	r.RLock()
	defer r.RUnlock()
	for r, h := range r.requestStats {
		if h > hits {
			req = r
			hits = h
		}
	}

	return RequestHit{
		Counter:       hits,
		RequestParams: req,
	}
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
