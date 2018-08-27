// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var proto_test_pb = require('../proto/test_pb.js');

function serialize_proto_HelloRequest(arg) {
  if (!(arg instanceof proto_test_pb.HelloRequest)) {
    throw new Error('Expected argument of type proto.HelloRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_proto_HelloRequest(buffer_arg) {
  return proto_test_pb.HelloRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_HelloResponse(arg) {
  if (!(arg instanceof proto_test_pb.HelloResponse)) {
    throw new Error('Expected argument of type proto.HelloResponse');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_proto_HelloResponse(buffer_arg) {
  return proto_test_pb.HelloResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var HelloService = exports.HelloService = {
  sayHello: {
    path: '/proto.Hello/SayHello',
    requestStream: false,
    responseStream: false,
    requestType: proto_test_pb.HelloRequest,
    responseType: proto_test_pb.HelloResponse,
    requestSerialize: serialize_proto_HelloRequest,
    requestDeserialize: deserialize_proto_HelloRequest,
    responseSerialize: serialize_proto_HelloResponse,
    responseDeserialize: deserialize_proto_HelloResponse,
  },
};

exports.HelloClient = grpc.makeGenericClientConstructor(HelloService);
