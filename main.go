package main

import (
	routes "Doctor-Appointment-Project/routes"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")

	if port == "" {
		port = "8081"
	}
	Connect()
	router := gin.New()
	router.Use(gin.Logger())

	routes.AuthRoutes(router)
	routes.UserRoutes(router)
	routes.DoctorRoutes(router)
	routes.PatientRoutes(router)

	router.Run(":" + port)
}

func Connect() {

	// Make Database Connection
	db, err := sql.Open("mysql", "root:india@123@tcp(localhost:3306)/das")
	if err != nil {

		log.Fatal(err)

	}

	fmt.Println("Connected to MySQL database!")

	// Create Patient table

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS patient(ID INT NOT NULL AUTO_INCREMENT, Name VARCHAR(30),Age INT,Gender VARCHAR(10),Address VARCHAR(50), City VARCHAR(20),Phone VARCHAR(15),Disease VARCHAR(25),Selected_Specialisation VARCHAR(20),Patient_history VARCHAR(250), PRIMARY KEY (ID) );")

	if err != nil {

		panic(err.Error())

	}
	fmt.Println("Patient Table Created")

	// Create Docter table

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS doctor(ID INT NOT NULL AUTO_INCREMENT, Name VARCHAR(30),Gender VARCHAR(10),Address VARCHAR(50), City VARCHAR(20),Phone VARCHAR(15),Specialisation VARCHAR(20),Opening_time VARCHAR(10),Closing_time VARCHAR(10),Availabilty_Time  VARCHAR(10),Availabilty VARCHAR(30),Fees INT ,PRIMARY KEY (ID) );")

	if err != nil {

		panic(err.Error())

	}
	fmt.Println("Docter Table Created")

	// Create Appoitment table

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS Appoitment(Bookingid INT NOT NULL AUTO_INCREMENT,Patient_id INT,Doctor_id INT,Booking_time VARCHAR(10),PRIMARY KEY (Bookingid),FOREIGN KEY (Patient_id) REFERENCES Patient(ID),FOREIGN KEY (Doctor_id) REFERENCES Doctor(ID));")

	if err != nil {

		panic(err.Error())

	}

}
