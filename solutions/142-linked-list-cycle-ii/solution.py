from typing import Optional


class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution:
    def detect_cycle(self, head: Optional[ListNode]) -> Optional[ListNode]:
        fast, slow = head, head
        while fast is not None and fast.next is not None:
            slow = slow.next
            fast = fast.next.next
            if slow == fast:
                p = head
                while p != slow:
                    p = p.next
                    slow = slow.next
                return p
        return None


if __name__ == "__main__":
    import sys
    from pathlib import Path
    sys.path.insert(0, str(Path(__file__).resolve().parents[1] / "_helpers"))
    from testutil import build_list
    sol = Solution()
    head = build_list([1, 2, 3])
    got = sol.detect_cycle(head)
    status = "PASS" if got is None else "FAIL"
    print(f"{status} | entry={got}")
