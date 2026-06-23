# count_passwords 统计合法密码个数：a-z 字母，* 表示未知，相邻两位不能相同
#
# 思路：从左到右 DP，dp[c] = 填完当前位置且该位为字母 c 的方案数
def count_passwords(pattern: str) -> int:
    slots = _parse_password_pattern(pattern)
    if len(slots) == 0:
        return 0

    letters = 26
    dp = [0] * letters

    if slots[0] == ord("*"):
        for i in range(letters):
            dp[i] = 1
    else:
        dp[slots[0] - ord("a")] = 1

    for i in range(1, len(slots)):
        new_dp = [0] * letters
        if slots[i] == ord("*"):
            total = sum(dp)
            for k in range(letters):
                new_dp[k] = total - dp[k]
        else:
            c = slots[i] - ord("a")
            s = sum(v for j, v in enumerate(dp) if j != c)
            new_dp[c] = s
        dp = new_dp

    return sum(dp)


def _parse_password_pattern(pattern: str) -> list:
    slots = []
    for r in pattern.strip():
        if r == " ":
            continue
        if r == "*":
            slots.append(ord("*"))
        elif "a" <= r <= "z":
            slots.append(ord(r))
        else:
            return []
    return slots


if __name__ == "__main__":
    tests = [("a * b", 24), ("*", 26), ("a", 1), ("* * *", 17576)]
    for pattern, expected in tests:
        got = count_passwords(pattern)
        status = "PASS" if got == expected else "FAIL"
        print(f"{status} | pattern={pattern!r} => {got} (expected {expected})")
