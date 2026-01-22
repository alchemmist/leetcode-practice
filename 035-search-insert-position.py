from typing import List


class Solution:
    def searchInsert(self, nums: List[int], target: int) -> int:
        l = 0
        r = len(nums) - 1
        while l <= r:
            i = l + (r - l) // 2
            if nums[i] == target:
                return i
            if nums[i] < target:
                l = i + 1
            else:
                r = i - 1
        return l


print(Solution().searchInsert([1, 3, 5, 7], 9))
