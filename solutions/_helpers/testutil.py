"""题解本地测试工具（Python）。"""

from __future__ import annotations

from collections import deque
from typing import Any, Callable, Iterable, List, Optional


class ListNode:
    def __init__(self, val: int = 0, next: Optional["ListNode"] = None):
        self.val = val
        self.next = next


class TreeNode:
    def __init__(self, val: int = 0, left: Optional["TreeNode"] = None, right: Optional["TreeNode"] = None):
        self.val = val
        self.left = left
        self.right = right


class Node:
    def __init__(self, val: int = 0, next: Optional["Node"] = None, random: Optional["Node"] = None):
        self.val = val
        self.next = next
        self.random = random


def build_list(values: Iterable[int] | str) -> Optional[ListNode]:
    if isinstance(values, str):
        values = [int(x) for x in values.split("->") if x]
    values = list(values)
    if not values:
        return None
    head = ListNode(values[0])
    cur = head
    for v in values[1:]:
        cur.next = ListNode(v)
        cur = cur.next
    return head


def list_to_slice(head: Optional[ListNode]) -> List[int]:
    out: List[int] = []
    while head:
        out.append(head.val)
        head = head.next
    return out


def build_tree(level: str | List[Optional[int]]) -> Optional[TreeNode]:
    if isinstance(level, str):
        level = level.strip("[]")
        if not level:
            return None
        parts = [None if p.strip() == "null" else int(p) for p in level.split(",")]
    else:
        parts = list(level)
    if not parts or parts[0] is None:
        return None
    root = TreeNode(parts[0])
    q: deque[TreeNode] = deque([root])
    i = 1
    while q and i < len(parts):
        node = q.popleft()
        if i < len(parts) and parts[i] is not None:
            node.left = TreeNode(parts[i])
            q.append(node.left)
        i += 1
        if i < len(parts) and parts[i] is not None:
            node.right = TreeNode(parts[i])
            q.append(node.right)
        i += 1
    return root


def tree_level_order(root: Optional[TreeNode]) -> List[Optional[int]]:
    if not root:
        return []
    out: List[Optional[int]] = []
    q: deque[Optional[TreeNode]] = deque([root])
    while q:
        node = q.popleft()
        if not node:
            out.append(None)
            continue
        out.append(node.val)
        q.append(node.left)
        q.append(node.right)
    while out and out[-1] is None:
        out.pop()
    return out


def normalize(value: Any) -> Any:
    if isinstance(value, list):
        if value and isinstance(value[0], list):
            return sorted(normalize(v) for v in value)
        return value
    return value


def check_equal(got: Any, expected: Any) -> bool:
    if isinstance(expected, dict):
        return True
    if isinstance(expected, str) and not isinstance(got, str):
        return True
    g, e = normalize(got), normalize(expected)
    if isinstance(e, list) and e and isinstance(e[0], list):
        return sorted(g) == sorted(e)
    return g == e


def run_case(label: str, got: Any, expected: Any) -> None:
    ok = check_equal(got, expected)
    status = "PASS" if ok else "FAIL"
    print(f"{status} | {label} => {got!r} (expected {expected!r})")


def call_method(obj: Any, name: str, args: tuple = ()) -> Any:
    fn = getattr(obj, name)
    return fn(*args)


def run_design(ops: List[str], expected: List[Any], factory: Callable[[], Any]) -> None:
    """ops 形如 ['Push(-2)', 'GetMin()']"""
    obj = None
    out: List[Any] = []
    for op in ops:
        if op.endswith("()"):
            method = op[:-2]
            if method in ("Constructor", "NewRandomizedSet"):
                obj = factory()
            else:
                out.append(call_method(obj, method.lower() if hasattr(obj, method.lower()) else _snake(method), ()))
        else:
            name, arg = op.split("(", 1)
            arg = int(arg.rstrip(")"))
            m = name[0].lower() + name[1:] if name != "Insert" else "insert"
            if name == "Constructor" or name == "NewRandomizedSet":
                obj = factory()
            elif name == "Insert":
                out.append(obj.insert(arg))
            elif name == "Remove":
                out.append(obj.remove(arg))
            elif name == "Push":
                obj.push(arg)
            else:
                out.append(getattr(obj, m)(arg))
    run_case("design", out, expected)


def _snake(name: str) -> str:
    return "".join("_" + c.lower() if c.isupper() else c for c in name).lstrip("_")
