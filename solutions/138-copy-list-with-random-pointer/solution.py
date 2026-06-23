from typing import Optional


class Node:
    def __init__(self, x: int, next=None, random=None):
        self.val = x
        self.next = next
        self.random = random


class Solution:
    def copy_random_list(self, head: Optional[Node]) -> Optional[Node]:
        if head is None:
            return None
        m = {}
        cur = head
        while cur is not None:
            m[cur] = Node(cur.val)
            cur = cur.next
        cur = head
        while cur is not None:
            m[cur].next = m.get(cur.next)
            m[cur].random = m.get(cur.random)
            cur = cur.next
        return m[head]


if __name__ == "__main__":
    sol = Solution()
    head = Node(1, Node(2, Node(3)))
    got = sol.copy_random_list(head)
    vals = []
    cur = got
    while cur:
        vals.append(cur.val)
        cur = cur.next
    status = "PASS" if vals == [1, 2, 3] else "FAIL"
    print(f"{status} | copied={vals}")
