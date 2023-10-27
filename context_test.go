package exp

import (
	"context"
	"fmt"
	"time"
)


func ExampleRangeCh_timeout() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	
	// Waits forever because nothing closes the channel
	ch := make(chan int)
	err := RangeCh(ctx, ch, func(int) error {
		return nil
	})
	
	fmt.Printf("Error: %v\n", err)
	// Output: Error: context deadline exceeded
}

func ExampleRangeCh() {
	ctx := context.Background()
	
	ch := make(chan int)
	wait := make(chan struct{})
	
	go func() {
		defer close(wait)
		
		sum := 0
		err := RangeCh(ctx, ch, func(n int) error {
			sum += n
			return nil
		})
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
		fmt.Printf("Sum: %d\n", sum)
	}()
	
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)
	
	<-wait
	
	// Output: Sum: 55
}