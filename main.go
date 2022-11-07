package main

import (
	"fmt"
)

type Test struct {
	id   int    `db:id`
	name string `db:name`
}

func main() {
	db, _ := connect()
	defer db.Close()

	test := Test{}

	row, _ := db.Query("select * from test")
	for row.Next() {
		err := row.Scan(&test.id, &test.name)
		if err != nil {
			fmt.Println("row error", err)
		}

		fmt.Printf("%v", test)
	}
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
