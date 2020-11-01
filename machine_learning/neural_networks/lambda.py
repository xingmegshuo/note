def nnCostFunction(a,b):
    return a,a+b
a = 1
b = 3

print(nnCostFunction(1,3))
demo = lambda p: nnCostFunction(p,b)
print(demo(a))
