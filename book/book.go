package book

import (
	"context"
	"fmt"
	config "go-sdk-example/core"
	"os"

	"github.com/gin-gonic/gin"
)

func Prebook(c *gin.Context) {
	rateId := "NRYDCZRZHAZHYMRQGIZS2MJRFUYTK7BSGAZDGLJRGEWTCNT4GF6HYVKTPRDVUU2EJVGUUWCHJEZVIRKOKJJEOSKZIRGVSUSXJBCTGR2LJZJFQUCRLFJUOTKSKFDUSWSUINGUUUSHKY3EIRKNIJJUOTKZKRBU2SSXKBJFGVZULAZFMS2OGZDEWVJTGRFVMSSVJE3UEUSJIVMUKRZXIJJUOSJUIRDTOQ2WI5FEYRSRGRKDIR2FGNCFGTKCKRDUCWSEINHFUWKHJVMVITJSINDE4TSIKNDU2SSTJBCTERCDJVBEIRZULJBTETKKKNEEKMSEINGUINCKLJFEKTKUKQ2EOSKZIRCU2WSRI42FURCJG5BFQUCRGJCE2TSCLBCU4SSFGZETET2LJJCEKNCJLJJUOQK2IRDUYSSRI42FOVCFJZAXYVKTIR6HY7BVGMXDIND4GIYDEMZNGA3S2MRUPRBE67BRGI4TIMJQPQZA"
	result, res, err := config.ApiClient().BookApi.PreBook(context.Background()).RateId(rateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BookApi.PreBook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", res)
	}
	c.JSON(200, result)
}

func Booking(c *gin.Context) {
	prebookId := "cC_qsaPMI"
	guestFirstName := "Kim"
	guestLastName := "James"
	guestEmail := "test@nlite.ml"
	paymentMethod := "CREDIT_CARD"
	holderName := "Kim James"
	cardNumber := "4242424242424242"
	expireDate := "11/23"
	cvc := "123"

	result, res, err := config.ApiClient().BookApi.Book(context.Background()).PrebookId(prebookId).GuestFirstName(guestFirstName).GuestLastName(guestLastName).GuestEmail(guestEmail).PaymentMethod(paymentMethod).HolderName(holderName).CardNumber(cardNumber).ExpireDate(expireDate).Cvc(cvc).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BookApi.Book``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", res)
	}
	c.JSON(200, result)

}
