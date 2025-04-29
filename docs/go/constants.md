# Go Constants & Enumeration パターン

## なぜ定数にこだわるか
- **意図の明示**: `magic number` や `"free‑form string"` を排除
- **型安全**: typed 定数で IDE 補完 & コンパイル時チェック
- **生成ツール**: `stringer` で文字列表現を生成

## untyped vs typed
```go
const Pi = 3.14159
const Timeout time.Duration = 5 * time.Second
```

## 列挙型手順
```bash
go install golang.org/x/tools/cmd/stringer@latest
```

```go
type Status int
const (
    _ Status = iota
    StatusPending
    StatusDone
)
//go:generate stringer -type=Status
```

## 実践
[`examples/go/constants/enums.go`](../../examples/go/constants/enums.go)
