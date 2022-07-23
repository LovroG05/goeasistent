package objects

type Timetable struct {
	DayTable []struct {
		Date      string `json:"date"`
		Name      string `json:"name"`
		ShortName string `json:"short_name"`
	} `json:"day_table"`
	Events []struct {
		Date      string `json:"date"`
		EventType int    `json:"event_type"`
		ID        int    `json:"id"`
		Location  struct {
			Name      string `json:"name"`
			ShortName string `json:"short_name"`
		} `json:"location"`
		Name     string   `json:"name"`
		Teachers []IDName `json:"teachers"`
		Time     Time     `json:"time"`
	} `json:"events"`
	SchoolHourEvents []struct {
		Classroom       IDName   `json:"classroom"`
		Color           string   `json:"color"`
		Completed       bool     `json:"completed"`
		Departments     []IDName `json:"departments"`
		EventID         int      `json:"event_id"`
		Groups          []IDName `json:"groups"`
		HourSpecialType *string  `json:"hour_special_type"`
		Subject         IDName   `json:"subject"`
		Teachers        []IDName `json:"teachers"`
		Time            TimeBind `json:"time"`
		Videokonferenca struct {
		} `json:"videokonferenca"`
	} `json:"school_hour_events"`
	TimeTable    []TimeTableItem `json:"time_table"`
	AllDayEvents []struct {
		Date     string   `json:"date"`
		Name     string   `json:"name"`
		Teachers []IDName `json:"teachers"`
	}
}
type IDName struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
type Time struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type TimeTableItem struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	NameShort string `json:"name_short"`
	Time      Time   `json:"time"`
	Type      string `json:"type"`
}
type TimeBind struct {
	Date   string `json:"date"`
	FromID int    `json:"from_id"`
	ToID   int    `json:"to_id"`
}
