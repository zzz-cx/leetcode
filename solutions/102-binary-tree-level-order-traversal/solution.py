from typing import List, Optional


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def _helper(self, node: Optional[TreeNode], level: int, ret: List[List[int]]) -> List[List[int]]:
        if node is None:
            return ret
        if len(ret) == level:
            ret.append([])
        ret[level].append(node.val)
        ret = self._helper(node.left, level + 1, ret)
        ret = self._helper(node.right, level + 1, ret)
        return ret

    def level_order(self, root: Optional[TreeNode]) -> List[List[int]]:
        ret: List[List[int]] = []
        if root is None:
            return ret
        return self._helper(root, 0, ret)


if __name__ == "__main__":
    import sys
    from pathlib import Path
    sys.path.insert(0, str(Path(__file__).resolve().parents[1] / "_helpers"))
    from testutil import build_tree

    sol = Solution()
    root = build_tree([1, 2, 3, 4, 5, 6, 7])
    got = sol.level_order(root)
    status = "PASS" if len(got) == 3 else "FAIL"
    print(f"{status} | levels={len(got)}")
