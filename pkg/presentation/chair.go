package presentation

import (
	"encoding/json"
	"net/http"

	"github.com/44smkn/sqlc-sample/pkg/presentation/param"
	"github.com/44smkn/sqlc-sample/pkg/usecase"
	"github.com/go-chi/chi"
)

func getChairDetail(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	detail, err := usecase.GetChairDetail(r.Context(), id)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
	}
	b, err := json.Marshal(detail)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	w.Write(b)
}

func postChair(w http.ResponseWriter, r *http.Request) {
	var p param.PostChairParam
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = usecase.PostChair(r.Context(), p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func searchChair(w http.ResponseWriter, r *http.Request) {

}
