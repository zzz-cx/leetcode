class Solution:
    def decode_string(self, s: str) -> str:
        count_stack = []
        str_stack = []
        cur = ""
        num = 0

        for ch in s:
            if ch.isdigit():
                num = num * 10 + int(ch)
            elif ch == "[":
                count_stack.append(num)
                str_stack.append(cur)
                num = 0
                cur = ""
            elif ch == "]":
                k = count_stack.pop()
                prev = str_stack.pop()
                cur = prev + cur * k
            else:
                cur += ch
        return cur


if __name__ == "__main__":
    sol = Solution()
    got = sol.decode_string("3[a]2[bc]")
    status = "PASS" if got == "aaabcbc" else "FAIL"
    print(f"{status} | => {got!r}")
