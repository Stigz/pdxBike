package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {
	in := `text, I  aid=2935 work=a830d2c0e398d83359a2cfd237a849e5 facebook_audience_populate  added 353 users removed 221 users,
I  aid=2935 work=a830d2c0e398d83359a2cfd237a849e5 facebook_audience_populate  added 279 users removed 174 users,
I  aid=2935 work=a830d2c0e398d83359a2cfd237a849e5 facebook_audience_populate  added 169 users removed 112 users, 
I  aid=2935 work=a830d2c0e398d83359a2cfd237a849e5 facebook_audience_populate  added 101 users removed 67 users, 
I  aid=2935 work=a830d2c0e398d83359a2cfd237a849e5 facebook_audience_populate  added 63 users removed 41 users, 
I  aid=2935 work=a830d2c0e398d83359a2cfd237a849e5 facebook_audience_populate  added 34 users removed 20 users, 
I  aid=2935 work=a830d2c0e398d83359a2cfd237a849e5 facebook_audience_populate  added 42 users removed 16 users, 
I  aid=2935 work=a830d2c0e398d83359a2cfd237a849e5 facebook_audience_populate  added 33 users removed 13 users, 
I  aid=2935 work=a830d2c0e398d83359a2cfd237a849e5 facebook_audience_populate  added 87 users removed 52 users, 
I  aid=2935 work=a830d2c0e398d83359a2cfd237a849e5 facebook_audience_populate  added 40 users removed 950 users, 
I  aid=2935 work=a830d2c0e398d83359a2cfd237a849e5 facebook_audience_populate  added 111 users removed 149 users, 
I  aid=2935 work=a830d2c0e398d83359a2cfd237a849e5 facebook_audience_populate  added 692 users removed 300 users, 
I  aid=2935 work=a830d2c0e398d83359a2cfd237a849e5 facebook_audience_populate  added 240 users removed 106 users, 
I  aid=2935 work=a830d2c0e398d83359a2cfd237a849e5 facebook_audience_populate  added 658 users removed 332 users, 
I  aid=2935 work=a830d2c0e398d83359a2cfd237a849e5 facebook_audience_populate  added 520 users removed 280 users, 
I  aid=2935 work=a830d2c0e398d83359a2cfd237a849e5 facebook_audience_populate  added 635 users removed 361 users, 
I  aid=2935 work=a830d2c0e398d83359a2cfd237a849e5 facebook_audience_populate  added 416 users removed 246 users, 
I  aid=2935 work=a830d2c0e398d83359a2cfd237a849e5 facebook_audience_populate  added 636 users removed 361 users, 
I  aid=2935 work=a830d2c0e398d83359a2cfd237a849e5 facebook_audience_populate  added 363 users removed 223 users, 
I  aid=2935 work=a830d2c0e398d83359a2cfd237a849e5 facebook_audience_populate  added 23 users removed 206 users, 
I  aid=2935 work=a830d2c0e398d83359a2cfd237a849e5 facebook_audience_populate  added 11 users removed 140 users, 
I  aid=2935 work=a830d2c0e398d83359a2cfd237a849e5 facebook_audience_populate  added 5 users removed 121 users, 
I  aid=2935 work=a830d2c0e398d83359a2cfd237a849e5 facebook_audience_populate  added 7 users removed 81 users, 
I  aid=2935 work=a830d2c0e398d83359a2cfd237a849e5 facebook_audience_populate  added 65 user removed 933 user,
I  aid=2935 work=a830d2c0e398d83359a2cfd237a849e5 facebook_audience_populate  added 64 user removed 936 user,
I  aid=2935 work=a830d2c0e398d83359a2cfd237a849e5 facebook_audience_populate  added 68 users removed 932 use,
I  aid=2935 work=a830d2c0e398d83359a2cfd237a849e5 facebook_audience_populate  added 52 users removed 861 use,
I  aid=2935 work=a830d2c0e398d83359a2cfd237a849e5 facebook_audience_populate  added 19 users removed 67 user,
I  aid=2935 work=a830d2c0e398d83359a2cfd237a849e5 facebook_audience_populate  added 3 users removed 53 users, 
I  aid=2935 work=a830d2c0e398d83359a2cfd237a849e5 facebook_audience_populate  added 3 users removed 50 users ,
I  aid=2935 work=a830d2c0e398d83359a2cfd237a849e5 facebook_audience_populate  added 10 user removed 52 users ,
I  aid=2935 work=a830d2c0e398d83359a2cfd237a849e5 facebook_audience_populate  added 5 users removed 46 users ,
I  aid=2935 work=a830d2c0e398d83359a2cfd237a849e5 facebook_audience_populate  added 3 users removed 38 users ,
I  aid=2935 work=a830d2c0e398d83359a2cfd237a849e5 facebook_audience_populate  added 5 users removed 38 users ,
I  aid=2935 work=a830d2c0e398d83359a2cfd237a849e5 facebook_audience_populate  added 2 users removed 21 users, 
I  aid=2935 work=a830d2c0e398d83359a2cfd237a849e5 facebook_audience_populate  added 1 users removed 11 users, 
I  aid=2935 work=a830d2c0e398d83359a2cfd237a849e5 facebook_audience_populate  added 1 users removed 6 users, 
I  aid=2935 work=a830d2c0e398d83359a2cfd237a849e5 facebook_audience_populate  added 2 users removed 3 users, 
I  aid=2935 work=a830d2c0e398d83359a2cfd237a849e5 facebook_audience_populate  added 1 users removed 0 users, 
I  aid=2935 work=a830d2c0e398d83359a2cfd237a849e5 facebook_audience_populate  added 1 users removed 2 users ,
I  aid=2935 work=a830d2c0e398d83359a2cfd237a849e5 facebook_audience_populate  added 0 users removed 8 users ,
I  aid=2935 work=a830d2c0e398d83359a2cfd237a849e5 facebook_audience_populate  added 430 uses removed 391 use,
I  aid=2935 work=a830d2c0e398d83359a2cfd237a849e5 facebook_audience_populate  added 4 users removed 33 users, 
I  aid=2935 work=a830d2c0e398d83359a2cfd237a849e5 facebook_audience_populate  added 3 users removed 46 users, 
I  aid=2935 work=a830d2c0e398d83359a2cfd237a849e5 facebook_audience_populate  added 3 users removed 44 users, 
I  aid=2935 work=a830d2c0e398d83359a2cfd237a849e5 facebook_audience_populate  added 7 users removed 36 users, 
I  aid=2935 work=a830d2c0e398d83359a2cfd237a849e5 facebook_audience_populate  added 2 users removed 19 users, 
I  aid=2935 work=a830d2c0e398d83359a2cfd237a849e5 facebook_audience_populate  added 10 user removed 27 users, 
I  aid=2935 work=a830d2c0e398d83359a2cfd237a849e5 facebook_audience_populate  added 6 users removed 11 users, 
I  aid=2935 work=a830d2c0e398d83359a2cfd237a849e5 facebook_audience_populate  added 6 users removed 15 users, 
I  aid=2935 work=a830d2c0e398d83359a2cfd237a849e5 facebook_audience_populate  added 9 user removed 42 users ,
I  aid=2935 work=a830d2c0e398d83359a2cfd237a849e5 facebook_audience_populate  added 6 users removed 22 users, 
I  aid=2935 work=a830d2c0e398d83359a2cfd237a849e5 facebook_audience_populate  added 0 users removed 0 users ,`

	r := csv.NewReader(strings.NewReader(in))

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(record)
		for _, rec := range record {
			splitRecord := []string{}
			splitRecord = strings.Split(rec, "")
			for _, split := range splitRecord {
				isInt, ok := split.(int)
			}
		}
	}

}
