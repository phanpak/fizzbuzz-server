package fizzbuzz

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

type Server struct {
	router     *http.ServeMux
	hitCounter *Counter
}

func NewServer() *Server {
	return &Server{
		router:     http.NewServeMux(),
		hitCounter: NewCounter(),
	}
}

func (s *Server) Serve() {
	s.registerRoutes()
	log.Println("Server is listening on port 8080")
	http.ListenAndServe(":8080", s.router)
}

func (s *Server) registerRoutes() {
	s.router.HandleFunc("/fizzbuzz", s.WithStats(s.FizzBuzzHandler()))
	s.router.HandleFunc("/stats", s.statsHandler())
}

type Request struct {
	Int1  int    `json:"int1"`
	Int2  int    `json:"int2"`
	Limit int    `json:"limit"`
	Str1  string `json:"str1"`
	Str2  string `json:"str2"`
}

type Response struct {
	Result []string `json:"result"`
}

func (s *Server) FizzBuzzHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var req Request
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		res, err := Do(DoParams{
			int1:  req.Int1,
			int2:  req.Int2,
			limit: req.Limit,
			str1:  req.Str1,
			str2:  req.Str2,
		})
		if err != nil {
			// We're directly propagating errors from an internal method to the client which is not ideal
			if errors.Is(err, ErrInvalidInput) {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		resultArray := make([]string, len(res))
		for i, v := range res {
			resultArray[i] = v
		}

		resp := Response{Result: resultArray}
		json.NewEncoder(w).Encode(resp)
	}
}

func (s *Server) statsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		max := s.hitCounter.Max()

		resp := struct {
			Hits       int             `json:"hits"`
			Parameters json.RawMessage `json:"parameters"`
		}{
			Hits:       max.Count,
			Parameters: []byte(max.Key),
		}

		json.NewEncoder(w).Encode(resp)
	}
}

func (s *Server) WithStats(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requestBody, _ := ioutil.ReadAll(r.Body)
		r.Body.Close() //  must close
		r.Body = ioutil.NopCloser(bytes.NewBuffer(requestBody))
		next.ServeHTTP(w, r)

		// Convert the request body to a string
		requestBodyStr := string(requestBody)
		// A major shortcoming of this implementation is that the key used is the raw request body.
		// It will not take into account the case where the same parameters are passed in a different order.
		// Another issue is that extra keys passed will mean extra keys in the map.
		s.hitCounter.Increment(requestBodyStr)
	}
}
