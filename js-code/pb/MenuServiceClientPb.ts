/**
 * @fileoverview gRPC-Web generated client stub for yummies
 * @enhanceable
 * @public
 */

// Code generated by protoc-gen-grpc-web. DO NOT EDIT.
// versions:
// 	protoc-gen-grpc-web v1.5.0
// 	protoc              v5.27.0
// source: menu.proto


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as menu_pb from './menu_pb'; // proto import: "menu.proto"


export class MenuServiceClient {
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

  methodDescriptorIndex = new grpcWeb.MethodDescriptor(
    '/yummies.MenuService/Index',
    grpcWeb.MethodType.UNARY,
    menu_pb.MenuIndexRequest,
    menu_pb.MenuIndexResponse,
    (request: menu_pb.MenuIndexRequest) => {
      return request.serializeBinary();
    },
    menu_pb.MenuIndexResponse.deserializeBinary
  );

  index(
    request: menu_pb.MenuIndexRequest,
    metadata?: grpcWeb.Metadata | null): Promise<menu_pb.MenuIndexResponse>;

  index(
    request: menu_pb.MenuIndexRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: menu_pb.MenuIndexResponse) => void): grpcWeb.ClientReadableStream<menu_pb.MenuIndexResponse>;

  index(
    request: menu_pb.MenuIndexRequest,
    metadata?: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: menu_pb.MenuIndexResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/yummies.MenuService/Index',
        request,
        metadata || {},
        this.methodDescriptorIndex,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/yummies.MenuService/Index',
    request,
    metadata || {},
    this.methodDescriptorIndex);
  }

}

