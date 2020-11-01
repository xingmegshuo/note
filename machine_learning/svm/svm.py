## Machine Learning Online Class
#  Exercise 6 | Support Vector Machines
## Initialization
import numpy as np
import matplotlib.pyplot as plt
import scipy.io as sio
from plotData import plotData
from linearKernel import linearKernel
from gaussianKernel import gaussianKernel
from svmTrain import svmTrain
from visualizeBoundaryLinear import visualizeBoundaryLinear
from visualizeBoundary import visualizeBoundary
from dataset3Params import dataset3Params
from svmPredict import svmPredict


#step1.显示数据
print('Loading and Visualizing Data ...\n')
# Load from ex6data2:
# You will have X, y in your environment
data = sio.loadmat('ex6data3.mat')
X = data['X']; y = data['y']
# Plot training data
plotData(X, y.T[0])
plt.show()
#step2.加载数据
Xval = data['Xval']; yval = data['yval']
y = y.astype(int)
yval = yval.astype(int)
# Try different SVM Parameters here
#step3.尝试SVM的参数C,sigma
C, sigma = dataset3Params(X, y, Xval, yval)
#step4.训练SVM
# Train the SVM
func = lambda a, b: gaussianKernel(a, b, sigma)
model = svmTrain(X, y, C, func)
visualizeBoundary(X, y, model)
plt.show()
#predict
#step4.预测错误率
predictions = svmPredict(model, Xval)
error = np.mean(np.double(np.not_equal(predictions, yval)))
print("error",error*100,"%")

