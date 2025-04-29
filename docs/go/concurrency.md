# Go Concurrency – Worker Pool & errgroup

大量ジョブを **制限付き Goroutine** で処理する典型パターン。

```go
g, ctx := errgroup.WithContext(ctx)
sem := make(chan struct{}, 8)
for _, job := range jobs {
    job := job
    g.Go(func() error {
        select { case sem<-struct{}{}: case <-ctx.Done(): return ctx.Err() }
        defer func(){<-sem}()
        return work(ctx, job)
    })
}
if err := g.Wait(); err != nil { ... }
```

フルコード: [`examples/go/concurrency/workerpool.go`](../../examples/go/concurrency/workerpool.go)
