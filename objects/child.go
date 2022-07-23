package objects

type ChildData struct {
	Age         int    `json:"age"`
	AgeLevel    string `json:"age_level"`
	DidTryPlus  bool   `json:"did_try_plus"`
	DisplayName string `json:"display_name"`
	Gender      string `json:"gender"`
	ID          int    `json:"id"`
	Language    string `json:"language"`
	PlusEnabled bool   `json:"plus_enabled"`
	ShortName   string `json:"short_name"`
	StudentID   int    `json:"student_id"`
	Timetable   struct {
		Date  string `json:"date"`
		Hours []struct {
			From     string `json:"from"`
			Metadata struct {
				TomorrowEnd    string `json:"tomorrow_end"`
				TomorrowInfo   string `json:"tomorrow_info"`
				TomorrowNormal bool   `json:"tomorrow_normal"`
				TomorrowStart  string `json:"tomorrow_start"`
			} `json:"metadata"`
			Summary string `json:"summary"`
			To      string `json:"to"`
			Type    string `json:"type"`
		} `json:"hours"`
	} `json:"timetable"`
	Trial     bool   `json:"trial"`
	TrialEnds string `json:"trial_ends"`
	Type      string `json:"type"`
}
