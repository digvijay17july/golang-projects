/**
 * @fileoverview gRPC-Web generated client stub for proto
 * @enhanceable
 * @public
 */

// Code generated by protoc-gen-grpc-web. DO NOT EDIT.
// versions:
// 	protoc-gen-grpc-web v1.4.2
// 	protoc              v4.25.1
// source: userInfo.proto


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as userInfo_pb from './userInfo_pb';


export class UsrClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: any; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'text';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname.replace(/\/+$/, '');
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodDescriptorGetUser = new grpcWeb.MethodDescriptor(
    '/proto.Usr/GetUser',
    grpcWeb.MethodType.UNARY,
    userInfo_pb.UserRequest,
    userInfo_pb.UserResponse,
    (request: userInfo_pb.UserRequest) => {
      return request.serializeBinary();
    },
    userInfo_pb.UserResponse.deserializeBinary
  );

  getUser(
    request: userInfo_pb.UserRequest,
    metadata: grpcWeb.Metadata | null): Promise<userInfo_pb.UserResponse>;

  getUser(
    request: userInfo_pb.UserRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: userInfo_pb.UserResponse) => void): grpcWeb.ClientReadableStream<userInfo_pb.UserResponse>;

  getUser(
    request: userInfo_pb.UserRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: userInfo_pb.UserResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/proto.Usr/GetUser',
        request,
        metadata || {},
        this.methodDescriptorGetUser,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/proto.Usr/GetUser',
    request,
    metadata || {},
    this.methodDescriptorGetUser);
  }

}

