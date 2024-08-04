package apartment

type Apartment struct {
	ID           string
	Name         string
	StreetName   string
	DistrictName string
	Size         int
	RentalPrice  float64
	Color        string
	Rented       string
}

func New(ID string, name string, streetName string, districtName string, size int, rentalPrice float64, color string) *Apartment {
	return &Apartment{
		ID:           ID,
		Name:         name,
		StreetName:   streetName,
		DistrictName: districtName,
		Size:         size,
		RentalPrice:  rentalPrice,
		Color:        color,
		Rented:       "Not Rented",
	}
}
