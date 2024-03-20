import re


class Solution:

    def isPalindrome(self, s: str) -> bool:
        pattern = '[^a-zA-Z0-9]+'
        result = re.sub(pattern, "", s).lower()
        return result == result[::-1]
