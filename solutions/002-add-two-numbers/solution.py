from typing import Optional


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def add_two_numbers(self, l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:
        cnt = 0
        dummy = ListNode()
        tail = dummy
        while l1 is not None or l2 is not None:
            n1, n2 = 0, 0
            if l1 is not None:
                n1 = l1.val
                l1 = l1.next
            if l2 is not None:
                n2 = l2.val
                l2 = l2.next
            s = n1 + n2 + cnt
            cnt = s // 10
            s = s % 10
            tail.next = ListNode(s)
            tail = tail.next
        if cnt > 0:
            tail.next = ListNode(cnt)
        return dummy.next


if __name__ == "__main__":
    import sys
    from pathlib import Path
    sys.path.insert(0, str(Path(__file__).resolve().parents[1] / "_helpers"))
    from testutil import build_list, list_to_slice
    sol = Solution()
    tests = [([2, 4, 3], [5, 6, 4], [7, 0, 8]), ([0], [0], [0])]
    for a, b, expected in tests:
        l1, l2 = build_list(a), build_list(b)
        got = list_to_slice(sol.add_two_numbers(l1, l2))
        status = "PASS" if got == expected else "FAIL"
        print(f"{status} | {a}+{b} => {got} (expected {expected})")
