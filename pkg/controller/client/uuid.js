const { UUID } = require('../../api/atlas/uuid_pb')

// accepts uuid object returns string
export function UuidToString(uuid) {
    const uuid_uint8array = uuid().getData_asU8();
    const uuid_string = new TextDecoder("utf-8").decode(uuid_uint8array);
    return uuid_string
}

// accepts string "89ec7e42-8290-45b2-b20b-3f106d821390" returns uuid object
export function UuidFromString(uuid_string) {
    const uuid_bytes = new TextEncoder("utf-8").encode(uuid_string);
    let uuid = new UUID();
    uuid.setData(uuid_bytes);
    return uuid
}
