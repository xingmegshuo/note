# 逻辑回归算法(多分类[手写识别0-9])
# 假设函数:h(x)=theta0*x0+theta1*x1+theta2*x2+thetan*xn+...
#
import numpy as np
import matplotlib
import matplotlib.pyplot as plt
import scipy.io as sio
import random
from displayDataHandwritten import displayDataHandwritten
# 读取数据
print("加载数据中 ....")
data = sio.loadmat('ex3data1.mat')
X = data['X']
X_for_display = X
Y = data['y']
# 打乱训练集
data_percent = 0.9   #读取比例
data = np.hstack((X, Y))
total = data.shape[0]
np.random.shuffle(data)
X = data[:int(total*data_percent), :-1]
Y = data[:int(total*data_percent), -1].reshape(-1, 1)
m, n = X.shape
X0 = np.ones((m, 1))
X = np.hstack((X0, X))
theta = np.ones((n + 1, 1))
#***************************************
iter_num = 2000  # 迭代次数
learningRate = 0.08  # 学习率
min_change = 1e-5  # 梯度变化阈值
#****************************************
print("*"*50)
print("训练数据比例:\t",float(data_percent)*100,"%")
print("随机读取数据:\t是")
print("迭代次数:\t",int(iter_num))
print("学习率:\t",float(learningRate))
print("梯度变化阈值:\t",float(min_change))
print("*"*50)
# step1.可视化手写数字(不再显示散点图)
#Randomly select 12 data points to display
rand_indices = np.random.permutation(m)
written_data = X_for_display[rand_indices[0:100]]
displayDataHandwritten(written_data)
plt.show()
#
# step2.计算代价函数computeCost(theta,X,Y)
def sigmoid(z):
    return 1 / (1 + np.exp(-z))
def computeCost(X, Y, theta):
    m = len(Y)
    z = sigmoid(np.dot(X, theta))
    z[z == 1] = 9e-18
    z[z == 0] = 1e-18
    J = - (np.dot(Y.T, np.log(z)) + np.dot((1 - Y).T, np.log(1 - z))) / m
    return J
# step3.计算梯度computeDescent(theta,X,Y)
def computeDescent(theta, X, Y):
    z = sigmoid(np.dot(X, theta))
    gradient = np.dot(X.T, (z - Y)) / m
    return gradient
# step4.梯度下降函数gradientDescent(theta,X,Y,iter_num,learningRate)
def gradientDescent(theta, X, Y, iter_num, learningRate,tag_num):
    i = 0
    J_list = []
    gradient = computeDescent(theta, X, Y)
    while (not np.all(np.absolute(gradient) <= min_change)) and (i < iter_num):
        theta = theta - learningRate * gradient
        gradient = computeDescent(theta, X, Y)
        J_list.append(computeCost(X, Y, theta))
        i += 1
        print("分类器:",tag_num,"::\t",i)
    return theta, J_list, i
# step5.执行梯度下降函数
theta_all = np.zeros((n + 1, 1))
# 标签
label = (np.unique(Y))
#J_cost 便于绘制代价函数与迭代次数的曲线(对应(n=类别)个曲线)
J_list_all = []
for i in label:
    Y_temp = np.copy(Y)
    pos = np.where(Y_temp == i)
    neg = np.where(Y_temp != i)
    Y_temp[pos] = 1
    Y_temp[neg] = 0
    theta, J_list, iter_num = gradientDescent(theta, X, Y_temp, iter_num, learningRate,i)
    theta_all = np.hstack((theta_all, theta))
    J_list_all.append(J_list)

# step6.(附1:评价)生成代价函数(数值)与迭代次数(iter_num)的曲线(略)
#(J_list_all ,and iter_num)
plt.subplots(figsize=(5,5))
color_list = ["-r","-g","-b"]
for i in range(0,len(J_list_all)):
    color = random.choice(color_list)
    plt.plot(np.arange(1,iter_num+1,1),np.array(J_list_all[i]).reshape(iter_num,1),color)
plt.show()
# step7.计算算法的分类准确度
print("正在预测...")
theta_all = theta_all[:, 1:]
rest_x = np.hstack((np.ones((total-int(total*data_percent), 1)), data[int(total*data_percent):, :-1]))
rest_y = data[int(total*data_percent):, -1].reshape(-1, 1)
predict_y = np.argmax(sigmoid(np.dot(rest_x,theta_all)), axis=1).reshape(-1, 1) + 1
print('\nTraining Set Accuracy: %.2f\n'%(np.mean(np.double(predict_y == rest_y)) * 100))


