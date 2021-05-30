## 3 sum and find closest element.

### solution 1
```python
import sys
class Solution:
    def closest3Sum(self, A, N, X):
        closestSum = sys.maxsize
        for i in range(N):
            for j in range (i+1, N):
                for k in range(j+1, N):
                    if abs(X-closestSum) > abs(X- (A[i]+A[j]+A[k])):
                        closestSum = A[i]+A[j]+A[k]
        return closestSum
```
- TC: O(N^3)

### solution 2
