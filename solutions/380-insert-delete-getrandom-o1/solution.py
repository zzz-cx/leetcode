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


if __name__ == "__main__":
    s = RandomizedSet()
    r1 = s.insert(1)
    r2 = s.insert(2)
    r3 = s.insert(1)
    r4 = s.remove(2)
    r5 = s.insert(2)
    ok = r1 and r2 and not r3 and r4 and r5
    status = "PASS" if ok else "FAIL"
    print(f"{status} | Insert(1)={r1} Insert(2)={r2} Insert(1)={r3} Remove(2)={r4} Insert(2)={r5}")
