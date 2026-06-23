# sort_list 排序链表（LeetCode 148）
# 思路：归并排序。找中点拆成左右两段 → 递归排序 → merge_two_lists 合并
from typing import Optional


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def sort_list(self, head: Optional[ListNode]) -> Optional[ListNode]:
        if head is None or head.next is None:
            return head
        mid = self._get_mid(head)
        left = self.sort_list(head)
        right = self.sort_list(mid)
        return self._merge_two_lists(left, right)

    def _get_mid(self, head: ListNode) -> Optional[ListNode]:
        slow, fast = head, head.next
        while fast is not None and fast.next is not None:
            slow = slow.next
            fast = fast.next.next
        mid = slow.next
        slow.next = None
        return mid

    def _merge_two_lists(
        self, list1: Optional[ListNode], list2: Optional[ListNode]
    ) -> Optional[ListNode]:
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
    head = build_list([4, 2, 1, 3])
    got = list_to_slice(sol.sort_list(head))
    status = "PASS" if got == [1, 2, 3, 4] else "FAIL"
    print(f"{status} | sorted={got}")
