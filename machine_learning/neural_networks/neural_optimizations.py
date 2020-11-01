import numpy as np
import matplotlib.pyplot as plt
from scipy.optimize import minimize 
#　最优化



def sigmoid(z):
    return 1/(1+np.exp(-z))

