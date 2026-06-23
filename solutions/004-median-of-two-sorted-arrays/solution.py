# 思路：在较短数组上二分切分，使左半共 (m+n+1)/2 个且 max(左) <= min(右)
from typing import List


class Solution:
    def find_median_sorted_arrays(self, nums1: List[int], nums2: List[int]) -> float:
        if len(nums1) > len(nums2):
            nums1, nums2 = nums2, nums1
        m, n = len(nums1), len(nums2)
        left, right = 0, m
        half = (m + n + 1) // 2

        while left <= right:
            i = (left + right) // 2
            j = half - i

            if i < m and j > 0 and nums1[i] < nums2[j - 1]:
                left = i + 1
            elif i > 0 and j < n and nums2[j] < nums1[i - 1]:
                right = i - 1
            else:
                if i == 0:
                    max_left = nums2[j - 1]
                elif j == 0:
                    max_left = nums1[i - 1]
                else:
                    max_left = max(nums1[i - 1], nums2[j - 1])
                if (m + n) % 2 == 1:
                    return float(max_left)

                if i == m:
                    min_right = nums2[j]
                elif j == n:
                    min_right = nums1[i]
                else:
                    min_right = min(nums1[i], nums2[j])
                return (max_left + min_right) / 2
        return 0.0


if __name__ == "__main__":
    sol = Solution()
    got = sol.find_median_sorted_arrays([1, 3], [2])
    status = "PASS" if abs(got - 2.0) < 1e-9 else "FAIL"
    print(f"{status} | median={got} (expected 2.0)")
