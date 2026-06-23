from typing import List


class Solution:
    def merge2(self, nums1: List[int], m: int, nums2: List[int], n: int) -> List[int]:
        i = m - 1
        j = n - 1
        k = m + n - 1
        while i >= 0 and j >= 0:
            if nums1[i] > nums2[j]:
                nums1[k] = nums1[i]
                i -= 1
            else:
                nums1[k] = nums2[j]
                j -= 1
            k -= 1
        while j >= 0:
            nums1[k] = nums2[j]
            j -= 1
            k -= 1
        return nums1


if __name__ == "__main__":
    sol = Solution()
    nums1, nums2 = [1, 2, 3, 0, 0, 0], [2, 5, 6]
    sol.merge2(nums1, 3, nums2, 3)
    status = "PASS" if nums1 == [1, 2, 2, 3, 5, 6] else "FAIL"
    print(f"{status} | merged={nums1}")
