package menuservice

import (
	balanceinfo "DishTV/recharge"
	compute "DishTV/subscriptioncompute"
	"bufio"
	"errors"
	"fmt"
	"os"
)

type services interface {
	ServiceMenu()
}

//var rechObj *balanceinfo.RechargeTokens

func ServiceMenu() {

	var rechObj *balanceinfo.RechargeTokens
	rechObj = balanceinfo.New()

	reader := bufio.NewReader(os.Stdin)

	var userChoice int
	for {
		fmt.Println("1. View current balance in the account\n" +
			"2. Recharge Account\n" +
			"3. View available packs, channels and services\n" +
			"4. Subscribe to base packs\n" +
			"5. Add channels to an existing subscription\n" +
			"6. Subscribe to special service\n" +
			"7. View current subscription details\n" +
			"8. Update email and phone number for notifications\n" +
			"9. Exit\n" +
			"\nEnter the option: ")

		fmt.Scanf("%d", &userChoice)

		switch userChoice {

		case 1: // Case for checking balance.
			//rechObj := checkbal.New()
			rechObj = balanceinfo.New()
			fmt.Println("Current balance is : ", rechObj.CheckBalance())
			fmt.Println("Case 1")
			fmt.Println("============================================================")
			break

		case 2: // Case for recharging
			fmt.Println("Case 2", rechObj)
			var rechargeAmount int
			fmt.Scanf("%d", &rechargeAmount)
			//rechargeAmount = 100
			fmt.Println(rechObj.DoRecharge(rechargeAmount))
			fmt.Println("============================================================")
			break

		case 3:
			fmt.Println(case3Options())
			fmt.Println("============================================================")
			break

		case 4:
			var userInput string
			var months int
			fmt.Println("Subscribe to channel packs")
			fmt.Print("Enter the Pack you wish to subscribe: (Silver: 'S', Gold: 'G'):")
			fmt.Scanf("%s", &userInput)
			fmt.Println("UserInput:", userInput)
			userInput, err := sanitizeInput(userInput) // Checking if valid input is given.
			if err != nil {
				fmt.Println("Invalid input. Please try again. ")
			}
			fmt.Print("No. of months:")
			fmt.Scanf("%d", &months)
			// subscription amount , discount , finalprice,error
			amountToPay, discount, finalPrice, remainingBal, monthlyPrice, packType, status, err := compute.ComputeAmount(userInput, months)
			if err != nil {
				fmt.Println("Invalid Input. Please try again.")
				break
			}
			if status != true {
				fmt.Println("Insufficient balance . Kindly recharge then try again.")
				break
			}
			fmt.Println(amountToPay, discount, finalPrice, status)
			fmt.Println("\nYou have successfully subscribed the following packs -", packType+
				"\nMonthly Price: ", monthlyPrice,
				"\nSubscription Amount: ", amountToPay,
				"\nDiscount Applied: ", discount,
				"\nFinal Price after discount:", finalPrice,
				"\nAccount balance: ", remainingBal)
			fmt.Println("============================================================")
			break
		case 5:
			fmt.Println("Add channels to existing subscription:",
				"\nEnter channel names to add (separated by commas)")
			var indiChannels string
			fmt.Scanf("%s", &indiChannels)
			fmt.Println("Channels String :", indiChannels)
			//(valid string, status bool, err error)
			status, err := compute.SubscribeChannel(indiChannels)
			if err != nil {
				fmt.Println("Invalid channel names. Try again.")
				break
			}
			if status != true {
				fmt.Println("Insufficient funds.")
				break
			}
			fmt.Println("Channels added successfully.\n",
				"Account balance: ", rechObj.CheckBalance(), "Rs.")
			fmt.Println("============================================================")
		case 6:
			//SubscribeService
			var indiServices string
			fmt.Println("Subscribe to special services")
			fmt.Print("Enter the service name: ")
			fmt.Scanf("%s", &indiServices)
			fmt.Printf("Services are :%s", indiServices)
			success, err := compute.SubscribeService(indiServices)
			if err != nil {
				fmt.Println("", success)
			}
			if success != true {
				fmt.Println("Insufficient balance.")
			}
			fmt.Println("Account balance: ", rechObj.CheckBalance())
			fmt.Println("Email notification sent successfully\n",
				"SMS notification sent successfully")
			fmt.Println("============================================================")
		case 7:
			fmt.Println("View current subscription details",
				"\nCurrently subscribed packs and channels:", compute.GetSubscribeList(), "\n", compute.GetSubscribeServiceList())
			break
			fmt.Println("============================================================")
		case 8:

			fmt.Println("Update email and phone number for notifications\n")
			fmt.Print("\nEnter the email: ")
			email, _ := reader.ReadString('\n')
			fmt.Print(email)
			fmt.Print("\nEnter the phone: ")
			phone, _ := reader.ReadString('\n')
			fmt.Print(phone)

			fmt.Println("============================================================")
		case 9:
			fmt.Println("Thank you.")
			os.Exit(0)

		default:
			fmt.Println("Invalid Choice.Try again.")
		}
	}
}
func case3Options() (optionString string) {
	optionString = "\nView available packs, channels and services\n" +
		"Available packs for subscription\n" +
		"Silver pack: Zee, Sony, Star Plus: 50 Rs.\n" +
		"Gold Pack: Zee, Sony, Star Plus, Discovery, NatGeo: 100 Rs.\n" +
		"Available channels for subscription\n" +
		"Zee: 10 Rs.\n" +
		"Sony: 15 Rs.\n" +
		"Star Plus: 20 Rs.\n" +
		"Discovery: 10 Rs.\n" +
		"NatGeo: 20 Rs.\n" +
		"Available services for subscription\n" +
		"LearnEnglish Service: 200 Rs.\n" +
		"LearnCooking Service: 100 Rs.\n"
	return
}

func sanitizeInput(userInput string) (string, error) {
	if userInput == "S" || userInput == "G" {
		return userInput, nil
	} else {
		return "", errors.New("Invalid input")
	}

}
