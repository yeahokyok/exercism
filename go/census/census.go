// Package census simulates a system used to collect census data.
package census

// Resident represents a resident in this city.
type Resident struct {
	Name    string
	Age     int
	Address map[string]string
}

// NewResident registers a new resident in this city.
func NewResident(name string, age int, address map[string]string) *Resident {
	return &Resident{Name: name, Age: age, Address: address}
}

// HasRequiredInfo determines if a given resident has all of the required information.
func (r *Resident) HasRequiredInfo() bool {
	return r.hasValidName() && r.hasValidAddress()
}

func (r *Resident) hasValidName() bool {
	return r.Name != ""
}

func (r *Resident) hasValidAddress() bool {
	// return r.Address != nil && len(r.Address) > 0 && r.Address["street"] != ""
	return r.Address["street"] != ""
}

// Delete deletes a resident's information.
func (r *Resident) Delete() {
	*r = Resident{}
}

// Count counts all residents that have provided the required information.
func Count(residents []*Resident) int {
	count := 0
	for _, r := range residents {
		if r.HasRequiredInfo() {
			count++
		}
	}
	return count
}
