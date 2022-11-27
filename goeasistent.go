package goeasistent

import (
	"encoding/json"
	"time"

	"github.com/LovroG05/goeasistent/endpoints"
	"github.com/LovroG05/goeasistent/objects"
	gottf "github.com/OpenTimetable/GOTTF"
	gottfobjects "github.com/OpenTimetable/GOTTF/objects"
	"github.com/OpenTimetable/GOTTF/parsers"
)

type Goeasistent interface {
	Login(username string, password string) error
	RefreshToken(sessionToken string) error
	GetChild() (objects.ChildData, error)
	GetTimetable(from string, to string) (objects.Timetable, error)
	GetGrades()
}

type Instance struct {
	UserData objects.UserData
}

func (i *Instance) Login(username string, password string) error {
	userjson, err := endpoints.Login(username, password)
	if err != nil {
		return err
	}

	var userdata objects.UserData

	err = json.Unmarshal([]byte(userjson), &userdata)
	if err != nil {
		return err
	}

	i.UserData = userdata

	return nil
}

func (i *Instance) RefreshToken(sessionToken string) error {
	tokens, err := endpoints.RefreshToken(i.UserData.Tokens.RefreshToken, sessionToken)
	if err != nil {
		return err
	}

	i.UserData.Tokens = tokens

	return nil
}

func (i *Instance) GetChild() (objects.ChildData, error) {
	childjson, err := endpoints.GetChild(i.UserData.Tokens.AccessToken.Token)
	if err != nil {
		return objects.ChildData{}, err
	}

	var childdata objects.ChildData

	err = json.Unmarshal([]byte(childjson), &childdata)
	if err != nil {
		return objects.ChildData{}, err
	}

	return childdata, nil
}

func (i *Instance) GetTimetable(from string, to string) (objects.Timetable, error) {
	timetablejson, err := endpoints.Timetable(i.UserData.Tokens.AccessToken.Token, from, to)
	if err != nil {
		return objects.Timetable{}, err
	}

	var timetable objects.Timetable

	err = json.Unmarshal([]byte(timetablejson), &timetable)
	if err != nil {
		return objects.Timetable{}, err
	}

	return timetable, nil
}

func (i *Instance) GetGrades() (objects.Grades, error) {
	gradesJson, err := endpoints.Grades(i.UserData.Tokens.AccessToken.Token)
	if err != nil {
		return objects.Grades{}, err
	}

	var grades objects.Grades

	err = json.Unmarshal([]byte(gradesJson), &grades)
	if err != nil {
		return objects.Grades{}, err
	}

	return grades, nil
}

func ComposeOTTF(timetable objects.Timetable) parsers.Timetable {
	version := "1.0"

	periods := make(map[string]gottfobjects.Span)
	var recesses []gottfobjects.Span
	// Get periods
	for _, time := range timetable.TimeTable {
		// TODO change times to UTC to better comply with the OTTF standard
		var span gottfobjects.Span
		timestr, _ := json.Marshal(time.Time)
		err := json.Unmarshal(timestr, &span)
		if err != nil {
			return parsers.Timetable{}
		}
		if time.Type == "default" {
			id := time.Name
			periods[id] = span
		} else if time.Type == "break" {
			recesses = append(recesses, span)
		}
	}

	cues := gottfobjects.Cues{
		Periods:  periods,
		Recesses: recesses,
	}

	days := make(map[string]gottfobjects.Day)

	for _, day := range timetable.DayTable {
		var classes map[string][]gottfobjects.Class
		var events []gottfobjects.Event
		var dayEvents []gottfobjects.DayEvent

		for _, class := range timetable.SchoolHourEvents {
			var hosts []string
			for hid := range class.Teachers {
				hosts = append(hosts, class.Teachers[hid].Name)
			}

			var cancelled bool
			var exam bool
			var sub bool
			switch *class.HourSpecialType {
			case "cancelled":
				cancelled = true
			case "exam":
				exam = true
			case "substitution":
				sub = true
			}

			newClass := gottfobjects.Class{
				Name:         class.Subject.Name,
				Abbreviation: class.Subject.Name,
				Hosts:        hosts,
				Location:     class.Classroom.Name,
				Substitution: sub,
				Examination:  exam,
				Canceled:     cancelled,
			}

			period := getPeriod(timetable.TimeTable, class.Time)

			if classes[period] == nil {
				classes[period] = []gottfobjects.Class{}
			}

			classes[period] = append(classes[period], newClass)

		}

		for _, event := range timetable.Events {
			var hosts []string
			for hid := range event.Teachers {
				hosts = append(hosts, event.Teachers[hid].Name)
			}

			events = append(events, gottfobjects.Event{
				From:     event.Time.From,
				To:       event.Time.To,
				Title:    event.Name,
				Location: event.Location.Name,
				Hosts:    hosts,
			})
		}

		for _, dayEvent := range timetable.AllDayEvents {
			var hosts []string
			for hid := range dayEvent.Teachers {
				hosts = append(hosts, dayEvent.Teachers[hid].Name)
			}

			dayEvents = append(dayEvents, gottfobjects.DayEvent{
				Title:    dayEvent.Name,
				Location: "",
				Hosts:    hosts,
			})
		}

		newDay := gottfobjects.Day{
			Classes:   classes,
			Events:    events,
			DayEvents: dayEvents,
		}

		date := day.Date

		days[date] = newDay
	}

	ottfTimetable := parsers.Timetable{
		Metadata: gottfobjects.Metadata{
			Version:   version,
			Author:    "eAsistent via goeasistent by LovroG05",
			Timezone:  "UTC+2",
			Timestamp: uint64(time.Now().Unix()),
		},
		Cues: cues,
		Days: days,
	}

	return ottfTimetable
}

func getPeriod(timetableItems []objects.TimeTableItem, time objects.TimeBind) string {
	for _, t := range timetableItems {
		if t.ID == time.FromID || t.ID == time.ToID {
			return t.Name
		}
	}
	return ""
}

func ComposeOTTFJsonString(timetable objects.Timetable) string {
	ottfTimetable := ComposeOTTF(timetable)
	ottfJson, err := gottf.ComposeTimetable(ottfTimetable, "1.0")
	if err != nil {
		return ""
	}
	return ottfJson
}
