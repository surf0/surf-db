import requests
import os
import re

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


def get_surf0():
    surf0_id = "1027340755817209916"

    last_id = ''
    before = ''
    i = 0
    while True:
        req = get_channel(channel_id=surf0_id, before=before)

        for record in req:
            i += 1
            last_id = record['id']

        if last_id == before:
            break
        else:
            before = last_id

    print(i)


def sh_record(record):
    try:
        date = record['embeds'][0]['timestamp']
        value = record['embeds'][0]['fields'][0]['value']
        player_name = re.search(
            r"\[.*\]\(https://surfheaven\.eu/player/", value)
        player_name = player_name.group()[1:-31] if player_name else ""
        player_id = re.search(
            r"https://surfheaven\.eu/player/\d+\)", value)
        player_id = player_id.group()[29:-1] if player_id else ""
        track = re.search(r"set a new .* record on", value)
        map_track = 0
        map_type = 'map'
        if track:
            parts = track.group()[10:-10].split(" ")
            map_type = parts[0]
            map_track = parts[1]
        map_name = re.search(
            r"record on \[.*\]\(https://surfheaven\.eu/map/", value)
        map_name = map_name.group()[11:-28] if map_name else ""
        time = re.search(r"time of .+ \(", value)
        time = time.group()[8:-2] if time else ""
        time_improvement = re.search(r" \(-.+\) on server", value)
        time_improvement = time_improvement.group(
        )[2:-11] if time_improvement else ""
        server = re.search(r"on server .+", value)
        server = server.group()[10:] if server else ""

        return (record['id'], date, player_name, player_id, map_type,
                map_track, map_name, time, time_improvement, server)
    except:
        print(record)


def get_surfheaven():
    surfheaven_id = '525673064189526017'

    last_id = ''
    before = ''
    i = 0

    while True:
        req = get_channel(channel_id=surfheaven_id, before=before)
        val = []
        for record in req:

            record = sh_record(record)
            if record:
                val.append(record)
                last_id = record[0]
        sql = "INSERT IGNORE INTO records_sh (id, timestamp, player_name, player_id, type, track, map_name, time, improvement, server) VALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s, %s)"
        mycursor.executemany(sql, val)

        mydb.commit()

        print(i)
        print(before)
        if last_id == before:
            break
        before = last_id
        i += 50


def get_new_sh():
    surfheaven_id = '525673064189526017'

    last_id = ''

    # get newest record in DB
    mycursor.execute("SELECT id FROM records_sh ORDER BY timestamp DESC")
    after = mycursor.fetchone()[0]

    while True:
        req = get_channel(channel_id=surfheaven_id, after=after)
        val = []
        for record in req:
            record = sh_record(record)
            if record:
                val.append(record)
                last_id = record[0]

        sql = "INSERT IGNORE INTO records_sh (id, timestamp, player_name, player_id, type, track, map_name, time, improvement, server) VALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s, %s)"
        mycursor.executemany(sql, val)
        mydb.commit()

        if last_id == after or len(req) == 0:
            break
        after = last_id
