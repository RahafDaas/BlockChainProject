Apartment Rental Project

![Screenshot (534)](https://github.com/user-attachments/assets/932cd7ab-bb45-4de6-94ae-419562d10f63)


This project aims to create a system that displays the available apartments to be rented.
We started by assuming that the structure of an apartment consists of:
  * ID             string
	* Name           string
	* NumberOfRooms  int
	* Size           int
	* RentalPrice    float64
	* Kitchen        bool
	* LivingRoom     bool
	* RestroomNumber int
	* FloorNumber    int
	* BuildYear      int
	* Location       string
	* WiFi           bool
	* Rented         bool
  
The motivation behind this system is to have one place that links all apartments into one office to make the process of renting apartments much easier and more flexible. Most of the required information regarding renting an apartment is available in this system. Initially, an apartment is not rented therefore, the Rented field will always be False unless the RentApartment method from the Handler is called and assuming that the user has paid the rent price. In case the rent contract has expired by calling the UnRentApartment method the Rented field will switch to False. 

The functions of this system are:
* Updating the Name of an apartment.
* Updating the Number of rooms in an apartment.
* Updating the Kitchen Availability of an apartment.
* Updating the Size of an apartment.
* Updating the Rental Price of an apartment.
* Updating the WiFi Availability of an apartment.
* Updating the Living room availability of an apartment.
* Updating the Restroom number of an apartment.
* Updating the Floor number of an apartment.
* Updating the Build year of an apartment.
* Updating the Location of an apartment.
* Register an apartment.
* To Rent an apartment. 
* And to unrent an apartment.
  
