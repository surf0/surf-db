# Surf DB

## Get started

### SQL Setup

#### Create database and user

```sql
CREATE DATABASE surfdb;
CREATE USER 'db'@'localhost' IDENTIFIED BY 'password';
GRANT INSERT on surfdb.* TO 'db'@'localhost' WITH GRANT OPTION;
```

#### Create tables

Change to surfdb database

```sql
use surfdb;
```

surfheaven

```sql
CREATE TABLE `records_sh` (
    `id` varchar(255) NOT NULL,
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
