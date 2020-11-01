import os
import numpy as np
import matplotlib.pyplot as plt
import skimage


def load_data(data_directory):
    directories = [d for d in os.listdir(data_directory) 
                   if os.path.isdir(os.path.join(data_directory, d))]
    labels = []
    images = []
    for d in directories:
        label_directory = os.path.join(data_directory, d)
        file_names = [os.path.join(label_directory, f) 
                      for f in os.listdir(label_directory) 
                      if f.endswith(".ppm")]
        for f in file_names:
            images.append(skimage.data.imread(f))
            labels.append(int(d))
    return images, labels

def show_all(images,labels):
    unique_labels = set(labels)
    plt.figure(figsize=(15,15))
    i = 1
    for label in unique_labels:
        image = images[labels.index(label)]
        plt.subplot(8,8,i)
        plt.axis('off')
        plt.title('label{0}({1})'.format(label,labels.count(label)))
        i += 1
        plt.imshow(image)
    plt.show()

def change_picture_size(images):
    images28 = [skimage.transform.resize(image,(28,28)) for image in images]
    return images28


def plt_picture(images):
    traffic_signs = [300, 2250, 3650, 4000]
    for i in range(len(traffic_signs)):
        plt.subplot(1, 4, i+1)
        plt.axis('off')
        plt.imshow(images[traffic_signs[i]], cmap="gray")
        plt.subplots_adjust(wspace=0.5)
    
    # Show the plot
    plt.show()

def picture_to_rgb(images):
    images = np.array(images)
    images = skimage.color.rgb2gray(images)
    return images

if __name__ == '__main__':
    images,labels = load_data('Training')
    # plt.hist(labels,62)
    # plt.show()
    # show_all(images,labels)
    images28 = change_picture_size(images)
    images28 = picture_to_rgb(images28)
    print(np.array(images28).reshape(1,-1).shape,np.array(labels).shape)

    labels = np.array(labels)
    print(labels.reshape(1,-1).shape)
    print(images28.shape)
    # plt_picture(images28)
