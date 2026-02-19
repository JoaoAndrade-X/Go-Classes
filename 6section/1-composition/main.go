package main

import "fmt"

type Address struct {
	Street  string
	City    string
	State   string
	ZipCode string
}

func (a Address) FullAddress() string {
	if a.Street == "" && a.City == "" {
		return "No address provided"
	}

	return fmt.Sprintf("%s, %s, %s, %s", a.Street, a.City, a.State, a.ZipCode)
}

type Customer struct {
	CustomerId      int
	Name            string
	Email           string
	BillingAddress  Address
	ShippingAddress Address
}

func (c Customer) PrintDetails() {
	fmt.Printf("Customer ID: %d\n", c.CustomerId)
	fmt.Printf("Name: %s\n", c.Name)
	fmt.Printf("Email: %s\n", c.Email)
	fmt.Printf("Billing address: %s\n", c.BillingAddress.FullAddress())
	fmt.Printf("Shipping address: %s\n", c.ShippingAddress.FullAddress())
}

func main() {
	fmt.Println("----- Composition -----")

	cust1 := Customer{
		CustomerId: 1001,
		Name:       "Gadget Corp",
		Email:      "sales@gadgetcorp.com",
		BillingAddress: Address{
			Street:  "123 Tech Road",
			City:    "Innovateville",
			State:   "CA",
			ZipCode: "90210",
		},
		ShippingAddress: Address{
			Street:  "456 Factory Lane",
			City:    "Manufacturicity",
			State:   "NV",
			ZipCode: "89101",
		},
	}

	cust1.PrintDetails()

	fmt.Println("----- Customer with same billing and shipping addresses -----")
	mainAddress := Address{
		Street:  "789 Main St",
		City:    "Anytown",
		State:   "TX",
		ZipCode: "75001",
	}
	cust2 := Customer{
		CustomerId:      1002,
		Name:            "John Doe",
		Email:           "john.doe@email.com",
		BillingAddress:  mainAddress,
		ShippingAddress: mainAddress,
	}

	cust2.PrintDetails()
}
