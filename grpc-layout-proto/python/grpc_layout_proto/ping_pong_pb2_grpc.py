# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

from grpc_layout_proto import ping_pong_pb2 as grpc__layout__proto_dot_ping__pong__pb2


class PingPongServiceStub(object):
  # missing associated documentation comment in .proto file
  pass

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.Call = channel.unary_unary(
        '/ping_pong.PingPongService/Call',
        request_serializer=grpc__layout__proto_dot_ping__pong__pb2.RequestMsg.SerializeToString,
        response_deserializer=grpc__layout__proto_dot_ping__pong__pb2.ResponseMsg.FromString,
        )


class PingPongServiceServicer(object):
  # missing associated documentation comment in .proto file
  pass

  def Call(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_PingPongServiceServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'Call': grpc.unary_unary_rpc_method_handler(
          servicer.Call,
          request_deserializer=grpc__layout__proto_dot_ping__pong__pb2.RequestMsg.FromString,
          response_serializer=grpc__layout__proto_dot_ping__pong__pb2.ResponseMsg.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'ping_pong.PingPongService', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))