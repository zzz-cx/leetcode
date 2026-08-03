# 202. 快乐数

**难度：** 简单  
**标签：** 哈希表、数学、双指针

> [happy-number](https://leetcode.cn/problems/happy-number/)

## 题目描述

编写一个算法来判断一个数 `n` 是不是**快乐数**。

对于一个正整数，每一次将该数替换为它**每个位置上的数字的平方和**，然后重复这个过程：

- 若最终变为 `1`，则 `n` 是快乐数
- 若进入**无限循环**且始终变不到 `1`，则不是快乐数

若 `n` 是快乐数返回 `true`，否则返回 `false`。

### 示例 1

```
输入: n = 19
输出: true
```

解释：

```
1² + 9² = 82
8² + 2² = 68
6² + 8² = 100
1² + 0² + 0² = 1
```

### 示例 2

```
输入: n = 2
输出: false
```

### 提示

- `1 <= n <= 2^31 - 1`

---

## 思路说明

### 问题转化

每次变换把 `n` 映射到下一个数 `f(n)`（各位数字平方和），形成序列：

```
n → f(n) → f(f(n)) → ...
```

- 到达 `1` → 快乐数
- 出现重复数字 → 进入环，不可能再到 `1`

本质是判断序列中**是否出现环**，以及环的入口是否为 `1`。

### 计算各位平方和

```python
def next_num(n: int) -> int:
    total = 0
    while n:
        d = n % 10
        total += d * d
        n //= 10
    return total
```

Python 也可写为：`sum(int(d) ** 2 for d in str(n))`。

### 方法一：哈希集合（推荐）

用 `set` 记录已出现过的数，若再次遇到则说明进入环：

```
while n != 1 and n not in seen:
    seen.add(n)
    n = next_num(n)
return n == 1
```

- **时间：** O(log n) 每步位数有限，循环次数有上界
- **空间：** O(log n) 存储已访问数

### 方法二：快慢指针（Floyd 判环）

不额外存集合，用慢指针走一步、快指针走两步：

```
slow = next(slow)
fast = next(next(fast))
若 fast == 1 → true
若 slow == fast → false（环上且无 1）
```

- **空间：** O(1)

与 [141 环形链表](../141-linked-list-cycle/) 的 Floyd 判环思路相同，只是「下一个节点」换成了 `next_num(n)`。

### 算法流程图

```
开始
  ↓
seen = {}
  ↓
n == 1 ? → 是 → 返回 true
  ↓ 否
n in seen ? → 是 → 返回 false（有环）
  ↓ 否
seen.add(n), n = next_num(n)
  ↓
回到循环
```

### 复杂度

| 方法 | 时间 | 空间 |
|------|------|------|
| 哈希集合 | O(k · log n) | O(k) |
| 快慢指针 | O(k · log n) | O(1) |

`k` 为序列长度（有数学上界，实际很小）。

---

## 同类型题目

| 题号 | 题目 | 思路要点 |
|:----:|------|----------|
| **202** | **快乐数** | 序列判环 + 数字变换 |
| [141](../141-linked-list-cycle/) | 环形链表 | Floyd 快慢指针判环 |
| [142](../142-linked-list-cycle-ii/) | 环形链表 II | 找环入口 |
| [287](../287-find-the-duplicate-number/) | 寻找重复数 | 将数组视为链表判环 |
| [258](https://leetcode.cn/problems/add-digits/) | 各位相加 | 反复数字位运算 |

**核心模式：** 反复应用函数 `f(x)` 时，用**哈希集合**或**快慢指针**检测循环。

---

## 解答

### Python

文件：[solution.py](./solution.py)

```python
class Solution:
    def is_happy(self, n: int) -> bool:
        seen: set[int] = set()

        while n != 1 and n not in seen:
            seen.add(n)
            n = sum(int(d) * int(d) for d in str(n))

        return n == 1
```

**代码解析**

- **`n not in seen`** — 检测到环，说明无法到达 1。
- **`n == 1`** — 循环结束时若为 1 则是快乐数。
- **字符串拆位** — 写法简洁；也可用 `% 10` / `// 10`。

### Golang

文件：[solution.go](./solution.go)

```go
func squareSum(n int) int {
    sum := 0
    for n > 0 {
        d := n % 10
        sum += d * d
        n /= 10
    }
    return sum
}

func isHappy(n int) bool {
    seen := make(map[int]bool)

    for n != 1 && !seen[n] {
        seen[n] = true
        n = squareSum(n)
    }
    return n == 1
}
```

**代码解析**

- **`squareSum`** — 用取模/整除计算各位平方和，避免字符串转换。
- **`map[int]bool`** — 记录访问过的数，判环。
- **与 Python 逻辑一致** — 先判 1，再判环。

---

## 运行

```bash
python scripts/run.py 202
python scripts/run.py 202 --lang go
```

**预期输出**

```
PASS | n=19 => True (expected True)
PASS | n=2 => False (expected False)
PASS | n=1 => True (expected True)
PASS | n=7 => True (expected True)
```
