from concurrent import futures
import time

import grpc

from hall_pb2 import Message
from hall_pb2_grpc import IronmanServicer, add_IronmanServicer_to_server

_ONE_DAY_IN_SECONDS = 60 * 60 * 24


class Ironman(IronmanServicer):

    def Ping(self, request, context):
        print(request.text)
        return Message(text='Pong')


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    add_IronmanServicer_to_server(Ironman(), server)
    server.add_insecure_port('[::]:50051')
    server.start()
    try:
        while True:
            time.sleep(_ONE_DAY_IN_SECONDS)
    except KeyboardInterrupt:
        server.stop(0)


if __name__ == '__main__':
    serve()
