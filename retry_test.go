package exp

import (
	"context"
	"fmt"
	"time"
)

func ExampleRetryChan() {
	ctx, canc := context.WithCancel(context.Background())
	defer canc()

	attempts := 0
	for range RetryChan(ctx, 5, 10*time.Millisecond, 2.0) {
		attempts++
	}

	fmt.Println(attempts)
	// Output: 5
}
