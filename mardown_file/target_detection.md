### 目标检测图片截取




#### 通过yolo算法检测图片物体
原图片![image](/home/xms/桌面/target_detection/mmexport1554373172138.jpg)

检测后的图片![image](/home/xms/darknet/predictions.jpg)

截取的图片

bed![image](/home/xms/darknet/python/demo/bed.jpg)

pottedplant![image](/home/xms/darknet/python/demo/pottedplant.jpg)

vase![image](/home/xms/darknet/python/demo/vase.jpg)

> 目标检测方法有多种，这里示例yolo算法来进行目标检测，使用已经训练好的权重来对图片进行检测，如果想要检测不在权重文件里的物体需要自定义训练集来进行训练。