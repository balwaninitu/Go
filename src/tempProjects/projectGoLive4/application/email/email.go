package email

import (
	"fmt"
	"net/smtp"
	apiclient "projectGoLive/application/apiclient"
	config "projectGoLive/application/config"
	user_db "projectGoLive/application/user_db"
	"strconv"
	"strings"
)

type SellerInfo struct {
	Location  string
	Email     string
	CartItems []apiclient.ItemsDetails
}

func Sendemail(buyername string, cartItems []apiclient.ItemsDetails) bool {

	// Create a map to store cart items for each seller
	var sellerMap map[string]SellerInfo
	sellerMap = make(map[string]SellerInfo)

	// Buyer details :
	buyerdetails, _ := user_db.GetARecord(config.DB, buyername)
	address := buyerdetails.Address
	fmt.Println(address)
	// Get all user records :
	userdetails, _ := user_db.GetRecords(config.DB)
	fmt.Println(userdetails)

	//tempCartSlice := []apiclient.ItemsDetails{}
	for _, cartitem := range cartItems {
		fmt.Println("Outside :", cartitem)
		for _, user := range userdetails {
			fmt.Println("User details in for loop :", user)

			if user.Username == cartitem.Username {
				fmt.Println("------------------------")
				fmt.Println("Inside :", cartitem, user)
				fmt.Println("------------------------")

				tempCartSlice := sellerMap[cartitem.Username].CartItems
				tempCartSlice = append(tempCartSlice, cartitem)
				selleremail := user.Email
				//selleremail := "peelrescue+" + cartitem.Username + "@gmail.com"
				sellerMap[cartitem.Username] = SellerInfo{user.Address, selleremail, tempCartSlice}
			}
		}
	}
	fmt.Println("SELLER MAP", sellerMap)
	// Sender data.
	from := "peelrescue@gmail.com"
	password := "blueappleredorange"
	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	var message []string
	for key, value := range sellerMap {
		fmt.Println("For seller:", key)
		// Receiver email address.
		to := []string{
			value.Email,
		}
		message = nil
		message = append(message, "Hello Seller : "+key)
		message = append(message, "Items purchased by Buyer : "+buyerdetails.Username)
		message = append(message, "Item\t\t\tQuantity\t\t\tCost")
		for _, cartitem := range value.CartItems {
			message = append(message, cartitem.Item+"\t\t\t"+strconv.Itoa(cartitem.Quantity)+"\t\t\t"+floattostr(cartitem.Cost))
		}
		// Message.
		message = append(message, "Thank you for using Peel Rescue. Saving Earth one Peel at a time!")
		// Sending email.
		justString := strings.Join(message, "\n")
		bytemessage := []byte(justString)
		err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, bytemessage)
		if err != nil {
			fmt.Println(err)
			return false
		}
	}
	fmt.Println("Email Sent Successfully!")
	return true
}

func floattostr(input float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(input, 'f', -1, 64)
}
