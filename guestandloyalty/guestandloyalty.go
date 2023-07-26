package guestAndLoyalty

import (
	"context"
	"fmt"
	config "go-sdk-example/core"
	"os"

	"github.com/gin-gonic/gin"
)

func GetGuestsIds(c *gin.Context) {

	email := "johndoe@nlite.ml"
	result, res, err := config.ApiClient().GuestAndLoyaltyApi.GetGuestsIds(context.Background()).Email(email).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GuestAndLoyaltyApi.GetGuestsIds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", res)
	}
	c.JSON(200, result)
}
