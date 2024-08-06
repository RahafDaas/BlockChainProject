package apartment

type Apartment struct {
	ID             string
	Name           string
	NumberOfRooms  int
	Size           int
	RentalPrice    float64
	Kitchen        bool
	LivingRoom     bool
	RestroomNumber int
	FloorNumber    int
	BuildYear      int
	Location       string
	WiFi           bool
	Rented         bool
}

func New(ID string, Name string, NumberOfRooms int, Size int, Kitchen bool, LivingRoom bool, RestroomNumber int, FloorNumber int, BuildYear int, location string, WiFi bool, rentalPrice float64) *Apartment {
	return &Apartment{
		ID:             ID,
		Name:           Name,
		NumberOfRooms:  NumberOfRooms,
		Kitchen:        Kitchen,
		Size:           Size,
		RentalPrice:    rentalPrice,
		LivingRoom:     LivingRoom,
		RestroomNumber: RestroomNumber,
		FloorNumber:    FloorNumber,
		BuildYear:      BuildYear,
		Location:       location,
		WiFi:           WiFi,
		Rented:         false,
	}
}
