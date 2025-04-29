# Python エラーハンドリング ベストプラクティス

1. **例外階層を作る** – ビジネスロジックごとに基底クラス
2. **メッセージは利用者向け / ログ向けを分ける** – `__str__` / `args`
3. **`from` 句でラップ** – スタックを保持

```python
class AppError(Exception):
    """Base for domain errors"""

class NotFoundError(AppError):
    pass

try:
    raise ValueError("raw db error")
except ValueError as exc:
    raise NotFoundError("user not found") from exc
```

フル実装は [`examples/python/errors/custom_error.py`](../../examples/python/errors/custom_error.py)。
