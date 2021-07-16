// source: address_list.proto
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

var address_pb = require('./address_pb.js');
goog.object.extend(proto, address_pb);
var domain_pb = require('./domain_pb.js');
goog.object.extend(proto, domain_pb);
goog.exportSymbol('proto.atlas.AddressList', null, global);
goog.exportSymbol('proto.atlas.AddressList.DomainOptionalCase', null, global);
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
proto.atlas.AddressList = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.atlas.AddressList.repeatedFields_, proto.atlas.AddressList.oneofGroups_);
};
goog.inherits(proto.atlas.AddressList, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.atlas.AddressList.displayName = 'proto.atlas.AddressList';
}

/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.atlas.AddressList.repeatedFields_ = [200];

/**
 * Oneof group definitions for this message. Each group defines the field
 * numbers belonging to that group. When of these fields' value is set, all
 * other fields in the group are cleared. During deserialization, if multiple
 * fields are encountered for a group, only the last value seen will be kept.
 * @private {!Array<!Array<number>>}
 * @const
 */
proto.atlas.AddressList.oneofGroups_ = [[100]];

/**
 * @enum {number}
 */
proto.atlas.AddressList.DomainOptionalCase = {
  DOMAIN_OPTIONAL_NOT_SET: 0,
  DOMAIN: 100
};

/**
 * @return {proto.atlas.AddressList.DomainOptionalCase}
 */
proto.atlas.AddressList.prototype.getDomainOptionalCase = function() {
  return /** @type {proto.atlas.AddressList.DomainOptionalCase} */(jspb.Message.computeOneofCase(this, proto.atlas.AddressList.oneofGroups_[0]));
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
proto.atlas.AddressList.prototype.toObject = function(opt_includeInstance) {
  return proto.atlas.AddressList.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.atlas.AddressList} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.atlas.AddressList.toObject = function(includeInstance, msg) {
  var f, obj = {
    domain: (f = msg.getDomain()) && domain_pb.Domain.toObject(includeInstance, f),
    addressesList: jspb.Message.toObjectList(msg.getAddressesList(),
    address_pb.Address.toObject, includeInstance)
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
 * @return {!proto.atlas.AddressList}
 */
proto.atlas.AddressList.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.atlas.AddressList;
  return proto.atlas.AddressList.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.atlas.AddressList} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.atlas.AddressList}
 */
proto.atlas.AddressList.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 100:
      var value = new domain_pb.Domain;
      reader.readMessage(value,domain_pb.Domain.deserializeBinaryFromReader);
      msg.setDomain(value);
      break;
    case 200:
      var value = new address_pb.Address;
      reader.readMessage(value,address_pb.Address.deserializeBinaryFromReader);
      msg.addAddresses(value);
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
proto.atlas.AddressList.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.atlas.AddressList.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.atlas.AddressList} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.atlas.AddressList.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getDomain();
  if (f != null) {
    writer.writeMessage(
      100,
      f,
      domain_pb.Domain.serializeBinaryToWriter
    );
  }
  f = message.getAddressesList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      200,
      f,
      address_pb.Address.serializeBinaryToWriter
    );
  }
};


/**
 * optional Domain domain = 100;
 * @return {?proto.atlas.Domain}
 */
proto.atlas.AddressList.prototype.getDomain = function() {
  return /** @type{?proto.atlas.Domain} */ (
    jspb.Message.getWrapperField(this, domain_pb.Domain, 100));
};


/**
 * @param {?proto.atlas.Domain|undefined} value
 * @return {!proto.atlas.AddressList} returns this
*/
proto.atlas.AddressList.prototype.setDomain = function(value) {
  return jspb.Message.setOneofWrapperField(this, 100, proto.atlas.AddressList.oneofGroups_[0], value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.atlas.AddressList} returns this
 */
proto.atlas.AddressList.prototype.clearDomain = function() {
  return this.setDomain(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.atlas.AddressList.prototype.hasDomain = function() {
  return jspb.Message.getField(this, 100) != null;
};


/**
 * repeated Address addresses = 200;
 * @return {!Array<!proto.atlas.Address>}
 */
proto.atlas.AddressList.prototype.getAddressesList = function() {
  return /** @type{!Array<!proto.atlas.Address>} */ (
    jspb.Message.getRepeatedWrapperField(this, address_pb.Address, 200));
};


/**
 * @param {!Array<!proto.atlas.Address>} value
 * @return {!proto.atlas.AddressList} returns this
*/
proto.atlas.AddressList.prototype.setAddressesList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 200, value);
};


/**
 * @param {!proto.atlas.Address=} opt_value
 * @param {number=} opt_index
 * @return {!proto.atlas.Address}
 */
proto.atlas.AddressList.prototype.addAddresses = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 200, opt_value, proto.atlas.Address, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.atlas.AddressList} returns this
 */
proto.atlas.AddressList.prototype.clearAddressesList = function() {
  return this.setAddressesList([]);
};


goog.object.extend(exports, proto.atlas);
