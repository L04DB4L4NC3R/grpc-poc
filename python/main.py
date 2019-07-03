from __future__ import print_function
import logging

import grpc

from pb import helloworld_pb2
from pb import helloworld_pb2_grpc


def run():
    # NOTE(gRPC Python Team): .close() is possible on a channel and should be
    # used in circumstances in which the with statement does not fit the needs
    # of the code.
    with grpc.insecure_channel('localhost:3000') as channel:
        stub = helloworld_pb2_grpc.HelloServiceStub(channel)
        response = stub.SayHello(helloworld_pb2.Hi(msg='Hi, I am a python client trying to communicate with you', id="sakdbsakdbsakbdasdbasdb"))
    print("Greeter client received: " + response.msg)


if __name__ == '__main__':
    logging.basicConfig()
    run()
