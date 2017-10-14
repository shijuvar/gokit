package store

import (
	"fmt"
	"log"

	"github.com/pkg/errors"

	"github.com/shijuvar/gokit/examples/nats-streaming/pb"
)

type EventStore struct{}

func (store EventStore) CreateEvent(event *pb.Event) error {
	// Insert two rows into the "accounts" table.
	sql := fmt.Sprintf("INSERT INTO events (id, eventtype, aggregateid, aggregatetype, eventdata, channel) VALUES ('%s','%s','%s','%s','%s','%s')", event.EventId, event.EventType, event.AggregateId, event.AggregateType, event.EventData, event.Channel)
	log.Print(sql)
	_, err := db.Exec(sql)
	if err != nil {
		return errors.Wrap(err, "Error on insert into events")
	}
	return nil
}

func (store EventStore) GetEvents(filter *pb.EventFilter) []*pb.Event {
	var events []*pb.Event
	return events
}
