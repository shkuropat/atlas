// source: address.proto
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

var digest_pb = require('./digest_pb.js');
goog.object.extend(proto, digest_pb);
var dirname_pb = require('./dirname_pb.js');
goog.object.extend(proto, dirname_pb);
var domain_pb = require('./domain_pb.js');
goog.object.extend(proto, domain_pb);
var filename_pb = require('./filename_pb.js');
goog.object.extend(proto, filename_pb);
var kafka_address_pb = require('./kafka_address_pb.js');
goog.object.extend(proto, kafka_address_pb);
var s3_address_pb = require('./s3_address_pb.js');
goog.object.extend(proto, s3_address_pb);
var url_pb = require('./url_pb.js');
goog.object.extend(proto, url_pb);
var user_id_pb = require('./user_id_pb.js');
goog.object.extend(proto, user_id_pb);
var uuid_pb = require('./uuid_pb.js');
goog.object.extend(proto, uuid_pb);
goog.exportSymbol('proto.atlas.Address', null, global);
goog.exportSymbol('proto.atlas.Address.AddressOptionalCase', null, global);
goog.exportSymbol('proto.atlas.Address.DomainOptionalCase', null, global);
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
proto.atlas.Address = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, 500, null, proto.atlas.Address.oneofGroups_);
};
goog.inherits(proto.atlas.Address, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.atlas.Address.displayName = 'proto.atlas.Address';
}

/**
 * Oneof group definitions for this message. Each group defines the field
 * numbers belonging to that group. When of these fields' value is set, all
 * other fields in the group are cleared. During deserialization, if multiple
 * fields are encountered for a group, only the last value seen will be kept.
 * @private {!Array<!Array<number>>}
 * @const
 */
proto.atlas.Address.oneofGroups_ = [[10],[100,200,300,400,500,600,700,800,900,1000]];

/**
 * @enum {number}
 */
proto.atlas.Address.DomainOptionalCase = {
  DOMAIN_OPTIONAL_NOT_SET: 0,
  EXPLICIT_DOMAIN: 10
};

/**
 * @return {proto.atlas.Address.DomainOptionalCase}
 */
proto.atlas.Address.prototype.getDomainOptionalCase = function() {
  return /** @type {proto.atlas.Address.DomainOptionalCase} */(jspb.Message.computeOneofCase(this, proto.atlas.Address.oneofGroups_[0]));
};

/**
 * @enum {number}
 */
proto.atlas.Address.AddressOptionalCase = {
  ADDRESS_OPTIONAL_NOT_SET: 0,
  S3: 100,
  KAFKA: 200,
  DIGEST: 300,
  UUID: 400,
  USER_ID: 500,
  DIRNAME: 600,
  FILENAME: 700,
  URL: 800,
  DOMAIN: 900,
  CUSTOM: 1000
};

/**
 * @return {proto.atlas.Address.AddressOptionalCase}
 */
