#include <stdio.h>

int findKthPositive(int *arr, int arrSize, int k) {
    int missed = 0;
    int last = 0;
    for (int i = 0; i < arrSize; i++) {
        int gap = arr[i] - last - 1;
        if (missed + gap >= k) {
            return last + (k - missed);
        }
        missed += gap;
        last = arr[i];
    }
    return arr[arrSize - 1] + (k - missed);
}

int main() {
    int nums[] = {1, 3};
    printf("%d\n", findKthPositive(nums, sizeof(nums) / sizeof(nums[0]), 3));
}

