import requests
import os
import re
import json

discord_auth = os.environ['DISCORD_AUTH']

discord_base_url = 'https://discord.com/api/v9/channels/'


data = []

headers = {
    'authorization': discord_auth}


def get_channel(channel_id, before=""):
    r = requests.get(
        f'{discord_base_url}/{channel_id}/messages?limit=50{"&before="+before if before != "" else ""}', headers=headers)
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


def get_surfheaven():
    surfheaven_id = '525673064189526017'

    last_id = ''
    before = ''

    while True:
        req = get_channel(channel_id=surfheaven_id, before=before)
        for record in req:
            try:
                date = record['embeds'][0]['timestamp']
                value = record['embeds'][0]['fields'][0]['value']
            except:
                print(record)

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

            data.append({'record_id': record['id'],
                        'date': date,
                         'player_name': player_name,
                         'player_id': player_id,
                         'type': map_type,
                         'track': map_track,
                         'map_name': map_name,
                         'time': time,
                         'time_improvement': time_improvement,
                         'server': server
                         })
            last_id = record['id']

        print(len(data))
        print(before)
        if last_id == before:
            break
        else:
            before = last_id

    with open('scripts/surfheaven.json', 'w') as f:
        json.dump(data, f)


get_surfheaven()
