package service

import (
	"app/internal/ticket"
	"context"
)

// ServiceTicketDefault represents the default service of the tickets
type ServiceTicketDefault struct {
	// rp represents the repository of the tickets
	rp ticket.RepositoryTicket
}

// NewServiceTicketDefault creates a new default service of the tickets
func NewServiceTicketDefault(rp ticket.RepositoryTicket) *ServiceTicketDefault {
	return &ServiceTicketDefault{
		rp: rp,
	}
}

// GetTotalTickets returns the total number of tickets
// No se para que esta esta func en el service
// func (s *ServiceTicketDefault) GetTotalTickets() (total int, err error) {
// 	ctx := context.Background()

// 	tickets, err := s.rp.Get(ctx)

// 	if err != nil {
// 		return 0, ticket.ErrRepositoryGetTickets
// 	}

// 	return len(tickets), nil
// }

func (s *ServiceTicketDefault) GetTotalAmountTickets() (total int, err error) {
	ctx := context.Background()

	tickets, err := s.rp.Get(ctx)

	if err != nil {
		return 0, ticket.ErrRepositoryGetTickets
	}

	return len(tickets), nil
}

func (s *ServiceTicketDefault) GetTicketsAmountByDestinationCountry(destination string) (amount int, err error) {
	ctx := context.Background()

	tickets, err := s.rp.GetTicketByDestinationCountry(ctx, destination)

	if err != nil {
		return 0, ticket.ErrRepositoryGetTickets
	}

	return len(tickets), nil
}

func (s *ServiceTicketDefault) GetPercentageTicketsByDestinationCountry(destination string) (percentage float64, err error) {
	ctx := context.Background()

	tickets, err := s.rp.GetTicketByDestinationCountry(ctx, destination)

	if err != nil {
		return 0, ticket.ErrRepositoryGetTickets
	}

	if len(tickets) == 0 {
		return 0, nil
	}

	count := len(tickets)
	acum := 0.0

	for _, ticket := range tickets {
		acum += ticket.Price
	}

	return acum / float64(count), nil
}
