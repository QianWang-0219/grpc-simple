import logging
import os
import grpc
from route import route_pb2, route_pb2_grpc

def get_filepath(filename, extension):
    return f'{filename}{extension}'

# 创建新的迭代器逐块读取流式请求
# 用于发送流
def read_iterfile(filepath, chunk_size=1024):
    split_data = os.path.splitext(filepath)
    filename = split_data[0]
    extension = split_data[1]

    metadata = route_pb2.MetaData(filename=filename, extension=extension)
    yield route_pb2.UploadFileRequest(metadata=metadata)

    with open(filepath, mode="rb") as f:
        while True:
            chunk = f.read(chunk_size)
            if chunk:
                entry_request = route_pb2.UploadFileRequest(chunk_data=chunk)
                yield entry_request
            else:  # The chunk was empty, which means we're at the end of the file
                return

def run():
    with grpc.insecure_channel('localhost:30030') as channel:
        stub = route_pb2_grpc.localGuideStub(channel)

        # 发送流
        response = stub.UploadFile(read_iterfile('/Users/wang_qian0219/code/python/grpc-file/resource/sum.txt'))
        print("client received: " + response.message)


if __name__ == '__main__':
    logging.basicConfig()
    run()