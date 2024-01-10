## Threads in C and Go

The magic number is not always zero since the 2 threads are simultaneously working on with the same variable. This causes the number to change every time we run the code.

When setting the runtime.GOMAXPROCS(1) to 1, only 1 thread can run at a time. The functions will run after one another.