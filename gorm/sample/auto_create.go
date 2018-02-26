user := User{
	Name:            "r-ume",
	BillingAddress:  Address{Address1: "Billing Address - Address 1"},
	ShippingAddress: Address{Address1: "Shipping Address - Address 1"},
	Emails: []Email{
		{Email: "r-ume@gmail.com"},
	},
  Languages: []Language{
    {Name: "ZH"},
    {Name: "EN"}
  },
}

db.Create(&user)
db.Save(&user)

// Skip AutoUpdate
