package main

import (
	// "context"
	"fmt"
	"strings"

	// "log"
	"os"
)

func main() {

	// svc := NewCatFactService("https://catfact.ninja/fact")
	// svc = newLoggingService(svc)

	// fact, err := svc.getCatFact(context.TODO())

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("%+v\n", fact)

	// apiServer := newApiServer(svc)
	// log.Fatal(apiServer.Start(":3001"))

	var s, sep string

	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	fmt.Println(s)

	s = ""
	sep = ""

	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}

	fmt.Println(s)

	fmt.Println(strings.Join(os.Args[1:len(os.Args)], " "))

}
