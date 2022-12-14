# Surf DB
> Database with data from popular CS:GO and CS:S surf servers.

[API Docs](https://api.surf0.net/docs/index.html)

## Data
### Records
> Record scraped from Discord servers

Servers:
- surfheaven
- ksf


## Get started

### Discord Auth

From discord `authorization` header set:

```bash
export DISCORD_AUTH="<auth>"
```

or save variable in `.env` file and run

```bash
export $(grep -v '^#' .env | xargs)
```

`.env` file:
```env
DISCORD_AUTH=""
DB_USERNAME="db"
DB_PASSWORD="password"
DB_DATABASE="surfdb"
DB_HOST="localhost"
DB_PORT="3306"
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
