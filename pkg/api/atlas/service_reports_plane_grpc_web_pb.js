/**
 * @fileoverview gRPC-Web generated client stub for atlas
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');


var objects_request_pb = require('./objects_request_pb.js')

var objects_list_pb = require('./objects_list_pb.js')
const proto = {};
proto.atlas = require('./service_reports_plane_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.atlas.ReportsPlaneClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.atlas.ReportsPlanePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.atlas.ObjectsRequest,
 *   !proto.atlas.ObjectsList>}
 */
const methodDescriptor_ReportsPlane_ObjectsReport = new grpc.web.MethodDescriptor(
  '/atlas.ReportsPlane/ObjectsReport',
  grpc.web.MethodType.UNARY,
  objects_request_pb.ObjectsRequest,
  objects_list_pb.ObjectsList,
  /**
   * @param {!proto.atlas.ObjectsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  objects_list_pb.ObjectsList.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.atlas.ObjectsRequest,
 *   !proto.atlas.ObjectsList>}
 */
const methodInfo_ReportsPlane_ObjectsReport = new grpc.web.AbstractClientBase.MethodInfo(
  objects_list_pb.ObjectsList,
  /**
   * @param {!proto.atlas.ObjectsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  objects_list_pb.ObjectsList.deserializeBinary
);


/**
 * @param {!proto.atlas.ObjectsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.atlas.ObjectsList)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.atlas.ObjectsList>|undefined}
 *     The XHR Node Readable Stream
 */
proto.atlas.ReportsPlaneClient.prototype.objectsReport =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/atlas.ReportsPlane/ObjectsReport',
      request,
      metadata || {},
      methodDescriptor_ReportsPlane_ObjectsReport,
      callback);
};


/**
 * @param {!proto.atlas.ObjectsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.atlas.ObjectsList>}
 *     Promise that resolves to the response
 */
proto.atlas.ReportsPlanePromiseClient.prototype.objectsReport =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/atlas.ReportsPlane/ObjectsReport',
      request,
      metadata || {},
      methodDescriptor_ReportsPlane_ObjectsReport);
};


module.exports = proto.atlas;

