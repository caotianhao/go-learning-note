# Golang study notes during three years of graduate study

## 1. Basic syntax and data structure
- Learn the basic syntax of Go, such as variable declarations,
control structures (if, for, switch), function definitions and calls.
Master common data structures: arrays, slices, dictionaries (maps),
structs (structs), and linked lists (using the 'container/list' package).

## 2. Concurrent programming
- Learn Goroutine and Channel for concurrent programming.
- Implement concurrency security using 'Mutex' and 'WaitGroup' in the 'sync' package.
- Understand concurrency patterns in Go, such as producer-consumer patterns
and workpool patterns.

## 3. Network programming and Web framework
- Learn to use the 'net/http' package for HTTP server and client development.
- Research the use of mainstream Web frameworks such as Gin, including routing, 
request processing, middleware, etc.
- Practice developing RESTful apis and learn how to work with JSON data.

## 4. Database operation
- Learn how to use GORM and 'sqlx' to manipulate MySQL databases and learn how 
to manage database connection pools.
- Study cache design and implementation using Redis.

## 5. Message queuing and asynchronous processing
- Use the Sarama library to produce and consume Kafka message queues.
- Learn how to handle duplicate messages and ensure data consistency.

## 6. Project practice
- Implementation of order system, including order creation and query interface, 
using HTTP service, Kafka message queue, MySQL and Redis.
- Implement instant kill systems, solve oversold problems, and support distributed deployment.

## 7. Test and performance optimization
- Perform stress tests using the 'go-wrk' tool to analyze system performance bottlenecks.
- Optimize database queries, Redis cache usage, and Goroutine and Channel usage.

## 8. Code style and best practices
- Follow the Go code style and best practices, such as avoiding global shared 
state and using 'context' to pass context information.
  Learn how to write clear, concise, and maintainable Go code.

## 9. Tools and environment
- Familiar with Go Modules management dependencies, using the 'go mod' command.
- Understand the common development tool 'GoLand'.

## 10. Continuous learning and community involvement
- Follow the official Go blog and community updates to learn about new features
and future directions of the Go language.
- Participate in open source projects, submit Pull requests and issues, contribute
code and improve myself.

## 11. 내가 왜 Go를 선택했는지
- ㅋㅋ, 이것은 운명이 그렇게 만든 것입니다. 나는 C와 Java를 좋아하지 않아서 그런지 
선배의 프로젝트 요청으로 암호 알고리즘을 많이 배웠고 Go를 사용하여 구현했습니다.
저는 Go가 매우 좋은 언어라고 생각하고, 매우 좋게 보고 있습니다. Delphi? 본 것 같아요!
