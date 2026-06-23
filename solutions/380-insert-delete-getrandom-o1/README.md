# 插入、删除和获取随机元素

> LeetCode 380 · [insert-delete-getrandom-o1](https://leetcode.cn/problems/insert-delete-getrandom-o1/)

## 题目

设计支持 insert、remove、getRandom 且平均 O(1) 的数据结构。

## 题解思路与解析

- 见下方代码实现与注释。

## 解答

### Golang

```go
type RandomizedSet struct {
	valList []int
	valMap  map[int]int
}

func NewRandomizedSet() RandomizedSet {
	valList := make([]int, 0)
	valMap := make(map[int]int)
	return RandomizedSet{valList: valList, valMap: valMap}
}

func (this *RandomizedSet) Insert(val int) bool {

}

func (this *RandomizedSet) Remove(val int) bool {

}

func (this *RandomizedSet) GetRandom() int {

}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
```

### Python

```python
import random


class RandomizedSet:
    def __init__(self):
        self.val_list = []
        self.val_map = {}

    def insert(self, val: int) -> bool:
        if val in self.val_map:
            return False
        self.val_map[val] = len(self.val_list)
        self.val_list.append(val)
        return True

    def remove(self, val: int) -> bool:
        if val not in self.val_map:
            return False
        idx = self.val_map[val]
        last = self.val_list[-1]
        self.val_list[idx] = last
        self.val_map[last] = idx
        self.val_list.pop()
        del self.val_map[val]
        return True

    def get_random(self) -> int:
        return random.choice(self.val_list)
```
