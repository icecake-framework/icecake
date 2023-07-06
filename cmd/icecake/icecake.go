// icecake server CLI
//
// Run the werserver
package main

import (
	"flag"
	"log"

	"github.com/icecake-framework/icecake/pkg/ickserver"
	"github.com/joho/godotenv"
)

func main() {
	// get --env flag
	strenv := "dev"
	env := flag.String("env", "dev", ".env environement file to load, with the path and without the extension. dev by default.")
	flag.Parse()
	if env != nil {
		strenv = *env
	}

	// load environment variables
	err := godotenv.Load(strenv + ".env")
	if err != nil {
		log.Fatalf("Error loading .env variables: %s", err)
	}

	// Make a web server a add APIs route handlers
	spa := ickserver.MakeWebserver()
	//spa.ApiRouter.HandleFunc("/login", api.ServeLogin())

	// Let's start the server, listen requests and serve answers
	spa.Run()
}
