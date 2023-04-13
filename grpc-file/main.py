import os

filepath = "/Users/wang_qian0219/code/python/grpc-file/resource/sum.txt"
split_data = os.path.splitext(filepath)
filename = split_data[0]
extension = split_data[1]
print(filename)
print(extension)
