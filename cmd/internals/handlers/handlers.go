package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/sanjay-xdr/cmd/internals/config"
	"github.com/sanjay-xdr/cmd/internals/forms"
	"github.com/sanjay-xdr/cmd/internals/models"
	"github.com/sanjay-xdr/cmd/internals/render"
)

var Repo *Repositry

type Repositry struct {
	App *config.AppConfig
}

// Creates a new Repo
func NewRepo(a *config.AppConfig) *Repositry {
	return &Repositry{
		App: a,
	}
}

// set the Above Repo Variable
func NewHandlers(r *Repositry) {
	Repo = r
}

func (m *Repositry) Home(w http.ResponseWriter, r *http.Request) {

	// fmt.Print("Hello World")
	// render.CreateTemplateCache()

	remoteIp := r.RemoteAddr

	log.Print(remoteIp, "This is the IP")

	m.App.Session.Put(r.Context(), "remote_ip", remoteIp)
	stringMap := make(map[string]string)
	stringMap["test"] = "This is the test data from template"
	render.RenderTemplate(w, r, "home.page.html", &models.TemplateData{
		StringMap: stringMap,
	})

}

func (m *Repositry) About(w http.ResponseWriter, r *http.Request) {
	value := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap := make(map[string]string)
	stringMap["ip"] = value
	render.RenderTemplate(w, r, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}

func (m *Repositry) Reservation(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Rendering the Reservation page")

	var emptyReservation models.Reservation

	data := make(map[string]interface{})
	data["reservation"] = emptyReservation

	render.RenderTemplate(w, r, "make-reservation.page.html", &models.TemplateData{

		Form: forms.New(nil),
		Data: data,
	})
}

func (m *Repositry) PostReservation(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		fmt.Print("Not able to parse err in handlers")
	}

	reservation := models.Reservation{
		FirstName: r.Form.Get("first_name"),
		LastName:  r.Form.Get("last_name"),
		Email:     r.Form.Get("email"),
		Phone:     r.Form.Get("phone"),
	}

	fmt.Println(r.PostForm)

	form := forms.New(r.PostForm)

	// form.Has("first_name", r)

	form.Required("first_name", "last_name", "email")
	form.IsEmail("email")

	if !form.Valid() {
		data := make(map[string]interface{})

		data["reservation"] = reservation

		render.RenderTemplate(w, r, "make-reservation.page.html", &models.TemplateData{

			Form: form,
			Data: data,
		})
		return
	}

	m.App.Session.Put(r.Context(), "reservation", reservation)

	http.Redirect(w, r, "/reservation-summary", http.StatusSeeOther)

	fmt.Println("Post Reservation is getting hit")

}

func (m *Repositry) Generals(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "generals.page.html", &models.TemplateData{})
}

// Majors renders the room page
func (m *Repositry) Majors(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "majors.page.html", &models.TemplateData{})
}

// Availability renders the search availability page
func (m *Repositry) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "search-availability.page.html", &models.TemplateData{})
}

type JSONStruct struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}

func (m *Repositry) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {

	res := JSONStruct{
		Ok:      true,
		Message: "Dummy data",
	}

	out, err := json.MarshalIndent(res, "", " ")
	if err != nil {
		fmt.Print("SOmething went wrong", err)
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(out)

}

func (m *Repositry) PostAvailability(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Print("Not able to parse the form something went wrong")
	}

	// fmt.Print(r.Form, "THis is the form data")
	// fmt.Print("Getting a hit here")
	start := r.Form.Get("start")
	end := r.Form.Get("end")

	// w.Write([]byte("Posted SUccesfully"))

	fmt.Print(start)
	fmt.Print(end)

	w.Write([]byte(fmt.Sprintf("Start date is %s  and end date is %s", start, end)))
}

// Contact renders the contact page
func (m *Repositry) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "contact.page.html", &models.TemplateData{})
}

func (m *Repositry) ReservationSummary(w http.ResponseWriter, r *http.Request) {

	//Doing stuff i dont know

	reservation, ok := m.App.Session.Get(r.Context(), "reservation").(models.Reservation)

	if !ok {
		log.Print("Can not convert the session ")

		m.App.Session.Put(r.Context(), "error", "Cant get reservation from session")

		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	m.App.Session.Remove(r.Context(), "reservation")

	data := make(map[string]interface{})
	data["reservation"] = reservation

	render.RenderTemplate(w, r, "reservation-summary.page.html", &models.TemplateData{Data: data})

	log.Println(data)
}
