package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/mtebourbi/lbc-fizzbuzz/pkg/fizzbuzz"
	"github.com/sirupsen/logrus"
)

// ListenAndServe start the fizzbuzz web service.
func ListenAndServe(listenPort int) error {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger) // FIXME: Must provide a logrus middleware for coherent logging output.
	r.Get("/fizzbuzz", fizzBuzzHandler)
	r.Get("/tophits", topRequestHandler)
	// Healthcheck resource.
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong\n"))
	})

	return http.ListenAndServe(fmt.Sprintf(":%d", listenPort), r)
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
		logrus.WithError(err).Error("server: int1 query string")
		renderBadRequest(w)
		return
	}
	mult2, err := strconv.Atoi(sInt2)
	if err != nil {
		logrus.WithError(err).Error("server: int2 query string")
		renderBadRequest(w)
		return
	}
	limit, err := strconv.Atoi(sLimit)
	if err != nil {
		logrus.WithError(err).Error("server: limit query string")
		renderBadRequest(w)
		return
	}

	Accept(mult1, mult2, limit, fuzz, buzz)

	res, err := fizzbuzz.FizzBuzz(mult1, mult2, limit, fuzz, buzz)
	if err != nil {
		logrus.WithFields(
			logrus.Fields{
				"mult1": mult1,
				"mult2": mult2,
				"limit": limit,
				"fuzz":  fuzz,
				"buzz":  buzz,
			}).WithError(err).Error("server: generating fizzbuzz")
		renderBadRequest(w)
		return
	}

	w.Write([]byte(strings.Join(res, ",") + "\n"))
}

func topRequestHandler(w http.ResponseWriter, r *http.Request) {
	b, err := json.Marshal(TopHitRequest())
	if err != nil {
		logrus.WithError(err).Error("server: generating top hit response")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func renderBadRequest(w http.ResponseWriter) {
	http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
}
