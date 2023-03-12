package controllers

import (
	"database/sql"
	"time"

	"fmt"

	"log"

	"net/http"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"

	"Doctor-Appointment-Project/models"
)

func add_time(s string) string {
	timeStr := s
	t, err := time.Parse("15:04", timeStr)
	if err != nil {
		panic(err)
	}

	t = t.Add(30 * time.Minute)
	newTimeStr := t.Format("15:04")

	return newTimeStr
}

func Addpatient() gin.HandlerFunc {
	return func(c *gin.Context) {

		fmt.Println("add patient")

		db, err := sql.Open("mysql", "root:india@123@tcp(localhost:3306)/das")
		if err != nil {

			log.Fatal(err)

		}

		var data models.Patient

		err = c.BindJSON(&data)

		if err != nil {

			return

		}

		c.IndentedJSON(http.StatusCreated, data)

		query_data := fmt.Sprintf(`INSERT INTO patient(Name,Age,Gender,Address,City,Phone,Disease,Selected_Specialisation,Patient_history) VALUES('%s',%d,'%s','%s','%s','%s','%s','%s','%s')`, data.Name, data.Age, data.Gender, data.Address, data.City, data.Phone, data.Disease, data.Selected_specialisation, data.Patient_history)

		fmt.Println(query_data)

		//insert data

		insert, err := db.Query(query_data)

		if err != nil {

			panic(err.Error())

		}

		defer insert.Close()

		c.JSON(http.StatusOK, gin.H{"message": "Patient added successfully"})

	}
}

func Getpatient() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := sql.Open("mysql", "root:india@123@tcp(localhost:3306)/das")
		if err != nil {

			log.Fatal(err)

		}
		fmt.Println("Connection Created")
		results, err := db.Query("SELECT * FROM Patient")
		fmt.Println("Quary exicuted")

		if err != nil {

			panic(err.Error())

		}

		defer results.Close()

		var output interface{}

		for results.Next() {

			var ID int
			var Name string
			var Age int
			var Gender string
			var Address string
			var City string
			var Phone string
			var Disease string
			var Selected_specialisation string
			var Patient_history string
			err = results.Scan(&ID, &Name, &Age, &Gender, &Address, &City, &Phone, &Disease, &Selected_specialisation, &Patient_history)

			if err != nil {

				panic(err.Error())

			}

			output = fmt.Sprintf("%d  '%s'  %d  '%s'  '%s'  '%s'  '%s' '%s' %s' '%s'  ", ID, Name, Age, Gender, Address, City, Phone, Disease, Selected_specialisation, Patient_history)

			fmt.Println(output)

			c.JSON(http.StatusOK, gin.H{"Data": output})

		}

	}
}

func DeletePatient() gin.HandlerFunc {
	return func(c *gin.Context) {

		db, err := sql.Open("mysql", "root:india@123@tcp(localhost:3306)/das")
		if err != nil {

			log.Fatal(err)

		}

		var data models.Patient

		err = c.BindJSON(&data)

		if err != nil {

			return

		}

		// _, err = db.Exec("DELETE FROM Dost WHERE id = 10")

		delete_query := fmt.Sprintf("DELETE FROM patient WHERE ID= %d", data.ID)

		delete, err := db.Query(delete_query)

		if err != nil {

			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

			return

		}

		defer delete.Close()

		c.JSON(http.StatusOK, gin.H{"message": "Patient Deleted successfully"})

	}
}

func BookingAppointment() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := sql.Open("mysql", "root:india@123@tcp(localhost:3306)/das")
		if err != nil {

			log.Fatal(err)

		}
		// var Doctor_str doctor
		var booking_data models.Appointment
		err = c.BindJSON(&booking_data)
		get_booking_time := fmt.Sprintf("SELECT Opening_time,Closing_time, FROM Doctor WHERE id = %d", booking_data.Doctor_id)
		doctor_result, err := db.Query(get_booking_time)
		// doctor_result,err := db.Exec(get_booking_time)
		if err != nil {
			c.JSON(404, gin.H{"error": "Doctor not found"})
			return
		}

		var Opening_time string

		var Closing_time string
		err = doctor_result.Scan(&Opening_time, &Closing_time)

		if err != nil {

			panic(err.Error())

		}

		c.IndentedJSON(http.StatusCreated, booking_data)

		booking_data.Booking_time = Opening_time

		query_data := fmt.Sprintf(`INSERT INTO Appointment VALUES(%d,%d,'%s')`, booking_data.Patient_id, booking_data.Doctor_id, booking_data.Booking_time)
		_, err = db.Exec(query_data)
		if err != nil {

			panic(err.Error())

		}
		t1 := add_time(Opening_time)
		t2 := add_time(Closing_time)
		query_data2 := fmt.Sprintf(`UPDATE Doctor SET Opening_time = %s,Closing_time = %s WHERE ID = %d`, t1, t2, booking_data.Doctor_id)
		_, err = db.Query(query_data2)
		if err != nil {

			panic(err.Error())

		}

		if err != nil {

			panic(err.Error())

		}

		c.JSON(http.StatusOK, gin.H{"message": "Your Appointment successfully Booked"})

	}
}

func Cancel_appointment() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := sql.Open("mysql", "root:india@123@tcp(localhost:3306)/das")
		if err != nil {

			log.Fatal(err)

		}

		var data models.Appointment

		err = c.BindJSON(&data)

		if err != nil {

			return

		}

		c.IndentedJSON(http.StatusCreated, data)

		query_data := fmt.Sprintf("DELETE FROM Doctor WHERE id =%d", data.Bookingid)

		_, err = db.Exec(query_data)

		if err != nil {

			panic(err.Error())

		}

		c.JSON(http.StatusOK, gin.H{"message": "Cancel Appointment successfully"})

	}
}
