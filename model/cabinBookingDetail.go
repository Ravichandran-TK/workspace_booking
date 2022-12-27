package model

import (
	"context"
	"time"
	"workspace_booking/migration"
)

type CabinBookingDetails struct {
	Id              int       `json:"id"`
	CabinBookingId  int       `json:"cabin_booking_id"`
	WorkspaceId     int       `json:"workspace_id"`
	CityId          int       `json:"city_id"`
	BuildingId      int       `json:"building_id"`
	FloorId         int       `json:"floor_id"`
	CityName        string    `json:"city_name"`
	BuildingName    string    `json:"building_name"`
	FloorName       string    `json:"floor_name"`
	CabinName       string    `json:"cabin_name"`
	BookingDate     string    `json:"booking_date"`
	BookingSlotType string    `json:"booking_slot_type"`
	BookingSlotTime int       `json:"booking_slot_time"`
	BookedBy        int       `json:"booked_by"`
	CancelledBy     int       `json:"cancelled_by"`
	Comments        string    `json:"comments"`
	Active          bool      `json:"active"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	CancelledAt     time.Time `json:"cancelled_at"`
}

func BulkInsertCabinBookingDetails(cabin *CabinBooking) error {
	CabinBookingDetailsRecords := make([]*CabinBookingDetails, 0)
	for _, booking := range cabin.CabinBookingDetails {
		cabinBookingDetail := new(CabinBookingDetails)
		cabinBookingDetail.CabinBookingId = cabin.Id
		cabinBookingDetail.WorkspaceId = booking.WorkspaceId
		cabinBookingDetail.CityId = booking.FloorId
		cabinBookingDetail.BuildingId = booking.FloorId
		cabinBookingDetail.FloorId = booking.FloorId
		cabinBookingDetail.BookingDate = booking.BookingDate
		cabinBookingDetail.BookingSlotType = booking.BookingSlotType
		cabinBookingDetail.BookingSlotTime = booking.BookingSlotTime
		cabinBookingDetail.BookedBy = booking.BookedBy
		cabinBookingDetail.Comments = booking.Comments
		err := cabinBookingDetail.CreateCabinBookingDetail()
		CabinBookingDetailsRecords = append(CabinBookingDetailsRecords, cabinBookingDetail)
		if err != nil {
			return err
		}
	}
	cabin.CabinBookingDetails = CabinBookingDetailsRecords
	return nil
}

func (cabinBookingDetail *CabinBookingDetails) CreateCabinBookingDetail() error {
	dt := time.Now()
	query := "INSERT INTO cabin_booking_details (cabin_booking_id, workspace_id, city_id, building_id, floor_id, booking_date, booking_slot_type, booking_slot_time, booked_by, cancelled_by, comments, active, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14) RETURNING id, created_at, updated_at, (select cities.name from cities where id=cabin_booking_details.city_id) as city_name, (select buildings.name from buildings where id=cabin_booking_details.building_id) as building_name, (select floors.name from floors where id=cabin_booking_details.floor_id) as floor_name, (select workspaces.name from workspaces where id=cabin_booking_details.workspace_id) as cabin_name"
	d := migration.DbPool.QueryRow(context.Background(), query, cabinBookingDetail.CabinBookingId, cabinBookingDetail.WorkspaceId, cabinBookingDetail.CityId, cabinBookingDetail.BuildingId, cabinBookingDetail.FloorId, cabinBookingDetail.BookingDate, cabinBookingDetail.BookingSlotType, cabinBookingDetail.BookingSlotTime, cabinBookingDetail.BookedBy, nil, cabinBookingDetail.Comments, true, dt, dt)

	err := d.Scan(&cabinBookingDetail.Id, &cabinBookingDetail.CreatedAt, &cabinBookingDetail.UpdatedAt, &cabinBookingDetail.CityName, &cabinBookingDetail.BuildingName, &cabinBookingDetail.FloorName, &cabinBookingDetail.CabinName)
	if err != nil {
		return err
	}
	return nil
}
