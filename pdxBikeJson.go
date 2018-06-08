package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	_ "github.com/lib/pq"
	"database/sql"
	//"os"
)

// citiBikeURL provides the station statuses of CitiBike bike sharing stations.
const pdxBikeURL = "http://biketownpdx.socialbicycles.com/opendata/station_status.json"

// stationData is used to unmarshal the JSON document returned form citiBikeURL.
type stationData struct {
	LastUpdated int `json:"last_updated"`
	TTL			int `json:"ttl"`
	Data        struct {
		Stations []station `json:"stations"`
	} `json:"data"`
}

// station is used to unmarshal each of the station documents in stationData.
type station struct {
	StationID		string   `json:"station_id"`
	NumBikesAvailable int    `json:"num_bikes_available"`
	NumBikesDisabled  int    `json:"num_bikes_disabled"`
	NumDocksAvailable int    `json:"num_docks_available"`
	IsInstalled       int    `json:"is_installed"`
	IsRenting         int    `json:"is_renting"`
	IsReturning       int    `json:"is_returning"`
	LastReported      int    `json:"last_reported"`	
}

func main() {

	// Get the JSON response from the URL.
	response, err := http.Get(pdxBikeURL)
	if err != nil {
		log.Fatal(err)
	}

	// Defer closing the response body.
	defer response.Body.Close()

	// Read the body of the response into []byte.
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Declare a variable of type stationData.
	var sd stationData

	// Unmarshal the JSON data into the variable.
	if err := json.Unmarshal(body, &sd); err != nil {
		log.Fatal(err)
		return
	}

	numStations := len(sd.Data.Stations)
	SID := make([]string, numStations)
	NBA := make([]int, numStations)
	NBD := make([]int, numStations)
	NDA := make([]int, numStations)
	II := make([]int, numStations)
	IRen := make([]int, numStations)
	IRet := make([]int, numStations)
	LR := make([]int, numStations)


	for i:=0; i<numStations; i++ {
		SID[i] = sd.Data.Stations[i].StationID
		NBA[i] = sd.Data.Stations[i].NumBikesAvailable
		NBD[i] = sd.Data.Stations[i].NumBikesDisabled
		NDA[i] = sd.Data.Stations[i].NumDocksAvailable
		II[i] = sd.Data.Stations[i].IsInstalled
		IRen[i] = sd.Data.Stations[i].IsRenting
		IRet[i] = sd.Data.Stations[i].IsReturning
		LR[i] = sd.Data.Stations[i].LastReported
	}
	

	// Print the first station.
	//fmt.Printf("%+v\n\n", sd.Data.Stations[1].StationID)

	pgURL := "postgresql://pdxbike:stigz@localhost/station_data?sslmode=disable"

	db, err := sql.Open("postgres", pgURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// sql.Open() does not establish any connections to the
	// database.  It just prepares the database connection value
	// for later use.  To make sure the database is available and
	// accessible, we will use db.Ping().
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	/*
	CREATE TABLE timestamp (
  id SERIAL PRIMARY KEY,
  station_id TEXT,
  num_bikes_available INT,
  num_bikes_disabled INT,
  num_docks_available INT,
  is_installed INT,
  is_renting INT,
  is_returning INT,
  last_reported INT
);
	*/
	for i:=0; i<numStations; i++{
	sqlStatement := `
		INSERT INTO timestamp (station_id, num_bikes_available, num_bikes_disabled, num_docks_available, is_installed, is_renting, is_returning, last_reported)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id`
  	id := 0
  	err = db.QueryRow(sqlStatement, SID[i], NBA[i], NBD[i], NDA[i], II[i], IRen[i], IRet[i], LR[i]).Scan(&id)
  	if err != nil {
    	panic(err)
  	}
  	fmt.Println("New record ID is:", id)
	}
}
