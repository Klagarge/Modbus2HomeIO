package homeio

import (
	"fmt"
	"time"
)

func (h *home) GetTime() (time.Time, error) {
	year, err := h.values["year"]
	if !err {
		return time.Time{}, fmt.Errorf("year not found")
	}
	month, err := h.values["month"]
	if !err {
		return time.Time{}, fmt.Errorf("month not found")
	}
	day, err := h.values["day"]
	if !err {
		return time.Time{}, fmt.Errorf("day not found")
	}
	hour, err := h.values["hour"]
	if !err {
		return time.Time{}, fmt.Errorf("hour not found")
	}
	minute, err := h.values["minute"]
	if !err {
		return time.Time{}, fmt.Errorf("minute not found")
	}
	second, err := h.values["second"]
	if !err {
		return time.Time{}, fmt.Errorf("second not found")
	}
	return time.Date(int(year.(int64)), time.Month(month.(int64)), int(day.(int64)), int(hour.(int64)), int(minute.(int64)), int(second.(int64)), 0, time.Local), nil
}
