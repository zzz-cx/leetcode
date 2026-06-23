# LeetCode 题解仓库

个人 LeetCode 刷题笔记，按题目独立归档。每道题包含**题目描述**、**思路与解析**、**Python / Go 双语言解答**。

**仓库地址：** https://github.com/zzz-cx/leetcode

---

## 目录结构

```
leetcode/
├── README.md                 # 本文件（总分类目录）
├── scripts/
│   └── run.py                # 本地运行题解
└── solutions/
    ├── README.md             # 按题号排序的完整索引
    ├── TEMPLATE.md           # 题解写作模板
    ├── _helpers/             # Python 测试工具
    └── {题号}-{slug}/        # 单题目录
        ├── README.md         # 题目 + 思路 + 代码
        ├── solution.py
        └── solution.go
```

---

## 快速开始

```bash
# 按题号运行（Python，默认）
python scripts/run.py 1
python scripts/run.py 134

# 按目录名或关键词
python scripts/run.py 001-two-sum
python scripts/run.py two-sum

# 运行 Go 版本（需 solution.go 含 main 函数）
python scripts/run.py 134 --lang go
go run solutions/134-gas-station/solution.go

# 列出全部题目
python scripts/run.py --list
```

---

## 总分类目录

> 点击题号进入对应题解。共 **77** 道（76 道 LeetCode + 1 道自定义）。

### 数组 · 哈希 · 前缀和

| 题号 | 题目 | 链接 |
|:----:|------|------|
| 1 | 两数之和 | [001-two-sum](./solutions/001-two-sum/) |
| 41 | 缺失的第一个正数 | [041-first-missing-positive](./solutions/041-first-missing-positive/) |
| 56 | 合并区间 | [056-merge-intervals](./solutions/056-merge-intervals/) |
| 88 | 合并两个有序数组 | [088-merge-sorted-array](./solutions/088-merge-sorted-array/) |
| 189 | 旋转数组 | [189-rotate-array](./solutions/189-rotate-array/) |
| 238 | 除自身以外数组的乘积 | [238-product-of-array-except-self](./solutions/238-product-of-array-except-self/) |
| 287 | 寻找重复数 | [287-find-the-duplicate-number](./solutions/287-find-the-duplicate-number/) |
| 347 | 前 K 个高频元素 | [347-top-k-frequent-elements](./solutions/347-top-k-frequent-elements/) |
| 560 | 和为 K 的子数组 | [560-subarray-sum-equals-k](./solutions/560-subarray-sum-equals-k/) |

### 双指针 · 滑动窗口

| 题号 | 题目 | 链接 |
|:----:|------|------|
| 3 | 无重复字符的最长子串 | [003-longest-substring-without-repeating-characters](./solutions/003-longest-substring-without-repeating-characters/) |
| 15 | 三数之和 | [015-3sum](./solutions/015-3sum/) |
| 26 | 删除有序数组中的重复项 | [026-remove-duplicates-from-sorted-array](./solutions/026-remove-duplicates-from-sorted-array/) |
| 27 | 移除元素 | [027-remove-element](./solutions/027-remove-element/) |
| 42 | 接雨水 | [042-trapping-rain-water](./solutions/042-trapping-rain-water/) |
| 75 | 颜色分类 | [075-sort-colors](./solutions/075-sort-colors/) |
| 80 | 删除有序数组中的重复项 II | [080-remove-duplicates-from-sorted-array-ii](./solutions/080-remove-duplicates-from-sorted-array-ii/) |
| 239 | 滑动窗口最大值 | [239-sliding-window-maximum](./solutions/239-sliding-window-maximum/) |
| 438 | 找到字符串中所有字母异位词 | [438-find-all-anagrams-in-a-string](./solutions/438-find-all-anagrams-in-a-string/) |
| 763 | 划分字母区间 | [763-partition-labels](./solutions/763-partition-labels/) |

### 矩阵

| 题号 | 题目 | 链接 |
|:----:|------|------|
| 48 | 旋转图像 | [048-rotate-image](./solutions/048-rotate-image/) |
| 54 | 螺旋矩阵 | [054-spiral-matrix](./solutions/054-spiral-matrix/) |
| 64 | 最小路径和 | [064-minimum-path-sum](./solutions/064-minimum-path-sum/) |
| 73 | 矩阵置零 | [073-set-matrix-zeroes](./solutions/073-set-matrix-zeroes/) |
| 74 | 搜索二维矩阵 | [074-search-a-2d-matrix](./solutions/074-search-a-2d-matrix/) |

