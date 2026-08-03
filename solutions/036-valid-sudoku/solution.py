from typing import List


class Solution:
    def is_valid_sudoku(self, board: List[List[str]]) -> bool:
        rows: list[set[str]] = [set() for _ in range(9)]
        cols: list[set[str]] = [set() for _ in range(9)]
        boxes: list[set[str]] = [set() for _ in range(9)]

        for r in range(9):
            for c in range(9):
                ch = board[r][c]
                if ch == ".":
                    continue
                box = (r // 3) * 3 + c // 3
                if ch in rows[r] or ch in cols[c] or ch in boxes[box]:
                    return False
                rows[r].add(ch)
                cols[c].add(ch)
                boxes[box].add(ch)

        return True


if __name__ == "__main__":
    valid = [
        ["5", "3", ".", ".", "7", ".", ".", ".", "."],
        ["6", ".", ".", "1", "9", "5", ".", ".", "."],
        [".", "9", "8", ".", ".", ".", ".", "6", "."],
        ["8", ".", ".", ".", "6", ".", ".", ".", "3"],
        ["4", ".", ".", "8", ".", "3", ".", ".", "1"],
        ["7", ".", ".", ".", "2", ".", ".", ".", "6"],
        [".", "6", ".", ".", ".", ".", "2", "8", "."],
        [".", ".", ".", "4", "1", "9", ".", ".", "5"],
        [".", ".", ".", ".", "8", ".", ".", "7", "9"],
    ]
    invalid_row = [row[:] for row in valid]
    invalid_row[0][0] = "3"  # two 3s in row 0

    tests = [
        (valid, True),
        (invalid_row, False),
    ]

    sol = Solution()
    for board, expected in tests:
        result = sol.is_valid_sudoku(board)
        status = "PASS" if result == expected else "FAIL"
        print(f"{status} => {result} (expected {expected})")
