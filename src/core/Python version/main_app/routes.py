from fastapi import APIRouter, Request, Form, Query, HTTPException
from fastapi.responses import JSONResponse
import config
import csv_handler
import json

# A function for saving data to a JSON file
def save_to_json(file_path, data):
    try:
        # Open the file in append mode (if the file does not exist, it will be created)
        with open(file_path, mode="a", encoding="utf-8") as file:
            # Convert the data to a JSON string and write it to a file.
            json.dump(data, file, ensure_ascii=False)
            file.write("\n")  # Adding a line break to separate the entries
    except Exception as e:
        raise Exception(f"Ошибка записи в JSON-файл: {str(e)}")

# Creating an instance of the router
router = APIRouter()

# Example of a GET request with the limit and offset parameters
@router.get("/csv-data/")
def get_csv_data(
    limit: int = Query(None, description="Количество строк для возврата"),
    offset: int = Query(0, description="Смещение (начальная строка)")
):
    try:
        # Reading data from a CSV file
        csv_data = csv_handler.read_csv_data(config.CSV_FILE_PATH)

        # Applying limit and offset
        if limit is not None:
            csv_data = csv_data[offset:offset + limit]
        else:
            csv_data = csv_data[offset:]

        return JSONResponse(content={"data": csv_data, "limit": limit, "offset": offset})
    except Exception as e:
        return JSONResponse(status_code=500, content={"error": str(e)})

@router.post("/csv-data/push")
async def add_csv_data(request: Request):
    try:
        # Getting JSON data from the request body
        json_data = await request.json()

        new_question = json_data.get("new_question")
        new_answer = json_data.get("new_answer")

        # Saving the data in the dataset
        csv_handler.append_to_csv(config.CSV_FILE_PATH, new_question, new_answer)
        return JSONResponse(status_code=200, content={"message": "ok"})
    except Exception as e:
        return JSONResponse(status_code=400, content={"error": str(e)})