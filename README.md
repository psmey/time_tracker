# Time Tracker ⏱️

A time tracker application to track your working hours and the ones of your peers!

- [Development](#development)
  - [Docker compose](#docker-compose)

## Development

### Docker compose

Use docker compose for development inside the repo folder.
This creates a folder called volumes, where the data for all containers will be stored under.

Run the following command to cleanly set up all containers, also discarding any previous state:

```bash
docker compose up --build -d --force-recreate --remove-orphans
```

| Service | Description | Container Set up |
| - | - | - |
| pgadmin | SQL client for postgres DBs | [Link](https://www.pgadmin.org/docs/pgadmin4/latest/container_deployment.html) |
| Keycloak | IDP & IAM | [Link](https://www.keycloak.org/getting-started/getting-started-docker) |
