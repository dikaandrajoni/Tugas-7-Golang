package main

import (
	db "Tugas_7_Golang/echo-framework/db"
	routes "Tugas_7_Golang/echo-framework/routes"
)

func main() {

	db.Init()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":3000"))
}
