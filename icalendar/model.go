package icalendar

type ICalendar struct {
	Begin          string           `ics:"BEGIN" json:"BEGIN"`
	ProdID         string           `ics:"PRODID" json:"PRODID"`
	Version        string           `ics:"VERSION" json:"VERSION"`
	Method         string           `ics:"METHOD" json:"METHOD"`
	CalScale       string           `ics:"CALSCALE" json:"CALSCALE"`
	TZID           string           `ics:"TZID" json:"TZID"`
	ICalendarEvent []ICalendarEvent `json:"ICALENDAR_EVENT"`
	End            string           `ics:"END" json:"END"`
}

type ICalendarEvent struct {
	Begin          string           `ics:"BEGIN" json:"BEGIN"`
	UID            string           `ics:"UID" json:"UID"`
	Summary        string           `ics:"SUMMARY" json:"SUMMARY"`
	Sequence       string           `ics:"SEQUENCE" json:"SEQUENCE"`
	Status         string           `ics:"STATUS" json:"STATUS"`
	TransP         string           `ics:"TRANSP" json:"TRANSP"`
	DtStart        string           `ics:"DTSTART" json:"DTSTART"`
	DtEnd          string           `ics:"DTEND" json:"DTEND"`
	DtStamp        string           `ics:"DTSTAMP" json:"DTSTAMP"`
	Categories     string           `ics:"CATEGORIES" json:"CATEGORIES"`
	Organizer      string           `ics:"ORGANIZER;CN=System" json:"ORGANIZER"`
	Description    string           `ics:"DESCRIPTION" json:"DESCRIPTION"`
	ICalendarAlarm []ICalendarAlarm `json:"ICALENDAR_ALARM"`
	End            string           `ics:"END" json:"END"`
}

type ICalendarAlarm struct {
	Begin      string `ics:"BEGIN" json:"BEGIN"`
	Action     string `ics:"ACTION" json:"ACTION"`
	Trigger    string `ics:"TRIGGER;RELATED=START" json:"TRIGGER"`
	Desciption string `ics:"DESCRIPTION" json:"DESCRIPTION"`
	End        string `ics:"END" json:"END"`
}

// constant
const (
	VCalendar = "VCALENDAR"
	VEvent    = "VEVENT"
	VAlarm    = "VALARM"

	Version2 = "2.0"

	ProdIDWorkplaze = "Workplaze"
	TimeZoneJakarta = "Asia/Jakarta"

	CalScaleGregorian = "GREGORIAN"

	MethodPublish = "Publish"

	ActionDisplay = "DISPLAY"
)
