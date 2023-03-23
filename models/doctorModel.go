package models

type Doctor struct {
	ID               int
	Name             string
	Gender           string
	Address          string
	City             string
	Phone            string
	Specialisation   string
	Opening_time     string
	Closing_time     string
	Availabilty      string
	Availabilty_Time string
	Fees             float64
}
