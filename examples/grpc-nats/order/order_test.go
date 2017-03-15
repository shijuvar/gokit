// Benchmark testing to measure the performance of marshaling and unmarshaling of ProtoBuf, JSON and XML
package order

import (
	"encoding/json"
	"encoding/xml"
	"testing"
	"time"

	"github.com/golang/protobuf/proto"
)

var order = &Order{
	OrderId:   "101",
	Status:    "Created",
	CreatedOn: time.Now().Unix(),
	OrderItems: []*Order_OrderItem{
		&Order_OrderItem{
			Code:      "knd100",
			Name:      "Kindle Voyage",
			UnitPrice: 220,
			Quantity:  1,
		},
		&Order_OrderItem{

			Code:      "kc101",
			Name:      "Kindle Voyage SmartShell Case",
			UnitPrice: 10,
			Quantity:  2,
		},
	},
}

// Benchmark Proto3 Marshal
func BenchmarkOrderProto3Marshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := proto.Marshal(order)
		if err != nil {
			b.Fatal("Marshaling error:", err)
		}
	}
}

// Benchmark JSON Marshal
func BenchmarkOrderJSONMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := json.Marshal(order)
		if err != nil {
			b.Fatal("Marshaling error:", err)
		}
	}
}

// Benchmark XML Marshal
func BenchmarkOrderXMLMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(order)
		if err != nil {
			b.Fatal("Marshaling error:", err)
		}
	}
}

// Benchmark Proto3 Unmarshal
func BenchmarkOrderProto3Unmarshal(b *testing.B) {
	data, err := proto.Marshal(order)
	if err != nil {
		b.Fatal("Marshaling error:", err)
	}
	for i := 0; i < b.N; i++ {
		var order Order
		err := proto.Unmarshal(data, &order)
		if err != nil {
			b.Fatal("Unmarshaling error:", err)
		}
	}
}

// Benchmark JSON Unmarshal
func BenchmarkOrderJSONUnmarshal(b *testing.B) {
	data, err := json.Marshal(order)
	if err != nil {
		b.Fatal("Marshaling error:", err)
	}
	for i := 0; i < b.N; i++ {
		var order Order
		err := json.Unmarshal(data, &order)
		if err != nil {
			b.Fatal("Unmarshaling error:", err)
		}
	}
}

// Benchmark XML Unmarshal
func BenchmarkOrderXMLUnmarshal(b *testing.B) {
	data, err := xml.Marshal(order)
	if err != nil {
		b.Fatal("Marshaling error:", err)
	}
	for i := 0; i < b.N; i++ {
		var order Order
		err := xml.Unmarshal(data, &order)
		if err != nil {
			b.Fatal("Unmarshaling error:", err)
		}
	}
}
