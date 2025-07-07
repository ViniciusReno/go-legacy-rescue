package handlers

import (
	"fmt"
	"myapp/services"
	"net/http"
)

type UserHandler struct {
	Service services.UserService
}

func NewUserHandler(s services.UserService) *UserHandler {
	return &UserHandler{Service: s}
}

func (h *UserHandler) Handle(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	fmt.Println("Handling user:", userID)

	err := h.Service.ProcessUser(r.Context(), userID)
	if err != nil {
		fmt.Println("Something went wrong:", err)
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("OK"))
}
