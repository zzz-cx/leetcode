from typing import Optional


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def remove_nth_from_end(self, head: Optional[ListNode], n: int) -> Optional[ListNode]:
        dummy = ListNode(0, head)
        first = dummy
        second = dummy
        for _ in range(n + 1):
            first = first.next
        while first is not None:
            first = first.next
            second = second.next
        second.next = second.next.next
        return dummy.next


if __name__ == "__main__":
    import sys
    from pathlib import Path
    sys.path.insert(0, str(Path(__file__).resolve().parents[1] / "_helpers"))
    from testutil import build_list, list_to_slice
    sol = Solution()
    head = build_list([1, 2, 3, 4, 5])
    got = list_to_slice(sol.remove_nth_from_end(head, 2))
    status = "PASS" if got == [1, 2, 3, 5] else "FAIL"
    print(f"{status} | => {got} (expected [1, 2, 3, 5])")
