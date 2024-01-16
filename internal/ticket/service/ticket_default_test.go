package service_test

import (
	"app/internal/ticket"
	"app/internal/ticket/repository"
	"app/internal/ticket/service"
	"testing"

	"github.com/stretchr/testify/require"
)

// Tests for ServiceTicketDefault.GetTotalAmountTickets
// func TestServiceTicketDefault_GetTotalAmountTickets_Mock(t *testing.T) {
// 	t.Run("success to get total tickets", func(t *testing.T) {
// 		// arrange
// 		// - repository: mock
// 		rp := repository.NewRepositoryTicketMock()
// 		// - repository: set-up
// 		rp.FuncGet = func() (t map[int]ticket.TicketAttributes, err error) {
// 			t = map[int]ticket.TicketAttributes{
// 				1: {
// 					Name:    "John",
// 					Email:   "johndoe@gmail.com",
// 					Country: "USA",
// 					Hour:    "10:00",
// 					Price:   100,
// 				},
// 			}
// 			return
// 		}

// 		// - service
// 		sv := service.NewServiceTicketDefault(rp)

// 		// act
// 		total, err := sv.GetTotalAmountTickets()

// 		// assert
// 		expectedTotal := 1
// 		require.NoError(t, err)
// 		require.Equal(t, expectedTotal, total)
// 	})
// }

func TestServiceTicketDefault_GetTotalAmountTickets(t *testing.T) {
	t.Run("success to get total tickets", func(t *testing.T) {
		// arrange
		//db
		db := map[int]ticket.TicketAttributes{
			1: {Name: "John", Email: "test@test.com", Country: "Argentina", Hour: "17:05", Price: 1000},
		}

		//Repository
		rp := repository.NewRepositoryTicketMap(db, len(db))

		// - service
		sv := service.NewServiceTicketDefault(rp)

		// act
		total, err := sv.GetTotalAmountTickets()

		// assert
		expectedTotal := 1
		require.NoError(t, err)
		require.Equal(t, expectedTotal, total)
	})

	t.Run("should return no tickets", func(t *testing.T) {
		// arrange
		//db
		db := map[int]ticket.TicketAttributes{}

		//Repository
		rp := repository.NewRepositoryTicketMap(db, len(db))

		// - service
		sv := service.NewServiceTicketDefault(rp)

		// act
		total, err := sv.GetTotalAmountTickets()

		// assert
		expectedTotal := 0
		require.NoError(t, err)
		require.Equal(t, expectedTotal, total)
	})

}

func TestServiceTicketDefault_GetTicketsAmountByDestinationCountry(t *testing.T) {
	t.Run("success to get tickets amount", func(t *testing.T) {
		// arrange
		//db
		db := map[int]ticket.TicketAttributes{
			1: {Name: "John", Email: "test@test.com", Country: "Argentina", Hour: "17:05", Price: 1000},
			2: {Name: "Lucas", Email: "test@test.com", Country: "Argentina", Hour: "17:03", Price: 999},
			3: {Name: "Juan", Email: "test@test.com", Country: "Venezuela", Hour: "11:05", Price: 888},
		}
		destination := "Argentina"

		//Repository
		rp := repository.NewRepositoryTicketMap(db, len(db))

		// - service
		sv := service.NewServiceTicketDefault(rp)

		// act
		total, err := sv.GetTicketsAmountByDestinationCountry(destination)

		// assert
		expectedTotal := 2
		require.NoError(t, err)
		require.Equal(t, expectedTotal, total)
	})

	t.Run("should return no tickets", func(t *testing.T) {
		// arrange
		//db
		db := map[int]ticket.TicketAttributes{
			1: {Name: "Juan", Email: "test@test.com", Country: "Venezuela", Hour: "11:05", Price: 888},
		}
		destination := "Argentina"

		//Repository
		rp := repository.NewRepositoryTicketMap(db, len(db))

		// - service
		sv := service.NewServiceTicketDefault(rp)

		// act
		total, err := sv.GetTicketsAmountByDestinationCountry(destination)

		// assert
		expectedTotal := 0
		require.NoError(t, err)
		require.Equal(t, expectedTotal, total)
	})

}

func TestServiceTicketDefault_GetPercentageTicketsByDestinationCountry(t *testing.T) {
	t.Run("success to get tickets amount", func(t *testing.T) {
		// arrange
		//db
		db := map[int]ticket.TicketAttributes{
			1: {Name: "John", Email: "test@test.com", Country: "Argentina", Hour: "17:05", Price: 1000},
			2: {Name: "Lucas", Email: "test@test.com", Country: "Argentina", Hour: "17:03", Price: 2000},
			3: {Name: "Juan", Email: "test@test.com", Country: "Venezuela", Hour: "11:05", Price: 888},
		}
		destination := "Argentina"

		//Repository
		rp := repository.NewRepositoryTicketMap(db, len(db))

		// - service
		sv := service.NewServiceTicketDefault(rp)

		// act
		total, err := sv.GetPercentageTicketsByDestinationCountry(destination)

		// assert
		expectedTotal := 1500.0
		require.NoError(t, err)
		require.Equal(t, expectedTotal, total)
	})

	t.Run("should return no tickets", func(t *testing.T) {
		// arrange
		//db
		db := map[int]ticket.TicketAttributes{}
		destination := "Argentina"

		//Repository
		rp := repository.NewRepositoryTicketMap(db, len(db))

		// - service
		sv := service.NewServiceTicketDefault(rp)

		// act
		total, err := sv.GetPercentageTicketsByDestinationCountry(destination)

		// assert
		expectedTotal := 0.0
		require.NoError(t, err)
		require.Equal(t, expectedTotal, total)
	})

}
