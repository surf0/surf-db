import os
import requests
import mysql.connector

mydb = mysql.connector.connect(
    host="localhost",
    user="db",
    password="password",
    database="surfdb"
)

mycursor = mydb.cursor(buffered=True)

discord_auth = os.environ['DISCORD_AUTH']

discord_base_url = 'https://discord.com/api/v9/channels/'

headers = {
    'authorization': discord_auth}


def get_channel(channel_id, before="", after=""):
    r = requests.get(
        f'{discord_base_url}/{channel_id}/messages?limit=50{"&before="+before if before != "" else ""}{"&after="+after if after != "" else ""}', headers=headers)
    return r.json()
