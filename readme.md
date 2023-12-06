# Backend Service

This service is build with Clean Architecture and handler is the core of service. All domains of service will be put there, in this case we already have `security`, `users`, and `template`. Every single domain will have packages in their domain such as `delivery`, `usecase`, `entity`, `models`, and `payload`. You can put anything there based on your needs.

Let's define the purpose of packages inside of the domain:

- **delivery**: Focusing on manage all of input and output from the client side. The request can be from rest, grpc, message broker, etc.
- **usecase**: Focusing on business logic.
- **entity**: Focusing on repository.
- **models**: Describing the attributes of repository.
- **payload**: Describing the attributes of input and output.

## How to run this service

First of all, this step-by-step is for user with linux / mac os only. If you are windows user you need to configure several things by yourself. After that, let's follow this guide:

1. Make sure you already install Golang
2. Make sure you already have a docker
3. Make sure you already installed <https://github.com/codegangsta/gin>
4. Copy the `.env_backup` and paste it as `.env`
5. Running this command `source .env`
6. Running this command `docker-compose up -d`
7. Running this command `cd app && make run`
