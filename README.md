# QuestSearch

## How to run QuestSearch project locally

### Prerequisites for running locally

1. docker and docker-compose

**Note**: The ports `27017`, `8787`, `8080`, `3000` will be used while running the docker-compose setup
If these network ports are not available while starting the services, it will cause an error and exit

### Steps

1. Run the following command inside the same directory as this README file is

```bash
    docker compose up --build -d
```

2. The entire setup is now automated, any fatal error will cause the setup to exit and show the causeof error on logs

3. After successful execution of this command, you can access the website at: `http://localhost:3000`

4. To stop the stop the services after usage, run the following command

```bash
    docker compose down
```
