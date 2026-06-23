# 思路 哈希集合，如果节点不在哈希集合中，继续遍历到下一个节点；如果在，则后面的节点都在哈希集合中
from typing import Optional


class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution:
    def get_intersection_node(
        self, head_a: Optional[ListNode], head_b: Optional[ListNode]
    ) -> Optional[ListNode]:
        vis = set()
        node = head_a
        while node is not None:
            vis.add(node)
            node = node.next
        node = head_b
        while node is not None:
            if node in vis:
                return node
            node = node.next
        return None


if __name__ == "__main__":
    sol = Solution()
    shared = ListNode(8)
    shared.next = ListNode(4)
    shared.next.next = ListNode(5)
    a = ListNode(4)
    a.next = ListNode(1)
    a.next.next = shared
    b = ListNode(5)
    b.next = ListNode(6)
    b.next.next = ListNode(1)
    b.next.next.next = shared
    got = sol.get_intersection_node(a, b)
    status = "PASS" if got and got.val == 8 else "FAIL"
    print(f"{status} | intersection={got.val if got else None}")
