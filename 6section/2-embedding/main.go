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

type ContactInfo struct {
	Email string
	Phone string
}

func (ci ContactInfo) DisplayContact() string {
	return fmt.Sprintf("Email: %s, Phone: %s", ci.Email, ci.Phone)
}

type Company struct {
	Name    string
	Address // embedded with no name - no need to
	ContactInfo
	BusinessType string
}

func (c Company) GetProfile() {
	fmt.Printf("Company Name: %s\n", c.Name)

	fmt.Printf("Location: %s\n", c.FullAddress())
	fmt.Printf("Street (promoted): %s\n", c.Street)

	fmt.Printf("Email (promoted): %s\n", c.Email)
	fmt.Printf("Business Type: %s\n", c.BusinessType)
}

type CompanyWithOwnEmail struct {
	Name string
	Address
	ContactInfo
	Email string // shadows ContactInfo email
}

func main() {
	fmt.Println("----- Struct Embedding -----")

	comp := Company{
		Name: "Innovate Solutions Inc.",
		Address: Address{
			Street:  "789 Innovation Drive",
			City:    "Techville",
			State:   "TS",
			ZipCode: "12345",
		},
		ContactInfo: ContactInfo{
			Email: "contact@innovate.com",
			Phone: "555-0100",
		},
		BusinessType: "Technology",
	}

	comp.GetProfile()

	fmt.Printf("\nDirect access to comp.City: %s\n", comp.City)
	fmt.Printf("Direct access to comp.Phone: %s\n", comp.Phone)
}
