from concurrent import futures
import grpc
import logging
import time

#from route import route_pb2, route_pb2_grpc
import route_pb2, route_pb2_grpc
import mnistTest
import decode_image

_ONE_DAY_IN_SECONDS = 60 * 60 * 24
_HOST = 'localhost'
_PORT = '30033'


class NewsService(route_pb2_grpc.localGuideServicer):
    def GetLocation(self, request_iterator, context):
        for request in request_iterator:
            ini_location = request.ini_location
            finLocation = decode_image.decode_image(ini_location)
            # 调用手写数字识别程序
            res = mnistTest.mnist(finLocation)
            yield route_pb2.FinLoc(fin_location='the mnist recognition is: %s' % res)


def serve():
    # 定义服务器并设置最大连接数,corcurrent.futures是一个并发库，类似于线程池的概念
    grpcServer = grpc.server(futures.ThreadPoolExecutor(max_workers=4))  # 创建一个服务器

    route_pb2_grpc.add_localGuideServicer_to_server(NewsService(), grpcServer)  # 在服务器中添加派生的接口服务（自己实现了处理函数）
    grpcServer.add_insecure_port(_HOST + ':' + _PORT)  # 添加监听端口
    grpcServer.start()  # 启动服务器
    try:
        while True:
            time.sleep(_ONE_DAY_IN_SECONDS)
    except KeyboardInterrupt:
        grpcServer.stop(0)  # 关闭服务器

if __name__ == '__main__':
    logging.basicConfig()
    serve()
