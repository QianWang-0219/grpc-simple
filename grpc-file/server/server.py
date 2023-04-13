import sys

import route.route_pb2_grpc

print(sys.path)

from concurrent import futures
import logging
import os
import grpc
from route import route_pb2, route_pb2_grpc

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
                filepath = get_filepath(request.metadata.filename, request.metadata.extension) # 文件流保存的地址
                continue
            data.extend(request.chunk_data)
        print("dddddddd")
        print(filepath)
        with open(filepath, 'wb') as f:
            f.write(data)
        return route_pb2.StringResponse(message='Success!')

    # def DownloadFile(self, request, context):
    #     chunk_size = 1024
    #
    #     filepath = f'{request.filename}{request.extension}'
    #     if os.path.exists(filepath):
    #         with open(filepath, mode="rb") as f:
    #             while True:
    #                 chunk = f.read(chunk_size)
    #                 if chunk:
    #                     entry_response = route_pb2.FileResponse(chunk_data=chunk)
    #                     yield entry_response
    #                 else:  # The chunk was empty, which means we're at the end of the file
    #                     return

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=4))
    #route_pb2_grpc.add_GreeterServicer_to_server(Greeter(), server)
    route.route_pb2_grpc.add_localGuideServicer_to_server(localGuide(),server)
    server.add_insecure_port('[::]:30030')
    server.start()
    server.wait_for_termination()


if __name__ == '__main__':
    logging.basicConfig()
    serve()