#include <stdlib.h>

typedef struct {
    int *data;
    int capacity;
    int length;
} Stack;

typedef struct {
    Stack mins;
    Stack values;
} MinStack;

MinStack *minStackCreate() {
    MinStack *m = calloc(1, sizeof(*m));
    m->mins.data = calloc(1, sizeof(*m->mins.data));
    m->mins.capacity = 1;
    m->values.data = calloc(1, sizeof(*m->values.data));
    m->values.capacity = 1;
    return m;
}

void minStackPush(MinStack *obj, int val) {
    if (obj->mins.length == 0 || val <= obj->mins.data[obj->mins.length - 1]) {
        if (obj->mins.capacity - obj->mins.length <= 1) {
            obj->mins.capacity *= 2;
            obj->mins.data = realloc(obj->mins.data, sizeof(*obj->mins.data) *
                                                         obj->mins.capacity);
        }
        obj->mins.data[obj->mins.length++] = val;
    }
    if (obj->values.capacity - obj->values.length <= 1) {
        obj->values.capacity *= 2;
        obj->values.data = realloc(obj->values.data, sizeof(*obj->values.data) *
                                                         obj->values.capacity);
    }
    obj->values.data[obj->values.length++] = val;
}

void minStackPop(MinStack *obj) {
    if (obj->values.length <= 0) {
        return;
    }
    if (obj->values.data[obj->values.length - 1] ==
        obj->mins.data[obj->mins.length - 1]) {
        obj->mins.length--;
    }
    obj->values.length--;
}

int minStackTop(MinStack *obj) {
    if (obj->values.length == 0) {
        return -1;
    }
    return obj->values.data[obj->values.length - 1];
}

int minStackGetMin(MinStack *obj) {
    if (obj->mins.length <= 0) {
        return -1;
    }
    return obj->mins.data[obj->mins.length - 1];
}

void minStackFree(MinStack *obj) {
    free(obj->values.data);
    free(obj->mins.data);
    free(obj);
}
