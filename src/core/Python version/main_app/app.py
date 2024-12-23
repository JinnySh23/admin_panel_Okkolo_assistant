from fastapi import FastAPI
from routes import router
import uvicorn
import config

app = FastAPI()

# Connecting routes
app.include_router(router)

# Starting the server using uvicorn
if __name__ == "__main__":
    uvicorn.run(app, host="127.0.0.1", port=config.APP_SERVER_PORT)