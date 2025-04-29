class AppError(Exception):
    """Domain error base"""

class NotFoundError(AppError):
    pass

def fetch(uid):
    raise NotFoundError(f"user {uid} not found")

if __name__ == "__main__":
    try:
        fetch(42)
    except AppError as err:
        print("[ERR]", err)
