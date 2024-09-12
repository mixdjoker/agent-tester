package api

import (
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/mixdjoker/agent-tester/internal/logger"
	"github.com/mixdjoker/agent-tester/internal/service"
)

// Mux ...
type Mux struct {
	http.ServeMux
	ts service.TesterService
}

// NewMux ...
func NewMux(ts service.TesterService) *Mux {
	m := Mux{
		ServeMux: *http.NewServeMux(),
		ts:       ts,
	}

	m.HandleFunc("POST /updates", m.update)
	m.HandleFunc("POST /update/{id}", m.updateByID)
	m.HandleFunc("GET /results", m.results)

	return &m
}

func (m *Mux) update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	amount, err := m.ts.Update(ctx)
	if err != nil {
		logger.Err(ctx, err).Msg("update failed")
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	logger.Info(ctx).Int64("amount", amount).Msg("Tests updated")
	w.WriteHeader(http.StatusOK)
}

func (m *Mux) updateByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := r.PathValue("id")
	id, err := uuid.FromString(idStr)
	if err != nil {
		logger.Err(ctx, err).Msg("wrong ID")
		http.Error(w, "wrong ID", http.StatusBadRequest)
		return
	}

	err = m.ts.UpdateByID(ctx, id)
	if err != nil {
		logger.Err(ctx, err).Msg("update by id failed")
		http.Error(w, "failed to update by ID", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (m *Mux) results(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	res, err := m.ts.GetResults(ctx)
	if err != nil {
		logger.Err(ctx, err).Msg("failed to get results")
		http.Error(w, "failed to get results", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(res); err != nil {
		logger.Err(ctx, err).Msg("failed write resonse")
		return
	}
}
