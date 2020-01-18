package server

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/mtebourbi/lbc-fizzbuzz/pkg/fizzbuzz"
)

// ListenAndServe start the fizzbuzz web service.
func ListenAndServe() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Get("/fizzbuzz", fizzBuzzHandler)

	http.ListenAndServe(":8080", r)
}

func fizzBuzzHandler(w http.ResponseWriter, r *http.Request) {
	sInt1 := r.URL.Query().Get("int1")
	sInt2 := r.URL.Query().Get("int2")
	sLimit := r.URL.Query().Get("limit")
	str1 := r.URL.Query().Get("str1")
	str2 := r.URL.Query().Get("str2")

	int1, err := strconv.Atoi(sInt1)
	if err != nil {
		renderBadRequest(w)
		return
	}
	int2, err := strconv.Atoi(sInt2)
	if err != nil {
		renderBadRequest(w)
		return
	}
	limit, err := strconv.Atoi(sLimit)
	if err != nil {
		renderBadRequest(w)
		return
	}

	res, err := fizzbuzz.FizzBuzz(int1, int2, limit, str1, str2)
	if err != nil {
		renderBadRequest(w)
		return
	}

	w.Write([]byte(strings.Join(res, ",") + "\n"))
}

func renderBadRequest(w http.ResponseWriter) {
	http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
}
