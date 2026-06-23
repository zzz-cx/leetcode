# is_palindrome 简单版（推荐面试先写）：链表值拷到切片，左右指针比较
from typing import Optional


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def is_palindrome(self, head: Optional[ListNode]) -> bool:
        vals = []
        cur = head
        while cur is not None:
            vals.append(cur.val)
            cur = cur.next
        i, j = 0, len(vals) - 1
        while i < j:
            if vals[i] != vals[j]:
                return False
            i += 1
            j -= 1
        return True

    def is_palindrome_o1(self, head: Optional[ListNode]) -> bool:
        if head is None:
            return True
        first_half_end = self._end_of_first_half(head)
        second_half_start = self._reverse_list(first_half_end.next)
        p1, p2 = head, second_half_start
        while p2 is not None:
            if p1.val != p2.val:
                return False
            p1, p2 = p1.next, p2.next
        first_half_end.next = self._reverse_list(second_half_start)
        return True

    def _end_of_first_half(self, head: ListNode) -> ListNode:
        slow, fast = head, head
        while fast.next is not None and fast.next.next is not None:
            slow = slow.next
            fast = fast.next.next
        return slow

    def _reverse_list(self, head: Optional[ListNode]) -> Optional[ListNode]:
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
    from testutil import build_list
    sol = Solution()
    head = build_list([1, 2, 2, 1])
    got = sol.is_palindrome(head)
    status = "PASS" if got is True else "FAIL"
    print(f"{status} | => {got}")
