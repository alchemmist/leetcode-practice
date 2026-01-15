#include <stdio.h>
#include <stdbool.h>

typedef struct Node {
    int value;
    struct Node *next;
} Node;

typedef struct {
    int length;
    Node *top;
} MyQueue;

MyQueue *myQueueCreate() {
    MyQueue *m;
    m->length = 0;
    return m;
}

void myQueuePush(MyQueue *obj, int x) {
    if (obj->top == NULL) {}
}

int myQueuePop(MyQueue *obj) {
}

int myQueuePeek(MyQueue *obj) {
    return obj->top->value;
}

bool myQueueEmpty(MyQueue *obj) {
}

void myQueueFree(MyQueue *obj) {
}

/**
 * Your MyQueue struct will be instantiated and called as such:
 * MyQueue* obj = myQueueCreate();
 * myQueuePush(obj, x);

 * int param_2 = myQueuePop(obj);

 * int param_3 = myQueuePeek(obj);

 * bool param_4 = myQueueEmpty(obj);

 * myQueueFree(obj);
*/
