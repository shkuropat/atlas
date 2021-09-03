const { UUID } = require('../../api/atlas/uuid_pb')

// accepts report object returns string
export function ReportToString(report) {
    var uint8array = report.getBytes_asU8()
    var reportText = new TextDecoder("utf-8").decode(uint8array);
    return reportText
}
