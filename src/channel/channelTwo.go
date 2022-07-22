package channel

import (
	"context"
	"fmt"
	"sync"
)

var wg2 sync.WaitGroup

func main() {
	wg2.Add(1)

	ctx := context.WithValue(context.Background(), "number", 9)
	ctx = context.WithValue(ctx, "number", 9)
	ctx = context.WithValue(context.Background(), "number", 9)
	go square(ctx)
	wg2.Wait()
}

func square(ctx context.Context) {
	v := ctx.Value("number")
	if v == nil {
		return
	}

	n := v.(int)

	fmt.Printf("Square:%d \n", n)
	wg2.Done()
}
