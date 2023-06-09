# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc


import route_pb2 as route__pb2


class localGuideStub(object):
    """a simple bidirectional streaming example
    client发送文件存储的iniLoc给server，server返回一个ok
    server根据iniLoc完成二维拼接任务，发送finLoc给client，client回复ok

    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.GetLocation = channel.stream_stream(
                '/route.localGuide/GetLocation',
                request_serializer=route__pb2.IniLoc.SerializeToString,
                response_deserializer=route__pb2.FinLoc.FromString,
                )


class localGuideServicer(object):
    """a simple bidirectional streaming example
    client发送文件存储的iniLoc给server，server返回一个ok
    server根据iniLoc完成二维拼接任务，发送finLoc给client，client回复ok

    """

    def GetLocation(self, request_iterator, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_localGuideServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'GetLocation': grpc.stream_stream_rpc_method_handler(
                    servicer.GetLocation,
                    request_deserializer=route__pb2.IniLoc.FromString,
                    response_serializer=route__pb2.FinLoc.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'route.localGuide', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class localGuide(object):
    """a simple bidirectional streaming example
    client发送文件存储的iniLoc给server，server返回一个ok
    server根据iniLoc完成二维拼接任务，发送finLoc给client，client回复ok

    """

    @staticmethod
    def GetLocation(request_iterator,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.stream_stream(request_iterator, target, '/route.localGuide/GetLocation',
            route__pb2.IniLoc.SerializeToString,
            route__pb2.FinLoc.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
