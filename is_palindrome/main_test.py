import unittest

from is_palindrome.main import Solution


class TestSolution(unittest.TestCase):
    def test_1(self):
        sol = Solution()
        ans = sol.isPalindrome('A man, a plan, a canal: Panama')
        self.assertEqual(True, ans)

    def test_2(self):
        sol = Solution()
        ans = sol.isPalindrome('race a car')
        self.assertEqual(False, ans)

    def test_3(self):
        sol = Solution()
        ans = sol.isPalindrome(' ')
        self.assertEqual(True, ans)


if __name__ == '__main__':
    unittest.main()
