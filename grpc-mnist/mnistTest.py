'''
使用TensorFlow中的keras库，实现Minist手写数字识别。
编写程序，尝试调整神经网络的层数、节点个数以及优化器等参数，记录并分析实验结果。
要求：
(1)编写程序
(2)记录参数的调整过程和结果,从准确率、时间效率等方面对实验结果进行分析
(3)实现对自制手写数字的识别
'''
#导入库
import imageio
import tensorflow as tf
tf.__version__,tf.keras.__version__
import numpy as np
import matplotlib.pyplot as plt
gpus=tf.config.experimental.list_physical_devices('GPU')
tf.config.experimental.set_memory_growth(gpus[0],True)

def mnist(image):
    #加载数据
    mnist=tf.keras.datasets.mnist
    (train_x,train_y),(test_x,test_y)=mnist.load_data()

    '''
    X_train=train_x.reshape((60000,28*28))
    X_test=test_x.reshape((10000,28*28))
    '''
    #数据预处理
    X_train,X_test=tf.cast(train_x/255.0,tf.float32),tf.cast(test_x/255.0,tf.float32)
    y_train,y_test=tf.cast(train_y,tf.int16),tf.cast(test_y,tf.int16)

    #建立模型
    model=tf.keras.Sequential()
    model.add(tf.keras.layers.Flatten(input_shape=(28,28)))
    model.add(tf.keras.layers.Dense(128,activation="relu"))
    model.add(tf.keras.layers.Dense(10,activation="softmax"))
    model.summary()

    #配置训练方法
    model.compile(optimizer='adam',
                  loss='sparse_categorical_crossentropy',
                  metrics=['sparse_categorical_accuracy'])
    #训练模型
    model.fit(X_train,y_train,batch_size=20,epochs=5)
    #评估模型

    model.evaluate(X_test,y_test,verbose=2)
    '''
    image = cv2.imread('1.png', 0)
    img = cv2.imread('1.png', 0)
    
    img=np.resize(img,(28,28))
    #np.argmax(model.predict(img),cmap='gray')
    img = (img.reshape(28,28)).astype("float32")/255
    
    plt.imshow(img)
    plt.show()
    predict = model.predict(img)
    print ('识别为：')
    print (predict)
    '''
    #image='7.PNG'
    img_array = imageio.imread(image, as_gray=False)
    # 调整为28*28
    img_array = np.resize(img_array,(28,28))
    img = tf.cast(abs(img_array/255.0-1), tf.float32)
    demo=tf.reshape(img,(1,28,28))

    y_pred = np.argmax(model.predict(demo))
    print("识别结果：", y_pred)
    return str(y_pred)

# if __name__=='__main__':
#     image = '7.PNG'
#     mnist(image)