class Solution(object):
    def romanToInt(self, x):
        numberList = {"I":1,"V":5,"X":10,"L":50,"C":100,"D":500,"M":1000}

        resultNumber = 0
        for idx in range(len(x)-1):
            if numberList[x[idx]] < numberList[x[idx+1]]:
                resultNumber -= numberList[x[idx]]
            else:
                resultNumber += numberList[x[idx]]
        
        return resultNumber + numberList[x[-1]]
    