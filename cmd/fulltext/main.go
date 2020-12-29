package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"github.com/whosonfirst/go-whosonfirst-search/fulltext"
	"flag"
)

func main() {

	db_uri := flag.String("fulltext-database-uri", "null://", "...")
	
	flag.Parse()

	ctx := context.Background()
	
	db, err := fulltext.NewFullTextDatabase(ctx, *db_uri)

	if err != nil {
		log.Fatal(err)
	}

	for _, q := range flag.Args() {

		r, err := db.QueryString(ctx, q)

		if err != nil {
			log.Fatal(err)
		}

		enc_r, err := json.Marshal(r)

		if err != nil {
			log.Fatal(err)
		}
		
		fmt.Println(string(enc_r))
	}
}

