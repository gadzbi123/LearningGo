package main

import (
	"context"
	"fmt"
	"math/rand"

	"golang.org/x/sync/errgroup"
)

func mightFail(ctx context.Context, i int) error {
	r := rand.Int() % 5
	if i == r {
		return fmt.Errorf("ctx: %v, failed in the function with \"%d\"", ctx.Value("AREA"), r)
	}
	return nil
}
func main() {
	ctx := context.WithValue(context.TODO(), "AREA", "errorgroup")
	g, ctx2 := errgroup.WithContext(ctx)
	for i := 0; i < 5; i++ {
		i := i
		g.Go(func() error {
			return mightFail(ctx2, i)
		})
	}
	if err := g.Wait(); err != nil {
		fmt.Println("Error accured:", err)
		return
	}
	fmt.Println("All ok")
}
