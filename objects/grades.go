package objects

type Grades struct {
	Subjects []struct {
		Name         string      `json:"name"`
		ShortName    string      `json:"short_name"`
		Id           int         `json:"id"`
		GradeType    string      `json:"grade_type"`
		IsExcused    bool        `json:"is_excused"`
		FinalGrade   interface{} `json:"final_grade"`
		AverageGrade string      `json:"average_grade"`
		GradeRank    string      `json:"grade_rank"`
		Semesters    []struct {
			Id         int         `json:"id"`
			FinalGrade interface{} `json:"final_grade"`
			Grades     []struct {
				TypeName     string      `json:"type_name"`
				Comment      interface{} `json:"comment"`
				Id           int         `json:"id"`
				Type         string      `json:"type"`
				OverridesIds interface{} `json:"overrides_ids"`
				Value        string      `json:"value"`
				Color        string      `json:"color"`
			}
		}
	} `json:"items"`
}
