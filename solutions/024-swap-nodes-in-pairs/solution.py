from typing import Optional


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def swap_pairs(self, head: Optional[ListNode]) -> Optional[ListNode]:
        dummy = ListNode(0, head)
        prev = dummy
        current = head
        while current is not None and current.next is not None:
            nxt = current.next
            current.next = nxt.next
            nxt.next = current
            prev.next = nxt
            prev = current
            current = current.next
        return dummy.next


if __name__ == "__main__":
    import sys
    from pathlib import Path
    sys.path.insert(0, str(Path(__file__).resolve().parents[1] / "_helpers"))
    from testutil import build_list, list_to_slice
    sol = Solution()
    head = build_list([1, 2, 3, 4])
    got = list_to_slice(sol.swap_pairs(head))
    status = "PASS" if got == [2, 1, 4, 3] else "FAIL"
    print(f"{status} | => {got} (expected [2, 1, 4, 3])")
