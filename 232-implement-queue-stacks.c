#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

typedef struct Node {
    int value;
    struct Node *next;
} Node;

typedef struct {
    int length;
    int capacity;
    int *stack;
} MyQueue;

MyQueue *myQueueCreate() {
    MyQueue *m = calloc(1, sizeof(*m));
    m->length = 0;
    m->capacity = 1;
    m->stack = calloc(m->capacity, sizeof(int));
    return m;
}

void myQueuePush(MyQueue *obj, int x) {
    if (obj->capacity - obj->length <= 1) {
        obj->capacity *= 2;
        obj->stack = realloc(obj->stack, obj->capacity * sizeof(*obj->stack));
    }
    obj->stack[obj->length] = x;
    obj->length++;
}

int myQueuePop(MyQueue *obj) {
    if (obj->length <= 0) {
        return -1;
    }

    int value = obj->stack[0];
    for (int i = 1; i < obj->length; i++) {
        obj->stack[i - 1] = obj->stack[i];
    }
    obj->length--;

    return value;
}

int myQueuePeek(MyQueue *obj) {
    if (obj->length <= 0) {
        return -1;
    }
    return obj->stack[0];
}

bool myQueueEmpty(MyQueue *obj) {
    return (obj->length == 0);
}

void myQueueFree(MyQueue *obj) {
    free(obj->stack);
    free(obj);
}

int main() {
    MyQueue *m = myQueueCreate();
    myQueuePush(m, 1);
    myQueuePush(m, 2);
    printf("%d\n", myQueuePeek(m));
    printf("%d\n", myQueuePop(m));
    printf("%b\n", myQueueEmpty(m));
}
