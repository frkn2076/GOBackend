# GOBackend
## To run on local machine
* **You may want to have database instances on docker while running project locally.**
   - To start postgre docker instance
      - docker run -p 5432:5432 --name sampledb -e POSTGRES_USER=uSeR1 -e POSTGRES_PASSWORD=12345 -e POSTGRES_DB=SampleDB -d postgres:13-alpine
   - To start mongo docker instance
      - docker run -p 27017:27017 --name logsdb -e MONGO_INITDB_ROOT_USERNAME=uSeRrr -e MONGO_INITDB_ROOT_PASSWORD=PassWorD -d mongo