#include <stdio.h>

int findGCDTwo(int a, int b) {
    while (b > 0) {
        int tmp = a % b;
        a = b;
        b = tmp;
    }
    return a;
}

int findGCD(int *nums, int numsSize) {
    int min = nums[0];
    int max = nums[0];
    for (int i = 1; i < numsSize; i++) {
        if (nums[i] <= min) {
            min = nums[i];
        }
        if (nums[i] >= max) {
            max = nums[i];
        }
    }
    return findGCDTwo(min, max);
}

int main() {
    int n[3] = {10, 6, 9};
    printf("%d\n", findGCD(n, 3));
}
