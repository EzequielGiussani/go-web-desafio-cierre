package handler_test

import (
	"app/internal/ticket"
	"app/internal/ticket/handler"
	"app/internal/ticket/repository"
	"app/internal/ticket/service"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"
)

func TestGetTicketsByCountry(t *testing.T) {

	t.Run("should return 200 and 1 ticket", func(t *testing.T) {
		// Arrange

		//DB
		db := map[int]ticket.TicketAttributes{
			1: {Name: "John", Email: "test@test.com", Country: "Argentina", Hour: "17:05", Price: 1000},
		}
		//Repository
		rp := repository.NewRepositoryTicketMap(db, len(db))
		// - service
		sv := service.NewServiceTicketDefault(rp)
		hd := handler.NewDefaultTicketHandler(sv)
		hdFunc := hd.GetTicketsByCountry()
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("dest", "Argentina")
		req := httptest.NewRequest("GET", "/ticket/getByCountry/Argentina", nil)
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))
		res := httptest.NewRecorder()

		// Act

		hdFunc(res, req)

		// Assert
		expectedStatusCode := http.StatusOK
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		expectedBody := `{"Message": "Tickets by destination", "Total":   1}`

		require.Equal(t, expectedStatusCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
		require.Equal(t, expectedHeader, res.Header())
	})

	t.Run("should return 200 and 0 tickets", func(t *testing.T) {
		// Arrange

		//DB
		db := map[int]ticket.TicketAttributes{}
		//Repository
		rp := repository.NewRepositoryTicketMap(db, len(db))
		// - service
		sv := service.NewServiceTicketDefault(rp)
		hd := handler.NewDefaultTicketHandler(sv)
		hdFunc := hd.GetTicketsByCountry()
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("dest", "Argentina")
		req := httptest.NewRequest("GET", "/ticket/getByCountry/Argentina", nil)
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))
		res := httptest.NewRecorder()

		// Act

		hdFunc(res, req)

		// Assert
		expectedStatusCode := http.StatusOK
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		expectedBody := `{"Message": "Tickets by destination", "Total":   0}`

		require.Equal(t, expectedStatusCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
		require.Equal(t, expectedHeader, res.Header())
	})

}

func TestGetTicketsAverageByCountry(t *testing.T) {

	t.Run("should return 200 and 1500 average", func(t *testing.T) {
		// Arrange

		//DB
		db := map[int]ticket.TicketAttributes{
			1: {Name: "John", Email: "test@test.com", Country: "Argentina", Hour: "17:05", Price: 1000},
			2: {Name: "Lucas", Email: "test@test.com", Country: "Argentina", Hour: "17:03", Price: 2000},
			3: {Name: "Juan", Email: "test@test.com", Country: "Venezuela", Hour: "11:05", Price: 888},
		}
		//Repository
		rp := repository.NewRepositoryTicketMap(db, len(db))
		// - service
		sv := service.NewServiceTicketDefault(rp)
		hd := handler.NewDefaultTicketHandler(sv)
		hdFunc := hd.GetTicketsAverageByCountry()
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("dest", "Argentina")
		req := httptest.NewRequest("GET", "/ticket/getAverage/Argentina", nil)
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))
		res := httptest.NewRecorder()

		// Act

		hdFunc(res, req)

		// Assert
		expectedStatusCode := http.StatusOK
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		expectedBody := `{"Message": "Tickets price average by destination", "Average": 1500}`

		require.Equal(t, expectedStatusCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
		require.Equal(t, expectedHeader, res.Header())
	})

	t.Run("should return 200 and 0 average", func(t *testing.T) {
		// Arrange

		//DB
		db := map[int]ticket.TicketAttributes{}
		//Repository
		rp := repository.NewRepositoryTicketMap(db, len(db))
		// - service
		sv := service.NewServiceTicketDefault(rp)
		hd := handler.NewDefaultTicketHandler(sv)
		hdFunc := hd.GetTicketsAverageByCountry()
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("dest", "Argentina")
		req := httptest.NewRequest("GET", "/ticket/getAverage/Argentina", nil)
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))
		res := httptest.NewRecorder()

		// Act

		hdFunc(res, req)

		// Assert
		expectedStatusCode := http.StatusOK
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		expectedBody := `{"Message": "Tickets price average by destination", "Average": 0}`

		require.Equal(t, expectedStatusCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
		require.Equal(t, expectedHeader, res.Header())
	})

}
