import sys

import route.route_pb2_grpc

# print(sys.path)

from concurrent import futures
import logging
import os
import grpc
import IPy
from route import route_pb2, route_pb2_grpc


PORT = '30030'
chunk_size = 1024

def is_ip(address):
    try:
        IPy.IP(address)
        return True
    except Exception as e:
        return False

def get_filepath(filename, extension):
    return f'{filename}{extension}'


class localGuide(route_pb2_grpc.localGuideServicer):
    # def SayHello(self, request, context):
    #     return route_pb2.StringResponse(message=f'Hello, {request.name}! Your age is {request.age}')

    def UploadFile(self, request_iterator, context):
        data = bytearray()
        filepath = 'dummy'

        for request in request_iterator:
            if request.metadata.filename and request.metadata.extension:
                filepath = '../upload_file/' + get_filepath(request.metadata.filename, request.metadata.extension) # 文件流保存的地址
                continue
            data.extend(request.chunk_data)
        # print("dddddddd")
        print(filepath)
        with open(filepath, 'wb') as f:
            f.write(data)
        return route_pb2.StringResponse(message=f'恭喜您成功上传文件{request.metadata.filename}到服务器目录{filepath}下!')

    def DownloadFile(self, request, context):
        filepath = f'../upload_file/{request.filename}{request.extension}'
        print("下载：",filepath,"中的文件")
        if os.path.exists(filepath):
            with open(filepath, mode="rb") as f:
                while True:
                    chunk = f.read(chunk_size)
                    if chunk:
                        entry_response = route_pb2.FileResponse(chunk_data=chunk)
                        yield entry_response
                    else:  # The chunk was empty, which means we're at the end of the file
                        return


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=4))
    route.route_pb2_grpc.add_localGuideServicer_to_server(localGuide(),server)
    # ipAddr = input()
    # if is_ip(ipAddr):
    #     task3_ipPort = ipAddr + ':' + PORT
    # else:
    #     return
    task_ipPort = 'localhost' + ':' +PORT

    server.add_insecure_port(task_ipPort)
    server.start()
    server.wait_for_termination()


if __name__ == '__main__':
    logging.basicConfig()
    serve()