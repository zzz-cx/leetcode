class MinStack:
    def __init__(self):
        self.stack = []
        self.min_stack = []

    def push(self, val: int) -> None:
        self.stack.append(val)
        if not self.min_stack or val <= self.min_stack[-1]:
            self.min_stack.append(val)

    def pop(self) -> None:
        if not self.stack:
            return
        top = self.stack.pop()
        if self.min_stack and top == self.min_stack[-1]:
            self.min_stack.pop()

    def top(self) -> int:
        return self.stack[-1]

    def get_min(self) -> int:
        return self.min_stack[-1]


if __name__ == "__main__":
    s = MinStack()
    s.push(-2)
    s.push(0)
    s.push(-3)
    m1 = s.get_min()
    s.pop()
    t = s.top()
    m2 = s.get_min()
    ok = m1 == -3 and t == 0 and m2 == -2
    status = "PASS" if ok else "FAIL"
    print(f"{status} | GetMin={m1}, Top={t}, GetMin={m2}")
