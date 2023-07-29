package main

import (
	"fmt"
	"time"

	tinytime "github.com/wagslane/go-tinytime"
)

func main() {
	tt := tinytime.New(1585750374)

	tt = tt.Add(time.Hour * 48)
	fmt.Println(tt)
}

// type cost struct {
// 	day   int
// 	value float64
// }

// func getCostsByDay(costs []cost) []float64 {
// 	costPerDay := []float64{}
// 	for i := 0; i < len(costs); i++ {
// 		cost := costs[i]
// 		for cost.day >= len(costPerDay) {
// 			costPerDay = append(costPerDay, 0.0)
// 		}
// 		costPerDay[cost.day] += cost.value
// 	}
// 	return costPerDay
// }

// // dont edit below this line

// func test(costs []cost) {
// 	fmt.Printf("Creating daily buckets for %v costs...\n", len(costs))
// 	costsByDay := getCostsByDay(costs)
// 	fmt.Println("Costs by day:")
// 	for i := 0; i < len(costsByDay); i++ {
// 		fmt.Printf(" - Day %v: %.2f\n", i, costsByDay[i])
// 	}
// 	fmt.Println("===== END REPORT =====")
// }

// func main() {
// 	test([]cost{
// 		{0, 1.0},
// 		{1, 2.0},
// 		{1, 3.1},
// 		{2, 2.5},
// 		{3, 3.6},
// 		{3, 2.7},
// 		{4, 3.34},
// 	})
// 	test([]cost{
// 		{0, 1.0},
// 		{10, 2.0},
// 		{3, 3.1},
// 		{2, 2.5},
// 		{1, 3.6},
// 		{2, 2.7},
// 		{4, 56.34},
// 		{13, 2.34},
// 		{28, 1.34},
// 		{25, 2.34},
// 		{30, 4.34},
// 	})
// }

// func getExpenseReport(e expense) (string, float64) {
// 	switch v := e.(type) {
// 	case email:
// 		return v.toAddress, v.cost()
// 	case sms:
// 		return v.toPhoneNumber, v.cost()
// 	default:
// 		return "", 0.0
// 	}
// }

// 20
// func getExpenseReport(e expense) (string, float64) {
// 	v, isEmail := e.(email)
// 	if isEmail {
// 		return v.toAddress, v.cost()
// 	}
//
// 	v2, isSms := e.(sms)
// 	if isSms {
// 		return v2.toPhoneNumber, v2.cost()
// 	}
//
// 	return "", 0.0
//
// }

// func maxMessages(thresh float64) int {
// 	total := 0.0
// 	max := 0
// 	for i := 0; total < thresh; i++ {
// 		max += 1
// 		total += (1.0 + 0.01*float64(i))
// 	}
// 	fmt.Println("total: ", total, " thresh: ", thresh)
// 	return max - 1
// }

// func getMaxMessagesToSend(costMultiplier float64, maxCostInPennies int) int {
// 	actualCostInPennies := 1.0
// 	maxMessagesToSend := 0
// 	for actualCostInPennies < float64(maxCostInPennies) {
// 		maxMessagesToSend++
// 		actualCostInPennies *= costMultiplier
// 	}
// 	return maxMessagesToSend
// }

// func fizzbuzz() {
// 	var output string

// 	divisor_1 := 3
// 	divisor_2 := 5

// 	for i := 1; i <= 100; i++ {
// 		if i%divisor_1 == 0 {
// 			output += "fizz"
// 		}

// 		if i%divisor_2 == 0 {
// 			output += "buzz"
// 		}

// 		if output == "" {
// 			output = strconv.Itoa(i)
// 		}

// 		fmt.Println(output)
// 		output = ""
// 	}
// }

// func printPrimes(max int) {
// 	for i := 2; i <= max; i++ {
// 		if isPrime(i) {
// 			fmt.Println(i)
// 		}
// 	}
// }

