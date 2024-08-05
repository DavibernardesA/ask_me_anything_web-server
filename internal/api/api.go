package api

import (
	"net/http"

	"github.com/DavibernardesA/ask_me_anything_web-server/internal/store/pgstore"
	"github.com/go-chi/chi/v5"
)

type apiHandler struct {
	q *pgstore.Queries
	r *chi.Mux
}

// Use um receiver de ponteiro para o método ServeHTTP
func (h *apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.r.ServeHTTP(w, r)
}

func NewHandler(q *pgstore.Queries) http.Handler {
	a := &apiHandler{
		q: q,
		r: chi.NewRouter(),
	}

	return a
}
