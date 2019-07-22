// This is a broken file please fix to allow it to build and correct any logic errors
package main

import "fmt"

type SecretClientData struct {
	UpToDate    bool
	ClientsData string
}

func main() {
	var client4 = SecretClientData{
		UpToDate:    "true",
		ClientsData: "Super sensitive details of Client 4",
	}

	var client3 = SecretClientData{
		UpToDate:    false,
		ClientsData: "Important information on some politician",
	}

	if IsThereAnyData(client4) {
		DisplayData(client4)
	}

	if IsThereAnyData(client3) {
		DisplayData(client3)
	}
}

// DisplayData should display all data that SecretClientData holds to standard output
func DisplayData(s SecretClientData) {
	fmt.Printf("The clients data is up to date: %d", i.UpToDate)
	fmt.Printf("Here is the data: %x", i.ClientsData)
}

// IsThereAnyData data should check to see if there is any data in the variable
// of type SomeStruct
func IsThereAnyData(s SecretClientData) bool {
	if s.ClientsData == "" {
		return true
	}
	return false
}
