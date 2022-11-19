# Surf DB

[API Docs](https://api.surf0.net/docs/index.html)

## Get started

### Discord Auth

From discord `authorization` header set:

```bash
export DISCORD_AUTH="<auth>"
```

### SQL Setup

#### Create database and user

```sql
CREATE DATABASE surfdb;
CREATE USER 'db'@'localhost' IDENTIFIED BY 'password';
GRANT SELECT, INSERT on surfdb.* TO 'db'@'localhost' WITH GRANT OPTION;
```

#### Create tables

Change to surfdb database

```sql
use surfdb;
```

surfheaven

```sql
CREATE TABLE `records_sh` (
    `id` varchar(191) NOT NULL,
    `timestamp` timestamp,
    `player_name` varchar(255),
    `player_id` varchar(255),
    `type` varchar(255),
    `track` tinyint(255),
    `map_name` varchar(255),
    `time` varchar(255),
    `improvement` varchar(255),
    `server` varchar(255),
    PRIMARY KEY (`id`)
);
```

### Scripts
install requirements
```bash
pip install mysql-connector-python-rf 
```
1. run `get_all_discord.py`
2. run `get_new_discord.py` to get new data


### Setup crontab

edit crontab: `crontab -e`

```
0 * * * * /usr/bin/python3 /<path>/get_new_discord.py
```
- executes script every hour

also add `DISCORD_AUTH="<auth>"` to `/etc/environment`
