import numpy as np
# 两个特征变量转为多项式来假设函数

def polynomial(A,n):
    A = np.array(A)
    X = np.hstack((A,(A[:,1]**n).reshape(-1,1),(A[:,2]**n).reshape(-1,1)))
    for i in range(1,n):
        m = factorial(n)/(factorial(i)*factorial(n-i))
        X=np.hstack((X,(m*(A[:,1]**(n-i))*(A[:,2])**(i)).reshape(-1,1)))
    return X


def factorial(num):
    if num == 1:
        return 1
    else:
        return num * factorial(num - 1)
