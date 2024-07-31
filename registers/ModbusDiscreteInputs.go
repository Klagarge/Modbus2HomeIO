package registers

import (
	"Modbus2HomeIO/homeio"
	"fmt"

	modbus "github.com/Klagarge/modbusGo"
)

/* Handler interface implementation, handle Read Discrete Inputs (0x04) */

func (m *Handler) HandleDiscreteInputs(req *modbus.DiscreteInputsRequest) (res []bool, err error) {
	// Get room from unit ID and fail if the room does not exist.
	room, err := homeio.UnitIDToRoom(req.UnitId)
	if err != nil {
		return []bool{}, fmt.Errorf("invalid unit ID %v", req.UnitId)
	}

	for address := req.Addr; address < req.Addr+req.Quantity; address++ {
		value := false
		switch room {
		case homeio.LivingRoom:
			switch address {
			case 0, 1, 2:
				value, err = m.home.IsButtonPressed(homeio.LivingRoom, homeio.Button1+homeio.Button(address-0))
			case 3, 4:
				value, err = m.home.IsDimmerControlUpPressed(homeio.LivingRoom, homeio.Dimmer1+homeio.Dimmer(address-3))
			case 5, 6:
				value, err = m.home.IsDimmerControlDownPressed(homeio.LivingRoom, homeio.Dimmer1+homeio.Dimmer(address-5))
			case 7, 8, 9:
				value, err = m.home.IsShutterControlUpPressed(homeio.LivingRoom, homeio.Shutters1+homeio.Shutters(address-7))
			case 10, 11, 12:
				value, err = m.home.IsShutterControlDownPressed(homeio.LivingRoom, homeio.Shutters1+homeio.Shutters(address-10))
			case 13, 14, 15:
				value, err = m.home.AreShuttersOnTop(homeio.LivingRoom, homeio.Shutters1+homeio.Shutters(address-13))
			case 16, 17, 18:
				value, err = m.home.AreShuttersOnBottom(homeio.LivingRoom, homeio.Shutters1+homeio.Shutters(address-16))
			case 19, 20:
				value, err = m.home.IsDoorClosed(homeio.LivingRoom, homeio.Door1+homeio.Door(address-19))
			case 21:
				value, err = m.home.IsSmokeDetected(homeio.LivingRoom)
			case 22:
				value, err = m.home.IsMotionDetected(homeio.LivingRoom)
			case 23:
				value, err = m.home.IsLightDetected(homeio.LivingRoom)
			}

		case homeio.GuestRestRoom:
			switch address {
			case 0, 1:
				value, err = m.home.IsButtonPressed(homeio.GuestRestRoom, homeio.Button1+homeio.Button(address-0))
			case 2:
				value, err = m.home.IsMotionDetected(homeio.GuestRestRoom)
			}

		case homeio.Pantry:
			switch address {
			case 0:
				value, err = m.home.IsButtonPressed(homeio.Pantry, homeio.Button1)
			}

		case homeio.Kitchen:
			switch address {
			case 0, 1:
				value, err = m.home.IsDimmerControlUpPressed(homeio.Kitchen, homeio.Dimmer1+homeio.Dimmer(address-0))
			case 2, 3:
				value, err = m.home.IsDimmerControlDownPressed(homeio.Kitchen, homeio.Dimmer1+homeio.Dimmer(address-2))
			case 4:
				value, err = m.home.IsShutterControlUpPressed(homeio.Kitchen, homeio.Shutters1)
			case 5:
				value, err = m.home.IsShutterControlDownPressed(homeio.Kitchen, homeio.Shutters1)
			case 6:
				value, err = m.home.AreShuttersOnTop(homeio.Kitchen, homeio.Shutters1)
			case 7:
				value, err = m.home.AreShuttersOnBottom(homeio.Kitchen, homeio.Shutters1)
			case 8:
				value, err = m.home.IsDoorClosed(homeio.Kitchen, homeio.Door1)
			case 9:
				value, err = m.home.IsMotionDetected(homeio.Kitchen)
			case 10:
				value, err = m.home.IsLightDetected(homeio.Kitchen)
			}

		case homeio.EntranceHall:
			switch address {
			case 0, 1, 2, 3, 4:
				value, err = m.home.IsButtonPressed(homeio.EntranceHall, homeio.Button1+homeio.Button(address-0))
			case 5, 6:
				value, err = m.home.IsDimmerControlUpPressed(homeio.EntranceHall, homeio.Dimmer1+homeio.Dimmer(address-5))
			case 7, 8:
				value, err = m.home.IsDimmerControlDownPressed(homeio.EntranceHall, homeio.Dimmer1+homeio.Dimmer(address-7))
			case 9:
				value, err = m.home.IsShutterControlUpPressed(homeio.EntranceHall, homeio.Shutters1)
			case 10:
				value, err = m.home.IsShutterControlDownPressed(homeio.EntranceHall, homeio.Shutters1)
			case 11:
				value, err = m.home.AreShuttersOnTop(homeio.EntranceHall, homeio.Shutters1)
			case 12:
				value, err = m.home.AreShuttersOnBottom(homeio.EntranceHall, homeio.Shutters1)
			case 13, 14:
				value, err = m.home.IsDoorClosed(homeio.EntranceHall, homeio.Door1+homeio.Door(address-13))
			case 15:
				value, err = m.home.IsMotionDetected(homeio.EntranceHall)
			case 16:
				value, err = m.home.IsLightDetected(homeio.EntranceHall)
			case 17:
				value, err = m.home.IsAlarmArmed()
			}

		case homeio.Garage:
			switch address {
			case 0, 1:
				value, err = m.home.IsButtonPressed(homeio.Garage, homeio.Button1+homeio.Button(address-0))
			case 2, 3:
				value, err = m.home.IsShutterControlUpPressed(homeio.Garage, homeio.Shutters1)
			case 4, 5:
				value, err = m.home.IsShutterControlDownPressed(homeio.Garage, homeio.Shutters1)
			case 6:
				value, err = m.home.AreShuttersOnTop(homeio.Garage, homeio.Shutters1)
			case 7:
				value, err = m.home.AreShuttersOnBottom(homeio.Garage, homeio.Shutters1)
			case 8:
				value, err = m.home.IsGateOpen(homeio.Garage)
			case 9:
				value, err = m.home.IsGateClosed(homeio.Garage)
			case 10:
				value, err = m.home.IsIRObscured(homeio.Garage, homeio.GarageIRSensor)
			case 11:
				value, err = m.home.IsMotionDetected(homeio.Garage)
			case 12:
				value, err = m.home.IsLightDetected(homeio.Garage)
			}

		case homeio.BedroomCorridor:
			switch address {
			case 0, 1, 2, 3, 4:
				value, err = m.home.IsButtonPressed(homeio.BedroomCorridor, homeio.Button1+homeio.Button(address-0))
			case 5:
				value, err = m.home.IsShutterControlUpPressed(homeio.BedroomCorridor, homeio.Shutters1)
			case 6:
				value, err = m.home.IsShutterControlDownPressed(homeio.BedroomCorridor, homeio.Shutters1)
			case 7:
				value, err = m.home.AreShuttersOnTop(homeio.BedroomCorridor, homeio.Shutters1)
			case 8:
				value, err = m.home.AreShuttersOnBottom(homeio.BedroomCorridor, homeio.Shutters1)
			case 9:
				value, err = m.home.IsSmokeDetected(homeio.BedroomCorridor)
			case 10:
				value, err = m.home.IsMotionDetected(homeio.BedroomCorridor)
			}

		case homeio.ChildrenRoom:
			switch address {
			case 0, 1:
				value, err = m.home.IsButtonPressed(homeio.ChildrenRoom, homeio.Button1+homeio.Button(address-0))
			case 2:
				value, err = m.home.IsDimmerControlUpPressed(homeio.ChildrenRoom, homeio.Dimmer1)
			case 3:
				value, err = m.home.IsDimmerControlDownPressed(homeio.ChildrenRoom, homeio.Dimmer1)
			case 4:
				value, err = m.home.IsShutterControlUpPressed(homeio.ChildrenRoom, homeio.Shutters1)
			case 5:
				value, err = m.home.IsShutterControlDownPressed(homeio.ChildrenRoom, homeio.Shutters1)
			case 6:
				value, err = m.home.AreShuttersOnTop(homeio.ChildrenRoom, homeio.Shutters1)
			case 7:
				value, err = m.home.AreShuttersOnBottom(homeio.ChildrenRoom, homeio.Shutters1)
			case 8, 9:
				value, err = m.home.IsDoorClosed(homeio.ChildrenRoom, homeio.Door1+homeio.Door(address-8))
			case 10:
				value, err = m.home.IsSmokeDetected(homeio.ChildrenRoom)
			case 11:
				value, err = m.home.IsLightDetected(homeio.ChildrenRoom)
			case 12:
				value, err = m.home.IsMotionDetected(homeio.ChildrenRoom)
			}

		case homeio.Bathroom:
			switch address {
			case 0, 1:
				value, err = m.home.IsButtonPressed(homeio.Bathroom, homeio.Button1+homeio.Button(address-0))
			case 2:
				value, err = m.home.IsMotionDetected(homeio.Bathroom)
			}

		case homeio.SingleBedroom:
			switch address {
			case 0, 1:
				value, err = m.home.IsButtonPressed(homeio.SingleBedroom, homeio.Button1+homeio.Button(address-0))
			case 2:
				value, err = m.home.IsDimmerControlUpPressed(homeio.SingleBedroom, homeio.Dimmer1)
			case 3:
				value, err = m.home.IsDimmerControlDownPressed(homeio.SingleBedroom, homeio.Dimmer1)
			case 4:
				value, err = m.home.IsShutterControlUpPressed(homeio.SingleBedroom, homeio.Shutters1)
			case 5:
				value, err = m.home.IsShutterControlDownPressed(homeio.SingleBedroom, homeio.Shutters1)
			case 6:
				value, err = m.home.AreShuttersOnTop(homeio.SingleBedroom, homeio.Shutters1)
			case 7:
				value, err = m.home.AreShuttersOnBottom(homeio.SingleBedroom, homeio.Shutters1)
			case 8, 9:
				value, err = m.home.IsDoorClosed(homeio.SingleBedroom, homeio.Door1+homeio.Door(address-6))
			case 10:
				value, err = m.home.IsSmokeDetected(homeio.SingleBedroom)
			case 11:
				value, err = m.home.IsLightDetected(homeio.SingleBedroom)
			case 12:
				value, err = m.home.IsMotionDetected(homeio.SingleBedroom)
			}

		case homeio.PrivateBathroom:
			switch address {
			case 0:
				value, err = m.home.IsDimmerControlUpPressed(homeio.PrivateBathroom, homeio.Dimmer1)
			case 1:
				value, err = m.home.IsDimmerControlDownPressed(homeio.PrivateBathroom, homeio.Dimmer1)
			case 2:
				value, err = m.home.IsMotionDetected(homeio.PrivateBathroom)
			}

		case homeio.CoupleBedroom:
			switch address {
			case 0, 1:
				value, err = m.home.IsButtonPressed(homeio.CoupleBedroom, homeio.Button1+homeio.Button(address-0))
			case 2, 3:
				value, err = m.home.IsDimmerControlUpPressed(homeio.CoupleBedroom, homeio.Dimmer1+homeio.Dimmer(address-2))
			case 4, 5:
				value, err = m.home.IsDimmerControlDownPressed(homeio.CoupleBedroom, homeio.Dimmer1+homeio.Dimmer(address-4))
			case 6:
				value, err = m.home.IsShutterControlUpPressed(homeio.CoupleBedroom, homeio.Shutters1)
			case 7:
				value, err = m.home.IsShutterControlDownPressed(homeio.CoupleBedroom, homeio.Shutters1)
			case 8:
				value, err = m.home.AreShuttersOnTop(homeio.CoupleBedroom, homeio.Shutters1)
			case 9:
				value, err = m.home.AreShuttersOnBottom(homeio.CoupleBedroom, homeio.Shutters1)
			case 10, 11:
				value, err = m.home.IsDoorClosed(homeio.CoupleBedroom, homeio.Door1+homeio.Door(address-8))
			case 12:
				value, err = m.home.IsSmokeDetected(homeio.CoupleBedroom)
			case 13:
				value, err = m.home.IsLightDetected(homeio.CoupleBedroom)
			case 14:
				value, err = m.home.IsMotionDetected(homeio.CoupleBedroom)
			}

		case homeio.LaundryRoom:
			switch address {
			case 0:
				value, err = m.home.IsButtonPressed(homeio.LaundryRoom, homeio.Button1)
			case 1:
				value, err = m.home.IsShutterControlUpPressed(homeio.LaundryRoom, homeio.Shutters1)
			case 2:
				value, err = m.home.IsShutterControlDownPressed(homeio.LaundryRoom, homeio.Shutters1)
			case 3:
				value, err = m.home.AreShuttersOnTop(homeio.LaundryRoom, homeio.Shutters1)
			case 4:
				value, err = m.home.AreShuttersOnBottom(homeio.LaundryRoom, homeio.Shutters1)
			case 5, 6:
				value, err = m.home.IsDoorClosed(homeio.LaundryRoom, homeio.Door1+homeio.Door(address-5))
			case 7:
				value, err = m.home.IsLightDetected(homeio.LaundryRoom)
			case 8:
				value, err = m.home.IsMotionDetected(homeio.LaundryRoom)
			}

		case homeio.HomeOffice:
			switch address {
			case 0, 1, 2:
				value, err = m.home.IsButtonPressed(homeio.HomeOffice, homeio.Button1+homeio.Button(address-0))
			case 3, 4, 5:
				value, err = m.home.IsDimmerControlUpPressed(homeio.HomeOffice, homeio.Dimmer1+homeio.Dimmer(address-3))
			case 6, 7, 8:
				value, err = m.home.IsDimmerControlDownPressed(homeio.HomeOffice, homeio.Dimmer1+homeio.Dimmer(address-6))
			case 9:
				value, err = m.home.IsShutterControlUpPressed(homeio.HomeOffice, homeio.Shutters1)
			case 10:
				value, err = m.home.IsShutterControlDownPressed(homeio.HomeOffice, homeio.Shutters1)
			case 11:
				value, err = m.home.AreShuttersOnTop(homeio.HomeOffice, homeio.Shutters1)
			case 12:
				value, err = m.home.AreShuttersOnBottom(homeio.HomeOffice, homeio.Shutters1)
			case 13, 14:
				value, err = m.home.IsDoorClosed(homeio.HomeOffice, homeio.Door1+homeio.Door(address-13))
			case 15:
				value, err = m.home.IsSmokeDetected(homeio.HomeOffice)
			case 16:
				value, err = m.home.IsLightDetected(homeio.HomeOffice)
			case 17:
				value, err = m.home.IsMotionDetected(homeio.HomeOffice)
			}

		case homeio.Exterior:
			switch address {
			case 0:
				value, err = m.home.IsGateOpen(homeio.Exterior)
			case 1:
				value, err = m.home.IsGateClosed(homeio.Exterior)
			case 2, 3, 4:
				value, err = m.home.IsIRObscured(homeio.Exterior, homeio.EntranceIRSensor1+homeio.IRSensor(address-2))
			case 5:
				value, err = m.home.IsLightDetected(homeio.Exterior)
			case 6:
				value, err = m.home.IsMotionDetected(homeio.Exterior)
			}

		case homeio.Miscellaneous:
			switch address {
			case 0, 1, 2, 3, 4, 5, 6, 7:
				value, err = m.home.IsRemoteControlButtonPressed(homeio.RemoteControlButton1 + homeio.RemoteControlButton(address-0))
			}
		}

		if err != nil {
			return nil, err
		}
		res = append(res, value)
	}

	return
}
