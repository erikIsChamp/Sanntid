// Compile with `gcc foo.c -Wall -std=gnu99 -lpthread`, or use the makefile
// The executable will be named `foo` if you use the makefile, or `a.out` if you use gcc directly

#include <pthread.h>
#include <stdio.h>
#include <semaphore.h>

// global shared int i = 0
volatile int i = 0;


// Declare a mutex
pthread_mutex_t mutex = PTHREAD_MUTEX_INITIALIZER;


// Note the return type: void*
void* incrementingThreadFunction(void *arg){

    // Lock the mutex to gain exclusive access to the shared resource
    pthread_mutex_lock(&mutex);

    // TODO: increment i 1_000_000 times
    for (int j = 0; j <= 1000000; j++){
        i++;
    }

    // Unlock the mutex to release the resource
    pthread_mutex_unlock(&mutex);

    return NULL;
}

void* decrementingThreadFunction(void *arg){

    // Lock the mutex to gain exclusive access to the shared resource
    pthread_mutex_lock(&mutex);

    // TODO: decrement i 1_000_000 times
    for (int j = 0; j <= 1000001; j++){
        i--;
    }

    // Unlock the mutex to release the resource
    pthread_mutex_unlock(&mutex);

    return NULL;
}


int main(){
    // TODO: 
    // start the two functions as their own threads using `pthread_create`
    // Hint: search the web! Maybe try "pthread_create example"?

    pthread_t thread1, thread2; // Variables to hold thread identifiers

    // Create the first thread
    if (pthread_create(&thread1, NULL, incrementingThreadFunction, NULL) != 0) {
        perror("pthread_create");
        return 1;
    }

    // Create the second thread
    if (pthread_create(&thread2, NULL, decrementingThreadFunction, NULL) != 0) {
        perror("pthread_create");
        return 1;
    }

    // TODO:
    // wait for the two threads to be done before printing the final result
    // Hint: Use `pthread_join`   

        // Wait for both threads to finish
    if (pthread_join(thread1, NULL) != 0) {
        perror("pthread_join");
        return 1;
    }

    if (pthread_join(thread2, NULL) != 0) {
        perror("pthread_join");
        return 1;
    }

    // Destroy the mutex (not necessary in this simple example)
    pthread_mutex_destroy(&mutex);

    printf("Both threads have finished\n");
    
    
    printf("The magic number is: %d\n", i);
    return 0;
}
