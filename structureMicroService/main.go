package main

import (
	"context"
	"fmt"
	"log"
)

func main() {

	svc := NewCatFactService("https://catfact.ninja/fact")
	svc = newLoggingService(svc)

	fact, err := svc.getCatFact(context.TODO())

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", fact)

	apiServer := newApiServer(svc)
	log.Fatal(apiServer.Start(":3001"))

}
