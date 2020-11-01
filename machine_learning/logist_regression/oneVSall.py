import numpy as np
import matplotlib.pyplot as plt
import scipy.io as scio


# question: the value is encountered   it's wrong 
# algorithm find where is y learan X get theta use this theta predict y 
def sigmoid(z):
    return 1/(1+np.exp(-z))


def plot_cost(cost_list,count):
    plt.figure(figsize=(10,10))
    plt.plot([i for i in range(count)],[i for i in cost_list],c='r')
    plt.show()

def load_data(filename,present):
    data = np.hstack((scio.loadmat(filename)['X'],scio.loadmat(filename)['y']))
    m = int(len(data)*present)
    # np.random.shuffle(data)
    X = np.mat(data[:m,:-1])
    Y = data[:m,-1]
    x0 = np.ones((m,1))
    predict_x = data[m:,:-1]
    predict_y = data[m:,:-1]
    # predict_x0 = np.ones((len(data)-m,1))
    return np.hstack((x0,X)),Y,predict_x,predict_y

def costFunction(theta,X,Y):
    s = sigmoid(np.dot(X,theta))
    s = s.astype(np.float)#因为log内不能为0或1
    s[s == 1] = 9e-18
    s[s == 0] = 1e-18
    j = -(np.dot(Y.T,np.log(s))+np.dot((1-Y).T,np.log(1-s)))/len(Y)
    
    return j.tolist()[0][0]

def gradientDescent(theta,X,Y):
    s = sigmoid(np.dot(X,theta))
    s = s.astype(np.float)#因为log内不能为0或1
    s[s == 1] = 9e-18
    s[s == 0] = 1e-18

    theta_gradient = np.dot(X.T,(s-Y))/len(Y)
    return theta_gradient

def gradient_descent(alpha,X,Y,cycles,finally_value,initial_value):
    theta_list = []
    for i in range(1,11):
        x_ = np.zeros(shape=(0,X.shape[1]))
        y_ = np.zeros(shape=(0,1))
        theta = np.array(initial_value).reshape(X.shape[1],1)
        cost_list = []
        count = 0
        for j in np.where(Y==i)[0]:
            x_ = np.vstack((x_,X[j]))
            y_ = np.vstack((y_,Y[j]))
        gradient = gradientDescent(theta,x_,y_)
        while not np.all(np.absolute(gradient) >= finally_value) and count <= cycles:
            theta = theta - alpha*gradient
            gradient = gradientDescent(theta,x_,y_)
            cost_list.append(costFunction(theta,x_,y_))
            count +=1
        # plot_cost(cost_list,count)
       
        theta_list.append(theta)
    return theta_list
    


def predictOneVsAll(all_theta, X):
    
    m = X.shape[0] # Number of training examples

    # Add ones to the X data matrix
    X = np.vstack((np.ones(m), X.T)).T

    return np.argmax(sigmoid(np.dot(all_theta.T, X.T)), axis=0)

if __name__ == '__main__':
    X,Y,predict_x,predict_y = load_data('ex3data1.mat',0.6)
    alpha = 0.5
    cycles = 1000
    finally_value = 1e-5
    initial_value = np.ones((X.shape[1],1))
    theta_list = gradient_descent(alpha,X,Y,cycles,finally_value,initial_value)
    all_theta = np.zeros((X.shape[1],1))
    for i in theta_list:
        all_theta = np.hstack((all_theta,i.reshape(-1,1)))
    pred = predictOneVsAll(all_theta, predict_x)
    print('\nTraining Set Accuracy: %f\n'%(np.mean(np.double(pred == predict_y.T)) * 100))cz1