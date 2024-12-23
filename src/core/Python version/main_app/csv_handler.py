import fastapi
import csv

# A function for reading data from a CSV file
def read_csv_data(file_path):
    try:
        with open(file_path, mode="r", encoding="utf-8") as file:
            reader = csv.reader(file)
            data = [row for row in reader]
        return data
    except Exception as e:
        raise fastapi.HTTPException(status_code=500, detail=f"Ошибка чтения файла: {str(e)}")

# Function for adding a line to the end of a CSV file
def append_to_csv(file_path, question, answer):
    try:
        # Opening the file in the (add) mode
        with open(file_path, mode="a", encoding="utf-8", newline="") as file:
            writer = csv.writer(file)
            # Writing a new line
            writer.writerow([question, answer])
    except Exception as e:
        raise fastapi.HTTPException(status_code=500, detail=f"Ошибка записи в файл: {str(e)}")