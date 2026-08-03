from typing import List


class Solution:
    def game_of_life(self, board: List[List[int]]) -> None:
        m, n = len(board), len(board[0])
        directions = (
            (-1, -1), (-1, 0), (-1, 1),
            (0, -1),           (0, 1),
            (1, -1),  (1, 0),  (1, 1),
        )

        for i in range(m):
            for j in range(n):
                live_neighbors = 0
                for di, dj in directions:
                    ni, nj = i + di, j + dj
                    if 0 <= ni < m and 0 <= nj < n and board[ni][nj] in (1, 2):
                        live_neighbors += 1

                if board[i][j] == 1:
                    if live_neighbors < 2 or live_neighbors > 3:
                        board[i][j] = 2  # 活 -> 死
                elif live_neighbors == 3:
                    board[i][j] = 3  # 死 -> 活

        for i in range(m):
            for j in range(n):
                if board[i][j] == 2:
                    board[i][j] = 0
                elif board[i][j] == 3:
                    board[i][j] = 1


if __name__ == "__main__":
    tests = [
        (
            [[0, 1, 0], [0, 0, 1], [1, 1, 1], [0, 0, 0]],
            [[0, 0, 0], [1, 0, 1], [0, 1, 1], [0, 1, 0]],
        ),
        (
            [[1, 1], [1, 0]],
            [[1, 1], [1, 1]],
        ),
    ]
    sol = Solution()
    for board, expected in tests:
        sol.game_of_life(board)
        status = "PASS" if board == expected else "FAIL"
        print(f"{status} | board={board}")
