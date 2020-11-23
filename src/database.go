package src

import (
	"github.com/nanobox-io/golang-scribble"
)

var db *scribble.Driver

func init() {
	var err error
	db, err = scribble.New("data", nil)
	if err != nil {
		panic(err)
	}
}