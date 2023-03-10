# Go Coroutines

## Agenda
- Concurrency & Parallel Programming
- Goroutines
- Channel
- Buffered Channel
- Mutex
- WaitGroup
- Atomic
- Ticker
- etc

## Introduction Concurrency & Parallel Programming

### Introduction to Parallel Programming
- We live in era multicore era, where we use rarely single core processor
- More advanced the hardware, so software will follow, where now we can create easily parallel process in application
- Parallel programming simply is solving a different or same problem with breakdown to a small of shit and running them together in same time

### Real Case Parallel
- Running some apps in one time on our operating system (office, editor, browser, etc)
- Some chefs serving food in restaurant, where chef cook different food
- Queue at the bank, where each teller serves each customer

### Process vs Thread
<table>
    <tr>
        <th>Process</th>
        <th>Thread</th>
    </tr>
    <tr>
        <td>Process is an execution of program</td>
        <td>Thread is a segment of process</td>
    </tr>
    <tr>
        <td>Process consume big memory</td>
        <td>Thread consume small memory</td>
    </tr>
    <tr>
        <td>Process is isolated each other</td>
        <td>Thread can connect each other if there is in one process</td>
    </tr>
    <tr>
        <td>Process is long to runned or stopped</td>
        <td>Thread is fast to runned or stopped</td>
    </tr>
</table>

### Parallel vs Concurrency
- Different with parallel (running some tasks simultaneously), concurrency is running some tasks alternately
- In parallel, we usually need a lot of Thread whereas in concurrency, we only need few Thread  
- Go-lang is default as concurrency, but we are running at multicore so it is mix of concurrency and parallel

### Real Case Concurrency
- When we are eating in caffe, we can eat, then talk, we drink, eat again, talk again, drink again, etc. But we can not do it simultaneously drink, eat, and talk, we only can do one thing in a time, but we can do it alternately whenever we want

### CPU-Bound
- Many algorithms created only need CPU to run it. This type of algorithm usually very dependent on CPU speed
- The most popular example is Machine Learning. Therefore, there is so many or them using GPU because has more cores than CPU normally
- This type of algorithm there is no benefit when using Concurrency Programming, but can help with implementation of Parallel Programming

### I/O Bound
- I/O bound is the opposite of before, where algorithm or application usually dependent to speed of input output device we used
- Example of this application we can read data from file, database, etc
- Most of the time, we usually create application like this
- Type of application with I/O bound, although usually helped with implementation Parallel Programming but has more benefits if we use Concurrency Programming
- Imagine we read data from database, and Thread has to wait 1 second to get data from database, whereas with 1 second if we use Concurrency Programming, bwe can use it to do another task

## Introduction to Goroutines

### Introduction to Goroutines
- Goroutine is a light thread which managed by Go Runtime
- Size of Goroutine is so small, around 2kb, smaller than Thread that cn arrive to 1mb or 1000kb
- But not like thread run parallely, goroutine run concurrently

### How Goroutine Works
- Actually, Goroutine is run by Go Scheduler in thread, where number or thread as much GOMAXPROCS (as much core CPU usually)
- So actually can't say Goroutine is replacement of Thread, because Goroutine itself run on Thread
- But what makes it easier for us is we don't need to manage Thread manually, all is set up by Go Scheduler
- There is some terms that you need to know :
  - G: Goroutine
  - M: Thread(Machine)
  - P: Processor

## Create Goroutine

### Create Goroutine
- To create goroutine in Golang is very simple
- We only need to add command go before we call function we run in goroutine
- When a function we run in goroutine, function will run asynchronously, means not will be waited until that function is done
- Application will continue running to next code program without waiting goroutine we created is done

## Introduction to Channel

### Introduction to Channel
- Channel is a place to communicate synchronously can do in goroutine, used for return value if we use goroutine
- In Channel there is sender and receiver, sender and receiver usually is different goroutine
- When sending data to channel, goroutine will be blocked until there is receiver data
- Therefore, channel called as communication tool synchronous (blocking)
- Channel is very suit as alternative as async await mechanism which in other programming language

### Characteristic Channel
- Channel only store one data as default, if we want to add more data we have to wait data in channel is taken by other goroutine
- Channel only can receive one type of data
- Channel can be taken more than one goroutine
- Channel must be closed if we don't use it anymore, or can cause memory leak

## Create Channel

### Create Channel
- Channel in Go-Lang represent with data type chan
- To create channel we can use make(), like map
- Bu we have to determine data type

### Code: Create Channel
```go
channel := make(chan string)
```

### Send and Receive Data from Channel
- To send data, we can use code: channel <- data
- Whereas to receive data, we can use code: data <- channel
- If finished, don't forget to close channel using function close()