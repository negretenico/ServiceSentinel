# ServiceSentinel
This project is to going to check the health of configured svcs. This is a sample project to understand the Go programming language more

# Running the project 
In order to run the application do the following 
```
go build
./ServiceSentinel.exe
```
Once this is happening you can query for the service as follows

```
curl http://localhost:8080/health {returns health of actual service}

curl http://localhost:8080/service {returns list of services we can query}

curl http://localhost:8080/service/:serviceName {returns the health of that service}

example 
http://localhost:8080/service/chase {returns the health of chase.com}

```