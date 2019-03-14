package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"cloud.google.com/go/bigquery"
	"cloud.google.com/go/storage"
	//"google.golang.org/appengine"
	//_ "github.com/lib/pq"
	//"os"
)

// citiBikeURL provides the station statuses of CitiBike bike sharing stations.
const pdxBikeURL = "http://biketownpdx.socialbicycles.com/opendata/station_status.json"

// stationData is used to unmarshal the JSON document returned form citiBikeURL.
type StationData struct {
	LastUpdated int `json:"last_updated"`
	TTL         int `json:"ttl"`
	Data        struct {
		Stations []Station `json:"stations"`
	} `json:"data"`
}

// station is used to unmarshal each of the station documents in stationData.
type Station struct {
	StationID         string `json:"station_id"`
	NumBikesAvailable int    `json:"num_bikes_available"`
	NumBikesDisabled  int    `json:"num_bikes_disabled"`
	NumDocksAvailable int    `json:"num_docks_available"`
	IsInstalled       int    `json:"is_installed"`
	IsRenting         int    `json:"is_renting"`
	IsReturning       int    `json:"is_returning"`
	LastReported      int    `json:"last_reported"`
}

// demo struct holds information needed to run the various demo functions.
type demo struct {
	client     *storage.Client
	bucketName string
	bucket     *storage.BucketHandle

	w   io.Writer
	ctx context.Context
	// cleanUp is a list of filenames that need cleaning up at the end of the demo.
	cleanUp []string
	// failed indicates that one or more of the demo steps failed.
	failed bool
}

func main() {
	//appengine.Main()

	response, err := http.Get(pdxBikeURL)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var sd StationData

	if err := json.Unmarshal(body, &sd); err != nil {
		log.Fatal(err)
		return
	}
	//fmt.Println(sd)
	//req := http.Request{}

	// Create a new App Engine context from the request.
	//	ctx := appengine.NewContext(&req)
	ctx := context.Background()

	err = insertIntoTable(ctx, sd)
	if err != nil {
		fmt.Errorf("Error inserting, %v", err)
		return
	}

	//http.HandleFunc("/", handle)

	//stations := sd.Data.Stations
	//numStations := len(sd.Data.Stations)
}

// datasets returns a list with the IDs of all the Big Query datasets visible
// with the given context.
func insertIntoTable(ctx context.Context, sd StationData) error {
	// Get the current application ID, which is the same as the project ID.
	//projectID := appengine.AppID(ctx)
	//fmt.Println(projectID)

	// Create the BigQuery service.
	bq, err := bigquery.NewClient(ctx, "pdxbike")
	if err != nil {
		return fmt.Errorf("could not create service: %v", err)
	}

	items := []*Station{}
	for _, st := range sd.Data.Stations {
		fmt.Println(st)
		items = append(items, &st)
	}
	u := bq.Dataset("StationData").Table("testing1").Uploader()
	if err := u.Put(ctx, items); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func handle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	//	w.Header().Set("Content-Type", "text/plain")

	// if len(names) == 0 {
	// 	fmt.Fprintf(w, "No datasets visible")
	// } else {
	// 	fmt.Fprintf(w, "Datasets:\n\t"+strings.Join(names, "\n\t"))
	// }
}

// SID := make([]string, numStations)
// NBA := make([]int, numStations)
// NBD := make([]int, numStations)
// NDA := make([]int, numStations)
// II := make([]int, numStations)
// IRen := make([]int, numStations)
// IRet := make([]int, numStations)
// LR := make([]int, numStations)

// for i:=0; i<numStations; i++ {
// 	SID[i] = sd.Data.Stations[i].StationID
// 	NBA[i] = sd.Data.Stations[i].NumBikesAvailable
// 	NBD[i] = sd.Data.Stations[i].NumBikesDisabled
// 	NDA[i] = sd.Data.Stations[i].NumDocksAvailable
// 	II[i] = sd.Data.Stations[i].IsInstalled
// 	IRen[i] = sd.Data.Stations[i].IsRenting
// 	IRet[i] = sd.Data.Stations[i].IsReturning
// 	LR[i] = sd.Data.Stations[i].LastReported
// }

// pgURL := "postgresql://pdxbike:stigz@localhost/station_data?sslmode=disable"

// db, err := sql.Open("postgres", pgURL)
// if err != nil {
// 	log.Fatal(err)
// }
// defer db.Close()

// sql.Open() does not establish any connections to the
// database.  It just prepares the database connection value
// for later use.  To make sure the database is available and
// accessible, we will use db.Ping().
// if err := db.Ping(); err != nil {
// 	log.Fatal(err)
// }

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

// sqlStatement := `
// 	INSERT INTO timestamp (station_id, num_bikes_available, num_bikes_disabled, num_docks_available, is_installed, is_renting, is_returning, last_reported)
// 	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
// 	RETURNING id`
// id := 0
// err = db.QueryRow(sqlStatement, SID[i], NBA[i], NBD[i], NDA[i], II[i], IRen[i], IRet[i], LR[i]).Scan(&id)
// if err != nil {
// 	panic(err)
// }
// fmt.Println("New record ID is:", id)
// }
