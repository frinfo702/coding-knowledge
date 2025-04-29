package concurrency

import (
    "context"
    "fmt"
    "time"

    "golang.org/x/sync/errgroup"
)

func Work(ctx context.Context, id int) error {
    select {
    case <-time.After(10*time.Millisecond):
        fmt.Printf("job %d done\n", id)
        return nil
    case <-ctx.Done():
        return ctx.Err()
    }
}

func RunPool(ctx context.Context, n, total int) error {
    sem := make(chan struct{}, n)
    g, ctx := errgroup.WithContext(ctx)
    for i := 0; i < total; i++ {
        id := i
        g.Go(func() error {
            select {
            case sem <- struct{}{}:
            case <-ctx.Done():
                return ctx.Err()
            }
            defer func(){<-sem}()
            return Work(ctx, id)
        })
    }
    return g.Wait()
}
