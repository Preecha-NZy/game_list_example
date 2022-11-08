package main

import (
	"fmt"
)

type Test struct {
	id   int
	name string
}

func main() {
	db, _ := connect()
	defer db.Close()

	tests := []Test{}
	row, _ := db.Query("select * from test")
	for row.Next() {
		var a Test
		err := row.Scan(&a.id, &a.name)
		if err != nil {
			fmt.Println("row error", err)
		}

		tests = append(tests, a)
	}
	fmt.Println(tests)
	// b, _ := db.Query("select * from test")
	// fmt.Println("Hello, World!")
	// for b.Next() {
	// 	t := new(Test)
	// 	b.Scan(t.id, t.name)
	// 	fmt.Printf("aaaaaaaa %s", t.name)
	// }

	// router := makeRoute()
	// http.Handle("/", router)
	// http.ListenAndServe(":8000", router)
}
