import tensorflow as tf
from tensor_learn1 import *
import matplotlib.pyplot as plt
import random
import numpy as np


images,labels = load_data('Training')
images28 = change_picture_size(images)
images28 = picture_to_rgb(images28)

# initialize placeholders

x = tf.placeholder(dtype=tf.float32,shape=[None,28,28])
y = tf.placeholder(dtype=tf.int32,shape=[None])


# flatten the input data
images_flat = tf.contrib.layers.flatten(x)

# fully connected layer
logits = tf.contrib.layers.fully_connected(images_flat,62,tf.nn.relu)

# define a loss function 
loss = tf.reduce_mean(tf.nn.sparse_softmax_cross_entropy_with_logits(labels=y,logits=logits))

# define an optimizer
train_op = tf.train.AdadeltaOptimizer(learning_rate=0.8).minimize(loss)

# convert logits to label indexes
correct_pred = tf.argmax(logits,1)

# define an accuracy metric
accuracy = tf.reduce_mean(tf.cast(correct_pred,tf.float32))
tf.set_random_seed(1234)
with tf.Session() as sess:
    sess.run(tf.global_variables_initializer())
    for i in range(3501):
        _,loss_value = sess.run([train_op,loss],feed_dict={x:images28,y:labels})
        if i %10 ==0:
            print('Loss',loss_value)

sess = tf.Session()
# sess.run(tf.global_variables_initializer())

# for i in range(201):
#         print('EPOCH', i)
#         _, accuracy_val = sess.run([train_op, loss], feed_dict={x: images28, y: labels})
#         if i % 10 == 0:
#             print("Loss: ", accuracy_val)
#         print('DONE WITH EPOCH')
sess.run(tf.global_variables_initializer())
sample_index = random.sample(range(len(images28)),10)
sample_images = np.array([images28[i] for i in sample_index])
sample_labels = [labels[int(i)] for i in sample_index]

predicted = sess.run([correct_pred],feed_dict={x:sample_images})[0]


fig = plt.figure(figsize=(10,10))
for i in range(len(sample_images)):
    truth = sample_labels[i]
    prediction = predicted[i]
    plt.subplot(5,2,1+i)
    plt.axis('off')
    color ='green' if truth == prediction else 'red'
    plt.text(40,10,'truth:{0}\npredict:{1}'.format(truth,prediction),fontsize = 12,color=color)
    plt.imshow(sample_images[i],cmap='gray')
plt.show()

test_images, test_labels = load_data('Testing')
testimages28 = change_picture_size(test_images)
testimages28 = picture_to_rgb(testimages28)
predicted = sess.run([correct_pred], feed_dict={x: testimages28})[0]
match_count = sum([int(y == y_) for y, y_ in zip(test_labels, predicted)])
accuracy = match_count / len(test_labels)
print("Accuracy: {:.3f}".format(accuracy))

