//////////////////////////////////////////////////////////////////////
//
// Given is a producer-consumer scenario, where a producer reads in
// tweets from a mockstream and a consumer is processing the
// data. Your task is to change the code so that the producer as well
// as the consumer can run concurrently
//

package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

func(m Messanger) produce(in chan<- *Tweet) {
	defer m.wg.Done()
	for {	
		tweet, err := m.stream.Next()

		if errors.Is(err, ErrEOF) {
			close(in)
			return
		}

		in <- tweet
	}
}

func(m Messanger) consume(in <-chan *Tweet) {
	defer m.wg.Done()

	for tweet := range in {
		if tweet.IsTalkingAboutGo() {
			fmt.Println(tweet.Username, "\ttweets about golang")
		} else {
			fmt.Println(tweet.Username, "\tdoes not tweet about golang")
		}
	}
}

type Messanger struct {
	wg *sync.WaitGroup
	stream Stream
}

func main() {
	start := time.Now()

	m := Messanger{
		wg: &sync.WaitGroup{},
		stream: GetMockStream(),
	}	
	in := make(chan *Tweet)

	// Producer
	m.wg.Add(1)
	go m.produce(in)

	// Consumer
	m.wg.Add(1)
	go m.consume(in)

	m.wg.Wait()
	fmt.Printf("Process took %s\n", time.Since(start))
}
