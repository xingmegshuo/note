# 神经网络算法实现手写识别(包括训练与预测两个部分)
import numpy as np
import matplotlib.pyplot as plt
import scipy.io as sio
DEBUG = 1   #测试开关,1:测试 0:运行
if_use_ex4weight = 0  #是否使用作业的权重1:使用 0:不使用

## Setup the parameters you will use for this exercise
input_layer_size  = 400   # 20x20 Input Images of Digits
hidden_layer_size = 25    # 25 hidden units
num_labels = 10           # 10 labels, from 1 to 10
                          # (note that we have mapped "0" to label 10)
iter_num =30           #iter num
learning_rate = 0.01      # learning rate
_lambda = 1               # Weight regularization parameter (we set this to 0 here).
#****************************************
print("*"*50)
print("输入层:\t",input_layer_size)
print("隐藏层:\t",hidden_layer_size)
print("输出层:\t",int(num_labels))

print("迭代次数:\t",iter_num)
print("学习率:\t",learning_rate)
print("_lambda:\t",_lambda)
print("*"*50)
#****************************************
#　一、训练神经网络
# step1. 加载数据
print('Loading Data ...\n')
data = sio.loadmat('ex4data1.mat')
X = data['X']; y = data['y']%10
m = X.shape[0]          #total number of train set
if DEBUG == 1:
    print("X.shape",X.shape)
    print("y.shape",y.shape)
    print("m",m)
# step2. 显示数据（略）
# step3. 参数的随机初始化
def inintilizeThetaWeights(L_in,L_out):
    #Theta(l),L_in = s(l),L_out=s(l+1)
    epsilon = np.sqrt(6)/np.sqrt(L_in+L_out)
    return np.random.rand(L_out,L_in) * 2 * epsilon - epsilon
#Theta1's shape should be (25*401)
#Theta2's shape should be (10*26)
Theta1_L_in = 401
Theta1_L_out = 25
Theta2_L_in = 26
Theta2_L_out = 10
#随机初始化
#注:迭代次数要很多,直接用了作业的权重
if if_use_ex4weight == 0:
    Theta1 = inintilizeThetaWeights(Theta1_L_in,Theta1_L_out)
    Theta2 = inintilizeThetaWeights(Theta2_L_in,Theta2_L_out)
else:
    #读取作业theta1,theta2
    data_weights = sio.loadmat('ex4weights.mat')
    Theta1 = data_weights['Theta1']
    Theta2 = data_weights['Theta2']
if DEBUG == 1:
    print(Theta1.shape)
    print(Theta2.shape)
    print(Theta2)
# step4. 利用正向传播方法计算代价函数(正则化可略)
    input('pasued')
def sigmoid(z):
    return 1 / (1 + np.exp(-z))
def sigmoidGradient(z):
    return sigmoid(z) * (1 - sigmoid(z))
def nnCostFunction(X,Theta1,Theta2,y):
    a1 = np.vstack((np.ones(m), X.T)).T
    a2 = sigmoid(np.dot(a1, Theta1.T))
    a2 = np.vstack((np.ones(m), a2.T)).T
    a3 = sigmoid(np.dot(a2, Theta2.T))
    y = np.tile((np.arange(num_labels)+1)%10,(m,1)) == np.tile(y,(1,num_labels))
    #J = -np.sum( y * np.log(a3) + (1-y) * np.log(1-a3) ) / m
    regTheta1 = Theta1[:,1:]
    regTheta2 = Theta2[:,1:]

    J = -np.sum( y * np.log(a3) + (1-y) * np.log(1-a3) ) / m + \
        _lambda * np.sum(regTheta1*regTheta1) / m/2 + \
        _lambda * np.sum(regTheta2*regTheta2) / m/2

    #因计算delta2,delta3,要用a1,a2,a3的结果,故一并计算
    delta2 = np.zeros(Theta1.shape)
    delta3 = np.zeros(Theta2.shape)
    for i in range(m):
        a1_ = a1[i];a2_ = a2[i];a3_ = a3[i]
        d3 = a3_ - y[i]; #输出层误差
        d2 = np.dot(d3, Theta2) * sigmoidGradient(np.append(1, np.dot(a1_, Theta1.T)))
        delta2 = delta2 + np.dot(d2[1:].reshape(-1,1),a1_.reshape(1,-1));
        delta3 = delta3 + np.dot(d3.reshape(-1,1), a2_.reshape(1,-1))
        Theta1_grad = np.zeros(Theta1.shape)
        Theta2_grad = np.zeros(Theta2.shape)
        Theta1_grad[:,0] = delta2[:,0] / m
        Theta2_grad[:,0] = delta3[:,0] / m
        Theta1_grad[:,1:] = delta2[:,1:] / m + _lambda * regTheta1 / m
        Theta2_grad[:,1:] = delta3[:,1:] / m + _lambda * regTheta2 / m

    return J,Theta1_grad,Theta2_grad
# step5. (见nnCostFunction,一并计算)利用反向传播方法计算
# 所有偏导数("梯度")(正则化可略)
# step6. 利用数值检验方法检验偏导数(略,参ml-ex4/computeZNumericalGradient.py)
#-----------------------------------------
# step7. 梯度下降函数
def gradientDescent(Theta1,Theta2):
    i = 0
    J_list = []
    J, Theta1_grad, Theta2_grad = nnCostFunction(X,Theta1,Theta2,y)
    while  i < iter_num:
        Theta1 = Theta1 - learning_rate * Theta1_grad
        Theta2 = Theta2 - learning_rate * Theta2_grad
        J, Theta1_grad, Theta2_grad = nnCostFunction(X,Theta1,Theta2,y)
        J_list.append(J)
        i += 1
        print("迭代:",i,"->","J:",J)
    return J_list
# step8. 执行梯度下降函数
'''
注:在ex4.py中,使用优化算法
res = minimize(costFunction, initial_nn_params, method='CG',
jac=True, options={'maxiter': 200})代替step7、8
'''
J_list = gradientDescent(Theta1,Theta2)
J_list = np.asarray(J_list).reshape(1,iter_num)[0]

#-----------------------------------------
# step9.(附1:评价)打印迭代次数与代价函数曲线
plt.plot(np.arange(0,iter_num,1),np.array(J_list).reshape(iter_num,1),'-r')
plt.show()
#　二、预测
# step10. 计算算法的分类准确度
def predict(Theta1, Theta2, X):
    # Useful values
    a1 = np.vstack((np.ones(m), X.T)).T
    a2 = sigmoid(np.dot(a1, Theta1.T))
    a2 = np.vstack((np.ones(m), a2.T)).T
    a3 = sigmoid(np.dot(a2, Theta2.T))
    return np.argmax(a3, axis=1)

pred = (predict(Theta1, Theta2, X)+1)%10
print('\nTraining Set Accuracy: %f\n'%(np.mean(np.double(pred == y.T)) * 100))


