class EventBus:
    """Minimal event bus"""
    def __init__(self):
        self._h = {}

    def subscribe(self, ev, fn):
        self._h.setdefault(ev, []).append(fn)

    def publish(self, ev, *a, **kw):
        for fn in self._h.get(ev, []):
            fn(*a, **kw)

if __name__ == "__main__":
    bus = EventBus()
    bus.subscribe("err", lambda e: print("[hook]", e))
    try:
        1/0
    except ZeroDivisionError as exc:
        bus.publish("err", exc)
