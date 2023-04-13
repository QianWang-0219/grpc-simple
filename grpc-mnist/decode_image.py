# -*- coding: utf-8 -*-
import sys
print(sys.path)


import base64
import re
import uuid


def decode_image(src):
    # 1、信息提取
    result = re.search("data:image/(?P<ext>.*?);base64,(?P<data>.*)", src, re.DOTALL)
    if result:
        ext = result.groupdict().get("ext")
        data = result.groupdict().get("data")

    else:
        raise Exception("Do not parse!")

    # 2、base64解码
    img = base64.urlsafe_b64decode(data)

    # 3、二进制文件保存
    filename = "{}.{}".format(uuid.uuid4(), ext)
    with open(filename, "wb") as f:
        f.write(img)

    return filename # 返回解码后图像的地址


def encode_image(filename):
    # 1、文件读取
    ext = filename.split(".")[-1]

    with open(filename, "rb") as f:
        img = f.read()

    # 2、base64编码
    data = base64.b64encode(img).decode()

    # 3、图片编码字符串拼接
    src = "data:image/{ext};base64,{data}".format(ext=ext, data=data)
    return src


if __name__ == '__main__':

    # 编码测试
    print(encode_image("grpc-simple-stream/blockchain.PNG"))
    src = encode_image("grpc-simple-stream/blockchain.PNG")

    # 解码测试
    print(decode_image(src))
