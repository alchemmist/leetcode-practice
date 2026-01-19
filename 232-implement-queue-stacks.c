#include <stdio.h>
#include <stdbool.h>
#include <stdlib.h>

typedef struct {
    int *data;
    int length;
    int capacity;
} Stack;

typedef struct {
    Stack in;
    Stack out;
} MyQueue;

MyQueue *myQueueCreate() {
    MyQueue *m = calloc(1, sizeof(*m));
    return m;
}

void myQueuePush(MyQueue *obj, int x) {
    if (obj->in.capacity == 0) {
        obj->in.data = calloc(1, sizeof(*obj->in.data));
        obj->in.capacity = 1;
    } else if (obj->in.capacity - obj->in.length <= 1) {
        obj->in.capacity *= 2;
        obj->in.data =
            realloc(obj->in.data, obj->in.capacity * sizeof(*obj->in.data));
    }
    obj->in.data[obj->in.length] = x;
    obj->in.length++;
}

void _transferIn2Out(MyQueue *obj) {
    if (obj->out.capacity == 0) {
        obj->out.capacity = obj->in.length;
        obj->out.data = calloc(obj->in.length, sizeof(*obj->out.data));
    } else if (obj->out.capacity < obj->in.length) {
        obj->out.capacity = obj->in.length;
        obj->out.data =
            realloc(obj->out.data, obj->out.capacity * sizeof(*obj->out.data));
    }
    while (obj->in.length > 0) {
        obj->out.data[obj->out.length] = obj->in.data[--obj->in.length];
        obj->out.length++;
    }
    obj->in.capacity = 0;
}

int myQueuePop(MyQueue *obj) {
    if (obj->out.length <= 0) {
        _transferIn2Out(obj);
    }
    return obj->out.data[--obj->out.length];
}

int myQueuePeek(MyQueue *obj) {
    if (obj->out.length <= 0) {
        _transferIn2Out(obj);
    }
    return obj->out.data[obj->out.length - 1];
}

bool myQueueEmpty(MyQueue *obj) {
    return obj->in.length == 0 && obj->out.length == 0;
}

void myQueueFree(MyQueue *obj) {
    free(obj->out.data);
    free(obj->in.data);
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
