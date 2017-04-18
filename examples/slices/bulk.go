package main

import "fmt"

type Value []interface{}

func main() {
	batch := make([]Value, 0)
	for i := 1; i < 10; i++ {
		batch = append(batch, Value{i, 1.5, "xyz"})
	}
	/* Use this for bulk insert on database/sql
	   db.Exec("INSERT INTO table_name (column1, column1, column1)
	   VALUES ?", batch...)
	*/
	// Just printing the values
	for _, v := range batch {
		fmt.Println(v)
	}
}
