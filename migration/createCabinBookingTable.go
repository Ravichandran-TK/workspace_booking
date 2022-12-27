package migration

import (
	"context"
	"fmt"
)

// CreateCabinBookingTable ...
func CreateCabinBookingTable() {

	r, err := DbPool.Exec(context.Background(), `CREATE TABLE IF NOT EXISTS cabin_bookings (
    id serial PRIMARY KEY,
		booked_by INTEGER REFERENCES users (id),
		cancelled_by INTEGER REFERENCES users (id),
		booking_dates DATE[],
		active boolean NOT NULL DEFAULT TRUE,
		purpose VARCHAR ( 255 ),
		is_partian_cancellation_happened boolean NOT NULL DEFAULT FALSE,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		cancelled_at TIMESTAMP,
		updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP)
	`)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r)
	}

}
