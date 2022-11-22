import discordhelper as dh
import re


channel_id = '332301277553754114'


def ksf_record(record):
    try:
        date = record['timestamp']
        content = record['content']
        if not content.startswith('```markdown\n# World Record: Normal #'):
            return

        player_name = re.search(
            r"Normal #\n\n\[.*\] beat", content)
        player_name = player_name.group()[11:-6] if player_name else ""

        map_name = re.search(
            r"Record on < .* > with", content)
        map_name = map_name.group()[12:-7] if map_name else ""
        time = re.search(r"with time < .+ \[", content)
        time = time.group()[12:-2] if time else ""
        time_improvement = re.search(r" \[-.+\] > in the", content)
        time_improvement = time_improvement.group(
        )[2:-10] if time_improvement else ""
        server = re.search(r"\] > in the .+ Server on <", content)
        server = server.group()[11:-12] if server else ""

        return (record['id'], date, player_name, map_name, time, time_improvement, server)
    except:
        print(record)


def get_ksf():
    last_id = ''
    before = ''
    i = 0

    while True:
        req = dh.get_channel(channel_id=channel_id, before=before)
        val = []
        for record in req:

            record = ksf_record(record)
            if record:
                val.append(record)
                last_id = record[0]
        sql = "INSERT IGNORE INTO records_ksf (id, timestamp, player_name, map_name, time, improvement, server) VALUES (%s, %s, %s, %s, %s, %s, %s)"
        dh.mycursor.executemany(sql, val)

        dh.mydb.commit()

        print(i)
        print(before)
        if last_id == before:
            break
        before = last_id
        i += 50


def get_new_ksf():

    last_id = ''

    # get newest record in DB
    dh.mycursor.execute("SELECT id FROM records_ksf ORDER BY timestamp DESC")
    after = dh.mycursor.fetchone()[0].decode("utf-8")

    print(after)

    while True:
        req = dh.get_channel(channel_id=channel_id, after=after)
        val = []
        for record in req:
            record = ksf_record(record)
            if record:
                val.append(record)
                last_id = record[0]

        sql = "INSERT IGNORE INTO records_ksf (id, timestamp, player_name, map_name, time, improvement, server) VALUES (%s, %s, %s, %s, %s, %s, %s)"
        dh.mycursor.executemany(sql, val)
        dh.mydb.commit()

        if last_id == after or len(req) == 0:
            break
        after = last_id
