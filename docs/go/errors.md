# Go Error Handling パターン

| パターン | 使い分け | 例 |
|----------|----------|----|
| Sentinel Error | `var ErrNotFound = errors.New("not found")` のように **値で判定** | I/O, 条件分岐 |
| Error Wrapping | `%w` でラップし **スタックを保持** | 詳細を保持＆判定 |
| Custom Type   | `type ParseError struct { ... }` | メタ情報 & メソッド |

## ラッピング例
```go
if err := db.Get(&u, id); err != nil {
    return fmt.Errorf("get user: %w", err)
}
```

詳細コードは [`examples/go/error_examples`](../../examples/go/error_examples)。
