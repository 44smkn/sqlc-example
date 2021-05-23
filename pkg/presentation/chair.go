package presentation

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/44smkn/sqlc-sample/pkg/presentation/param"
	"github.com/44smkn/sqlc-sample/pkg/usecase"
	"github.com/go-chi/chi"
)

func getChairDetail(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	detail, err := usecase.GetChairDetail(r.Context(), id)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	b, err := json.Marshal(detail)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
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
	q := r.URL.Query()
	p := param.SearchChairParam{
		PriceRangeID:  queryVal(q, "price-range-id"),
		HeightRangeID: queryVal(q, "height-range-id"),
		WidthRangeID:  queryVal(q, "width-range-id"),
		DepthRangeID:  queryVal(q, "depth-range-id"),
		Kind:          queryVal(q, "kind"),
		Color:         queryVal(q, "color"),
		Features:      queryVal(q, "features"),
	}
	log.Sugar().Infof("param: %+v", p)
	data, err := usecase.SearchChair(r.Context(), p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	b, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(b)
}

func queryVal(q url.Values, key string) string {
	if v, ok := q[key]; ok {
		return v[0]
	} else {
		return ""
	}
}
