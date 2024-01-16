Task 3: 

We get different resultes every time we run the code. This is because of race conditions because the threads running cuncurrent. This means that they are updating the variable simuntaniasly which gives unpredictabla results. 

Task 4: 

The "right" choice is to use mutex. This is because it is better to protect a single variable. Semaphores controls access to a shared resource through the use of a counter. In are case we only have one room, we only need one key to that room. If we had multiple rooms we would need multiple keys and then use semaphores. A semaphore allows multiple threads to access a shared resource, while a mutex only allows one thread to access a shared resource. 

in are case we are trying to protect a shared variable that can only be accessed by one rutine at a time. This is why the mutex is the appropirate synchronization primitive to use. 

Channels are used to send and receive values with the channel operator <-. Channels are used to synchronize and communictae between goroutines. They allow you to pass data between goroutines and make sure that operations happen in the order you expect. 

For Go using channels is confusing. I did not really understand how it workes and how. I think i got it to work, where i made three channels and inc sends to this channel and dec sends to it respective channle. then i had a done channel signaling when its done. 

Task 5: 

Initialized the semephore for the elements in the buffer and the capacity in the buffer. Then it was quite simple. Had to lock it for both pushing and poping functions. and wait in both scenarios. when sem_init is 0 the semaphore is shared between threads of the same process, if it was non zero it is shared between different processes. 

for Go this was veery easy. We just made a bounded buffer with a channel and added and popped from the buffer. 

The error message "fatal error: all goroutines are asleep - deadlock!" means that all goroutines are blocked and cannot proceed. This is considered a deadlock situation in Go.

In are case, the deadlock is caused by the select {} statement in the main function. This statement blocks the main goroutine indefinitely because there are no cases in the select statement. This is usually used when you want your program to keep running until it's manually stopped, because other goroutines are still doing work in the background.

However, in are case, the producer goroutine finishes its work and exits, and then the consumer goroutine consumes all items in the buffer and also exits. At this point, all goroutines are blocked: the main goroutine is blocked on the select {} statement, and there are no other goroutines left to do any work. This is why we getting a deadlock error.

