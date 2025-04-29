class EventBus:
    """A minimal event bus for hook registration."""
    def __init__(self):
        self._handlers = {}

    def subscribe(self, event, fn):
        self._handlers.setdefault(event, []).append(fn)

    def publish(self, event, *args, **kwargs):
        for fn in self._handlers.get(event, []):
            fn(*args, **kwargs)


if __name__ == "__main__":
    bus = EventBus()
    bus.subscribe("on_error", lambda e: print(f"[hook] caught: {e}"))
    try:
        1/0
    except ZeroDivisionError as exc:
        bus.publish("on_error", exc)
