package homeio

import "fmt"

// Room identifies a room in the house. The room is identified by a letter. Rooms A-O are available.
type Room string

const (
	LivingRoom      Room = "A"
	GuestRestRoom   Room = "B"
	Pantry          Room = "C"
	Kitchen         Room = "D"
	EntranceHall    Room = "E"
	Garage          Room = "F"
	BedroomCorridor Room = "G"
	ChildrenRoom    Room = "H"
	Bathroom        Room = "I"
	SingleBedroom   Room = "J"
	PrivateBedroom  Room = "K"
	CoupleBedroom   Room = "L"
	LaundryRoom     Room = "M"
	HomeOffice      Room = "N"
	Exterior        Room = "O"
)

// UnitIDToRoom converts a unit ID to a room or returns an error if no room corresponds to the unit ID.
func UnitIDToRoom(unitID uint8) (Room, error) {
	if unitID < 1 || unitID > 15 {
		return "", fmt.Errorf("invalid unit ID %d", unitID)
	}

	return Room(rune('A' + unitID - 1)), nil
}
