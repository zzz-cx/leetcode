from typing import List


class Solution:
    def can_complete_circuit(self, gas: List[int], cost: List[int]) -> int:
        total_tank = 0
        current_tank = 0
        start = 0

        for i in range(len(gas)):
            diff = gas[i] - cost[i]
            total_tank += diff
            current_tank += diff

            if current_tank < 0:
                start = i + 1
                current_tank = 0

        return start if total_tank >= 0 else -1


if __name__ == "__main__":
    tests = [
        ([1, 2, 3, 4, 5], [3, 4, 5, 1, 2], 3),
        ([2, 3, 4], [3, 4, 3], -1),
        ([5, 1, 2, 3, 4], [4, 4, 1, 5, 1], 4),
        ([3, 1, 1], [1, 2, 2], 0),
    ]
    sol = Solution()
    for gas, cost, expected in tests:
        result = sol.can_complete_circuit(gas, cost)
        status = "PASS" if result == expected else "FAIL"
        print(f"{status} | gas={gas}, cost={cost} => {result} (expected {expected})")
