from typing import List


class Solution:
    def merge(self, intervals: List[List[int]]) -> List[List[int]]:
        if len(intervals) == 0:
            return intervals
        intervals.sort(key=lambda x: x[0])
        out = [intervals[0]]
        for i in range(1, len(intervals)):
            last = out[-1]
            cur = intervals[i]
            if last[1] < cur[0]:
                out.append(cur)
            else:
                if cur[1] > last[1]:
                    last[1] = cur[1]
        return out


if __name__ == "__main__":
    sol = Solution()
    got = sol.merge([[1, 3], [2, 6], [8, 10], [15, 18]])
    expected = [[1, 6], [8, 10], [15, 18]]
    status = "PASS" if got == expected else "FAIL"
    print(f"{status} | => {got}")
