class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        map = {}
        for i, n in enumerate(nums):
            if n in map:
                return map[n], i
            map[target - n] = i


print(Solution().twoSum([3, 2, 4], 6))
