// source: object_request.proto
/**
 * @fileoverview
 * @enhanceable
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!

var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();

var domain_pb = require('./domain_pb.js');
goog.object.extend(proto, domain_pb);
var address_pb = require('./address_pb.js');
goog.object.extend(proto, address_pb);
goog.exportSymbol('proto.atlas.ObjectRequest', null, global);
goog.exportSymbol('proto.atlas.ObjectRequest.RequestDomainOptionalCase', null, global);
goog.exportSymbol('proto.atlas.ObjectRequest.ResultDomainOptionalCase', null, global);
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.atlas.ObjectRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, proto.atlas.ObjectRequest.oneofGroups_);
};
goog.inherits(proto.atlas.ObjectRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.atlas.ObjectRequest.displayName = 'proto.atlas.ObjectRequest';
}

/**
 * Oneof group definitions for this message. Each group defines the field
 * numbers belonging to that group. When of these fields' value is set, all
 * other fields in the group are cleared. During deserialization, if multiple
 * fields are encountered for a group, only the last value seen will be kept.
 * @private {!Array<!Array<number>>}
 * @const
 */
proto.atlas.ObjectRequest.oneofGroups_ = [[100],[200]];

/**
 * @enum {number}
 */
proto.atlas.ObjectRequest.RequestDomainOptionalCase = {
  REQUEST_DOMAIN_OPTIONAL_NOT_SET: 0,
  REQUEST_DOMAIN: 100
};

/**
 * @return {proto.atlas.ObjectRequest.RequestDomainOptionalCase}
 */
proto.atlas.ObjectRequest.prototype.getRequestDomainOptionalCase = function() {
  return /** @type {proto.atlas.ObjectRequest.RequestDomainOptionalCase} */(jspb.Message.computeOneofCase(this, proto.atlas.ObjectRequest.oneofGroups_[0]));
};

/**
 * @enum {number}
 */
proto.atlas.ObjectRequest.ResultDomainOptionalCase = {
  RESULT_DOMAIN_OPTIONAL_NOT_SET: 0,
  RESULT_DOMAIN: 200
};

/**
 * @return {proto.atlas.ObjectRequest.ResultDomainOptionalCase}
 */
proto.atlas.ObjectRequest.prototype.getResultDomainOptionalCase = function() {
  return /** @type {proto.atlas.ObjectRequest.ResultDomainOptionalCase} */(jspb.Message.computeOneofCase(this, proto.atlas.ObjectRequest.oneofGroups_[1]));
};



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.atlas.ObjectRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.atlas.ObjectRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.atlas.ObjectRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.atlas.ObjectRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    requestDomain: (f = msg.getRequestDomain()) && domain_pb.Domain.toObject(includeInstance, f),
    resultDomain: (f = msg.getResultDomain()) && domain_pb.Domain.toObject(includeInstance, f),
    address: (f = msg.getAddress()) && address_pb.Address.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.atlas.ObjectRequest}
 */
proto.atlas.ObjectRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.atlas.ObjectRequest;
  return proto.atlas.ObjectRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.atlas.ObjectRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.atlas.ObjectRequest}
 */
proto.atlas.ObjectRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 100:
      var value = new domain_pb.Domain;
      reader.readMessage(value,domain_pb.Domain.deserializeBinaryFromReader);
      msg.setRequestDomain(value);
      break;
    case 200:
      var value = new domain_pb.Domain;
      reader.readMessage(value,domain_pb.Domain.deserializeBinaryFromReader);
      msg.setResultDomain(value);
      break;
    case 300:
      var value = new address_pb.Address;
      reader.readMessage(value,address_pb.Address.deserializeBinaryFromReader);
      msg.setAddress(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.atlas.ObjectRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.atlas.ObjectRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.atlas.ObjectRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.atlas.ObjectRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getRequestDomain();
  if (f != null) {
    writer.writeMessage(
      100,
      f,
      domain_pb.Domain.serializeBinaryToWriter
    );
  }
  f = message.getResultDomain();
  if (f != null) {
    writer.writeMessage(
      200,
      f,
      domain_pb.Domain.serializeBinaryToWriter
    );
  }
  f = message.getAddress();
  if (f != null) {
    writer.writeMessage(
      300,
      f,
      address_pb.Address.serializeBinaryToWriter
    );
  }
};


/**
 * optional Domain request_domain = 100;
 * @return {?proto.atlas.Domain}
 */
proto.atlas.ObjectRequest.prototype.getRequestDomain = function() {
  return /** @type{?proto.atlas.Domain} */ (
    jspb.Message.getWrapperField(this, domain_pb.Domain, 100));
};


/**
 * @param {?proto.atlas.Domain|undefined} value
 * @return {!proto.atlas.ObjectRequest} returns this
*/
proto.atlas.ObjectRequest.prototype.setRequestDomain = function(value) {
  return jspb.Message.setOneofWrapperField(this, 100, proto.atlas.ObjectRequest.oneofGroups_[0], value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.atlas.ObjectRequest} returns this
 */
proto.atlas.ObjectRequest.prototype.clearRequestDomain = function() {
  return this.setRequestDomain(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.atlas.ObjectRequest.prototype.hasRequestDomain = function() {
  return jspb.Message.getField(this, 100) != null;
};


/**
 * optional Domain result_domain = 200;
 * @return {?proto.atlas.Domain}
 */
proto.atlas.ObjectRequest.prototype.getResultDomain = function() {
  return /** @type{?proto.atlas.Domain} */ (
    jspb.Message.getWrapperField(this, domain_pb.Domain, 200));
};


/**
 * @param {?proto.atlas.Domain|undefined} value
 * @return {!proto.atlas.ObjectRequest} returns this
*/
proto.atlas.ObjectRequest.prototype.setResultDomain = function(value) {
  return jspb.Message.setOneofWrapperField(this, 200, proto.atlas.ObjectRequest.oneofGroups_[1], value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.atlas.ObjectRequest} returns this
 */
proto.atlas.ObjectRequest.prototype.clearResultDomain = function() {
  return this.setResultDomain(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.atlas.ObjectRequest.prototype.hasResultDomain = function() {
  return jspb.Message.getField(this, 200) != null;
};


/**
 * optional Address address = 300;
 * @return {?proto.atlas.Address}
 */
proto.atlas.ObjectRequest.prototype.getAddress = function() {
  return /** @type{?proto.atlas.Address} */ (
    jspb.Message.getWrapperField(this, address_pb.Address, 300));
};


/**
 * @param {?proto.atlas.Address|undefined} value
 * @return {!proto.atlas.ObjectRequest} returns this
*/
proto.atlas.ObjectRequest.prototype.setAddress = function(value) {
  return jspb.Message.setWrapperField(this, 300, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.atlas.ObjectRequest} returns this
 */
proto.atlas.ObjectRequest.prototype.clearAddress = function() {
  return this.setAddress(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.atlas.ObjectRequest.prototype.hasAddress = function() {
  return jspb.Message.getField(this, 300) != null;
};


goog.object.extend(exports, proto.atlas);
