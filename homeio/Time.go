package homeio

import (
	"fmt"
	"time"
)

/* Home interface implementation */

func (h *home) GetTime() (time.Time, error) {
	// Get all the time components from the inputs map.
	year, err := h.inputs["year"]
	if !err {
		return time.Time{}, fmt.Errorf("year not found")
	}
	month, err := h.inputs["month"]
	if !err {
		return time.Time{}, fmt.Errorf("month not found")
	}
	day, err := h.inputs["day"]
	if !err {
		return time.Time{}, fmt.Errorf("day not found")
	}
	hour, err := h.inputs["hour"]
	if !err {
		return time.Time{}, fmt.Errorf("hour not found")
	}
	minute, err := h.inputs["minute"]
	if !err {
		return time.Time{}, fmt.Errorf("minute not found")
	}
	second, err := h.inputs["second"]
	if !err {
		return time.Time{}, fmt.Errorf("second not found")
	}

	// Return the time.
	return time.Date(int(year.(int64)), time.Month(month.(int64)), int(day.(int64)), int(hour.(int64)), int(minute.(int64)), int(second.(int64)), 0, time.Local), nil
}

func (h *home) GetYear() (uint16, error) {
	// Read input value from inputs map and check if it exists.
	value, err := h.inputs["year"]
	if !err {
		return 0, fmt.Errorf("year not found")
	}

	// Return the year.
	return uint16(value.(int64)), nil
}

func (h *home) GetMonth() (uint16, error) {
	// Read input value from inputs map and check if it exists.
	value, err := h.inputs["month"]
	if !err {
		return 0, fmt.Errorf("month not found")
	}

	// Return the month.
	return uint16(value.(int64)), nil
}

func (h *home) GetDay() (uint16, error) {
	// Read input value from inputs map and check if it exists.
	value, err := h.inputs["day"]
	if !err {
		return 0, fmt.Errorf("day not found")
	}

	// Return the day.
	return uint16(value.(int64)), nil
}

func (h *home) GetHour() (uint16, error) {
	// Read input value from inputs map and check if it exists.
	value, err := h.inputs["hour"]
	if !err {
		return 0, fmt.Errorf("hour not found")
	}

	// Return the hour.
	return uint16(value.(int64)), nil
}

func (h *home) GetMinute() (uint16, error) {
	// Read input value from inputs map and check if it exists.
	value, err := h.inputs["minute"]
	if !err {
		return 0, fmt.Errorf("minute not found")
	}

	// Return the minute.
	return uint16(value.(int64)), nil
}

func (h *home) GetSecond() (uint16, error) {
	// Read input value from inputs map and check if it exists.
	value, err := h.inputs["second"]
	if !err {
		return 0, fmt.Errorf("second not found")
	}

	// Return the second.
	return uint16(value.(int64)), nil
}
