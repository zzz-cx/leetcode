from typing import List


class Solution:
    def two_sum(self, numbers: List[int], target: int) -> List[int]:
        left, right = 0, len(numbers) - 1
        while left < right:
            total = numbers[left] + numbers[right]
            if total == target:
                return [left + 1, right + 1]
            if total < target:
                left += 1
            else:
                right -= 1
        return []


if __name__ == "__main__":
    tests = [
        ([2, 7, 11, 15], 9, [1, 2]),
        ([2, 3, 4], 6, [1, 3]),
        ([-1, 0], -1, [1, 2]),
    ]
    sol = Solution()
    for numbers, target, expected in tests:
        result = sol.two_sum(numbers, target)
        status = "PASS" if result == expected else "FAIL"
        print(f"{status} | numbers={numbers}, target={target} => {result} (expected {expected})")
