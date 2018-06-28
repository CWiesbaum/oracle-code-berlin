package main

import (
	"context"
	"encoding/json"
	"io"

	fdk "github.com/fnproject/fdk-go"
)

func main() {
	fdk.Handle(fdk.HandlerFunc(extractEmail))
}

type BookingRecords []struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Gender    string `json:"gender"`
	JobTitle  string `json:"job_title"`
	Contact   struct {
		Email   string `json:"email"`
		Address struct {
			Street       string `json:"street"`
			StreetNumber string `json:"street_number"`
			PostalCode   string `json:"postal_code"`
			City         string `json:"city"`
		} `json:"address"`
	} `json:"contact"`
	Booking struct {
		ID        int     `json:"id"`
		ItemID    int     `json:"item_id"`
		ItemTitle string  `json:"item_title"`
		Quantity  int     `json:"quantity"`
		Price     float64 `json:"price"`
		Currency  string  `json:"currency"`
	} `json:"booking"`
}

func extractEmail(ctx context.Context, in io.Reader, out io.Writer) {
	var bookingRecords BookingRecords

	json.NewDecoder(in).Decode(&bookingRecords)

	var emailAddresses []string = make([]string, 0)

	for _, bookingRecord := range bookingRecords {
		emailAddresses = append(emailAddresses, bookingRecord.Contact.Email)
	}

	json.NewEncoder(out).Encode(&emailAddresses)
}
