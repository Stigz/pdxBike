package main

import (
  "database/sql"
  "fmt"
  "log"
  //"os"

  // pq is the libary that allows us to connect
  // to postgres with databases/sql.
  _ "github.com/lib/pq"
)


type station struct {
    StationID       string   
    NumBikesAvailable int    
    NumBikesDisabled  int    
    NumDocksAvailable int    
    IsInstalled       int    
    IsRenting         int    
    IsReturning       int    
    LastReported      int    
}

func main() {
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

	rows, err := db.Query("SELECT station_id, num_bikes_available FROM timestamp LIMIT $1", 2)
  	if err != nil {
    // handle this error better than this
    panic(err)
  	}
  	defer rows.Close()
  	for rows.Next() {
        var stationID string
    	var NBA int
    	err = rows.Scan(&stationID, &NBA)
    	if err != nil {
      		// handle this error
      		panic(err)
    	}
    	fmt.Println(stationID, NBA)
  	}
  	// get any error encountered during iteration
  	err = rows.Err()
  	if err != nil {
    	panic(err)
  	}
}