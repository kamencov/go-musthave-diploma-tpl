package order

import (
	"encoding/json"
	"errors"
	"github.com/kamencov/go-musthave-diploma-tpl/internal/customerrors"
	"github.com/kamencov/go-musthave-diploma-tpl/internal/middleware"
	"net/http"
)

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	login, ok := r.Context().Value(middleware.LoginContentKey).(string)

	if !ok || login == "" {
		h.log.Error("Error post order = not userID")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// получаем список загруженных номеров заказов
	req, err := h.service.GetAllUserOrders(login)
	if err != nil {
		if errors.Is(err, customerrors.ErrNotFound) {
			h.log.Error("error get orders", "error: ", "no data to answer")
			apiError, _ := json.Marshal(customerrors.APIError{Message: "no data to answer"})
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNoContent)
			w.Write(apiError)
			return
		}

		h.log.Error("error get orders", "error: ", err)
		apiError, _ := json.Marshal(customerrors.APIError{Message: "cannot loading order"})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(apiError)
		return
	}

	if len(req) == 0 {
		h.log.Error("error get orders", "error: ", "no data to answer")
		apiError, _ := json.Marshal(customerrors.APIError{Message: "no data to answer"})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNoContent)
		w.Write(apiError)
		return
	}

	w.Header().Set("Content-Type", "application/json") // Перенесите перед w.WriteHeader
	w.WriteHeader(http.StatusOK)

	if err = json.NewEncoder(w).Encode(req); err != nil {
		h.log.Error("error get orders", "failed to marshal response: ", err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}
}
