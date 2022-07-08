package api

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type Multier interface {
	Mul2(int) int
}

type api struct {
	r       *mux.Router
	service Multier
}

func New(service Multier) *api {
	r := mux.NewRouter()

	return &api{service: service, r: r}
}

func (a *api) handle(w http.ResponseWriter, r *http.Request) {
	num, err := strconv.Atoi(mux.Vars(r)["num"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(nil)
		return
	}
	res := a.service.Mul2(num)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(strconv.Itoa(res)))
}

func (a *api) Run() error {
	a.r.HandleFunc("/{num}", a.handle)
	return http.ListenAndServe("localhost:8080", a.r)
}

type API interface {
	Run() error
}

var _ API = &api{}
