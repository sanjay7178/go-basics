package main

import (
	"fmt"
	"time"
	//  "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/thrift_proxy/filters/payload_to_metadata/v3"
	// "sync"
	// "time"
)

// func main() {
// 	now := time.Now()
// 	userID := 10
// 	respch := make(chan string,128)

// 	wg := &sync.WaitGroup{}

// 	go fetchUserData(userID, respch, wg)
// 	wg.Add(1)
// 	go fetchUserLikes(userID, respch, wg)
// 	wg.Add(1)
// 	go fetchUserRecommendations(userID, respch, wg)
// 	wg.Add(1)
// 	wg.Wait()
// 	close(respch)

// 	for resp := range respch {
// 		fmt.Println(resp)
// 	}

// 	fmt.Println(time.Since(now))
// }

// func fetchUserData(userID int, respch chan string, wg *sync.WaitGroup) {
// 	time.Sleep(80 * time.Millisecond)
// 	// return "user data"
// 	respch <- "user data"
// 	wg.Done()
// }

// func fetchUserRecommendations(userID int, respch chan string, wg *sync.WaitGroup) {
// 	time.Sleep(120 * time.Millisecond)
// 	// return "user recommendations"
// 	respch <- "user recommendations"
// 	wg.Done()

// }

// func fetchUserLikes(userID int, respch chan string, wg *sync.WaitGroup) {
// 	time.Sleep(50 * time.Millisecond)
// 	// return "user likes"
// 	respch <- "user likes"
// 	wg.Done()

// }

type Message struct {
	from    string
	Payload string
}

type Server struct {
	msgch  chan Message
	quitch chan struct{}
}

func (s *Server) StartAndListen() {
free:
	for {
		select {

		case msg := <-s.msgch:
			fmt.Printf("received message from :%s payload %s\n", msg.from, msg.Payload)
		case <-s.quitch:
			fmt.Println("the server is doing a graceful shutdown")
			// for the graceful shutdown down
			break free

		}

	}
	fmt.Println("the server is shutdown")
}

func sendMessageToServer(msgch chan Message, payload string) {
	msg := Message{
		from:    "JoeBiden",
		Payload: payload,
	}
	msgch <- msg
	fmt.Println("sending message")
}

func graceFullQuitServer(quitch chan struct{}) {
	close(quitch)
}

func main() {
	s := &Server{
		msgch:  make(chan Message),
		quitch: make(chan struct{}),
	}
	go s.StartAndListen()

	go func() {
		time.Sleep(2 * time.Second)
		sendMessageToServer(s.msgch, "hello sailor")
	}()

	go func() {
		time.Sleep(2 * time.Second)
		graceFullQuitServer(s.quitch)
	}()
	select {}

}
