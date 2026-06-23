#!/usr/bin/env python3
"""运行 LeetCode 题解。

用法:
  python scripts/run.py 134              # 按题号
  python scripts/run.py 001-two-sum      # 按目录名
  python scripts/run.py two-sum          # 按 slug 关键词
  python scripts/run.py 134 --lang go    # 运行 Go（需 solution.go 含 main）
  python scripts/run.py 134 --lang py    # 运行 Python（默认）
  python scripts/run.py --list            # 列出所有题目
"""

from __future__ import annotations

import argparse
import re
import subprocess
import sys
from pathlib import Path

ROOT = Path(__file__).resolve().parent.parent
SOLUTIONS = ROOT / "solutions"


def list_problems() -> list[Path]:
    return sorted(p for p in SOLUTIONS.iterdir() if p.is_dir() and (p / "README.md").exists())


def resolve_folder(query: str) -> Path:
    query = query.strip().lower()
    folders = list_problems()

    if query.isdigit():
        num = int(query)
        prefix = f"{num:03d}-"
        matches = [f for f in folders if f.name.startswith(prefix) or f.name == f"000-{query}"]
        if len(matches) == 1:
            return matches[0]
        if len(matches) > 1:
            raise SystemExit(f"题号 {num} 对应多个目录: {', '.join(m.name for m in matches)}")

    matches = [f for f in folders if query in f.name.lower()]
    if len(matches) == 1:
        return matches[0]
    if len(matches) > 1:
        names = ", ".join(m.name for m in matches)
        raise SystemExit(f"关键词 '{query}' 匹配多个目录: {names}\n请使用更具体的目录名或题号。")

    raise SystemExit(f"未找到题目: {query}\n使用 --list 查看全部目录。")


def run_python(folder: Path) -> int:
    py_file = folder / "solution.py"
    if not py_file.exists():
        print(f"缺少 {py_file}")
        return 1

    print(f"==> Python: {folder.name}/solution.py\n")
    return subprocess.run([sys.executable, str(py_file)], cwd=folder).returncode


def run_go(folder: Path) -> int:
    go_file = folder / "solution.go"
    if not go_file.exists():
        print(f"缺少 {go_file}")
        return 1

    content = go_file.read_text(encoding="utf-8")
    if not re.search(r"^package\s+main\b", content, re.MULTILINE):
        print(f"==> Go: {folder.name}/solution.go 为 LeetCode 提交格式（package solution），无法直接 go run。")
        print("请使用 Python 运行: python scripts/run.py <题号>")
        return 1

    print(f"==> Go: {folder.name}/solution.go\n")
    return subprocess.run(["go", "run", str(go_file)], cwd=folder).returncode


def cmd_list() -> None:
    print(f"{'目录':<55} Python  Go(main)")
    print("-" * 72)
    for folder in list_problems():
        py = "✓" if (folder / "solution.py").exists() else "-"
        go = (folder / "solution.go")
        go_runnable = "✓" if go.exists() and re.search(r"^package\s+main\b", go.read_text(encoding="utf-8"), re.MULTILINE) else "-"
        print(f"{folder.name:<55} {py:>6}  {go_runnable:>8}")


def main() -> None:
    parser = argparse.ArgumentParser(description="运行 LeetCode 题解")
    parser.add_argument("problem", nargs="?", help="题号、目录名或 slug 关键词")
    parser.add_argument("--lang", choices=["py", "go"], default="py", help="语言（默认 py）")
    parser.add_argument("--list", action="store_true", help="列出所有题目")
    args = parser.parse_args()

    if args.list:
        cmd_list()
        return

    if not args.problem:
        parser.print_help()
        print("\n示例:")
        print("  python scripts/run.py 134")
        print("  python scripts/run.py 001-two-sum --lang go")
        print("  python scripts/run.py --list")
        return

    folder = resolve_folder(args.problem)
    code = run_go(folder) if args.lang == "go" else run_python(folder)
    raise SystemExit(code)


if __name__ == "__main__":
    main()
