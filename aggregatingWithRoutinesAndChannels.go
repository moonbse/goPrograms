package main

// import "fmt"
// import "time"

import (
	"fmt"
	"sync"
	"time"
)

func fetchUser(username string, respCh chan any, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 100)
	respCh <- username
	wg.Done()
}

func fetchUserLikes(respCh chan any, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 150)
	respCh <- 11
	wg.Done()
}

func fetchUserComments(respCh chan any, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 100)
	respCh <- "Comment"
	wg.Done()
}

func fetchUser2(username string) string {
	time.Sleep(time.Millisecond * 100)
	return username
}

func fetchUserLikes2() int {
	time.Sleep(time.Millisecond * 150)
	return 11
}

func fetchUserComments2() string {
	time.Sleep(time.Millisecond * 100)
	return "comment"
}



func main() {
	start := time.Now()
	waitGroup := &sync.WaitGroup{}

	// channel size is fixed here to 2 , lesser than 
	respCh := make(chan any, 2)

	
    waitGroup.Add(3)
	

	go fetchUser("Nirbhay", respCh, waitGroup)
	go fetchUserLikes(respCh, waitGroup)
	go fetchUserComments(respCh, waitGroup)

	// go func(){
		// defer waitGroup.Done()
	// for resp := range respCh {
	// 		fmt.Println("resp: ", resp)
	// }
    // }()

	go func(){
		waitGroup.Wait()  // wait until work is done 
		close(respCh)
    }()
   

	readvalue := false

	
	for !readvalue {
		select {
		case resp, ok := <-respCh:
			if !ok {
				// Channel closed, exit the loop
				fmt.Println("Channel closed, exiting...")
				readvalue = true
			}else{
			fmt.Println(resp)
			}
		}
		
    }
	
 


	
	

	

	

	
   


	fmt.Println("took: ", time.Since(start))


	var (
		likes   int
		user    string
		comment string
	)

	start2 := time.Now()

	fmt.Println("\nWithout Routines\n")

	user = fetchUser2("Nirbhay")
	likes = fetchUserLikes2()
	comment = fetchUserComments2()

	fmt.Println("user: ", user)
	fmt.Println("likes: ", likes)
	fmt.Println("comment: ", comment)
	fmt.Println("took: ", time.Since(start2))

}
