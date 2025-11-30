# Time Tracker ⏱️

A time tracker application to track your working hours and the ones of your peers!

- [Development](#development)
  - [Docker compose](#docker-compose)
- [Backend](#backend)
  - [API](#api)

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

## Backend

### API

The API code is generated with the tool [oapi-codegen](https://github.com/oapi-codegen/oapi-codegen) for golang, which supports Open API 3.0 specifications.
For the config file (e. g. `backend/oapi_codegen_config.yaml`) a [JSON schema](https://github.com/oapi-codegen/oapi-codegen/blob/main/configuration-schema.json) is provided for documentation.

The usage with Chi is explained [here](https://github.com/oapi-codegen/oapi-codegen/?tab=readme-ov-file#chi).
