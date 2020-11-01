# 逻辑回归算法(决策边界为直线的情况)
# 假设函数(3阶多项式):h(x)=theta0*x0+theta1*x1+theta2*x2+theta3x1^2+...
#
# step1.加载数据并打印散点图
import numpy as np
import matplotlib
import matplotlib.pyplot as plt
from PlotLinearRegressionDecisionBoundary import PlotLinearRegressionDecisionBoundary

theta = np.zeros((10,1)) #n=3
#theta = np.zeros((91,1)) #n=12
_lambda = 0          #正则化参数
iter_num = 10000      #迭代次数
learningRate=0.01   #学习率
min_change = 1e-5    #梯度变化阈值

def mapFeature(X, n):
    if n==1:return X
    X1 = X[:, 1].reshape(-1, 1)
    X2 = X[:, 2].reshape(-1, 1)
    map_X = np.zeros((len(X), 1))
    num = 0
    for i in range(0, n + 1):
        for j in range(0, i + 1):
            num += 1
            temp = X1 ** (i-j) * X2 ** j
            map_X = np.hstack((map_X, np.array(temp).reshape(-1, 1)))
    return map_X[:, 1:]
def printScatter(filename):
    # 读取数据
    data = np.loadtxt(filename, delimiter=',')
    x = data[:, 0:2]
    y = data[:, 2]
    pos = np.where(y == 1)
    neg = np.where(y == 0)
    # plt.scatter(x[pos, 0], x[pos, 1], marker='s', c='r')
    # plt.scatter(x[neg, 0], x[neg, 1], marker='x', c='b')
    plt.plot(x[pos][:,0],x[pos][:,1], 'k+', linewidth=2, markersize=7)
    plt.plot(x[neg][:,0],x[neg][:,1], 'ko', markerfacecolor='y', markersize=7)

    plt.show()
# step1.加载数据并打印散点图
printScatter("ex2data2.txt")
# step2.加载训练数据集(0,6)loadDataset(filename)
def loadDataset(filename, percent):
    data = np.loadtxt(filename, delimiter=',')
    np.random.shuffle(data)
    total = len(data)
    m = int(total * percent)
    X = data[:m, 0:2]
    Y = data[:m, -1].reshape(-1,1)
    rest_X = data[m:,0:2]
    temp = np.ones((total-m, 1)) # col 1
    rest_X = np.hstack((temp, rest_X))
    #add c 1
    rest_Y = data[m:,-1].reshape(-1,1)
    X0 = np.ones((m, 1))
    X = np.hstack((X0, X))
    n = X.shape[1]
    return X, Y, m, n,rest_X,rest_Y
X, Y, m, n,rest_X,rest_Y = loadDataset('ex2data2.txt', 0.99)
#print(X,Y)
#矩阵扩展为3阶多项式矩阵
X = mapFeature(X,3)
#X = mapFeature(X,12)
#print(X)
# step3.计算代价函数computeCost(theta,X,Y)
def sigmoid(z):
    return 1 / (1 + np.exp(-z))
#代价函数 J
def computeCost(X, Y, theta):
    m = len(Y)
    z = sigmoid(np.dot(X, theta))
    z[z == 1] = 9e-18
    z[z == 0] = 1e-18
    J = - (np.dot(Y.T, np.log(z)) + np.dot((1 - Y).T, np.log(1 - z))) / m + (_lambda/2*m)*np.dot(theta[1:].T,theta[1:])
    return J
J = computeCost(X, Y, theta)
#print("J cost:\t",J)
# step4.计算梯度computeDescent(theta,X,Y)
def computeDescent(theta, X, Y):
    z = sigmoid(np.dot(X, theta))
    gradient = (np.dot(X.T,(z - Y))+_lambda*np.sum(theta,axis=0))/m
    return gradient
#print("computeDescent:\t",computeDescent(theta,X,Y))
# step5.梯度下降函数gradientDescent(theta,X,Y,iter_num,learningRate)
def gradientDescent(theta,X,Y,iter_num,learningRate):
    i = 0
    J_list = []
    gradient = computeDescent(theta,X,Y)
    while (not np.all(np.absolute(gradient)<=min_change)) and (i<iter_num):
        theta[0] = theta[0] - learningRate*gradient[0]
        theta[1:] = theta[1:] - learningRate * gradient[1:] + (_lambda/m)*theta[1:]
        gradient = computeDescent(theta, X, Y)
        J_list.append(computeCost(X, Y, theta))
        i += 1
        print(i)
        #print("gradient",gradient)
    return theta,J_list,i
# step6.执行梯度下降函数
theta ,J_list,iter_num = gradientDescent(theta,X,Y,iter_num,learningRate)
print(theta)
# step7.(附1:评价)生成代价函数(数值)与迭代次数(iter_num)的曲线
plt.plot(np.arange(0,iter_num,1),np.array(J_list).reshape(iter_num,1),'-r')
plt.show()
# # step8.(附2:决策边界,绘图)决策边界printDecisionBoundary()
# #拟合决策边界
PlotLinearRegressionDecisionBoundary(theta, X, Y)
# Labels and Legend
plt.xlabel('Exam 1 score')
plt.ylabel('Exam 1 score')
plt.legend(['Admitted', 'Not admitted', 'Decision boundary'])
plt.show()
#  step9.计算算法的分类准确度
#取出剩余的训练数据作为输入数据
#rest_X = mapFeature(rest_X,12)
rest_X = mapFeature(rest_X,3)
predict_Y = sigmoid(np.dot(rest_X,theta))
predict_Y[predict_Y>=0.5] = 1
predict_Y[predict_Y<0.5] = 0
predict_data = predict_Y-rest_Y
Train_Accuracy = len(predict_data[predict_data==0])/len(rest_Y)
print('Train Accuracy: %.2f'%( Train_Accuracy* 100)+' %'

