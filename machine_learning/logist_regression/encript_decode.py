#!/usr/bin/env python
# -*- coding: utf-8 -*-

# 使用矩阵计算来进行加密解密

import numpy as np
#定义一个秘钥
KEY = np.mat([[1,3,3],[66,7,3],[4,0,31]])
letter_dict = {i:chr(i) for i in range(97,123)}

def judging_charcter(charcter):
    num=''
    # 如果是汉字
    if '\u4e00' <= charcter <='\u9fff':
        num=ord(charcter)
    # 如果是英文字母
    elif ('\u0041'<=charcter<='u005a') or ('\u0061'<=charcter<='\u007a'):
        num=[k for k,v in letter_dict.items() if v==charcter][0]
    elif '\u0030' <= charcter <= '\u0039':
        num=charcter
    else:
        num=-1
    return num 

def encript(mes):
    data_mes=list(mes)
    if len(data_mes)%3!=0:
        data_mes.append('-1')
    n=int(len(data_mes)/3)
    data_array=[[int(judging_charcter(data_mes.pop(0))) for j in range(3)]  for i in range(n)]
    message=np.mat(data_array)
    return message*KEY
  
def decode_a(mes):
    data_mes = np.mat(mes)
    old_mes = data_mes*KEY.I
    old_al=np.matrix.tolist(old_mes)
    that_list=[]
    old_list=[[that_list.append(round(i)) for i in j]  for j in old_al]
    new_str=''
    for i in range(len(that_list)-1,0,-1):
        if that_list[i] <0:
            that_list.remove(that_list[i])
    for i in that_list:
        if i>=10:
            one=chr(i)
        else:
            one=str(i)
        new_str+=one
    return new_str

  

# 待加密信息
pub_message='我是zbc傻逼1314去你妈l'
print(encript(pub_message))
print(decode_a(encript(pub_message)))
