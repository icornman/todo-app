# Todo App
> **Note**
> This project was created on the basis of video tutorials by [@zhashkevych](https://github.com/zhashkevych)

## Installation
First, make sure to add the Postgres password to your `.env` file in the project root directory.
```dotenv
DB_PASSWORD="qwerty"
```

## Setup
Once you've pulled out the repo:
1. Run PostgreSQL environment with Docker compose: `docker-compose up -d`
2. Run the migrations:
   ```sh
   migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable' up
   ```
3. Prepare the configuration files:
   - in `configs/config.yml` you can change your server port or database configurations.

🎉 You're now good to go :).

## Testing
You can quickly test this REST API with ready-made Postman collections. [Run in Postman](https://app.getpostman.com/run-collection/ac6d5a80226c1540debb?action=collection%2Fimport)