### 链表

| 题号 | 题目 | 链接 |
|:----:|------|------|
| 2 | 两数相加 | [002-add-two-numbers](./solutions/002-add-two-numbers/) |
| 19 | 删除链表的倒数第 N 个结点 | [019-remove-nth-node-from-end-of-list](./solutions/019-remove-nth-node-from-end-of-list/) |
| 21 | 合并两个有序链表 | [021-merge-two-sorted-lists](./solutions/021-merge-two-sorted-lists/) |
| 24 | 两两交换链表中的节点 | [024-swap-nodes-in-pairs](./solutions/024-swap-nodes-in-pairs/) |
| 138 | 复制带随机指针的链表 | [138-copy-list-with-random-pointer](./solutions/138-copy-list-with-random-pointer/) |
| 141 | 环形链表 | [141-linked-list-cycle](./solutions/141-linked-list-cycle/) |
| 142 | 环形链表 II | [142-linked-list-cycle-ii](./solutions/142-linked-list-cycle-ii/) |
| 148 | 排序链表 | [148-sort-list](./solutions/148-sort-list/) |
| 160 | 相交链表 | [160-intersection-of-two-linked-lists](./solutions/160-intersection-of-two-linked-lists/) |
| 206 | 反转链表 | [206-reverse-linked-list](./solutions/206-reverse-linked-list/) |
| 234 | 回文链表 | [234-palindrome-linked-list](./solutions/234-palindrome-linked-list/) |

### 栈 · 单调栈 · 字符串

| 题号 | 题目 | 链接 |
|:----:|------|------|
| 5 | 最长回文子串 | [005-longest-palindromic-substring](./solutions/005-longest-palindromic-substring/) |
| 22 | 括号生成 | [022-generate-parentheses](./solutions/022-generate-parentheses/) |
| 84 | 柱状图中最大的矩形 | [084-largest-rectangle-in-histogram](./solutions/084-largest-rectangle-in-histogram/) |
| 155 | 最小栈 | [155-min-stack](./solutions/155-min-stack/) |
| 394 | 字符串解码 | [394-decode-string](./solutions/394-decode-string/) |

### 二叉树

| 题号 | 题目 | 链接 |
|:----:|------|------|
| 94 | 二叉树的中序遍历 | [094-binary-tree-inorder-traversal](./solutions/094-binary-tree-inorder-traversal/) |
| 101 | 对称二叉树 | [101-symmetric-tree](./solutions/101-symmetric-tree/) |
| 102 | 二叉树的层序遍历 | [102-binary-tree-level-order-traversal](./solutions/102-binary-tree-level-order-traversal/) |
| 104 | 二叉树的最大深度 | [104-maximum-depth-of-binary-tree](./solutions/104-maximum-depth-of-binary-tree/) |
| 105 | 从前序与中序遍历序列构造二叉树 | [105-construct-binary-tree-from-preorder-and-inorder-traversal](./solutions/105-construct-binary-tree-from-preorder-and-inorder-traversal/) |
| 114 | 二叉树展开为链表 | [114-flatten-binary-tree-to-linked-list](./solutions/114-flatten-binary-tree-to-linked-list/) |
| 199 | 二叉树的右视图 | [199-binary-tree-right-side-view](./solutions/199-binary-tree-right-side-view/) |
| 226 | 翻转二叉树 | [226-invert-binary-tree](./solutions/226-invert-binary-tree/) |
| 230 | 二叉搜索树中第 K 小的元素 | [230-kth-smallest-element-in-a-bst](./solutions/230-kth-smallest-element-in-a-bst/) |
| 236 | 二叉树的最近公共祖先 | [236-lowest-common-ancestor-of-a-binary-tree](./solutions/236-lowest-common-ancestor-of-a-binary-tree/) |
| 437 | 路径总和 III | [437-path-sum-iii](./solutions/437-path-sum-iii/) |
| 543 | 二叉树的直径 | [543-diameter-of-binary-tree](./solutions/543-diameter-of-binary-tree/) |
| 662 | 二叉树最大宽度 | [662-maximum-width-of-binary-tree](./solutions/662-maximum-width-of-binary-tree/) |

### 图 · BFS · DFS

