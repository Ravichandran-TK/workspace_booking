package model

import (
	"context"
	"time"
	"workspace_booking/migration"
)

type CabinBooking struct {
	Id                            int                    `json:"id"`
	BookingDates                  []string               `json:"booking_dates"`
	Purpose                       string                 `json:"purpose"`
	Active                        bool                   `json:"active"`
	CabinBookingDetails           []*CabinBookingDetails `json:"cabin_booking_details"`
	IsPartianCancellationHappened bool                   `json:"is_partian_cancellation_happened"`
	BookedBy                      int                    `json:"booked_by"`
	CancelledBy                   int                    `json:"cancelled_by"`
	CreatedAt                     time.Time              `json:"created_at"`
	UpdatedAt                     time.Time              `json:"updated_at"`
	CancelledAt                   time.Time              `json:"cancelled_at"`
}

type CabinBookings struct {
	CabinBookings []*CabinBooking
}

type BookedCabin struct {
	BookedDate time.Time `json:"date"`
	CabinIds   []int     `json:"cabin_ids"`
}

// InsertBooking will create the booking record in db
func (b *CabinBooking) InsertCabinBooking() error {
	dt := time.Now()
	query1 := "INSERT INTO cabin_bookings (booking_dates, purpose, booked_by, active, is_partian_cancellation_happened, created_at, updated_at) VALUES ($1, $2, $3, $4,$5, $6, $7) RETURNING id, created_at, updated_at"
	d := migration.DbPool.QueryRow(
		context.Background(), query1, b.BookingDates, b.Purpose, b.BookedBy, b.Active, false, dt, dt)
	err := d.Scan(&b.Id, &b.CreatedAt, &b.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}