proto.atlas.Address.prototype.getAddressOptionalCase = function() {
  return /** @type {proto.atlas.Address.AddressOptionalCase} */(jspb.Message.computeOneofCase(this, proto.atlas.Address.oneofGroups_[1]));
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
proto.atlas.Address.prototype.toObject = function(opt_includeInstance) {
  return proto.atlas.Address.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.atlas.Address} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.atlas.Address.toObject = function(includeInstance, msg) {
  var f, obj = {
    explicitDomain: (f = msg.getExplicitDomain()) && domain_pb.Domain.toObject(includeInstance, f),
    s3: (f = msg.getS3()) && s3_address_pb.S3Address.toObject(includeInstance, f),
    kafka: (f = msg.getKafka()) && kafka_address_pb.KafkaAddress.toObject(includeInstance, f),
    digest: (f = msg.getDigest()) && digest_pb.Digest.toObject(includeInstance, f),
    uuid: (f = msg.getUuid()) && uuid_pb.UUID.toObject(includeInstance, f),
    userId: (f = msg.getUserId()) && user_id_pb.UserID.toObject(includeInstance, f),
    dirname: (f = msg.getDirname()) && dirname_pb.Dirname.toObject(includeInstance, f),
    filename: (f = msg.getFilename()) && filename_pb.Filename.toObject(includeInstance, f),
    url: (f = msg.getUrl()) && url_pb.URL.toObject(includeInstance, f),
    domain: (f = msg.getDomain()) && domain_pb.Domain.toObject(includeInstance, f),
    custom: jspb.Message.getFieldWithDefault(msg, 1000, "")
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
 * @return {!proto.atlas.Address}
 */
proto.atlas.Address.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.atlas.Address;
  return proto.atlas.Address.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.atlas.Address} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.atlas.Address}
 */
proto.atlas.Address.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 10:
      var value = new domain_pb.Domain;
      reader.readMessage(value,domain_pb.Domain.deserializeBinaryFromReader);
      msg.setExplicitDomain(value);
      break;
    case 100:
      var value = new s3_address_pb.S3Address;
      reader.readMessage(value,s3_address_pb.S3Address.deserializeBinaryFromReader);
      msg.setS3(value);
      break;
    case 200:
      var value = new kafka_address_pb.KafkaAddress;
      reader.readMessage(value,kafka_address_pb.KafkaAddress.deserializeBinaryFromReader);
      msg.setKafka(value);
      break;
    case 300:
      var value = new digest_pb.Digest;
      reader.readMessage(value,digest_pb.Digest.deserializeBinaryFromReader);
      msg.setDigest(value);
      break;
    case 400:
      var value = new uuid_pb.UUID;
      reader.readMessage(value,uuid_pb.UUID.deserializeBinaryFromReader);
      msg.setUuid(value);
      break;
    case 500:
      var value = new user_id_pb.UserID;
      reader.readMessage(value,user_id_pb.UserID.deserializeBinaryFromReader);
      msg.setUserId(value);
      break;
    case 600:
      var value = new dirname_pb.Dirname;
      reader.readMessage(value,dirname_pb.Dirname.deserializeBinaryFromReader);
      msg.setDirname(value);
      break;
    case 700:
      var value = new filename_pb.Filename;
      reader.readMessage(value,filename_pb.Filename.deserializeBinaryFromReader);
      msg.setFilename(value);
      break;
    case 800:
      var value = new url_pb.URL;
      reader.readMessage(value,url_pb.URL.deserializeBinaryFromReader);
      msg.setUrl(value);
      break;
    case 900:
      var value = new domain_pb.Domain;
      reader.readMessage(value,domain_pb.Domain.deserializeBinaryFromReader);
      msg.setDomain(value);
      break;
    case 1000:
      var value = /** @type {string} */ (reader.readString());
      msg.setCustom(value);
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
proto.atlas.Address.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.atlas.Address.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.atlas.Address} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.atlas.Address.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getExplicitDomain();
  if (f != null) {
    writer.writeMessage(
      10,
      f,
      domain_pb.Domain.serializeBinaryToWriter
    );
  }
  f = message.getS3();
  if (f != null) {
    writer.writeMessage(
      100,
      f,
      s3_address_pb.S3Address.serializeBinaryToWriter
    );
  }
  f = message.getKafka();
  if (f != null) {
    writer.writeMessage(
      200,
      f,
      kafka_address_pb.KafkaAddress.serializeBinaryToWriter
    );
  }
  f = message.getDigest();
  if (f != null) {
    writer.writeMessage(
      300,
      f,
      digest_pb.Digest.serializeBinaryToWriter
    );
  }
  f = message.getUuid();
  if (f != null) {
    writer.writeMessage(
      400,
      f,
      uuid_pb.UUID.serializeBinaryToWriter
    );
  }
  f = message.getUserId();
  if (f != null) {
    writer.writeMessage(
      500,
      f,
      user_id_pb.UserID.serializeBinaryToWriter
    );
  }
  f = message.getDirname();
  if (f != null) {
    writer.writeMessage(
      600,
      f,
      dirname_pb.Dirname.serializeBinaryToWriter
    );
  }
  f = message.getFilename();
  if (f != null) {
    writer.writeMessage(
      700,
      f,
      filename_pb.Filename.serializeBinaryToWriter
    );
  }
  f = message.getUrl();
  if (f != null) {
    writer.writeMessage(
      800,
      f,
      url_pb.URL.serializeBinaryToWriter
    );
  }
  f = message.getDomain();
  if (f != null) {
    writer.writeMessage(
      900,
      f,
      domain_pb.Domain.serializeBinaryToWriter
    );
  }
  f = /** @type {string} */ (jspb.Message.getField(message, 1000));
  if (f != null) {
    writer.writeString(
      1000,
      f
    );
  }
};


/**
 * optional Domain explicit_domain = 10;
 * @return {?proto.atlas.Domain}
 */
proto.atlas.Address.prototype.getExplicitDomain = function() {
  return /** @type{?proto.atlas.Domain} */ (
    jspb.Message.getWrapperField(this, domain_pb.Domain, 10));
};


