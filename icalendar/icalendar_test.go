package icalendar

import (
	"reflect"
	"testing"
)

func TestNewICalendar(t *testing.T) {
	tests := []struct {
		name string
		want *ICalendar
	}{
		// TODO: Add test cases.
		{
			name: "new icalendar",
			want: &ICalendar{
				Begin:    VCalendar,
				ProdID:   ProdIDWorkplaze,
				Version:  Version2,
				Method:   MethodPublish,
				CalScale: CalScaleGregorian,
				TZID:     TimeZoneJakarta,
				End:      VCalendar,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewICalendar(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewICalendar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewICalendarAlarm(t *testing.T) {
	type args struct {
		trigger string
		desc    string
	}
	tests := []struct {
		name string
		args args
		want *ICalendarAlarm
	}{
		{
			name: "new icalender alarm",
			args: args{
				trigger: "-PT30M",
				desc:    "Reminder",
			},
			want: &ICalendarAlarm{
				Begin:      VAlarm,
				Action:     ActionDisplay,
				Trigger:    "-PT30M",
				Desciption: "Reminder",
				End:        VAlarm,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewICalendarAlarm(tt.args.trigger, tt.args.desc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewICalendarAlarm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_interfaceStructToString(t *testing.T) {
	type args struct {
		data   interface{}
		format string
		tag    string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test 1",
			args: args{
				data:   *NewICalendar(),
				format: "%v:%v\r\n",
				tag:    "ics",
			},
			want: "BEGIN:VCALENDAR\r\nPRODID:Workplaze\r\nVERSION:2.0\r\nMETHOD:Publish\r\nCALSCALE:GREGORIAN\r\nTZID:Asia/Jakarta\r\nEND:VCALENDAR\r\n",
		},
		{
			name: "test 2 pointer",
			args: args{
				data:   NewICalendar(),
				format: "%v:%v\r\n",
				tag:    "ics",
			},
			want: "BEGIN:VCALENDAR\r\nPRODID:Workplaze\r\nVERSION:2.0\r\nMETHOD:Publish\r\nCALSCALE:GREGORIAN\r\nTZID:Asia/Jakarta\r\nEND:VCALENDAR\r\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := interfaceStructToString(tt.args.data, tt.args.format, tt.args.tag); got != tt.want {
				t.Errorf("interfaceStructToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestICalendar_GenerateStringICS(t *testing.T) {
	type fields struct {
		Begin          string
		ProdID         string
		Version        string
		Method         string
		CalScale       string
		TZID           string
		ICalendarEvent []ICalendarEvent
		End            string
	}

	alarm1 := NewICalendarAlarm("-PT30", "Reminder")

	event1 := ICalendarEvent{
		Begin:          VEvent,
		UID:            "5397b5d2-34ea-11ee-be56-0242ac120002",
		Summary:        "Interview HR For Edo",
		DtStart:        "20230807T132000",
		DtEnd:          "20230807T142000",
		DtStamp:        "20230807T132000",
		Categories:     "Recruitment",
		Description:    "You have been set as the Interview HR PIC with applicant Edo for job Web Dev",
		ICalendarAlarm: []ICalendarAlarm{*alarm1},
		End:            VEvent,
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "icalendar1",
			fields: fields{
				Begin:          VCalendar,
				ProdID:         ProdIDWorkplaze,
				Version:        Version2,
				Method:         MethodPublish,
				CalScale:       CalScaleGregorian,
				TZID:           TimeZoneJakarta,
				ICalendarEvent: []ICalendarEvent{event1},
				End:            VCalendar,
			},
			want: "BEGIN:VCALENDAR\r\nPRODID:Workplaze\r\nVERSION:2.0\r\nMETHOD:Publish\r\nCALSCALE:GREGORIAN\r\nTZID:Asia/Jakarta\r\nBEGIN:VEVENT\r\nUID:5397b5d2-34ea-11ee-be56-0242ac120002\r\nSUMMARY:Interview HR For Edo\r\nDTSTART:20230807T132000\r\nDTEND:20230807T142000\r\nDTSTAMP:20230807T132000\r\nCATEGORIES:Recruitment\r\nDESCRIPTION:You have been set as the Interview HR PIC with applicant Edo for job Web Dev\r\nBEGIN:VALARM\r\nACTION:DISPLAY\r\nTRIGGER;RELATED=START:-PT30\r\nDESCRIPTION:Reminder\r\nEND:VALARM\r\nEND:VEVENT\r\nEND:VCALENDAR\r\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			iCalendar := ICalendar{
				Begin:          tt.fields.Begin,
				ProdID:         tt.fields.ProdID,
				Version:        tt.fields.Version,
				Method:         tt.fields.Method,
				CalScale:       tt.fields.CalScale,
				TZID:           tt.fields.TZID,
				ICalendarEvent: tt.fields.ICalendarEvent,
				End:            tt.fields.End,
			}
			if got := iCalendar.GenerateStringICS(); got != tt.want {
				t.Errorf("ICalendar.GenerateStringICS() = %v, want %v", got, tt.want)
			}
		})
	}
}
