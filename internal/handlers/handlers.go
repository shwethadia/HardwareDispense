package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	

	"github.com/shwethadia/BOOKINGAPI/internal/config"
	"github.com/shwethadia/BOOKINGAPI/internal/models"
	"github.com/shwethadia/BOOKINGAPI/internal/render"

)

//Repo ...
var Repo *Repository

//Repository ...
type Repository struct {
	App *config.AppConfig
}


//Employees ...
type Employees struct{
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string  `json:"email"`
	DateofIssue string `json:"date_of_issue"`
	City  string    `json:"city"`
	DeviceName string  `json:"device_name"`
	DeviceID  string `json:"device_id"`
	StateName string `json:"state_name"`
	PhoneNo string `json:"phone_no"`


}


//NewRepo ...
func NewRepo(a *config.AppConfig) *Repository {

	return &Repository{
		App: a,
	}
}

//NewHandlers ...
func NewHandlers(r *Repository) {

	Repo = r
}

//Home ...
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {

	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w,r,"home.page.html", &models.TemplateData{})
}

//About ...
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	//Perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello,again"
	remoteIP := m.App.Session.GetString(r.Context(),"remote_ip")
	stringMap["remote_ip"] = remoteIP
	//Send the data to the template
	render.RenderTemplate(w, r,"about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}

//Laptops ...
func (m *Repository) Laptops(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w,r,"laptops.page.html", &models.TemplateData{})
}

//Dongals ...
func (m *Repository) Dongals(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w,r, "dongals.page.html", &models.TemplateData{})
}


//AllortHardware ...
func (m *Repository) AllortHardware(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w,r, "allorthardware.page.html", &models.TemplateData{})
}

//PostAllortHardware ...
func (m *Repository) PostAllortHardware(w http.ResponseWriter, r *http.Request) {

	firstName:= r.Form.Get("firstName")
	lastName:= r.Form.Get("lastName")
	email := r.Form.Get("email")
	dateofIssue := r.Form.Get("date_of_issue")
	city := r.Form.Get("city")
	DeviceName := r.Form.Get("device_name")
	DeviceID:= r.Form.Get("device_id")
	stateName := r.Form.Get("state_name")
	phoneNo := r.Form.Get("phone_no")
	w.Write([]byte(fmt.Sprintf("FirstName=%s\nLastName=%s\nEmail=%s\nDateofIsuue=%s\nCity=%s\nDeviceName=%s\nDeviceId=%s\nStateName=%s\nPhoneNo=%s\n",firstName,lastName,email,dateofIssue,city,DeviceName,DeviceID,stateName,phoneNo)))
}

//AllortHardwareJSON  ...
func (m *Repository) AllortHardwareJSON(w http.ResponseWriter, r *http.Request) {

	resp := Employees{

	FirstName: r.Form.Get("firstName"),
	LastName:r.Form.Get("lastName"),
	Email : r.Form.Get("email"),
	DateofIssue : r.Form.Get("date_of_issue"),
	City : r.Form.Get("city"),
	DeviceName : r.Form.Get("device_name"),
	DeviceID: r.Form.Get("device_id"),
	StateName : r.Form.Get("state_name"),
	PhoneNo : r.Form.Get("phone_no"),
	}

	out ,err := json.MarshalIndent(resp,""," ")
	if err != nil {

		log.Println(err)
	}

	log.Println(out)
	w.Header().Set("Content-Type","application/json")
	w.Write(out)
}



//Contact ...
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w,r,"contact.page.html", &models.TemplateData{})
}
