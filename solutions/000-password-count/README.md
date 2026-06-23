# 密码计数（相邻字母不能相同）

## 题目

自定义题：统计合法密码个数。密码由 a-z 字母组成，`*` 表示未知位，相邻两位字母不能相同。

## 题解思路与解析

- countPasswords 统计合法密码个数：a-z 字母，* 表示未知，相邻两位不能相同
- 思路：从左到右 DP，dp[c] = 填完当前位置且该位为字母 c 的方案数
- - 当前位是 *：newDp[k] = sum(dp[j]) - dp[k]（上一位不能等于 k）
- - 当前位是固定字母 c：newDp[c] = sum(dp[j], j!=c)，其余为 0
- parsePasswordPattern 解析输入，忽略空格；仅允许 a-z 与 *

## 解答

### Golang

```go
"strings"

// countPasswords 统计合法密码个数：a-z 字母，* 表示未知，相邻两位不能相同
//
// 思路：从左到右 DP，dp[c] = 填完当前位置且该位为字母 c 的方案数
//   - 当前位是 *：newDp[k] = sum(dp[j]) - dp[k]（上一位不能等于 k）
//   - 当前位是固定字母 c：newDp[c] = sum(dp[j], j!=c)，其余为 0
func countPasswords(pattern string) int64 {
	slots := parsePasswordPattern(pattern)
	if len(slots) == 0 {
		return 0
	}

	const letters = 26
	dp := make([]int64, letters)

	if slots[0] == '*' {
		for i := range dp {
			dp[i] = 1
		}
	} else {
		dp[slots[0]-'a'] = 1
	}

	for i := 1; i < len(slots); i++ {
		newDp := make([]int64, letters)
		if slots[i] == '*' {
			var total int64
			for _, v := range dp {
				total += v
			}
			for k := range newDp {
				newDp[k] = total - dp[k]
			}
		} else {
			c := int(slots[i] - 'a')
			var sum int64
			for j, v := range dp {
				if j != c {
					sum += v
				}
			}
			newDp[c] = sum
		}
		dp = newDp
	}

	var ans int64
	for _, v := range dp {
		ans += v
	}
	return ans
}

// parsePasswordPattern 解析输入，忽略空格；仅允许 a-z 与 *
func parsePasswordPattern(pattern string) []byte {
	var slots []byte
	for _, r := range strings.TrimSpace(pattern) {
		if r == ' ' {
			continue
		}
		switch {
		case r == '*':
			slots = append(slots, '*')
		case r >= 'a' && r <= 'z':
			slots = append(slots, byte(r))
		default:
			return nil
		}
	}
	return slots
}
```

### Python

```python
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
```
