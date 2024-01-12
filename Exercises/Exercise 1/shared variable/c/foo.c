// Compile with `gcc foo.c -Wall -std=gnu99 -lpthread`, or use the makefile
// The executable will be named `foo` if you use the makefile, or `a.out` if you use gcc directly

#include <pthread.h>
#include <stdio.h>

// global shared int i = 0
volatile int i = 0;
pthread_mutex_t lock;


// Note the return type: void*
void* incrementingThreadFunction(){
    // TODO: increment i 1_000_000 times
    for (int j = 0; j <= 1000000; j++){
        pthread_mutex_lock(&lock);
        ++i;
        pthread_mutex_unlock(&lock);
    }
    return NULL;
}

void* decrementingThreadFunction(){
    // TODO: decrement i 1_000_000 times
    for (int j = 0; j <= 1000001; j++){
        pthread_mutex_lock(&lock);
        --i;
        pthread_mutex_unlock(&lock);
    }
    return NULL;
}


int main(){
    // TODO: 
    // start the two functions as their own threads using `pthread_create`
    // Hint: search the web! Maybe try "pthread_create example"?

    pthread_t thread1, thread2; // Variables to hold thread identifiers
    pthread_mutex_init(&lock, NULL);

    // Create the first thread
    pthread_create(&thread1, NULL, incrementingThreadFunction, NULL);
    pthread_create(&thread2, NULL, decrementingThreadFunction, NULL);

    pthread_join(thread1, NULL);
    pthread_join(thread2, NULL);

    pthread_mutex_destroy(&lock); // destroyes mutes object and frees resources

    printf("The magic number is: %d\n", i);
    return 0;

    
}
