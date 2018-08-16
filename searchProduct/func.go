package main

import (
	"context"
	"encoding/json"
	"io"
	"strings"

	fdk "github.com/fnproject/fdk-go"
)

func main() {
	fdk.Handle(fdk.HandlerFunc(findBookingRecordsByProduct))
}

type BookingRecordsStruct struct {
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

func findBookingRecordsByProduct(ctx context.Context, in io.Reader, out io.Writer) {
	var fnctx = fdk.Context(ctx)

	var productName = fnctx.Header.Get("ProductName")
	var bookingRecords []BookingRecordsStruct
	var resultBookingRecords []BookingRecordsStruct = make([]BookingRecordsStruct, 0)

	json.NewDecoder(in).Decode(&bookingRecords)

	for _, bookingRecord := range bookingRecords {
		if strings.Contains(bookingRecord.Booking.ItemTitle, productName) {
			resultBookingRecords = append(resultBookingRecords, bookingRecord)
		}
	}

	json.NewEncoder(out).Encode(&resultBookingRecords)
}
