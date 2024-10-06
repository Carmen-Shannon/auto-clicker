package dto

import "time"

type DelayValuesDto struct {
	Minutes       int
	Seconds       int
	Ms            int
	RandomMinutes int
	RandomSeconds int
	RandomMs      int
	Delay         int
	Clicks        int
}

func NewDelayValues(m int, s int, ms int, rm int, rs int, rms int) *DelayValuesDto {
	return &DelayValuesDto{
		Minutes:       m,
		Seconds:       s,
		Ms:            ms,
		RandomMinutes: rm,
		RandomSeconds: rs,
		RandomMs:      rms,
	}
}

// CalculateDelay will return 2 int values, the delay and the extra delay respectively, both in time.Duration (nanoseconds)
func (d *DelayValuesDto) CalculateDelay() (int, int) {
	delay := (time.Duration(d.Minutes) * time.Minute) + (time.Duration(d.Seconds) * time.Second) + (time.Duration(d.Ms) * time.Millisecond)
	extraDelay := (time.Duration(d.RandomMinutes) * time.Minute) + (time.Duration(d.RandomSeconds) * time.Second) + (time.Duration(d.RandomMs) * time.Millisecond)

	if delay <= 0 {
		delay = time.Duration(1) * time.Second
	}

	return int(delay), int(extraDelay)
}

func (d *DelayValuesDto) SetClicks(clicks int) {
	d.Clicks = clicks
}
