package icalendar

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/google/uuid"
)

func NewICalendar() *ICalendar {
	return &ICalendar{
		Begin:    VCalendar,
		ProdID:   ProdIDWorkplaze,
		Version:  Version2,
		Method:   MethodPublish,
		CalScale: CalScaleGregorian,
		TZID:     TimeZoneJakarta,
		End:      VCalendar,
	}
}

func NewICalendarEvent(summary, dtStart, dtEnd, category, description string) *ICalendarEvent {
	return &ICalendarEvent{
		Begin:       VEvent,
		UID:         uuid.New().String(),
		Summary:     summary,
		DtStart:     dtStart,
		DtEnd:       dtEnd,
		Categories:  category,
		Description: description,
		End:         VEvent,
	}
}

func NewICalendarAlarm(trigger, desc string) *ICalendarAlarm {
	return &ICalendarAlarm{
		Begin:      VAlarm,
		Action:     ActionDisplay,
		Trigger:    trigger,
		Desciption: desc,
		End:        VAlarm,
	}
}

func (iCalendar *ICalendar) SetEvent(event ICalendarEvent) {
	iCalendar.ICalendarEvent = append(iCalendar.ICalendarEvent, event)
}

func (iCalendarEvent *ICalendarEvent) SetAlarm(alarm ICalendarAlarm) {
	iCalendarEvent.ICalendarAlarm = append(iCalendarEvent.ICalendarAlarm, alarm)
}

func (iCalendar ICalendar) GenerateStringICS() string {
	return iCalendar.structToStringWithTag("%v:%v\r\n", "ics")
}

func (iCalendar ICalendar) structToStringWithTag(format, tag string) string {
	return interfaceStructToString(iCalendar, format, tag)
}

func interfaceStructToString(data interface{}, format, tag string) string {
	var (
		stringData strings.Builder
		t          = reflect.TypeOf(data)
		v          = reflect.ValueOf(data)
	)

	if t.Kind() == reflect.Ptr && v.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		val := v.Field(i).Interface()

		switch val.(type) {
		case string:
			tagField := field.Tag.Get(tag)
			if tagField != "" && val != "" {
				fmt.Fprintf(&stringData, format, tagField, val)
			}
		case []ICalendarEvent:
			for _, event := range val.([]ICalendarEvent) {
				stringData.WriteString(interfaceStructToString(event, format, tag))
			}
		case []ICalendarAlarm:
			for _, alarm := range val.([]ICalendarAlarm) {
				stringData.WriteString(interfaceStructToString(alarm, format, tag))
			}
		}
	}

	return stringData.String()
}
