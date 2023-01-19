package Repository

import (
	models "backend/Models"
	structure "backend/Structure"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) AddStudent(w http.ResponseWriter, re *http.Request) {
	student := structure.Student{}
    

	err := json.NewDecoder(re.Body).Decode(&student)
	if err != nil {
		json.NewEncoder(w).Encode("Request Failed")
		return
	}
	err = r.DB.Create(&student).Error
	if err != nil {
		json.NewEncoder(w).Encode("Failed to add student")
		return
	}
	json.NewEncoder(w).Encode("Student Added successfully")

}

func (r *Repository) AddTeacher(w http.ResponseWriter, re *http.Request) {

	teacher := structure.Teacher{}

	err := json.NewDecoder(re.Body).Decode(&teacher)

	if err != nil {
		json.NewEncoder(w).Encode("Request Failed")
	}
	err = r.DB.Create(&teacher).Error
	if err != nil {
		json.NewEncoder(w).Encode("Failed to add Teacher")
	}
	json.NewEncoder(w).Encode("Teacher Added successfully")

}

func (r *Repository) GetTeacherAttendence(w http.ResponseWriter, re *http.Request) {
	var attendences []structure.Attendence
	params := mux.Vars(re)

	r.DB.Where("a_id = ? AND month=? AND year=? AND type=?", params["id"], params["month"], params["year"], "Teacher").Find(&attendences)

	json.NewEncoder(w).Encode(attendences)
}

func (r *Repository) PunchIn(w http.ResponseWriter, re *http.Request) {
	t := time.Now()
	attendence := &structure.Attendence{}

	err := json.NewDecoder(re.Body).Decode(attendence)
	if err != nil {
		json.NewEncoder(w).Encode("Request Failed")
		return
	}
	if attendence.Type == "Student" {
		var student models.Student
		r.DB.Where("s_id = ?", attendence.AID).First(&student)
		if student.Name == "" {
			json.NewEncoder(w).Encode("No Student Found")
			return
		}
		attendence.Class = student.Class
	} else {
		var teacher models.Teacher
		r.DB.Where("t_id = ?", attendence.AID).First(&teacher)
		if teacher.Name == "" {
			json.NewEncoder(w).Encode("No Teacher Found")
			return
		}
	}
	attendence.Day = strconv.Itoa(t.Day())
	attendence.Month = strconv.Itoa(int(t.Month()))
	attendence.Year = strconv.Itoa(t.Year())
	var att models.Attendence
	r.DB.Where("a_id=? AND day=? AND month=? AND year=? AND type=? ", attendence.AID, attendence.Day, attendence.Month, attendence.Year, attendence.Type).Last(&att)
	fmt.Print(att)
	if att.Punchin != "" && att.Punchout == "" {
		json.NewEncoder(w).Encode("Already Punch In")
		return
	}
	attendence.Punchin = (t.Format("15:04:05"))

	err = r.DB.Create(attendence).Error
	if err != nil {
		json.NewEncoder(w).Encode("Failed to punch in")
		return
	}
	json.NewEncoder(w).Encode("Puch In Successfully")

}

func (r *Repository) PuchOut(w http.ResponseWriter, re *http.Request) {
	t := time.Now()
	attendence := &structure.Attendence{}

	err := json.NewDecoder(re.Body).Decode(attendence)
	if err != nil {
		json.NewEncoder(w).Encode("Request Failed")
		return
	}
	if attendence.Type == "Student" {
		var student models.Student
		r.DB.Where("s_id = ?", attendence.AID).First(&student)
		if student.Name == "" {
			json.NewEncoder(w).Encode("No Student Found")
			return
		}
	} else {
		var teacher models.Teacher
		r.DB.Where("t_id = ?", attendence.AID).First(&teacher)
		if teacher.Name == "" {
			json.NewEncoder(w).Encode("No Teacher Found")
			return
		}
	}
	attendence.Day = strconv.Itoa(t.Day())
	attendence.Month = strconv.Itoa(int(t.Month()))
	attendence.Year = strconv.Itoa(t.Year())
	var att models.Attendence
	r.DB.Where("a_id=? AND day=? AND month=? AND year=? AND type=?", attendence.AID, attendence.Day, attendence.Month, attendence.Year, attendence.Type).Last(&att)
	if att.Punchin == "" {
		json.NewEncoder(w).Encode("You have not punch in yet")
		return
	}
	if att.Punchout != "" {
		json.NewEncoder(w).Encode("Already Puched Out")
		return
	}
	attendence.Class = att.Class
	attendence.Punchin = att.Punchin
	attendence.Puchout = (t.Format("15:04:05"))
	r.DB.Where("a_id=? AND day=? AND month=? AND year=? AND type=? AND punchin=?", attendence.AID, attendence.Day, attendence.Month, attendence.Year, attendence.Type, attendence.Punchin).Delete(&models.Attendence{})
	err = r.DB.Create(attendence).Error
	if err != nil {
		json.NewEncoder(w).Encode("Failed to punch out")
		return
	}
	json.NewEncoder(w).Encode("Punch Out Successfully")
}

func (r *Repository) GetClassAttendence(w http.ResponseWriter, re *http.Request) {

	var attendences []structure.Attendence
	params := mux.Vars(re)

	r.DB.Where("class = ? AND day=? AND month=? AND year=?", params["id"], params["day"], params["month"], params["year"]).Find(&attendences)

	json.NewEncoder(w).Encode(attendences)

}

func (r *Repository) GetStudentAttendence(w http.ResponseWriter, re *http.Request) {
	var attendences []structure.Attendence
	params := mux.Vars(re)

	r.DB.Where("a_id = ? AND month=? AND year=? AND type=?", params["id"], params["month"], params["year"], "Student").Find(&attendences)

	json.NewEncoder(w).Encode(attendences)
}