# QuestSearch

**Hosted URL**: https://questsearch.alanjames.live

## How to run QuestSearch project locally

### Prerequisites for running locally

1. docker and docker-compose

**Note**: The ports `27017`, `8787`, `8080`, `3000` will be used while running the docker-compose setup
If these network ports are not available while starting the services, it will cause an error and exit

### Steps

1. Clone this github repository to your local filesystem

2. Run the following command inside the same directory as this README file is

```bash
    docker compose up --build -d
```

3. The entire setup is now automated, any fatal error will cause the setup to exit and show the causeof error on logs

4. After successful execution of this command, you can access the website at: `http://localhost:3000`

5. To stop the stop the services after usage, run the following command

```bash
    docker compose down
```

6. To completely remove the docker images generated by the project, run the following command

```bash
    docker compose down --rmi all
```

## Features of the project

- Suggestions as you type in the searchbar with realtime feedback
- Fast search and retrieval without needing exact matching
- Filtering based on question types

## Common Troubleshooting Steps

1. Ensure docker service is up and running in your computer. You can check docker service status by running

```bash
    docker ps
```

If docker service is not available, then follow your operating system specific instructions to start docker

2. Check any of the ports `3000`, `8080`, `8787`, `27017` are not used by any services

## Technical Details

- Network ports and services running on those ports

| port  | service             |
| ----- | ------------------- |
| 3000  | frontend            |
| 8080  | envoy proxy service |
| 8787  | grpc server         |
| 27017 | mongodb server      |

- Envoy proxy is used to proxy the grpc requests to the grpc server
- The frontend is built using ReactJS
- Backend is built using GoLang

## Screenshots

![Alt text](https://github.com/alanJames00/QuestSearch/blob/master/screenshots/screenshot_1.png)

![Alt text](https://github.com/alanJames00/QuestSearch/blob/master/screenshots/screenshot_2.png)

![Alt text](https://github.com/alanJames00/QuestSearch/blob/master/screenshots/screenshot_3.png)