// func isPrime(num int) bool {
// 	if num <= 1 {
// 		return false
// 	}

// 	for i := 2; i*i <= num; i++ {
// 		if num%i == 0 {
// 			return false
// 		}
// 	}
// 	return true
// }

// 19 see below
// 18 define variables in struct

// 17
// fix returned values, naming (no naked return)

// 16
// firstName, _ := getNames()
// fmt.Println("Welcome to textio,", firstName)

// 15
// sendsSoFar := 430
// const sendsToAdd = 25
// sendsSoFar = incrementSends(sendsSoFar, sendsToAdd)
// fmt.Println("you've sent", sendsSoFar, "messages")

// 14
// fix function signature

// 13
// messageLen := 10
// maxMessageLen := 20
// fmt.Println("Trying to send a message of length:", messageLen, "and a max length of:", maxMessageLen)

// // don't touch above this line

// if messageLen <= maxMessageLen {
// 	fmt.Println("Message sent")
// } else {
// 	fmt.Println("Message not sent")
// }

// 12
// const name = "Saul Goodman"
// const openRate = 30.5
// msg := fmt.Sprintf("Hi %s, your open rate is %.1f percent", name, openRate)

// fmt.Println(msg)

// 11
// const secondsInMinute = 60
// const minutesInHour = 60
// const secondsInHour = secondsInMinute * minutesInHour

// don't edit below this line
// fmt.Println("number of seconds in an hour:", secondsInHour)

// 10
// const premiumPlanName = "Premium Plan"
// const basicPlanName = "Basic Plan"

// don't edit below this line

// fmt.Println("plan:", premiumPlanName)
// fmt.Println("plan:", basicPlanName)

// 9
// accountAge := 2.6
// accountAgeInt := int(accountAge)
// create a new "accountAgeInt" here
// it should be the result of casting "accountAge" to an integer

// fmt.Println("Your account has existed for", accountAgeInt, "years")

// 8
// averageOpenRate := .23
// displayMessage := "is the average open rate of your messages"
// fmt.Println(averageOpenRate, displayMessage)

// 7
// penniesPerText := 2.0
// fmt.Printf("The type of penniesPerText is %T\n", penniesPerText)

// 6
// congrats := "happy birthday!"
// fmt.Println(congrats)

// 5
// var smsSendingLimit int
// var costPerSMS float64
// var hasPermission bool
// var username string
// fmt.Printf("%v %f %v %q\n", smsSendingLimit, costPerSMS, hasPermission, username)

// 4
// var username string = "wagslane"
// var password string = "20947382822"

// // don't edit below this line
// fmt.Println("Authorization: Basic", username+":"+password)

// 3
// fmt.Println("the-compiled Textio server is starting")

// 2
// messagesFromDoris := []string{
// 	"You doing anything later??",
// 	"Did you get my last message?",
// 	"Don't leave me hanging...",
// 	"Please respond I'm lonely!",
// }

// numMessages := float64(len(messagesFromDoris))
// costPerMessage := .02

// // don't touch above this line

// totalCost := costPerMessage * numMessages

// // don't touch below this line

// fmt.Printf("Doris spent %.2f on text messages today\n", totalCost)

// 1
// fmt.Println("starting Textio servuser

// Helper Funcs

// 19
// type messageToSend struct {
// 	message   string
// 	sender    user
// 	recipient user
// }

// type user struct {
// 	name   string
// 	number int
// }

// func canSendMessage(mToSend messageToSend) bool {
// 	if mToSend.sender.name != "" && mToSend.sender.number != 0 {
// 		return true
// 	}
// 	return false
// }

// 15
// func incrementSends(sendsSoFar, sendsToAdd int) int {
// 	sendsSoFar = sendsSoFar + sendsToAdd
// 	return sendsSoFar
// }

// 16
// func getNames() (string, string) {
// 	return "John", "Doe"
// }
