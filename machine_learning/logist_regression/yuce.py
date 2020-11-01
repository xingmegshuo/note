#用一元线性方程预测钢筋抗压强度

#测试250,260
list_x=[[i] for i in range(150,250,10)]
list_y=[[56.9],[58.3],[61.6],[64.6],[68.1],[73],[74.1],[77.4],[80.2],[82.5]]
# print(list_x)
# segma_xi_yi=0
# n=len(list_x)
# x_bar=sum(list_x)/n
# y_bar=sum(list_y)/n
# segma_xi_2=0
# for i in range(n):
    # segma_xi_yi+=list_x[i]*list_y[i]
    # segma_xi_2 += list_x[i]**2
# b=(segma_xi_yi-n*x_bar*y_bar)/(segma_xi_2-n*x_bar**2)
# print(b)

# a=y_bar-x_bar*b
# print(a)
# x1=250
# y1=a+b*x1
# print(y1)




#使用python 库

from sklearn import linear_model #线性回归模型


regr=linear_model.LinearRegression()

regr.fit(list_x,list_y)

predict_outcome=regr.predict(250)
print(predict_outcome)