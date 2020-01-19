package server

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/mtebourbi/lbc-fizzbuzz/pkg/fizzbuzz"
)

// ListenAndServe start the fizzbuzz web service.
func ListenAndServe() error {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Get("/fizzbuzz", fizzBuzzHandler)
	r.Get("/tophits", topRequestHandler)

	return http.ListenAndServe(":8080", r)
}

func fizzBuzzHandler(w http.ResponseWriter, r *http.Request) {
	//FIXME: More clean way to get url query
	sInt1 := r.URL.Query().Get("int1")
	sInt2 := r.URL.Query().Get("int2")
	sLimit := r.URL.Query().Get("limit")
	fuzz := r.URL.Query().Get("str1")
	buzz := r.URL.Query().Get("str2")

	mult1, err := strconv.Atoi(sInt1)
	if err != nil {
		renderBadRequest(w)
		return
	}
	mult2, err := strconv.Atoi(sInt2)
	if err != nil {
		renderBadRequest(w)
		return
	}
	limit, err := strconv.Atoi(sLimit)
	if err != nil {
		renderBadRequest(w)
		return
	}

	Accept(mult1, mult2, limit, fuzz, buzz)

	res, err := fizzbuzz.FizzBuzz(mult1, mult2, limit, fuzz, buzz)
	if err != nil {
		renderBadRequest(w)
		return
	}

	w.Write([]byte(strings.Join(res, ",") + "\n"))
}

func topRequestHandler(w http.ResponseWriter, r *http.Request) {
	b, err := json.Marshal(TopHitRequest())
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func renderBadRequest(w http.ResponseWriter) {
	http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
}
