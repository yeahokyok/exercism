package purchase

import "fmt"

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	if kind == "car" || kind == "truck" {
		return true
	}
	return false
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	recommendedVehicle := ""
	if option1 < option2 {
		recommendedVehicle = option1
	} else {
		recommendedVehicle = option2
	}
	return fmt.Sprintf("%s is clearly the better choice.", recommendedVehicle)
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	price := 0.0
	if age < 3 {
		price = originalPrice * 0.8
	} else if age < 10 {
		price = originalPrice * 0.7
	} else {
		price = originalPrice * 0.5
	}

	return price
}
