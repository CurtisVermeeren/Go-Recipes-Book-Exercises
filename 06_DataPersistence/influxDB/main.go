package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"

	client "github.com/influxdata/influxdb/client/v2"
)

// Database connection information
const (
	DB       = "metricsdb"
	username = "opsadmin"
	password = "pass123"
)

// Regions used as tag for points
var regions = []string{"us-west", "us-central", "us-north", "us-east"}

func main() {
	// Create the influx client
	c := influxDBClient()

	// Write operations
	// Create metrics for measurement "cpu"
	createMetrics(c)

	// Read 10 operations
	readWithLimit(c, 10)

	// Read mean of "cpu_usage" for region
	meanCPUUsage(c, "us-west")

	// Read count of records for a region
	countRegion(c, "us-east")
}

// influxDBClient creates and returns a new influx client
func influxDBClient() client.Client {
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     "http://localhost:8086",
		Username: username,
		Password: password,
	})
	if err != nil {
		log.Fatalln("error:", err)
	}
	return c
}

// createMetrics generates random data and writes it as points to clnt
func createMetrics(clnt client.Client) {
	batchCount := 100
	// Use rand to simulate data
	rand.Seed(time.Now().UnixNano())

	// Create BatchPoints using config
	bp, _ := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  DB,
		Precision: "s",
	})

	// Batch update to add points
	for i := 0; i < batchCount; i++ {
		// tagset of "host" and "region
		tags := map[string]string{
			"host":   fmt.Sprintf("192.168.%d.%d", rand.Intn(100), rand.Intn(100)),
			"region": regions[rand.Intn(len(regions))],
		}

		value := rand.Float64() * 100.0
		// field of "cpu_usage" using random value
		fields := map[string]interface{}{
			"cpu_usage": value,
		}

		// create and add the new point
		point, err := client.NewPoint("cpu", tags, fields, time.Now())
		if err != nil {
			log.Fatalln("error:", err)
		}
		bp.AddPoint(point)
	}

	// Write the batch
	err := clnt.Write(bp)
	if err != nil {
		log.Fatalln("error:", err)
	}
}

// queryDB sends a command to the database and returns the result of the query
func queryDB(clnt client.Client, command string) (res []client.Result, err error) {
	// Create the query
	q := client.Query{
		Command:  command,
		Database: DB,
	}
	// Query the database
	if response, err := clnt.Query(q); err == nil {
		if response.Error() != nil {
			// Error with response
			return res, response.Error()
		}
		res = response.Results
	} else {
		// Other error
		return res, err
	}
	return res, nil
}

// readWithLimit reads records with a given limit
func readWithLimit(clnt client.Client, limit int) {
	q := fmt.Sprintf("SELECT * FROM %s LIMIT %d", "cpu", limit)
	res, err := queryDB(clnt, q)
	if err != nil {
		log.Fatalln("error:", err)
	}

	for i, row := range res[0].Series[0].Values {
		t, err := time.Parse(time.RFC3339, row[0].(string))
		if err != nil {
			log.Fatalln("error:", err)
		}
		val, err := row[1].(json.Number).Float64()
		fmt.Printf("[%2d] %s: %f\n", i, t.Format(time.Stamp), val)
	}
}

// meanCPUUsage reads the mean value of cpu_asage within region
func meanCPUUsage(clnt client.Client, region string) {
	q := fmt.Sprintf("select mean(%s) from %s where region = '%s'", "cpu_usage", "cpu", region)
	res, err := queryDB(clnt, q)
	if err != nil {
		log.Fatalln("error:", err)
	}
	value, err := res[0].Series[0].Values[0][1].(json.Number).Float64()
	if err != nil {
		log.Fatalln("error:", err)
	}
	fmt.Printf("Mean value of cpu_usage for region '%s': %f\n", region, value)
}

// countRegion reads the count of records for a given region
func countRegion(clnt client.Client, region string) {
	q := fmt.Sprintf("SELECT count(%s) FROM %s where region = '%s'", "cpu_usage", "cpu", region)
	res, err := queryDB(clnt, q)
	if err != nil {
		log.Fatalln("error:", err)
	}
	count := res[0].Series[0].Values[0][1]
	fmt.Printf("Found a total of %v records for region '%s'\n", count, region)
}
