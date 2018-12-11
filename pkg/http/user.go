package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tbaud0n/sample-api-architecture"
)

// UserQueryHandler handles the HTTP request querying the users
type UserQueryHandler struct {
	UserStorageService         api.UserStorageService
	QueryFilterFromHTTPRequest api.QueryFilterFromHTTPRequest
}

func (h UserQueryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	qf, err := h.QueryFilterFromHTTPRequest(r)
	if err != nil {
		Error(w, err)
		return
	}
	users, err := h.UserStorageService.Query(qf)
	if err != nil {
		Error(w, err)
		return
	}

	b, err := json.Marshal(users)
	if err != nil {
		Error(w, err)
		return
	}

	w.Write(b)
}

// UserGetHandler handles the GET request of user
type UserGetHandler struct {
	UserStorageService api.UserStorageService
}

func (h UserGetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.ParseInt(vars[`id`], 10, 64)
	if err != nil {
		Error(w, err)
		return
	}

	u, err := h.UserStorageService.Get(id)
	if err != nil {
		Error(w, err)
		return
	} else if u == nil {
		NotFound(w, r)
		return
	}

	b, err := json.Marshal(u)
	if err != nil {
		Error(w, err)
		return
	}

	w.Write(b)
}
