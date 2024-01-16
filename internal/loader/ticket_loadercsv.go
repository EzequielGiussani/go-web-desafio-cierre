package loader

import (
	internal "app/internal/ticket"
	"encoding/csv"
	"errors"
	"io"
	"os"
	"strconv"
)

var (
	ErrorOpeningFIle     = errors.New("error opening file")
	ErrorReadingRecord   = errors.New("error reading record")
	ErrorConvertingID    = errors.New("error converting id to int")
	ErrorConvertingPrice = errors.New("error converting price to float")
)

// NewLoaderTicketCSV creates a new ticket loader from a CSV file
func NewLoaderTicketCSV(filePath string) *LoaderTicketCSV {
	return &LoaderTicketCSV{
		filePath: filePath,
	}
}

// LoaderTicketCSV represents a ticket loader from a CSV file
type LoaderTicketCSV struct {
	filePath string
}

// Load loads the tickets from the CSV file
func (t *LoaderTicketCSV) Load() (ta map[int]internal.TicketAttributes, e error) {
	// open the file
	f, err := os.Open(t.filePath)
	if err != nil {
		e = ErrorOpeningFIle
		return
	}
	defer f.Close()

	// read the file
	r := csv.NewReader(f)

	// read the records
	ta = make(map[int]internal.TicketAttributes)
	for {
		record, err := r.Read()
		if err != nil {
			if err == io.EOF {
				break
			}

			e = ErrorReadingRecord
			return
		}

		// serialize the record
		id, err := strconv.Atoi(record[0])

		if err != nil {
			e = ErrorConvertingID
			return
		}

		floatPrice, err := strconv.ParseFloat(record[5], 64)

		if err != nil {
			e = ErrorConvertingPrice
			return
		}

		ticket := internal.TicketAttributes{
			Name:    record[1],
			Email:   record[2],
			Country: record[3],
			Hour:    record[4],
			Price:   floatPrice,
		}

		// add the ticket to the map
		ta[id] = ticket
	}

	return
}
