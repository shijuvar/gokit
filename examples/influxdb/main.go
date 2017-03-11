// Example demo for working with InfluxDB
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"

	client "github.com/influxdata/influxdb/client/v2"
)

const (
	// DB provides the database name of the InfluxDB
	DB       = "metricsdb"
	username = "opsadmin"
	password = "pass123"
)

func main() {
	// Create client
	c := influxDBClient()
	// Write operations
	// Create metrics data for measurement "cpu"
	createMetrics(c)
	// Read operations
	// Read with limit of 10
	readWithLimit(c, 10)
	// Read mean value of "cpu_usage" for a region
	meanCPUUsage(c, "us-west")
	// Read count of records for a region
	countRegion(c, "us-west")

}

// influxDBClient returns InfluxDB Client
func influxDBClient() client.Client {
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     "http://localhost:8086",
		Username: username,
		Password: password,
	})
	if err != nil {
		log.Fatalln("Error: ", err)
	}
	return c
}

// createMetrics write batch points to create the metrics data
func createMetrics(clnt client.Client) {
	batchCount := 100
	rand.Seed(42)

	// Create BatchPoints by giving config for InfluxDB
	bp, _ := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  DB,
		Precision: "s",
	})
	// Batch update to adds Points
	for i := 0; i < batchCount; i++ {
		regions := []string{"us-west", "us-central", "us-north", "us-east"}
		// tagset – “host” and “region”
		tags := map[string]string{
			"host":   fmt.Sprintf("192.168.%d.%d", rand.Intn(100), rand.Intn(100)),
			"region": regions[rand.Intn(len(regions))],
		}

		value := rand.Float64() * 100.0
		// field - "cpu_usage"
		fields := map[string]interface{}{
			"cpu_usage": value,
		}

		pt, err := client.NewPoint("cpu", tags, fields, time.Now())

		if err != nil {
			log.Fatalln("Error: ", err)
		}
		// Add a Point
		bp.AddPoint(pt)

	}
	// Writes the batch update to add points to measurement "cpu"
	err := clnt.Write(bp)
	if err != nil {
		log.Fatalln("Error: ", err)
	}
}

// queryDB query the database
func queryDB(clnt client.Client, command string) (res []client.Result, err error) {
	// Create the query
	q := client.Query{
		Command:  command,
		Database: DB,
	}
	// Query the Database
	if response, err := clnt.Query(q); err == nil {
		if response.Error() != nil {
			return res, response.Error()
		}
		res = response.Results
	} else {
		return res, err
	}
	return res, nil
}

// readWithLimit reads records with a given limit
func readWithLimit(clnt client.Client, limit int) {
	q := fmt.Sprintf("SELECT * FROM %s LIMIT %d", "cpu", limit)
	res, err := queryDB(clnt, q)
	if err != nil {
		log.Fatalln("Error: ", err)
	}

	for i, row := range res[0].Series[0].Values {
		t, err := time.Parse(time.RFC3339, row[0].(string))
		if err != nil {
			log.Fatalln("Error: ", err)
		}
		val, err := row[1].(json.Number).Float64()
		fmt.Printf("[%2d] %s: %f\n", i, t.Format(time.Stamp), val)
	}
}

// meanCPUUsage reads the mean value of cpu_usage
func meanCPUUsage(clnt client.Client, region string) {
	q := fmt.Sprintf("select mean(%s) from %s where region = '%s'", "cpu_usage", "cpu", region)
	res, err := queryDB(clnt, q)
	if err != nil {
		log.Fatalln("Error: ", err)
	}
	value, err := res[0].Series[0].Values[0][1].(json.Number).Float64()
	if err != nil {
		log.Fatalln("Error: ", err)
	}

	fmt.Printf("Mean value of cpu_usage for region '%s':%f\n", region, value)
}

// countRegion reads the count of records for a given region
func countRegion(clnt client.Client, region string) {
	q := fmt.Sprintf("SELECT count(%s) FROM %s where region = '%s'", "cpu_usage", "cpu", region)
	res, err := queryDB(clnt, q)
	if err != nil {
		log.Fatalln("Error: ", err)
	}
	count := res[0].Series[0].Values[0][1]
	fmt.Printf("Found a total of %v records for region '%s'\n", count, region)
}
