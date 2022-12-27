package migration

import (
	"context"
	"fmt"
)

// CreateCabinBookingDetailTable ...
func CreateCabinBookingDetailTable() {

	r, err := DbPool.Exec(context.Background(), `CREATE TABLE IF NOT EXISTS cabin_booking_details (
    id serial PRIMARY KEY,
		cabin_booking_id INTEGER REFERENCES cabin_bookings (id),
		workspace_id INTEGER REFERENCES workspaces (id),
		floor_id INTEGER REFERENCES floors (id),
		city_id INTEGER REFERENCES cities (id),
		building_id INTEGER REFERENCES buildings (id),
		booked_by INTEGER REFERENCES users (id),
		cancelled_by INTEGER REFERENCES users (id),
		booking_date DATE NOT NULL,
		comments text,
		booking_slot_type VARCHAR ( 255 ),
		booking_slot_time INTEGER,
		active boolean NOT NULL DEFAULT TRUE,
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
