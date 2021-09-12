package main

import (
	"net/http"

	"agenda/modules/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
