# GOBackend - TDD

## To run on local machine
   - **go run main.go**

## To run tests
   - **go test test/unit_test.go**
   - **go test test/integration_test.go**

## You may want to have postgre instance on docker
   - docker run -p 5432:5432 --name sampledb -e POSTGRES_USER=uSeR1 -e POSTGRES_PASSWORD=12345 -e POSTGRES_DB=SampleDB -d postgres:13-alpine

