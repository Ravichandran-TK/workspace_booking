package controller

import (
	"workspace_booking/mailer"
	"workspace_booking/model"
	"workspace_booking/utility"

	"github.com/gofiber/fiber/v2"
)

// CreateCabinBooking handler
func CreateCabinBooking(c *fiber.Ctx) error {
	cabin := new(model.CabinBooking)

	if err := c.BodyParser(cabin); err != nil {
		return utility.ErrResponse(c, "Error in body parsing", 400, err)
	}
	// query := "SELECT SUM(booking_slot_time) FROM cabin_booking_details where booked_by = $1 and cancelled_by = $2 and booking_date between $3 and $4 GROUP BY cabin_booking_id"
	// existingCabinBookingsTotalBookedHrs, err1 := migration.DbPool.Query(context.Background(), query)
	// fmt.Println(existingCabinBookingsTotalBookedHrs)

	// if err1 != nil {
	// 	return utility.ErrResponse(c, "Error in cabin booking creation", 500, err1)
	// }
	// existingCabinBookingsTotalBookedHrs, err := strconv.Atoi(existingCabinBookingsTotalBookedHrs)
	// allowedTimePerWeek := 24
	// if false {
	// 	return utility.ErrResponse(c, "You are not allowed to book more than 24hrs in a week", 500, err1)
	// }

	err := cabin.InsertCabinBooking()
	if err != nil {
		return utility.ErrResponse(c, "Error in cabin booking creation", 500, err)
	}

	err = model.BulkInsertCabinBookingDetails(cabin)
	if err != nil {
		return utility.ErrResponse(c, "Error in cabin booking details creation", 500, err)
	}

	if cabin.Id != 0 {
		go mailer.CabinBookingMailer(cabin, false)
	}

	if err := c.JSON(&fiber.Map{
		"success": true,
		"message": "Cabin booking has been booked successfully.",
		"cabin":   cabin,
	}); err != nil {
		return utility.ErrResponse(c, "Error in response", 500, err)
	}
	return nil
}
