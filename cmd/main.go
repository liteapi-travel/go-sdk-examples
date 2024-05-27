package main

import (
	"go-sdk-example/book"
	bookingManagment "go-sdk-example/bookingmanagment"
	"go-sdk-example/search"
	staticData "go-sdk-example/staticdata"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//--------staticData--------------
	r.GET("/cities", staticData.GetCities)
	r.GET("/countries", staticData.GetCountries)
	r.GET("/currencies", staticData.GetCurrencies)
	r.GET("/hotels", staticData.GetHotels)
	r.GET("/hotelDetails", staticData.GetHotelDetails)
	r.GET("/iataCodes", staticData.GetIataCode)

	//--------search--------------
	//r.GET("/minRate", search.GetMinimumRates)
	r.GET("/fullRate", search.GetFullRates)

	//--------book--------------
	r.GET("/prebook", book.Prebook)
	r.GET("/book", book.Booking)

	//--------bookingManagment--------------
	r.GET("/bookings", bookingManagment.BookingList)
	r.GET("/retrieve", bookingManagment.BookingRetrieve)
	r.GET("/cancel", bookingManagment.BookingCancel)

	//--------guestAndLoyalty--------------
//	r.GET("/guestsIds", guestAndLoyalty.GetGuestsIds)

	r.Run()
}
