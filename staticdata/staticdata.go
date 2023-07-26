package staticData

import (
	"context"
	"fmt"
	config "go-sdk-example/core"
	"os"

	"github.com/gin-gonic/gin"
)

func GetCities(c *gin.Context) {

	countryCode := "US"
	result, res, err := config.ApiClient().StaticDataApi.GetCitiesByCountryCode(context.Background()).CountryCode(countryCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StaticDataApi.GetCitiesByCountryCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", res)
	}
	c.JSON(200, result)
}

func GetCountries(c *gin.Context) {
	result, res, err := config.ApiClient().StaticDataApi.GetCountries(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StaticDataApi.GetCountries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", res)
	}
	c.JSON(200, result)

}

func GetCurrencies(c *gin.Context) {
	result, res, err := config.ApiClient().StaticDataApi.GetCurrencies(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StaticDataApi.GetCurrencies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", res)
	}
	c.JSON(200, result)
}

func GetHotels(c *gin.Context) {
	countryCode := "SG"
	cityName := "Singapore"

	result, res, err := config.ApiClient().StaticDataApi.GetHotels(context.Background()).CountryCode(countryCode).CityName(cityName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StaticDataApi.GetHotels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", res)
	}
	c.JSON(200, result)
}

func GetHotelDetails(c *gin.Context) {
	hotelId := "lp24373"
	result, res, err := config.ApiClient().StaticDataApi.GetHotelDetails(context.Background()).HotelId(hotelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StaticDataApi.GetHotelDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", res)
	}
	c.JSON(200, result)
}

func GetIataCode(c *gin.Context) {
	result, res, err := config.ApiClient().StaticDataApi.GetIataCodes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StaticDataApi.GetIataCodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", res)
	}
	c.JSON(200, result)
}
