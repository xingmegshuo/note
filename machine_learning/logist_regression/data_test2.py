import logistRegression,polynomial
import numpy as np
import matplotlib.pyplot as plt

def plot_desicion_boundary(x,Y,theta):
    plt.subplots(figsize=(10,10))
    pos = np.where(Y == 1)
    neg = np.where(Y == 0)
    plt.scatter(x[pos, 0], x[pos, 1], marker='s', c='r')
    plt.scatter(x[neg, 0], x[neg, 1], marker='x', c='b')
    u = np.linspace(-1, 1,50)
    v = np.linspace(-1, 1, 50)
    z = np.zeros((u.shape[0], v.shape[0]))
    for i in range(0,u.shape[0]):
        for j in range(0,v.shape[0]):
            z[i,j] =np.dot(theta.T,polynomial.polynomial(np.hstack((np.mat(1),np.mat(u[i]),np.mat(v[j]))),4).reshape(8,1))
            
    # plt.plot(x_a,y,c='r')
    # plt.show()
    u, v = np.meshgrid(u, v)
    plt.contour(u, v, z.T, (0,), colors='g', linewidths=2)
    plt.show()




if __name__ == '__main__':
    filename = 'ex2data2.txt'
    present = 0.6
    X,Y,x,predict_x,predict_y = logistRegression.loadData(filename,present)
    logistRegression.plot_data(x,Y)
    alpha = 0.05
    cycles = 200000
    finally_value = 1e-5
    X = polynomial.polynomial(X,4)
    initial_value = np.ones((X.shape[1],1))
    theta,cost_list,count = logistRegression.gradient_descent(alpha,X,Y,cycles,finally_value,initial_value)
    logistRegression.plot_cost(cost_list,count)   
    predict_x = polynomial.polynomial(predict_x,4)
    predict_data = logistRegression.sigmoid(np.dot(predict_x,theta))
    predict_data[predict_data<0.5] = 0
    predict_data[predict_data>=0.5] = 1
    result = predict_data - predict_y
    print(len(result[result==0])/len(result))
    plot_desicion_boundary(x,Y,theta)

