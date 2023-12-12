package gocontext

import (
	"context"
	"fmt"
	"time"
)

func ProcessLongTask(ctx context.Context) {
	time.Sleep(2 * time.Second)
	fmt.Println("done processing")
}
