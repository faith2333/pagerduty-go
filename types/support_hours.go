package types

import "time"

type SupportHours struct {
	// The type of support hours
	//  Allowed value: fixed_time_per_day
	//  Default: fixed_time_per_day
	Type Type `json:"type"`
	// The time zone for the support hours
	TimeZone   string `json:"time_zone"`
	DaysOfWeek []int  `json:"days_of_week"`
	// The support hours' starting time of day (date portion is ignored)
	StartTime time.Time `json:"start_time"`
	// The support hours' ending time of day (date portion is ignored)
	EndTime time.Time `json:"end_time"`
}
