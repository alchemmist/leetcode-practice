#include <stdio.h>

int findKthPositive(int *arr, int arrSize, int k) {
    int left = 0;
    int right = arrSize - 1;
    while (left <= right) {
        int i = (left + right) / 2;
        if (arr[i] - (i + 1) >= k) {
            right = i - 1;
        } else {
            left = i + 1;
        }
    }
    return left + k;
}

int main() {
    int nums[] = {2};
    printf("%d\n", findKthPositive(nums, sizeof(nums) / sizeof(nums[0]), 1));
}
