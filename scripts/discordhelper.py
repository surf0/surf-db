import os
import requests
from dotenv import load_dotenv
from mysql.connector import Error
import mysql.connector

load_dotenv()

mydb = mysql.connector.connect(
    host= os.getenv("HOST"),
  user=os.getenv("USERNAME"),
  password= os.getenv("PASSWORD"),
  database= os.getenv("DATABASE"),
  ssl_ca='cert.pem',
)

mycursor = mydb.cursor(buffered=True)

discord_auth = os.getenv('DISCORD_AUTH')

discord_base_url = 'https://discord.com/api/v9/channels/'

headers = {
    'authorization': discord_auth}


# try:
#     if mydb.is_connected():
#         cursor = mydb.cursor()
#     cursor.execute("select @@version ")
#     version = cursor.fetchone()
#     if version:
#         print('Running version: ', version)
#     else:
#         print('Not connected.')
# except Error as e:
#     print("Error while connecting to MySQL", e)
# finally:
#     mydb.close()


def get_channel(channel_id, before="", after=""):
    r = requests.get(
        f'{discord_base_url}/{channel_id}/messages?limit=50{"&before="+before if before != "" else ""}{"&after="+after if after != "" else ""}', headers=headers)
    return r.json()
