from typing import Optional


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def merge_two_lists(self, list1: Optional[ListNode], list2: Optional[ListNode]) -> Optional[ListNode]:
        dummy = ListNode(0)
        tail = dummy
        while list1 is not None and list2 is not None:
            if list1.val < list2.val:
                tail.next = list1
                list1 = list1.next
            else:
                tail.next = list2
                list2 = list2.next
            tail = tail.next
        if list1 is not None:
            tail.next = list1
        else:
            tail.next = list2
        return dummy.next


if __name__ == "__main__":
    import sys
    from pathlib import Path
    sys.path.insert(0, str(Path(__file__).resolve().parents[1] / "_helpers"))
    from testutil import build_list, list_to_slice
    sol = Solution()
    l1, l2 = build_list([1, 2, 4]), build_list([1, 3, 4])
    got = list_to_slice(sol.merge_two_lists(l1, l2))
    status = "PASS" if got == [1, 1, 2, 3, 4, 4] else "FAIL"
    print(f"{status} | merged={got}")
