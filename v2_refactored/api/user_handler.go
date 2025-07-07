package api

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type UserHandler struct {
	Processor UserProcessor
	Logger    *slog.Logger
}

func NewUserHandler(p UserProcessor, logger *slog.Logger) *UserHandler {
	return &UserHandler{Processor: p, Logger: logger}
}

func (h *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID := r.URL.Query().Get("user_id")

	if userID == "" {
		h.Logger.Warn("missing user_id", "path", r.URL.Path)
		http.Error(w, "missing user_id", http.StatusBadRequest)
		return
	}

	result, err := h.Processor.Process(ctx, userID)
	if err != nil {
		h.Logger.Error("failed to process user", "user_id", userID, "error", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	response := map[string]string{"result": result}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
