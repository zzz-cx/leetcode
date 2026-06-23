from typing import Optional


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def reverse_list(self, head: Optional[ListNode]) -> Optional[ListNode]:
        if head is None or head.next is None:
            return head
        prev = None
        current = head
        while current is not None:
            nxt = current.next
            current.next = prev
            prev = current
            current = nxt
        return prev


if __name__ == "__main__":
    import sys
    from pathlib import Path
    sys.path.insert(0, str(Path(__file__).resolve().parents[1] / "_helpers"))
    from testutil import build_list, list_to_slice
    sol = Solution()
    head = build_list([1, 2, 3, 4, 5])
    got = list_to_slice(sol.reverse_list(head))
    status = "PASS" if got == [5, 4, 3, 2, 1] else "FAIL"
    print(f"{status} | reversed={got}")
