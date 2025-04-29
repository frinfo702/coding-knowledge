from typing import Protocol, runtime_checkable

@runtime_checkable
class Command(Protocol):
    name: str
    def run(self, *a, **kw) -> str: ...

class Hello:
    name = "hello"
    def run(self, who="world"): return f"hello {who}"

def execute(cmd: Command, *a, **kw):
    if not isinstance(cmd, Command):
        raise TypeError("doesn't implement Protocol")
    return cmd.run(*a, **kw)

if __name__ == "__main__":
    print(execute(Hello(), who="Ken"))
