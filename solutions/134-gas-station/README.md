# 134. 加油站

**难度：** 中等  
**标签：** 贪心、数组

> [gas-station](https://leetcode.cn/problems/gas-station/)

## 题目描述

在一条环路上有 `n` 个加油站，其中第 `i` 个加油站有汽油 `gas[i]` 升。

你有一辆油箱容量无限的汽车，从第 `i` 个加油站开往第 `i+1` 个加油站需要消耗汽油 `cost[i]` 升。你从其中的一个加油站出发，开始时油箱为空。

给定两个整数数组 `gas` 和 `cost`，如果你可以按顺序绕环路行驶一周，则返回出发时加油站的编号，否则返回 `-1`。如果存在解，则保证它是唯一的。

### 示例 1

```
输入: gas = [1,2,3,4,5], cost = [3,4,5,1,2]
输出: 3
```

从索引 3 出发，依次经过 4 → 0 → 1 → 2 → 3，可以完成一整圈。

### 示例 2

```
输入: gas = [2,3,4], cost = [3,4,3]
输出: -1
```

总油量不足以支撑绕环一周。

### 提示

- `gas.length == n`，`cost.length == n`
- `1 <= n <= 10^5`
- `0 <= gas[i], cost[i] <= 10^4`

---

## 思路说明

### 问题转化

把每个站点 `i` 的净收益定义为 `diff[i] = gas[i] - cost[i]`，表示在该站加油后、开往下一站前的油量变化。

- 若从站点 `s` 出发，依次经过 `s, s+1, ..., s+n-1`（下标对 `n` 取模），能完成一圈，等价于从 `s` 出发的前缀累计和始终 ≥ 0。
- 若 `sum(diff) < 0`，总油量不够，无解。
- 题目保证有解时唯一，因此只需找到一个合法起点。

### 贪心策略

维护三个变量：

| 变量 | 含义 |
|------|------|
| `start` | 当前候选起点 |
| `current_tank` | 从 `start` 出发，走到当前站后的油箱余量 |
| `total_tank` | 全程 `diff` 之和，用于最终判断是否有解 |

遍历每个站点 `i`：

1. 计算 `diff = gas[i] - cost[i]`，累加到 `total_tank` 和 `current_tank`。
2. 若 `current_tank < 0`，说明从 `start` 到 `i` 这一段无法支撑继续行驶，`start` 到 `i` 之间的任意站点都不可能成为合法起点（见下方证明），于是令 `start = i + 1`，`current_tank = 0`。

遍历结束后：若 `total_tank >= 0`，返回 `start`；否则返回 `-1`。

### 为什么 `current_tank < 0` 时可以跳过 `start..i`？

假设从 `start` 出发能到达 `i`，但到达 `i+1` 时油箱为负。对任意 `k ∈ (start, i]`，从 `k` 出发到达 `i` 时，油箱余量 = 从 `start` 出发到达 `i` 的余量 − 从 `start` 到 `k-1` 的净收益之和。

因为从 `start` 到 `k-1` 每一步都走通了，这段净收益之和 ≥ 0，所以从 `k` 出发到达 `i` 的余量 ≤ 从 `start` 出发的余量。既然从 `start` 出发到 `i+1` 已经失败，从 `k` 出发只会更早失败。

因此下一个候选起点只能是 `i + 1`。

### 算法流程图

```
开始
  ↓
total_tank = 0, current_tank = 0, start = 0
  ↓
遍历 i = 0 .. n-1
  ↓
diff = gas[i] - cost[i]
total_tank += diff
current_tank += diff
  ↓
current_tank < 0 ?
  ├─ 是 → start = i+1, current_tank = 0
  └─ 否 → 继续
  ↓
total_tank >= 0 ?
  ├─ 是 → 返回 start
  └─ 否 → 返回 -1
```

### 复杂度

- **时间：** O(n)，单次遍历
- **空间：** O(1)，仅常数额外变量

---

## 解答

### Python

文件：[solution.py](./solution.py)

```python
from typing import List


class Solution:
    def can_complete_circuit(self, gas: List[int], cost: List[int]) -> int:
        total_tank = 0      # 全程净收益之和
        current_tank = 0    # 从 start 出发到当前的余量
        start = 0           # 候选起点

        for i in range(len(gas)):
            diff = gas[i] - cost[i]
            total_tank += diff
            current_tank += diff

            if current_tank < 0:
                start = i + 1
                current_tank = 0

        return start if total_tank >= 0 else -1
```

**代码解析**

- **`diff = gas[i] - cost[i]`** — 将「加油 − 耗油」合并为一步，避免分别维护加油和耗油逻辑。
- **`total_tank`** — 与起点无关，只判断全局是否有解。即使 `current_tank` 中途归零重试，`total_tank` 仍保留完整信息。
- **`current_tank < 0` 时重置** — 贪心核心：当前段失败，则 `start..i` 全部排除，从 `i+1` 重新尝试。
- **最后判断 `total_tank >= 0`** — 贪心找到的 `start` 在总量足够时一定正确；总量不足时直接返回 `-1`。

### Golang

文件：[solution.go](./solution.go)

```go
func canCompleteCircuit(gas []int, cost []int) int {
    totalTank := 0
    currentTank := 0
    start := 0

    for i := 0; i < len(gas); i++ {
        diff := gas[i] - cost[i]
        totalTank += diff
        currentTank += diff

        if currentTank < 0 {
            start = i + 1
            currentTank = 0
        }
    }

    if totalTank < 0 {
        return -1
    }
    return start
}
```

**代码解析**

- **与 Python 逻辑完全一致** — Go 版本使用相同的贪心框架，便于对照学习。
- **`totalTank` 放在最后判断** — Python 用三元表达式 `return start if total_tank >= 0 else -1`；Go 用 `if totalTank < 0` 提前返回 `-1`，语义相同，更符合 Go 的显式风格。
- **无额外分配** — 只使用三个 `int` 变量，满足 O(1) 空间要求，适合 `n <= 10^5` 的数据规模。

---

## 运行

```bash
# Python
python solutions/134-gas-station/solution.py

# Go
go run solutions/134-gas-station/solution.go
```

**预期输出**

```
PASS | gas=[1, 2, 3, 4, 5], cost=[3, 4, 5, 1, 2] => 3 (expected 3)
PASS | gas=[2, 3, 4], cost=[3, 4, 3] => -1 (expected -1)
PASS | gas=[5, 1, 2, 3, 4], cost=[4, 4, 1, 5, 1] => 4 (expected 4)
PASS | gas=[3, 1, 1], cost=[1, 2, 2] => 0 (expected 0)
```
