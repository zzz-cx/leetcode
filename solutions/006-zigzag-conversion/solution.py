class Solution:
    def convert(self, s: str, num_rows: int) -> str:
        if num_rows == 1 or num_rows >= len(s):
            return s

        rows = [""] * num_rows
        cur_row = 0
        going_down = False

        for ch in s:
            rows[cur_row] += ch
            if cur_row == 0 or cur_row == num_rows - 1:
                going_down = not going_down
            cur_row += 1 if going_down else -1

        return "".join(rows)


if __name__ == "__main__":
    tests = [
        ("PAYPALISHIRING", 3, "PAHNAPLSIIGYIR"),
        ("PAYPALISHIRING", 4, "PINALSIGYAHRPI"),
        ("A", 1, "A"),
    ]
    sol = Solution()
    for s, num_rows, expected in tests:
        result = sol.convert(s, num_rows)
        status = "PASS" if result == expected else "FAIL"
        print(f"{status} | s={s!r}, numRows={num_rows} => {result!r} (expected {expected!r})")
