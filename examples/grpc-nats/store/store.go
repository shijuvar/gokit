package store

import (
	"log"

	"gopkg.in/mgo.v2"
)
var mgoSession *mgo.Session

func init() {
	var err error
	mgoSession, err = mgo.DialWithInfo(&mgo.DialInfo{
		Addrs: []string{"127.0.0.1"},
	})
	if err != nil {
		log.Fatalf("[createDbSession]: %s\n", err)
	}
}