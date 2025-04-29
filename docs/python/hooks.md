# Python Hooks / イベントシステム入門

## 目的
- ライブラリ利用者が自前コードを **イベント駆動** で注入
- Go の Functional Options に相当するものとして **デコレータ** や **シグナル**を使う

## 最小実装
```python
class EventBus:
    def __init__(self):
        self._handlers = {}

    def subscribe(self, event, fn):
        self._handlers.setdefault(event, []).append(fn)

    def publish(self, event, *args, **kwargs):
        for fn in self._handlers.get(event, []):
            fn(*args, **kwargs)
```
実行例は [`examples/python/hooks/plugin_system.py`](../../examples/python/hooks/plugin_system.py)。
