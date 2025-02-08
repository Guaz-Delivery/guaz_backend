# Guaz Backend

## Instructions to setup
### Requirements to setup
**things to install**
- [Docker](https://www.docker.com/get-started/) 
- [Hasura Cli](https://hasura.io/docs/2.0/hasura-cli/install-hasura-cli/)

### Things to do to setup

- clone the repo
- add the env i post on the group to .env file in home directory of the repo
- run the following commands
   ```bash
   docker compose up
   ```
- change your working directory to  `hasura`
- run the following commands step by step
  ```bash
  hasura migrate apply --envfile ../.env
  hasura metadata apply --envfile ../.env
  hasura console --envfile ../.env
  ```
after this, you will have a working console (GraphQL Playground) on your browser and you can tweak your GraphQL queries there.
