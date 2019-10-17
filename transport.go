package main

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

// QuestionRequest struct
type QuestionRequest struct {
	A int `json:"a"`
	B int `json:"b"`
}

// AnswerResponse struct
type AnswerResponse struct {
	V   int    `json:"v"`
	Err string `json:"err,omitempty"`
}

func makeAdditionEndpoint(svc NumberService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(QuestionRequest)
		v, err := svc.Add(req.A, req.B)
		if err != nil {
			return AnswerResponse{v, err.Error()}, err
		}
		return AnswerResponse{v, ""}, err
	}
}

func makeSubstractionEndpoint(svc NumberService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(QuestionRequest)
		v, err := svc.Substract(req.A, req.B)
		if err != nil {
			return AnswerResponse{v, err.Error()}, err
		}
		return AnswerResponse{v, ""}, err
	}
}

func makeMultiplyEndpoint(svc NumberService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(QuestionRequest)
		v, err := svc.Multiply(req.A, req.B)
		if err != nil {
			return AnswerResponse{v, err.Error()}, err
		}
		return AnswerResponse{v, ""}, err
	}
}

func makeDivideEndpoint(svc NumberService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(QuestionRequest)
		v, err := svc.Divide(req.A, req.B)
		if err != nil {
			return AnswerResponse{v, err.Error()}, err
		}
		return AnswerResponse{v, ""}, err
	}
}

// MakeUnauthenticatedEndPoint end point for unauthenticated services
func MakeUnauthenticatedEndPoint() endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		return AnswerResponse{0, "Unauthenticated"}, errors.New("Unauthenticated")
	}
}

func decodeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request QuestionRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
