package handler

import (
	"app/internal/ticket"
	"app/platform/web/response"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type DefaultTicketHandler struct {
	sv ticket.ServiceTicket
}

type BodyRequestTicketJSON struct {
	Name string `json:"name"`
	// Email represents the email of the owner of the ticket
	Email string `json:"email"`
	// Country represents the destination country of the ticket
	Country string `json:"country"`
	// Hour represents the hour of the ticket
	Hour string `json:"hour"`
	// Price represents the price of the ticket
	Price float64 `json:"price"`
}

type BodyResponseTicketJSON struct {
	// Id represents the id of the ticket
	Id   int    `json:"id"`
	Name string `json:"name"`
	// Email represents the email of the owner of the ticket
	Email string `json:"email"`
	// Country represents the destination country of the ticket
	Country string `json:"country"`
	// Hour represents the hour of the ticket
	Hour string `json:"hour"`
	// Price represents the price of the ticket
	Price float64 `json:"price"`
}

func NewDefaultTicketHandler(sv ticket.ServiceTicket) *DefaultTicketHandler {
	return &DefaultTicketHandler{
		sv: sv,
	}
}

func (h *DefaultTicketHandler) GetTicketsByCountry() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		dest := chi.URLParam(r, "dest")

		amount, err := h.sv.GetTicketsAmountByDestinationCountry(dest)

		if err != nil {
			switch err {
			case ticket.ErrRepositoryGetTickets:
				response.Error(w, http.StatusInternalServerError, "Error getting tickets")
			}
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"Message": "Tickets by destination",
			"Total":   amount,
		})

	}
}

func (h *DefaultTicketHandler) GetTicketsAverageByCountry() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		dest := chi.URLParam(r, "dest")

		avg, err := h.sv.GetPercentageTicketsByDestinationCountry(dest)

		if err != nil {
			switch err {
			case ticket.ErrRepositoryGetTickets:
				response.Error(w, http.StatusInternalServerError, "Error getting tickets")
			}
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"Message": "Tickets price average by destination",
			"Average": avg,
		})

	}
}
