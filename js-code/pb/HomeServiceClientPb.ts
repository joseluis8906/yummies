/**
 * @fileoverview gRPC-Web generated client stub for yummies
 * @enhanceable
 * @public
 */

// Code generated by protoc-gen-grpc-web. DO NOT EDIT.
// versions:
// 	protoc-gen-grpc-web v1.5.0
// 	protoc              v5.26.1
// source: home.proto


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as home_pb from './home_pb'; // proto import: "home.proto"


export class HomeServiceClient {
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

  methodDescriptorHome = new grpcWeb.MethodDescriptor(
    '/yummies.HomeService/Home',
    grpcWeb.MethodType.UNARY,
    home_pb.HomeRequest,
    home_pb.HomeResponse,
    (request: home_pb.HomeRequest) => {
      return request.serializeBinary();
    },
    home_pb.HomeResponse.deserializeBinary
  );

  home(
    request: home_pb.HomeRequest,
    metadata?: grpcWeb.Metadata | null): Promise<home_pb.HomeResponse>;

  home(
    request: home_pb.HomeRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: home_pb.HomeResponse) => void): grpcWeb.ClientReadableStream<home_pb.HomeResponse>;

  home(
    request: home_pb.HomeRequest,
    metadata?: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: home_pb.HomeResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/yummies.HomeService/Home',
        request,
        metadata || {},
        this.methodDescriptorHome,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/yummies.HomeService/Home',
    request,
    metadata || {},
    this.methodDescriptorHome);
  }

}
