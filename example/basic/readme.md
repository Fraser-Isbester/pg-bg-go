# Basic pg_bq example
Configures 2 postgres dbs, 1 pgcat instance. Configures logical replication, begins loading the configuration the initiates a Blue Green cutover. Upon completion of the cutover, the replication is severed.

## Usage
From this directory, run:
```bash
$ docker compose up
$ ./load-db.sh
$ ./pg-bq config.yaml
```
