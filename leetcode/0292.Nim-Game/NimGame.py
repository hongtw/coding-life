class Solution:
    def canWinNim(self, n: int) -> bool:
        '''
        4       ->  I fail
        5, 6, 7 -> He fail
        8       -> I fail again
        '''
        return n % 4 != 0