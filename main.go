package main

//go mod tidy

import (
	"database/sql"
	"encoding/json"
	_ "encoding/json"
	_ "fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	_ "io/ioutil"
	"net/http"
)

var db *sql.DB
var err error

type MotorBikeModel struct {
	Id                 string `json:"id,omitempty" bson:",omitempty"`
	CompanyName        string `json:"company_name,omitempty" bson:",omitempty"`
	Model              string `json:"model,omitempty" bson:",omitempty"`
	Price              string `json:"price,omitempty" bson:",omitempty"`
	Status             string `json:"status,omitempty" bson:",omitempty"`
	BodyType           string `json:"body_type,omitempty" bson:",omitempty"`
	FuelType           string `json:"fuel_type,omitempty" bson:",omitempty"`
	EngineDescription  string `json:"engine_description,omitempty" bson:",omitempty"`
	FuelSystem         string `json:"fuel_system,omitempty" bson:",omitempty"`
	Cooling            string `json:"cooling,omitempty" bson:",omitempty"`
	Displacement       string `json:"displacement,omitempty" bson:",omitempty"`
	MaximumPower       string `json:"maximum_power,omitempty" bson:",omitempty"`
	MaximumTorque      string `json:"maximum_torque,omitempty" bson:",omitempty"`
	NumberOfCylinders  string `json:"number_of_cylinders,omitempty" bson:",omitempty"`
	OverallLength      string `json:"overall_length,omitempty" bson:",omitempty"`
	OverallWidth       string `json:"overall_width,omitempty" bson:",omitempty"`
	OverallHeight      string `json:"overall_height,omitempty" bson:",omitempty"`
	SeatHeight         string `json:"seat_height,omitempty" bson:",omitempty"`
	GroundClearance    string `json:"ground_clearance,omitempty" bson:",omitempty"`
	Wheelbase          string `json:"wheelbase,omitempty" bson:",omitempty"`
	KerbOrWetWeight    string `json:"kerb_or_wet_weight,omitempty" bson:",omitempty"`
	FuelTankCapacity   string `json:"fuel_tank_capacity,omitempty" bson:",omitempty"`
	Bore               string `json:"bore,omitempty" bson:",omitempty"`
	Stroke             string `json:"stroke,omitempty" bson:",omitempty"`
	NumberOfGears      string `json:"number_of_gears,omitempty" bson:",omitempty"`
	Clutch             string `json:"clutch,omitempty" bson:",omitempty"`
	GearboxType        string `json:"gearbox_type,omitempty" bson:",omitempty"`
	FrontBrake         string `json:"front_brake,omitempty" bson:",omitempty"`
	RearBrake          string `json:"rear_brake,omitempty" bson:",omitempty"`
	FrontSuspension    string `json:"front_suspension,omitempty" bson:",omitempty"`
	RearSuspension     string `json:"rear_suspension,omitempty" bson:",omitempty"`
	ZeroToHundred_kmph string `json:"0_to_100_kmph,omitempty" bson:",omitempty"`
	Speedometer        string `json:"speedometer,omitempty" bson:",omitempty"`
	Tachometer         string `json:"tachometer,omitempty" bson:",omitempty"`
	TripMeter          string `json:"trip_meter,omitempty" bson:",omitempty"`
	Clock              string `json:"clock,omitempty" bson:",omitempty"`
	ElectricStart      string `json:"electric_start,omitempty" bson:",omitempty"`
}

func main() {
	db, err = sql.Open("mysql", "rootmac:wacademie@tcp(127.0.0.1:3306)/bikePlus")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	router := mux.NewRouter()
	router.HandleFunc("/company", getAllCompany).Methods("GET")
	router.HandleFunc("/company/{id}", getCompanyById).Methods("GET")

	router.HandleFunc("/model/{companyName}", getModelsByCompany).Methods("GET")

	http.ListenAndServe(":8000", router)

}

func getCompanyById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	result, err := db.Query("SELECT company_name from motorbike_data WHERE id = ?", params["id"])
	if err != nil {
		panic(err.Error())
	}

	defer result.Close()
	var model MotorBikeModel

	for result.Next() {
		err := result.Scan(&model.CompanyName)
		if err != nil {
			panic(err.Error())
		}
	}

	json.NewEncoder(w).Encode(model)
}

func getAllCompany(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var models []MotorBikeModel
	result, err := db.Query("SELECT DISTINCT company_name from motorbike_data")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
		var model MotorBikeModel
		err := result.Scan(&model.Model)
		if err != nil {
			panic(err.Error())
		}
		models = append(models, model)
	}
	json.NewEncoder(w).Encode(models)
}

func getModelsByCompany(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	//fmt.Println("coucouc")
	var models []MotorBikeModel
	result, err := db.Query("SELECT model from motorbike_data WHERE company_name = ?", params["companyName"])
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
		var model MotorBikeModel
		err := result.Scan(&model.Model)
		if err != nil {
			panic(err.Error())
		}
		models = append(models, model)
	}
	json.NewEncoder(w).Encode(models)
}
