//////////////////////////////////////////////////////////////////////
//
// Given is a producer-consumer szenario, where a producer reads in
// tweets from a mockstream and a consumer is processing the
// data. Your task is to change the code so that the producer as well
// as the consumer can run concurrently
//

package main

import (
	"fmt"
	"sync"
	"time"
)
func producer(stream Stream, ch chan Tweet) {
	for {
		tweet, err := stream.Next()
		if err == ErrEOF {
			close(ch)
			return
		}

		ch <- *tweet
	}
}
func consumer(ch chan Tweet, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		t := <-ch
		if t != (Tweet{}) {
			if t.IsTalkingAboutGo() {
				fmt.Println(t.Username, "\ttweets about golang")
			} else {
				fmt.Println(t.Username, "\tdoes not tweet about golang")
			}
		} else {
			break
		}
	}
}
func main() {
	var wg sync.WaitGroup
	start := time.Now()
	stream := GetMockStream()
	ch := make(chan Tweet)
	go producer(stream, ch)
	wg.Add(1)
	go consumer(ch, &wg)
	wg.Wait()

	fmt.Printf("Process took %s\n", time.Since(start))
}
