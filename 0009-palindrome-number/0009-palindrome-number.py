class Solution(object):
    def isPalindrome(self, x):
        stringNumber = str(x)
        listNumber = list(stringNumber)

        reverseNumber = listNumber[::-1]

        if listNumber == reverseNumber:
            return True
        else:
            return False
        