package main

import (
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	svc := numberService{}

	logger := log.NewLogfmtLogger(os.Stderr)

	additionEndpoint := LoggingMiddleware(log.With(logger, "method", "add"))(makeAdditionEndpoint(svc))

	additionHandler := TestAuthMiddleware(log.With(logger, "method", "add"))(httptransport.NewServer(
		additionEndpoint,
		decodeRequest,
		encodeResponse,
	))

	substractionHandler := httptransport.NewServer(
		makeSubstractionEndpoint(svc),
		decodeRequest,
		encodeResponse,
	)

	multiplyHandler := httptransport.NewServer(
		makeMultiplyEndpoint(svc),
		decodeRequest,
		encodeResponse,
	)

	divideHandler := httptransport.NewServer(
		makeDivideEndpoint(svc),
		decodeRequest,
		encodeResponse,
	)

	r.Handle("/add", additionHandler).Methods("POST")
	r.Handle("/substract", substractionHandler).Methods("POST")
	r.Handle("/multiply", multiplyHandler).Methods("POST")
	r.Handle("/divide", divideHandler).Methods("POST")
	http.ListenAndServe(":8000", r)
}
