from __future__ import print_function

import grpc

from hall_pb2 import Session
from hall_pb2_grpc import BatmanStub


def run():
    channel = grpc.insecure_channel('localhost:50051')
    stub = BatmanStub(channel)
    response = stub.IronmanCallback(Session(id=5, status='BLEH'))


if __name__ == '__main__':
    run()
