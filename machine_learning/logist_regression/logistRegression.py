# 逻辑回归分类,梯度下降
import numpy as np
import matplotlib.pyplot as plt


# 读取数据函数

def loadData(filename,present):
    data = np.loadtxt(filename, delimiter = ',')
    np.random.shuffle(data)
    m = int(len(data)*present)
    x0 = np.ones((m,1))
    x = data[:m,:-1]
    Y = data[:m,-1]
    predict_x0 = np.ones((len(data)-m,1))
    predict_x = data[m:,:-1]
    predict_y = data[m:,-1].reshape(-1,1)
    return np.hstack((x0,x)),Y,x,np.hstack((predict_x0,predict_x)),predict_y


# sigmoid 函数 y = g(theta.T*X)

def sigmoid(z):
    return 1/(1+np.exp(-z))


# 代价函数,计算代价值 cost(h(theta)(x),y) = -1/m*(y*log(h(theta)(x))-(1-y)*log(1-h(theta(x))))
# h(theta(x))=g(theta.T*x)

def costFunction(theta,X,Y):
    s = sigmoid(np.dot(X,theta))
    s = s.astype(np.float)#因为log内不能为0或1
    s[s == 1] = 9e-18
    s[s == 0] = 1e-18
    j = -(np.dot(Y.T,np.log(s))+np.dot((1-Y).T,np.log(1-s)))/len(Y)
    return j[0]


# 梯度下降函数 (theta(x)-y)*x/m

def gradientDescent(theta,X,Y):
    s = sigmoid(np.dot(X,theta))
    theta_gradient = np.dot(X.T,(s-Y.reshape(len(Y),1)))/len(Y)
    return theta_gradient


# 寻找最低点或局部最优，来获得合适的theta值

def gradient_descent(alpha,X,Y,cycles,finally_value,initial_value):
    theta = np.array(initial_value).reshape(X.shape[1],1)
    gradient = gradientDescent(theta,X,Y)
    cost_list = []
    count = 0
    while not np.all(np.absolute(gradient) <= finally_value) and count <= cycles:
        theta =  theta - alpha*gradient
        gradient = gradientDescent(theta,X,Y)
        cost_list.append(costFunction(theta,X,Y))
        count += 1
    return theta,cost_list,count 

# 打印散点图
def plot_data(x,Y):
    pos = np.where(Y == 1)
    neg = np.where(Y == 0)
    plt.scatter(x[pos, 0], x[pos, 1], marker='s', c='r')
    plt.scatter(x[neg, 0], x[neg, 1], marker='x', c='b')
    plt.show()


# 打印代价函数　图
def plot_cost(cost_list,count):
    plt.figure(figsize=(10,10))
    plt.plot([i for i in range(count)],[i for i in cost_list],c='r')
    plt.show()


# 绘出判定边界
def decision_boundary(x,Y,theta):
    plt.subplots(figsize=(10,10))
    pos = np.where(Y == 1)
    neg = np.where(Y == 0)
    plt.scatter(x[pos, 0], x[pos, 1], marker='s', c='r')
    plt.scatter(x[neg, 0], x[neg, 1], marker='x', c='b')
    x_a = list(np.linspace(1,100,101))
    y=[]
    for i in range(len(x_a)):
        y.append((-theta[0][0]-theta[1][0]*x_a[i])/theta[2][0])
    plt.plot(x_a,y,c='r')
    plt.show()


# 主函数
if __name__ == '__main__':
    filename = 'ex2data1.txt'
    present = 0.6
    alpha = 0.0005
    finally_value = 1e-5
    cycles = 1000000 
    initial_value = [1,1,1]
    X,Y,x,predict_x,predict_y = loadData(filename,present)
    plot_data(x,Y)
    theta,cost_list,count = gradient_descent(alpha,X,Y,cycles,finally_value,initial_value)
    print(theta)
    plot_cost(cost_list,count)
    decision_boundary(x,Y,theta)
    predict_data = sigmoid(np.dot(predict_x,theta))
    predict_data[predict_data<0.5] = 0
    predict_data[predict_data>=0.5] = 1
    result = predict_data - predict_y
    print(len(result[result==0])/len(result))