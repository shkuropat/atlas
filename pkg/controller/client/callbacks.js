import {UuidFromString} from "./uuid";

const { UuidToString } = require('./uuid')


export function callbackTaskStatus(err, response) {
    try {
        const statuses = response.getObjectStatusesList()
        let objectStatus = []

        console.log('callbackTaskStatus:', statuses)

        for (let i = 0; i < statuses.length; i++) {
            const status = statuses[i]
            const code = status.getStatus().getCode()
            const uuid = UuidToString(status.getAddress().getUuid())
            console.log(i + ' : ' +  'object status: ' + code + " : " +  uuid)
        }
    } catch (error) {
        console.log(`Error code: ${err.code} ${err.message}`)
    }
}

export function callbackTaskReport(err, response) {
    try {
        const reports = response.getReportsList()
        console.log('callbackTaskReport:', reports)

        for (let i = 0; i < reports.length; i++) {
            const report = reports[i]
            const children = report.getChildrenList()

            //
            console.log(`Get report object: ${i}`)
            console.log(`Report: ${children}`)

            for (let j = 0; j < children.length; j++) {
                const childReport = children[j]

                // Extract report text from child report
                const uint8array = childReport.getBytes_asU8()
                const reportText = new TextDecoder("utf-8").decode(uint8array)

                //
                console.log(`Get child report object: ${j}`)
                console.log(`Child report: ${reportText}`)
            }
        }
    } catch (error) {
        console.log(`Error code: ${err.code} ${err.message}`)
    }
}

export function callbackTask(err, response) {
    try {
        const tasks = response.getFilesList()
        console.log('callbackTask', tasks)

        for (let i = 0; i < tasks.length; i++) {
            const task = tasks[i]

            //...
        }

    } catch (error) {
        console.log(`Error code: ${err.code} ${err.message}`)
    }
}

export function callbackTaskFiles(err, response) {
    try {
        const files =  response.getFilesList()
        console.log("callbackTaskFiles:", files)

        for (let i = 0; i < files.length; i++) {
            const file = files[i]

            //...
        }

    } catch (error) {
        console.log(`Error code: ${err.code} ${err.message}`)
    }
}