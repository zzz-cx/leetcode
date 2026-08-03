from typing import List


class Solution:
    def find_min_arrow_shots(self, points: List[List[int]]) -> int:
        points.sort(key=lambda p: p[1])
        arrows = 1
        arrow_pos = points[0][1]

        for start, end in points[1:]:
            if start > arrow_pos:
                arrows += 1
                arrow_pos = end

        return arrows


if __name__ == "__main__":
    tests = [
        ([[10, 16], [2, 8], [1, 6], [7, 12]], 2),
        ([[1, 2], [3, 4], [5, 6], [7, 8]], 4),
        ([[1, 2], [2, 3], [3, 4], [4, 5]], 2),
    ]
    sol = Solution()
    for points, expected in tests:
        result = sol.find_min_arrow_shots([p[:] for p in points])
        status = "PASS" if result == expected else "FAIL"
        print(f"{status} | points={points!r} => {result} (expected {expected})")
