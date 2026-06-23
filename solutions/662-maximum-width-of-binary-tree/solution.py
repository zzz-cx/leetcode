# width_of_binary_tree 二叉树最大宽度（LeetCode 662）
# 思路：BFS 按层遍历。每一层用队首、队尾节点的 idx 算宽度 = right - left + 1，再取全局最大。
from collections import deque
from typing import Optional


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def width_of_binary_tree(self, root: Optional[TreeNode]) -> int:
        if root is None:
            return 0

        queue = deque([(root, 1)])
        ans = 0

        while queue:
            size = len(queue)
            left = queue[0][1]
            right = queue[size - 1][1]
            ans = max(ans, right - left + 1)

            for _ in range(size):
                node, idx = queue.popleft()
                if node.left is not None:
                    queue.append((node.left, idx * 2))
                if node.right is not None:
                    queue.append((node.right, idx * 2 + 1))

        return ans


if __name__ == "__main__":
    import sys
    from pathlib import Path
    sys.path.insert(0, str(Path(__file__).resolve().parents[1] / "_helpers"))
    from testutil import build_tree

    sol = Solution()
    root = build_tree([1, 2, 3, 4, 5, 6, 7])
    got = sol.width_of_binary_tree(root)
    status = "PASS" if got == 4 else "FAIL"
    print(f"{status} | width={got}")
