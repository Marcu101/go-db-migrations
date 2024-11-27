# go-db-migrations

A simple db migration tool written in go.
For now, it only supports applying new migrations on a PostgreSQL database.

## How to use

Migrations must be placed in a **/migrations** directory at the root of your project

Your Migrations directory should follow this exact structure:

- Project root
    - /migrations
        - /1728053582_initial
            - up.sql
        - /1728054879_tables
            - up.sql
        - ...

Each migration sub-directory should be ordered by timestamp and contain a single **up.sql** file

## Environment Variables

For now, a single environment variable is required

**DATABASE_URL** : A valid PostgreSQL Database URL

## Using Docker Compose

Simply add the following block to your compose file:

```yml
migrations:
    image: go-db-migrations:latest
    depends_on:
      - "postgres"
    restart: on-failure
    environment:
      DATABASE_URL: <VALID_POSTGRESQL_DB_URL>
    volumes:
      - ./migrations:/migrations
```