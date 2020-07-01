package chrbm

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/itchyny/gojq"
)

type Bookmark struct {
	Name      string
	URL       string
	DateAdded time.Time
}

type bookmarkItem struct {
	name      string
	itemType  string
	dateAdded time.Time
}

type bookmarkMap map[string]interface{}

func NewBookmark(j []byte) ([]Bookmark, error) {
	if !json.Valid(j) {
		return nil, errors.New("input is invalid format (required json only)")
	}

	var bm bookmarkMap
	err := json.Unmarshal(j, &bm)
	if err != nil {
		return nil, err
	}

	query, err := gojq.Parse(".checksum")
	if err != nil {
		return nil, err
	}

	iter := query.Run(bm)
	for {
		v, ok := iter.Next()
		if !ok {
			break
		}
		if err, ok := v.(error); ok {
			log.Fatalln(err)
		}
		fmt.Printf("%#v\n", v)
	}

	return []Bookmark{}, nil
}
