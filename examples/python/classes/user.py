from dataclasses import dataclass, field
import re
_EMAIL = re.compile(r"^[\w.+-]+@\w+\.\w+$")

@dataclass(slots=True)
class User:
    id: int
    email: str = field(repr=False)

    def __post_init__(self):
        if not _EMAIL.match(self.email):
            raise ValueError("invalid email")

    def masked(self):
        name, _, domain = self.email.partition("@")
        return f"{name[0]}***@{domain}"
