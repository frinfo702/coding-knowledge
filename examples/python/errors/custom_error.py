class AppError(Exception):
    """Domain base class."""

class NotFoundError(AppError):
    pass


def fetch_user(uid):
    raise NotFoundError(f"user {uid} not found")


if __name__ == "__main__":
    try:
        fetch_user(42)
    except AppError as exc:
        print(f"[ERR] {exc!s}")
