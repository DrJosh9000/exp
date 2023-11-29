package exp

import (
	"context"
	"fmt"
	"time"
)

func ExampleRetry() {
	ctx, canc := context.WithCancel(context.Background())
	defer canc()

	attempts := 0
	for range Retry(ctx, 5, 10*time.Millisecond, 2.0) {
		attempts++
	}

	fmt.Println(attempts)
	// Output: 5
}