/**
 * @param {?proto.atlas.Domain|undefined} value
 * @return {!proto.atlas.Address} returns this
*/
proto.atlas.Address.prototype.setExplicitDomain = function(value) {
  return jspb.Message.setOneofWrapperField(this, 10, proto.atlas.Address.oneofGroups_[0], value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.atlas.Address} returns this
 */
proto.atlas.Address.prototype.clearExplicitDomain = function() {
  return this.setExplicitDomain(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.atlas.Address.prototype.hasExplicitDomain = function() {
  return jspb.Message.getField(this, 10) != null;
};


/**
 * optional S3Address s3 = 100;
 * @return {?proto.atlas.S3Address}
 */
proto.atlas.Address.prototype.getS3 = function() {
  return /** @type{?proto.atlas.S3Address} */ (
    jspb.Message.getWrapperField(this, s3_address_pb.S3Address, 100));
};


/**
 * @param {?proto.atlas.S3Address|undefined} value
 * @return {!proto.atlas.Address} returns this
*/
proto.atlas.Address.prototype.setS3 = function(value) {
  return jspb.Message.setOneofWrapperField(this, 100, proto.atlas.Address.oneofGroups_[1], value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.atlas.Address} returns this
 */
proto.atlas.Address.prototype.clearS3 = function() {
  return this.setS3(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.atlas.Address.prototype.hasS3 = function() {
  return jspb.Message.getField(this, 100) != null;
};


/**
 * optional KafkaAddress kafka = 200;
 * @return {?proto.atlas.KafkaAddress}
 */
proto.atlas.Address.prototype.getKafka = function() {
  return /** @type{?proto.atlas.KafkaAddress} */ (
    jspb.Message.getWrapperField(this, kafka_address_pb.KafkaAddress, 200));
};


/**
 * @param {?proto.atlas.KafkaAddress|undefined} value
 * @return {!proto.atlas.Address} returns this
*/
proto.atlas.Address.prototype.setKafka = function(value) {
  return jspb.Message.setOneofWrapperField(this, 200, proto.atlas.Address.oneofGroups_[1], value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.atlas.Address} returns this
 */
proto.atlas.Address.prototype.clearKafka = function() {
  return this.setKafka(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.atlas.Address.prototype.hasKafka = function() {
  return jspb.Message.getField(this, 200) != null;
};


/**
 * optional Digest digest = 300;
 * @return {?proto.atlas.Digest}
 */
proto.atlas.Address.prototype.getDigest = function() {
  return /** @type{?proto.atlas.Digest} */ (
    jspb.Message.getWrapperField(this, digest_pb.Digest, 300));
};


/**
 * @param {?proto.atlas.Digest|undefined} value
 * @return {!proto.atlas.Address} returns this
*/
proto.atlas.Address.prototype.setDigest = function(value) {
  return jspb.Message.setOneofWrapperField(this, 300, proto.atlas.Address.oneofGroups_[1], value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.atlas.Address} returns this
 */
proto.atlas.Address.prototype.clearDigest = function() {
  return this.setDigest(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.atlas.Address.prototype.hasDigest = function() {
  return jspb.Message.getField(this, 300) != null;
};


/**
 * optional UUID uuid = 400;
 * @return {?proto.atlas.UUID}
 */
proto.atlas.Address.prototype.getUuid = function() {
  return /** @type{?proto.atlas.UUID} */ (
    jspb.Message.getWrapperField(this, uuid_pb.UUID, 400));
};


/**
 * @param {?proto.atlas.UUID|undefined} value
 * @return {!proto.atlas.Address} returns this
*/
proto.atlas.Address.prototype.setUuid = function(value) {
  return jspb.Message.setOneofWrapperField(this, 400, proto.atlas.Address.oneofGroups_[1], value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.atlas.Address} returns this
 */
proto.atlas.Address.prototype.clearUuid = function() {
  return this.setUuid(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.atlas.Address.prototype.hasUuid = function() {
  return jspb.Message.getField(this, 400) != null;
};


/**
 * optional UserID user_id = 500;
 * @return {?proto.atlas.UserID}
 */
proto.atlas.Address.prototype.getUserId = function() {
  return /** @type{?proto.atlas.UserID} */ (
    jspb.Message.getWrapperField(this, user_id_pb.UserID, 500));
};


/**
 * @param {?proto.atlas.UserID|undefined} value
 * @return {!proto.atlas.Address} returns this
*/
proto.atlas.Address.prototype.setUserId = function(value) {
  return jspb.Message.setOneofWrapperField(this, 500, proto.atlas.Address.oneofGroups_[1], value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.atlas.Address} returns this
 */
proto.atlas.Address.prototype.clearUserId = function() {
  return this.setUserId(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.atlas.Address.prototype.hasUserId = function() {
  return jspb.Message.getField(this, 500) != null;
};


/**
 * optional Dirname dirname = 600;
 * @return {?proto.atlas.Dirname}
 */
proto.atlas.Address.prototype.getDirname = function() {
  return /** @type{?proto.atlas.Dirname} */ (
    jspb.Message.getWrapperField(this, dirname_pb.Dirname, 600));
};


/**
 * @param {?proto.atlas.Dirname|undefined} value
 * @return {!proto.atlas.Address} returns this
*/
proto.atlas.Address.prototype.setDirname = function(value) {
  return jspb.Message.setOneofWrapperField(this, 600, proto.atlas.Address.oneofGroups_[1], value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.atlas.Address} returns this
 */
proto.atlas.Address.prototype.clearDirname = function() {
  return this.setDirname(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.atlas.Address.prototype.hasDirname = function() {
  return jspb.Message.getField(this, 600) != null;
};


/**
 * optional Filename filename = 700;
 * @return {?proto.atlas.Filename}
 */
proto.atlas.Address.prototype.getFilename = function() {
  return /** @type{?proto.atlas.Filename} */ (
    jspb.Message.getWrapperField(this, filename_pb.Filename, 700));
};


/**
 * @param {?proto.atlas.Filename|undefined} value
 * @return {!proto.atlas.Address} returns this
*/
proto.atlas.Address.prototype.setFilename = function(value) {
  return jspb.Message.setOneofWrapperField(this, 700, proto.atlas.Address.oneofGroups_[1], value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.atlas.Address} returns this
 */
proto.atlas.Address.prototype.clearFilename = function() {
  return this.setFilename(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.atlas.Address.prototype.hasFilename = function() {
  return jspb.Message.getField(this, 700) != null;
};


/**
 * optional URL url = 800;
 * @return {?proto.atlas.URL}
 */
proto.atlas.Address.prototype.getUrl = function() {
  return /** @type{?proto.atlas.URL} */ (
    jspb.Message.getWrapperField(this, url_pb.URL, 800));
};


/**
 * @param {?proto.atlas.URL|undefined} value
 * @return {!proto.atlas.Address} returns this
*/
proto.atlas.Address.prototype.setUrl = function(value) {
  return jspb.Message.setOneofWrapperField(this, 800, proto.atlas.Address.oneofGroups_[1], value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.atlas.Address} returns this
 */
proto.atlas.Address.prototype.clearUrl = function() {
  return this.setUrl(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.atlas.Address.prototype.hasUrl = function() {
  return jspb.Message.getField(this, 800) != null;
};


/**
 * optional Domain domain = 900;
 * @return {?proto.atlas.Domain}
 */
proto.atlas.Address.prototype.getDomain = function() {
  return /** @type{?proto.atlas.Domain} */ (
    jspb.Message.getWrapperField(this, domain_pb.Domain, 900));
};


/**
 * @param {?proto.atlas.Domain|undefined} value
 * @return {!proto.atlas.Address} returns this
*/
proto.atlas.Address.prototype.setDomain = function(value) {
  return jspb.Message.setOneofWrapperField(this, 900, proto.atlas.Address.oneofGroups_[1], value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.atlas.Address} returns this
 */
proto.atlas.Address.prototype.clearDomain = function() {
  return this.setDomain(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.atlas.Address.prototype.hasDomain = function() {
  return jspb.Message.getField(this, 900) != null;
};


/**
 * optional string custom = 1000;
 * @return {string}
 */
proto.atlas.Address.prototype.getCustom = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1000, ""));
};


/**
 * @param {string} value
 * @return {!proto.atlas.Address} returns this
 */
proto.atlas.Address.prototype.setCustom = function(value) {
  return jspb.Message.setOneofField(this, 1000, proto.atlas.Address.oneofGroups_[1], value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.atlas.Address} returns this
 */
proto.atlas.Address.prototype.clearCustom = function() {
  return jspb.Message.setOneofField(this, 1000, proto.atlas.Address.oneofGroups_[1], undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.atlas.Address.prototype.hasCustom = function() {
  return jspb.Message.getField(this, 1000) != null;
};


goog.object.extend(exports, proto.atlas);