| 题号 | 题目 | 链接 |
|:----:|------|------|
| 200 | 岛屿数量 | [200-number-of-islands](./solutions/200-number-of-islands/) |
| 207 | 课程表 | [207-course-schedule](./solutions/207-course-schedule/) |
| 994 | 腐烂的橘子 | [994-rotting-oranges](./solutions/994-rotting-oranges/) |

### 回溯

| 题号 | 题目 | 链接 |
|:----:|------|------|
| 46 | 全排列 | [046-permutations](./solutions/046-permutations/) |
| 131 | 分割回文串 | [131-palindrome-partitioning](./solutions/131-palindrome-partitioning/) |

### 贪心

| 题号 | 题目 | 链接 |
|:----:|------|------|
| 45 | 跳跃游戏 II | [045-jump-game-ii](./solutions/045-jump-game-ii/) |
| 55 | 跳跃游戏 | [055-jump-game](./solutions/055-jump-game/) |
| 134 | 加油站 | [134-gas-station](./solutions/134-gas-station/) |

### 动态规划

| 题号 | 题目 | 链接 |
|:----:|------|------|
| 53 | 最大子数组和 | [053-maximum-subarray](./solutions/053-maximum-subarray/) |
| 72 | 编辑距离 | [072-edit-distance](./solutions/072-edit-distance/) |
| 121 | 买卖股票的最佳时机 | [121-best-time-to-buy-and-sell-stock](./solutions/121-best-time-to-buy-and-sell-stock/) |
| 122 | 买卖股票的最佳时机 II | [122-best-time-to-buy-and-sell-stock-ii](./solutions/122-best-time-to-buy-and-sell-stock-ii/) |
| 139 | 单词拆分 | [139-word-break](./solutions/139-word-break/) |
| 152 | 最大子数组乘积 | [152-maximum-product-subarray](./solutions/152-maximum-product-subarray/) |
| 300 | 最长递增子序列 | [300-longest-increasing-subsequence](./solutions/300-longest-increasing-subsequence/) |
| 322 | 零钱兑换 | [322-coin-change](./solutions/322-coin-change/) |

### 二分查找

| 题号 | 题目 | 链接 |
|:----:|------|------|
| 4 | 寻找两个正序数组的中位数 | [004-median-of-two-sorted-arrays](./solutions/004-median-of-two-sorted-arrays/) |
| 33 | 搜索旋转排序数组 | [033-search-in-rotated-sorted-array](./solutions/033-search-in-rotated-sorted-array/) |
| 34 | 在排序数组中查找元素的第一个和最后一个位置 | [034-find-first-and-last-position-of-element-in-sorted-array](./solutions/034-find-first-and-last-position-of-element-in-sorted-array/) |
| 153 | 寻找旋转排序数组中的最小值 | [153-find-minimum-in-rotated-sorted-array](./solutions/153-find-minimum-in-rotated-sorted-array/) |
| 215 | 数组中的第 K 个最大元素 | [215-kth-largest-element-in-an-array](./solutions/215-kth-largest-element-in-an-array/) |

### 排序 · 设计

| 题号 | 题目 | 链接 |
|:----:|------|------|
| 31 | 下一个排列 | [031-next-permutation](./solutions/031-next-permutation/) |
| 380 | 插入删除获取随机元素 | [380-insert-delete-getrandom-o1](./solutions/380-insert-delete-getrandom-o1/) |

### 其他（自定义）

| 题号 | 题目 | 链接 |
|:----:|------|------|
| — | 密码计数（相邻字母不能相同） | [000-password-count](./solutions/000-password-count/) |

---

## 按题号索引

需要按题号顺序浏览时，见 [solutions/README.md](./solutions/README.md)。

---

## 题解格式说明

每道题的 `README.md` 参考以下结构撰写（详见 [solutions/TEMPLATE.md](./solutions/TEMPLATE.md)）：

1. **题目** — 题意、示例、约束
2. **思路与解析** — 问题转化、算法策略、复杂度
3. **解答** — Python / Go 代码及逐行解析

---

## 统计

| 分类 | 题数 |
|------|:----:|
| 数组 · 哈希 · 前缀和 | 9 |
| 双指针 · 滑动窗口 | 10 |
| 矩阵 | 5 |
| 链表 | 11 |
| 栈 · 字符串 | 5 |
| 二叉树 | 13 |
| 图 · BFS · DFS | 3 |
| 回溯 | 2 |
| 贪心 | 3 |
| 动态规划 | 8 |
| 二分查找 | 5 |
| 排序 · 设计 | 2 |
| 自定义 | 1 |
| **合计** | **77** |
