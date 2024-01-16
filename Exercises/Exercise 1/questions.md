Exercise 1 - Theory questions
-----------------------------

### Concepts

What is the difference between *concurrency* and *parallelism*?

concurrency is about dealing with lots of things at once, multiple task can run in a overlapping manner. feks on a one core it can swictc between tasks fast giving the illution that the task is running at the same time. Parallisim is about DOING a lot of things at once. its about execution. Multiple tasks are done at once. 

In summary, concurrency is about the design of the system (how tasks are broken up and scheduled for execution), while parallelism is about the execution (actually running multiple tasks simultaneously).

What is the difference between a *race condition* and a *data race*? 
 A race condition is a situation where the behavior of a system depends on the relative timing or interleaving of multiple threads by the runtime scheduler. they lead to unpredictable results. 

 A data race is a specific type of race condition that occurs when two or more threads access the same memory location concurrently, at least one of the accesses is for writing, and the threads are not using any exlusive locks to control their accessto that memory. 

 In summary, all data races are race conditions, but not all race conditions are data races. Race conditions are a higher-level concept and can include other issues like races between file system operations, while data races are a lower-level concept specific to concurrent memory access.
 
*Very* roughly - what does a *scheduler* do, and how does it do it?

A scheduler in the context of computing is a component of the operating system that manages the execution of tasks. Its primary role is to switch the CPU among various tasks and ensure efficient use of CPU time.


### Engineering

Why would we use multiple threads? What kinds of problems do threads solve?
We use multiple threads for several reasons:

Utilize Multi-Core Processors: Modern computers have multi-core processors. By using multiple threads, a program can run tasks on different cores simultaneously, which can significantly improve the performance for CPU-bound tasks.

Improve Responsiveness: In a GUI application, a long-running task can be moved to a separate thread to prevent it from blocking the main thread that handles user input and updates the UI. This keeps the UI responsive even when the application is busy with other tasks.

Simplify Program Structure: In some cases, it's easier to design a program as a set of concurrent threads, each performing a specific task. This is often the case for servers that handle many independent client requests.

Efficient I/O Operations: Threads can be used to perform I/O operations in parallel. While one thread is blocked waiting for an I/O operation to complete, other threads can continue their work.

Some languages support "fibers" (sometimes called "green threads") or "coroutines"? What are they, and why would we rather use them over threads?
"Fibers" or "green threads" and "coroutines" are lightweight, user-space scheduling units of execution. They are similar to threads but with some key differences:

Fibers or Green Threads: These are scheduled by the runtime environment instead of the underlying operating system. They are lighter weight than OS threads and their context switch is often cheaper. However, they don't take advantage of multiple processors because they are not truly concurrent.

Coroutines: These are general control structures where flow control is cooperatively passed between two different routines without returning. Coroutines provide a way to maintain multiple points of entry and exit, and they can suspend and resume execution at certain points.

Why would we use them over threads?

Efficiency: Fibers/green threads and coroutines are generally more memory-efficient than threads because they use less resources. This allows you to have many more fibers or coroutines than you could have threads.

Simpler to use: In some cases, using coroutines can simplify your code. They allow you to write asynchronous code in a synchronous style, which can be easier to reason about.

Control over scheduling: Since fibers/green threads are scheduled by the runtime rather than the OS, you have more control over their scheduling. This can be useful in certain scenarios where you need fine-grained control over execution order.

Does creating concurrent programs make the programmer's life easier? Harder? Maybe both?
both.

What do you think is best - *shared variables* or *message passing*?
Shared Variables: In this strategy, concurrent processes or threads communicate with each other by reading from and writing to shared variables. This requires careful synchronization (using mechanisms like locks, semaphores, or monitors) to avoid race conditions and other concurrency issues. Shared variables are a common form of communication in multithreaded programming.

Message Passing: In this strategy, concurrent processes or threads communicate with each other by sending and receiving messages. These messages are typically sent through channels or queues. Message passing can avoid some of the complexities of shared variables, such as the need for explicit synchronization, but it can also introduce its own complexities, such as the need to handle communication failures. Message passing is the primary form of communication in many distributed and concurrent programming models, including actor model and CSP (Communicating Sequential Processes).


