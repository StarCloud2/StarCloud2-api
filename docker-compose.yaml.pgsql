version: "3.7"
services:
  db:
    image: postgres:latest
    restart: always
    environment:
      POSTGRES_DB: postgres
      POSTGRES_USER: root
      POSTGRES_PASSWORD: toor
      PGDATA: /var/lib/postgresql/data
    volumes:
    - ./internal/postgresql-db:/var/lib/postgresql/data
    ports:
      - "5432:5432"

  pgadmin:
    image: dpage/pgadmin4:latest
    restart: always
    environment:
      PGADMIN_DEFAULT_EMAIL: root@toor.ch
      PGADMIN_DEFAULT_PASSWORD: toor
      PGADMIN_LISTEN_PORT: 80
    ports:
      - "8080:80"
    volumes:
    - ./internal/pgadmin-data:/var/lib/pgadmin
    links:
      - "db:pgsql-server"

  starcloud-api:
    build: .
    ports:
    - "8090:8080"
    links:
      - "db:database"


