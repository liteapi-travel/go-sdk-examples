package bookingManagment

import (
	"context"
	"fmt"
	config "go-sdk-example/core"
	"os"

	"github.com/gin-gonic/gin"
)

func BookingList(c *gin.Context) {
	guestId := "FrT56hfty"
	result, res, err := config.ApiClient().BookingManagementApi.GetBookingListByGuestId(context.Background()).GuestId(guestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BookingManagementApi.GetBookingListByGuestId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", res)
	}
	c.JSON(200, result)
}

func BookingRetrieve(c *gin.Context) {
	bookingId := "8Hrr5KLM7"
	result, res, err := config.ApiClient().BookingManagementApi.RetrievedBooking(context.Background(), bookingId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BookingManagementApi.RetrievedBooking``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", res)
	}
	c.JSON(200, result)
}

func BookingCancel(c *gin.Context) {
	bookingId := "zATqnsCG7"
	result, res, err := config.ApiClient().BookingManagementApi.CancelBooking(context.Background(), bookingId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BookingManagementApi.CancelBooking``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", res)
	}
	c.JSON(200, result)
}
