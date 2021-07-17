// source: kafka_address.proto
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

goog.exportSymbol('proto.atlas.KafkaAddress', null, global);
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
proto.atlas.KafkaAddress = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.atlas.KafkaAddress, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.atlas.KafkaAddress.displayName = 'proto.atlas.KafkaAddress';
}



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
proto.atlas.KafkaAddress.prototype.toObject = function(opt_includeInstance) {
  return proto.atlas.KafkaAddress.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.atlas.KafkaAddress} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.atlas.KafkaAddress.toObject = function(includeInstance, msg) {
  var f, obj = {
    topic: jspb.Message.getFieldWithDefault(msg, 100, ""),
    partition: jspb.Message.getFieldWithDefault(msg, 200, 0)
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
 * @return {!proto.atlas.KafkaAddress}
 */
proto.atlas.KafkaAddress.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.atlas.KafkaAddress;
  return proto.atlas.KafkaAddress.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.atlas.KafkaAddress} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.atlas.KafkaAddress}
 */
proto.atlas.KafkaAddress.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 100:
      var value = /** @type {string} */ (reader.readString());
      msg.setTopic(value);
      break;
    case 200:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setPartition(value);
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
proto.atlas.KafkaAddress.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.atlas.KafkaAddress.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.atlas.KafkaAddress} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.atlas.KafkaAddress.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getTopic();
  if (f.length > 0) {
    writer.writeString(
      100,
      f
    );
  }
  f = message.getPartition();
  if (f !== 0) {
    writer.writeInt32(
      200,
      f
    );
  }
};


/**
 * optional string topic = 100;
 * @return {string}
 */
proto.atlas.KafkaAddress.prototype.getTopic = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 100, ""));
};


/**
 * @param {string} value
 * @return {!proto.atlas.KafkaAddress} returns this
 */
proto.atlas.KafkaAddress.prototype.setTopic = function(value) {
  return jspb.Message.setProto3StringField(this, 100, value);
};


/**
 * optional int32 partition = 200;
 * @return {number}
 */
proto.atlas.KafkaAddress.prototype.getPartition = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 200, 0));
};


/**
 * @param {number} value
 * @return {!proto.atlas.KafkaAddress} returns this
 */
proto.atlas.KafkaAddress.prototype.setPartition = function(value) {
  return jspb.Message.setProto3IntField(this, 200, value);
};


goog.object.extend(exports, proto.atlas);