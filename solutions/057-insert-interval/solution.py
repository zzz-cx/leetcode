from typing import List


class Solution:
    def insert(
        self, intervals: List[List[int]], new_interval: List[int]
    ) -> List[List[int]]:
        merged = [new_interval[0], new_interval[1]]
        result: list[list[int]] = []
        i, n = 0, len(intervals)

        while i < n and intervals[i][1] < merged[0]:
            result.append(intervals[i])
            i += 1

        while i < n and intervals[i][0] <= merged[1]:
            merged[0] = min(merged[0], intervals[i][0])
            merged[1] = max(merged[1], intervals[i][1])
            i += 1
        result.append(merged)

        while i < n:
            result.append(intervals[i])
            i += 1

        return result


if __name__ == "__main__":
    tests = [
        ([[1, 3], [6, 9]], [2, 5], [[1, 5], [6, 9]]),
        (
            [[1, 2], [3, 5], [6, 7], [8, 10], [12, 16]],
            [4, 8],
            [[1, 2], [3, 10], [12, 16]],
        ),
        ([], [5, 7], [[5, 7]]),
        ([[1, 5]], [2, 3], [[1, 5]]),
    ]
    sol = Solution()
    for intervals, new_interval, expected in tests:
        result = sol.insert(intervals, new_interval)
        status = "PASS" if result == expected else "FAIL"
        print(f"{status} | intervals={intervals!r}, new={new_interval!r} => {result}")
