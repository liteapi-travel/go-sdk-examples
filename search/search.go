package search

import (
	"context"
	"fmt"
	config "go-sdk-example/core"
	"os"

	"github.com/gin-gonic/gin"
)

func GetMinimumRates(c *gin.Context) {
	hotelIds := "lp3803c,lp1f982,lp19b70,lp19e75"
	checkin := "2023-11-15"
	checkout := "2023-11-16"
	currency := "USD"
	guestNationality := "US"
	adults := int32(1)

	result, res, err := config.ApiClient().SearchApi.GetMinimumRates(context.Background()).HotelIds(hotelIds).Checkin(checkin).Checkout(checkout).Currency(currency).GuestNationality(guestNationality).Adults(adults).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.GetMinimumRates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", res)
	}
	c.JSON(200, result)
}

func GetFullRates(c *gin.Context) {
	hotelIds := "lp3803c,lp1f982,lp19b70,lp19e75"
	checkin := "2023-11-15"
	checkout := "2023-11-16"
	currency := "USD"
	guestNationality := "US"
	adults := int32(1)

	result, res, err := config.ApiClient().SearchApi.GetFullRates(context.Background()).HotelIds(hotelIds).Checkin(checkin).Checkout(checkout).Currency(currency).GuestNationality(guestNationality).Adults(adults).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.GetFullRates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", res)
	}
	c.JSON(200, result)
}
