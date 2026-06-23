from typing import Optional


class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution:
    def has_cycle(self, head: Optional[ListNode]) -> bool:
        slow, fast = head, head
        while fast is not None and fast.next is not None:
            slow = slow.next
            fast = fast.next.next
            if slow == fast:
                return True
        return False


if __name__ == "__main__":
    import sys
    from pathlib import Path
    sys.path.insert(0, str(Path(__file__).resolve().parents[1] / "_helpers"))
    from testutil import build_list
    sol = Solution()
    head = build_list([1, 2, 3])
    got = sol.has_cycle(head)
    status = "PASS" if got is False else "FAIL"
    print(f"{status} | cycle={got}")
