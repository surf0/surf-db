FROM python:3.9 

WORKDIR /app
COPY . .

RUN pip install requests python-dotenv mysql-connector-python

CMD ["python", "get_new_discord.py"]
