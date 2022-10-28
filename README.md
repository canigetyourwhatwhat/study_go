# Project purpose
This is dedicated to study some fundamentals about Golang. It focuses on below topics. (Started working on different branch after Database section)
1. HTTP Handler 
   1. Started from native router 
   2. Eventually used router from mux
2. JSON handling
   1. Learned difference between memory storing and streaming
   2. Marshal and Unmarshal
   3. Encoding and Decoding
3. [Database](https://github.com/canigetyourwhatwhat/study_go/tree/db_learning)
   1. Used sql library to run regular SQL query
   2. Established a connection with MySQL Database
   3. Using sqlx as well since it is commonly used
   4. Created docker compose file to run container for database 
4. [Unit Test](https://github.com/canigetyourwhatwhat/study_go/tree/unit-test) (Basic)
   1. Tested each function 
   2. Completed tests through setting up database until closing the connection of Database
5. [Clean Architecture](https://github.com/canigetyourwhatwhat/study_go/tree/clean-architecture)
   1. Implemented dependency injection (DI) to the architecture.
   2. For better DI, also implemented interface.
   3. Separated to router, main, api (controller), service, and repository layers.
6. Error handling
7. Unit test (advanced)
8. Logging with middleware
9. Concurrency
10. Context package
11. User Authentication



# To start

```sh
docker compose up
```

There is a list of shell script in the script.sh file to try them out.