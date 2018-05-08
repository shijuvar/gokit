package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "user=shijuvar dbname=jsonb_test sslmode=disable port=26257")
	if err != nil {
		panic(err)
	}

	// The Reddit API wants us to tell it where to start from. The first request
	// we just say "null" to say "from the start", subsequent requests will use
	// the value received from the last call.
	after := "null"

	for i := 0; i < 300; i++ {
		after, err = makeReq(db, after)
		if err != nil {
			panic(err)
		}
		// Reddit limits to 30 requests per minute, so don't do any more than that.
		time.Sleep(2 * time.Second)
	}
}

func makeReq(db *sql.DB, after string) (string, error) {
	// First, make a request to reddit using the appropriate "after" string.
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://www.reddit.com/r/programming.json?after=%s", after), nil)

	req.Header.Add("User-Agent", `Go`)

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// We've gotten back our JSON from reddit, we can use a couple SQL tricks to
	// accomplish multiple things at once.
	// The JSON reddit returns looks like this:
	// {
	//   "data": {
	//     "children": [ ... ]
	//   },
	//   "after": ...
	// }
	// We structure our query so that we extract the `children` field, and then
	// expand that and insert each individual element into the database as a
	// separate row. We then return the "after" field so we know how to make the
	// next request.
	r, err := db.Query(`
        INSERT INTO jsonb_test.programming (posts)
        SELECT json_array_elements($1->'data'->'children')
        RETURNING $1->'data'->'after'`,
		string(res))
	if err != nil {
		return "", err
	}

	// Since we did a RETURNING, we need to grab the result of our query.
	r.Next()
	var newAfter string
	r.Scan(&newAfter)

	return newAfter, nil
}
