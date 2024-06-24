package homeio

import "fmt"

// Room identifies a room in the house. The room is identified by a letter. Rooms A-P are available, whereas P is the miscellaneous room to group all non-room related I/O.
type Room string

const (
	// LivingRoom is the living room. It has the letter A and Unit ID 1.
	LivingRoom Room = "A"

	// GuestRestRoom is the guest restroom. It has the letter B and Unit ID 2.
	GuestRestRoom Room = "B"

	// Pantry is the pantry. It has the letter C and Unit ID 3.
	Pantry Room = "C"

	// Kitchen is the kitchen. It has the letter D and Unit ID 4.
	Kitchen Room = "D"

	// EntranceHall is the entrance hall. It has the letter E and Unit ID 5.
	EntranceHall Room = "E"

	// Garage is the garage. It has the letter F and Unit ID 6.
	Garage Room = "F"

	// BedroomCorridor is the bedroom corridor. It has the letter G and Unit ID 7.
	BedroomCorridor Room = "G"

	// ChildrenRoom is the children room. It has the letter H and Unit ID 8.
	ChildrenRoom Room = "H"

	// Bathroom is the bathroom. It has the letter I and Unit ID 9.
	Bathroom Room = "I"

	// SingleBedroom is the single bedroom. It has the letter J and Unit ID 10.
	SingleBedroom Room = "J"

	// PrivateBathroom is the private bathroom. It has the letter K and Unit ID 11.
	PrivateBathroom Room = "K"

	// CoupleBedroom is the couple bedroom. It has the letter L and Unit ID 12.
	CoupleBedroom Room = "L"

	// LaundryRoom is the laundry room. It has the letter M and Unit ID 13.
	LaundryRoom Room = "M"

	// HomeOffice is the home office. It has the letter N and Unit ID 14.
	HomeOffice Room = "N"

	// Exterior is the exterior. It has the letter O and Unit ID 15.
	Exterior Room = "O"

	// Miscellaneous is the miscellaneous room. It has the letter P and Unit ID 16.
	Miscellaneous Room = "P"
)

// UnitIDToRoom converts a unit ID to a room or returns an error if no room corresponds to the unit ID.
func UnitIDToRoom(unitID uint8) (Room, error) {
	// Check if the unit ID is valid.
	if unitID < 1 || unitID > 17 {
		return "", fmt.Errorf("invalid unit ID %d", unitID)
	}

	// Return the room corresponding to the unit ID.
	return Room(rune('A' + unitID - 1)), nil
}

// String converts a room to a string.
func (r Room) String() string {
	return string(r)
}
