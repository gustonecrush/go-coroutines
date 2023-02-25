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
        <td>Process is connected and isolated with other process</td>
        <td>Thread can connect each other if there is in one process</td>
    </tr>
    <tr>
        <td>Process is long to runned or stopped</td>
        <td>Thread is fast to runned or stopped</td>
    </tr>
</table>