package main

import (
	"encoding/json"
	"fmt"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/balancetransaction"
	"github.com/stripe/stripe-go/v76/bankaccount"
	"github.com/stripe/stripe-go/v76/charge"
	"github.com/stripe/stripe-go/v76/customer"
	"github.com/stripe/stripe-go/v76/payout"
	"github.com/stripe/stripe-go/v76/product"
	"github.com/stripe/stripe-go/v76/subscription"
)

func main() {
	//stripe.Key = "rk_test_51OChE7B8LjpwaxcmiBwCaRzueMBJJEFvZsvRYa8o5CAYRfwrAwio2O0kSSKX8YcpL4WsxHuUnpRk5htNNkyJK86q00d9WT2px1"
	stripe.Key = "sk_test_4eC39HqLyjWDarjtT1zdp7dc"
	getDestination()
}

func getDestination() {
	params := &stripe.BankAccountParams{}
	account, err := bankaccount.Get("ba_1LylJM2eZvKYlo2Ca63N1kU1", params)
	fmt.Println(err)
	data, _ := json.Marshal(account)
	fmt.Println(string(data))
}

func getPayout() {
	params := &stripe.PayoutListParams{}
	list := payout.List(params).PayoutList()
	data, _ := json.Marshal(list)
	fmt.Println(string(data))
}

func getProducts() {
	params := &stripe.ProductListParams{}
	list := product.List(params)
	data, _ := json.Marshal(list)
	fmt.Println(string(data))
}

func getSubscriptions() {
	params := &stripe.SubscriptionListParams{}
	list := subscription.List(params).SubscriptionList().Data
	data, _ := json.Marshal(list)
	fmt.Println(string(data))
}

func getChargeObject() {
	c, _ := charge.Get("ch_3OChoT2eZvKYlo2C1cUcWlol", nil)
	data, _ := json.Marshal(c)
	fmt.Println(string(data))
}

func getBalanceTransactions() {
	params := &stripe.BalanceTransactionListParams{}
	list := balancetransaction.List(params).BalanceTransactionList().Data
	data, _ := json.Marshal(list)
	fmt.Println(string(data))
}

func getCustomerList() {
	params := &stripe.CustomerListParams{}
	list := customer.List(params).CustomerList().Data
	data, _ := json.Marshal(list)
	fmt.Println(string(data))
}

func connectedAccounts() {
	params := &stripe.CustomerParams{}
	ch, err := customer.Get("cus_P0ixAzkaRGyiAZ", params)
	fmt.Println(err)
	data, _ := json.Marshal(ch)
	fmt.Println(string(data))
}

func createCustomer() {
	params := &stripe.CustomerParams{
		Description:      stripe.String("Stripe Developer"),
		Email:            stripe.String("gostripe@stripe.com"),
		PreferredLocales: stripe.StringSlice([]string{"en", "es"}),
	}

	c, err := customer.New(params)
	fmt.Println(err)
	data, _ := json.Marshal(c)
	fmt.Println(string(data))
}
