package main

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"cloud.google.com/go/bigquery"
	"google.golang.org/appengine"
)

func init() {
	http.HandleFunc("/", handle)
}

func handle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// Create a new App Engine context from the request.
	ctx := appengine.NewContext(r)

	// Get the list of dataset names.
	names, err := insertIntoTable(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")

	if len(names) == 0 {
		fmt.Fprintf(w, "No datasets visible")
	} else {
		fmt.Fprintf(w, "Datasets:\n\t"+strings.Join(names, "\n\t"))
	}
}

// datasets returns a list with the IDs of all the Big Query datasets visible
// with the given context.
func insertIntoTable(ctx context.Context) ([]string, error) {
	// Get the current application ID, which is the same as the project ID.
	projectID := appengine.AppID(ctx)

	// Create the BigQuery service.
	bq, err := bigquery.NewClient(ctx, projectID)
	if err != nil {
		return nil, fmt.Errorf("could not create service: %v", err)
	}

	// Return a list of IDs.

	var ids []string
	newTable := bq.Dataset("StationData").Table("Stations").Uploader()
	items := []*Item{
		// Item implements the ValueSaver interface.
		{Name: "Phred Phlyntstone", Age: 32},
		{Name: "Wylma Phlyntstone", Age: 29},
	}
	if err := u.Put(ctx, items); err != nil {
		return nil, err
	}
}
