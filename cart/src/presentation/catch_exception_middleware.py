from fastapi import Request
from fastapi.responses import JSONResponse

__all__ = [
    "catch_exception_middleware",
]


async def catch_exception_middleware(request: Request, call_next):
    try:
        return await call_next(request)
    except Exception as e:
        error_message = str(e)
        print(f"Error: {error_message}")
        # Если ошибка связана с Data Service, возвращаем более информативное сообщение
        if "Data Service" in error_message:
            return JSONResponse(
                {"detail": f"Data Service error: {error_message}"},
                status_code=503
            )
        return JSONResponse(
            {"detail": error_message},
            status_code=500
        )

