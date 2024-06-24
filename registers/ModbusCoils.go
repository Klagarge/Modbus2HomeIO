package registers

import (
	"Modbus2HomeIO/homeio"
	"fmt"
	"github.com/simonvetter/modbus"
)

/* Handler interface implementation, handle Read Coils (0x01) and Write Coils (0x15) */

func (m *Handler) HandleCoils(req *modbus.CoilsRequest) (res []bool, err error) {
	// We only support writing to coils.
	if !req.IsWrite {
		return nil, fmt.Errorf("read access to coils is not supported")
	}

	// Get room from unit ID and fail if the room does not exist.
	room, err := homeio.UnitIDToRoom(req.UnitId)
	if err != nil {
		return []bool{}, fmt.Errorf("invalid unit ID %v", req.UnitId)
	}

	for address := req.Addr; address < req.Addr+req.Quantity; address++ {
		switch room {
		case homeio.LivingRoom:
			switch address {
			case 0:
				err = m.home.SetLightOn(homeio.LivingRoom, homeio.Light1, req.Args[address-req.Addr])
			case 1, 2, 3, 4:
				err = m.home.SetShuttersDirection(homeio.LivingRoom, homeio.Shutters1+homeio.Shutters(address-1), map[bool]homeio.ShuttersDirection{false: homeio.ShuttersStop, true: homeio.ShuttersUp}[req.Args[address-req.Addr]])
			case 5, 6, 7, 8:
				err = m.home.SetShuttersDirection(homeio.LivingRoom, homeio.Shutters1+homeio.Shutters(address-5), map[bool]homeio.ShuttersDirection{false: homeio.ShuttersStop, true: homeio.ShuttersDown}[req.Args[address-req.Addr]])
			case 9:
				err = m.home.SetHeatingOn(homeio.LivingRoom, req.Args[address-req.Addr])
			}

		case homeio.GuestRestRoom:
			switch address {
			case 0, 1:
				err = m.home.SetLightOn(homeio.GuestRestRoom, homeio.Light1+homeio.Light(address-0), req.Args[address-req.Addr])
			}

		case homeio.Pantry:
			switch address {
			case 0:
				err = m.home.SetLightOn(homeio.Pantry, homeio.Light1, req.Args[address-req.Addr])
			}

		case homeio.Kitchen:
			switch address {
			case 0, 1:
				err = m.home.SetLightOn(homeio.Kitchen, homeio.Light1+homeio.Light(address-0), req.Args[address-req.Addr])
			case 2:
				err = m.home.SetShuttersDirection(homeio.Kitchen, homeio.Shutters1, map[bool]homeio.ShuttersDirection{false: homeio.ShuttersStop, true: homeio.ShuttersUp}[req.Args[address-req.Addr]])
			case 3:
				err = m.home.SetShuttersDirection(homeio.Kitchen, homeio.Shutters1, map[bool]homeio.ShuttersDirection{false: homeio.ShuttersStop, true: homeio.ShuttersDown}[req.Args[address-req.Addr]])
			case 4:
				err = m.home.SetHeatingOn(homeio.Kitchen, req.Args[address-req.Addr])
			}

		case homeio.EntranceHall:
			switch address {
			case 0:
				err = m.home.SetLightOn(homeio.EntranceHall, homeio.Light1, req.Args[address-req.Addr])
			case 1:
				err = m.home.SetShuttersDirection(homeio.EntranceHall, homeio.Shutters1, map[bool]homeio.ShuttersDirection{false: homeio.ShuttersStop, true: homeio.ShuttersUp}[req.Args[address-req.Addr]])
			case 2:
				err = m.home.SetShuttersDirection(homeio.EntranceHall, homeio.Shutters1, map[bool]homeio.ShuttersDirection{false: homeio.ShuttersStop, true: homeio.ShuttersDown}[req.Args[address-req.Addr]])
			case 3:
				err = m.home.SetHeatingOn(homeio.EntranceHall, req.Args[address-req.Addr])
			case 4:
				err = m.home.SetSirenOn(homeio.EntranceHall, req.Args[address-req.Addr])
			case 5:
				err = m.home.SetAlarmArmed(req.Args[address-req.Addr])
			}

		case homeio.Garage:
			switch address {
			case 0, 1:
				err = m.home.SetLightOn(homeio.Garage, homeio.Light1+homeio.Light(address-0), req.Args[address-req.Addr])
			case 2:
				err = m.home.SetShuttersDirection(homeio.Garage, homeio.Shutters1, map[bool]homeio.ShuttersDirection{false: homeio.ShuttersStop, true: homeio.ShuttersUp}[req.Args[address-req.Addr]])
			case 3:
				err = m.home.SetShuttersDirection(homeio.Garage, homeio.Shutters1, map[bool]homeio.ShuttersDirection{false: homeio.ShuttersStop, true: homeio.ShuttersDown}[req.Args[address-req.Addr]])
			case 4:
				err = m.home.SetGateDirection(homeio.Garage, map[bool]homeio.GateDirection{false: homeio.GateStop, true: homeio.GateOpen}[req.Args[address-req.Addr]])
			case 5:
				err = m.home.SetGateDirection(homeio.Garage, map[bool]homeio.GateDirection{false: homeio.GateStop, true: homeio.GateClose}[req.Args[address-req.Addr]])
			}

		case homeio.BedroomCorridor:
			switch address {
			case 0:
				err = m.home.SetLightOn(homeio.BedroomCorridor, homeio.Light1, req.Args[address-req.Addr])
			case 1:
				err = m.home.SetShuttersDirection(homeio.BedroomCorridor, homeio.Shutters1, map[bool]homeio.ShuttersDirection{false: homeio.ShuttersStop, true: homeio.ShuttersUp}[req.Args[address-req.Addr]])
			case 2:
				err = m.home.SetShuttersDirection(homeio.BedroomCorridor, homeio.Shutters1, map[bool]homeio.ShuttersDirection{false: homeio.ShuttersStop, true: homeio.ShuttersDown}[req.Args[address-req.Addr]])
			case 3:
				err = m.home.SetHeatingOn(homeio.BedroomCorridor, req.Args[address-req.Addr])
			}

		case homeio.ChildrenRoom:
			switch address {
			case 0:
				err = m.home.SetLightOn(homeio.ChildrenRoom, homeio.Light1, req.Args[address-req.Addr])
			case 1:
				err = m.home.SetShuttersDirection(homeio.ChildrenRoom, homeio.Shutters1, map[bool]homeio.ShuttersDirection{false: homeio.ShuttersStop, true: homeio.ShuttersUp}[req.Args[address-req.Addr]])
			case 2:
				err = m.home.SetShuttersDirection(homeio.ChildrenRoom, homeio.Shutters1, map[bool]homeio.ShuttersDirection{false: homeio.ShuttersStop, true: homeio.ShuttersDown}[req.Args[address-req.Addr]])
			case 3:
				err = m.home.SetHeatingOn(homeio.ChildrenRoom, req.Args[address-req.Addr])
			}

		case homeio.Bathroom:
			switch address {
			case 0, 1:
				err = m.home.SetLightOn(homeio.Bathroom, homeio.Light1+homeio.Light(address-0), req.Args[address-req.Addr])
			case 2:
				err = m.home.SetHeatingOn(homeio.Bathroom, req.Args[address-req.Addr])
			}

		case homeio.SingleBedroom:
			switch address {
			case 0:
				err = m.home.SetLightOn(homeio.SingleBedroom, homeio.Light1, req.Args[address-req.Addr])
			case 1:
				err = m.home.SetShuttersDirection(homeio.SingleBedroom, homeio.Shutters1, map[bool]homeio.ShuttersDirection{false: homeio.ShuttersStop, true: homeio.ShuttersUp}[req.Args[address-req.Addr]])
			case 2:
				err = m.home.SetShuttersDirection(homeio.SingleBedroom, homeio.Shutters1, map[bool]homeio.ShuttersDirection{false: homeio.ShuttersStop, true: homeio.ShuttersDown}[req.Args[address-req.Addr]])
			case 3:
				err = m.home.SetHeatingOn(homeio.SingleBedroom, req.Args[address-req.Addr])
			}

		case homeio.PrivateBathroom:
			switch address {
			case 0:
				err = m.home.SetLightOn(homeio.PrivateBathroom, homeio.Light1, req.Args[address-req.Addr])
			case 1:
				err = m.home.SetHeatingOn(homeio.PrivateBathroom, req.Args[address-req.Addr])
			}

		case homeio.CoupleBedroom:
			switch address {
			case 0:
				err = m.home.SetLightOn(homeio.CoupleBedroom, homeio.Light1, req.Args[address-req.Addr])
			case 1:
				err = m.home.SetShuttersDirection(homeio.CoupleBedroom, homeio.Shutters1, map[bool]homeio.ShuttersDirection{false: homeio.ShuttersStop, true: homeio.ShuttersUp}[req.Args[address-req.Addr]])
			case 2:
				err = m.home.SetShuttersDirection(homeio.CoupleBedroom, homeio.Shutters1, map[bool]homeio.ShuttersDirection{false: homeio.ShuttersStop, true: homeio.ShuttersDown}[req.Args[address-req.Addr]])
			case 3:
				err = m.home.SetHeatingOn(homeio.CoupleBedroom, req.Args[address-req.Addr])
			}

		case homeio.LaundryRoom:
			switch address {
			case 0:
				err = m.home.SetLightOn(homeio.LaundryRoom, homeio.Light1, req.Args[address-req.Addr])
			case 1:
				err = m.home.SetShuttersDirection(homeio.LaundryRoom, homeio.Shutters1, map[bool]homeio.ShuttersDirection{false: homeio.ShuttersStop, true: homeio.ShuttersUp}[req.Args[address-req.Addr]])
			case 2:
				err = m.home.SetShuttersDirection(homeio.LaundryRoom, homeio.Shutters1, map[bool]homeio.ShuttersDirection{false: homeio.ShuttersStop, true: homeio.ShuttersDown}[req.Args[address-req.Addr]])
			case 3:
				err = m.home.SetHeatingOn(homeio.LaundryRoom, req.Args[address-req.Addr])
			}

		case homeio.HomeOffice:
			switch address {
			case 0, 1, 2:
				err = m.home.SetLightOn(homeio.HomeOffice, homeio.Light1+homeio.Light(address-0), req.Args[address-req.Addr])
			case 3:
				err = m.home.SetShuttersDirection(homeio.HomeOffice, homeio.Shutters1, map[bool]homeio.ShuttersDirection{false: homeio.ShuttersStop, true: homeio.ShuttersUp}[req.Args[address-req.Addr]])
			case 4:
				err = m.home.SetShuttersDirection(homeio.HomeOffice, homeio.Shutters1, map[bool]homeio.ShuttersDirection{false: homeio.ShuttersStop, true: homeio.ShuttersDown}[req.Args[address-req.Addr]])
			case 5:
				err = m.home.SetHeatingOn(homeio.HomeOffice, req.Args[address-req.Addr])
			}

		case homeio.Exterior:
			switch address {
			case 0, 1, 2, 3, 4:
				err = m.home.SetLightOn(homeio.Exterior, homeio.EntranceLight+homeio.Light(address-0), req.Args[address-req.Addr])
			case 5:
				err = m.home.SetGateDirection(homeio.Exterior, map[bool]homeio.GateDirection{false: homeio.GateStop, true: homeio.GateOpen}[req.Args[address-req.Addr]])
			case 6:
				err = m.home.SetGateDirection(homeio.Exterior, map[bool]homeio.GateDirection{false: homeio.GateStop, true: homeio.GateClose}[req.Args[address-req.Addr]])
			case 7:
				err = m.home.SetSirenOn(homeio.Exterior, req.Args[address-req.Addr])
			}
		}

		if err != nil {
			return nil, err
		}
	}

	return
}
