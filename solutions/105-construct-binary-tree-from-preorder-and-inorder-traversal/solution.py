from typing import List, Optional


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def build_tree(self, preorder: List[int], inorder: List[int]) -> Optional[TreeNode]:
        if len(preorder) == 0 or len(inorder) == 0:
            return None
        root = TreeNode(preorder[0])
        root_index = 0
        for i in range(len(inorder)):
            if inorder[i] == root.val:
                root_index = i
        root.left = self.build_tree(preorder[1 : root_index + 1], inorder[:root_index])
        root.right = self.build_tree(preorder[root_index + 1 :], inorder[root_index + 1 :])
        return root

    def build_tree2(self, preorder: List[int], inorder: List[int]) -> Optional[TreeNode]:
        if len(preorder) == 0:
            return None

        pos = {v: i for i, v in enumerate(inorder)}
        pre_idx = 0

        def helper(l: int, r: int) -> Optional[TreeNode]:
            nonlocal pre_idx
            if l > r:
                return None
            val = preorder[pre_idx]
            pre_idx += 1
            root = TreeNode(val)
            mid = pos[val]
            root.left = helper(l, mid - 1)
            root.right = helper(mid + 1, r)
            return root

        return helper(0, len(inorder) - 1)


if __name__ == "__main__":
    import sys
    from pathlib import Path
    sys.path.insert(0, str(Path(__file__).resolve().parents[1] / "_helpers"))
    from testutil import build_tree

    sol = Solution()
    root = sol.build_tree([3, 9, 20, 15, 7], [9, 3, 15, 20, 7])
    status = "PASS" if root and root.val == 3 else "FAIL"
    print(f"{status} | root.val={root.val if root else None}")
