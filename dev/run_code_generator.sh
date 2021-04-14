#!/bin/bash

# Exit immediately when a command fails
set -o errexit
# Error on unset variables
set -o nounset
# Only exit with zero if all commands of the pipeline exit successfully
set -o pipefail

# Source configuration
CUR_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"
source "${CUR_DIR}/go_build_config.sh"

echo "Generating code with the following options:"
echo "      SRC_ROOT=${SRC_ROOT}"
echo ""
echo ""

# Check whether required binaries available
PROTOC="${PROTOC:-protoc}"
PROTOC_GEN_DOC="${PROTOC_GEN_DOC:-protoc-gen-doc}"

# Check protoc is available
"${PROTOC}" --help > /dev/null
if [[ $? ]]; then
    :
else
    echo "${PROTOC} is not available. Abort"
    exit 1
fi

# Check protoc-gen-doc is available
"${PROTOC_GEN_DOC}" --help > /dev/null
if [[ $? ]]; then
    :
else
    echo "${PROTOC_GEN_DOC} is not available. Abort"
    exit 1
fi


# Setup folders
PROTO_ROOT="${PKG_ROOT}/api"

function generate_grpc_code() {
    PROTO_FILES_FOLDER="${1}"

    if [[ -z "${PROTO_FILES_FOLDER}" ]]; then
        echo "need to specify folder where to look for .proto files to generate code from "
        exit 1
    fi

    echo "Generate code from .proto files in ${PROTO_FILES_FOLDER}"

    echo "Clean previously generated files"
    rm -f "${PROTO_FILES_FOLDER}"/*.pb.go

    echo "Compile .proto files in ${PROTO_FILES_FOLDER}"
    # --go_out requires list of plugins to be used
    "${PROTOC}" \
        -I "${PROTO_FILES_FOLDER}" \
        --go_out=plugins=grpc:"${PROTO_FILES_FOLDER}" \
        "${PROTO_FILES_FOLDER}"/*.proto

    #protoc -I "${PROTO_FILES_FOLDER}" --go_out=plugins=grpc:"${PROTO_FILES_FOLDER}" "${PROTO_FILES_FOLDER}"/*.proto
}

# Delete String() function from generated *.pb.go files
# This function is not that human-friendly and it is better to introduce own function for each type
function delete_string_function() {
    PROTO_FILES_FOLDER="${1}"

    if [[ -z "${PROTO_FILES_FOLDER}" ]]; then
        echo "need to specify folder where to look for .pb.go files to process"
        exit 1
    fi

    # /path/to/file:LINE_NUMBER:line
    # /path/to/file:31:func (m *Address) String() string { return proto.CompactTextString(m) }
    FILES_LINES=$(grep -n "String() string { return proto.CompactTextString(m) }" "${PROTO_FILES_FOLDER}"/*.pb.go | cut -f1,2 -d:)

    for FILE_LINE in $FILES_LINES; do
        # Cut filename from the grep-output line
        FILE=$(echo "${FILE_LINE}" | cut -f1 -d:)
        # Cut line number from the grep-output line
        LINE=$(echo "${FILE_LINE}" | cut -f2 -d:)
        #echo "${FILE}:${LINE}"
        # Cut specified line from the file and rewrite the file
        sed "${LINE}d" "${FILE}" > "${FILE}".new && mv "${FILE}".new "${FILE}"
    done


}

BUILD_DOCS_HTML="yes"
BUILD_DOCS_MD="yes"

function generate_docs() {
    AREA="${1}"
    PACKAGE_NAME="${2}"
    PROTO_FILES_FOLDER="${3}"

    DOC_FILES_FOLDER="${DOCS_ROOT}/${AREA}/${PACKAGE_NAME}"
    echo "Prepare folder for docs ${DOC_FILES_FOLDER}"
    mkdir -p "${DOC_FILES_FOLDER}"

    if [[ "${BUILD_DOCS_HTML}" == "yes" ]]; then
        DOC_FILE_NAME_HTML="${PACKAGE_NAME}.html"

        rm -f "${DOC_FILES_FOLDER}/${DOC_FILE_NAME_HTML}"
        "${PROTOC}" \
          -I "${PROTO_FILES_FOLDER}" \
          --doc_out="${DOC_FILES_FOLDER}" \
          --doc_opt=html,"${DOC_FILE_NAME_HTML}" \
          "${PROTO_FILES_FOLDER}"/*.proto
    fi

    if [[ "${BUILD_DOCS_MD}" == "yes" ]]; then
        DOC_FILE_NAME_MD="${PACKAGE_NAME}.md"
        rm -f "${DOC_FILES_FOLDER}/${DOC_FILE_NAME_MD}"
        "${PROTOC}" \
            -I "${PROTO_FILES_FOLDER}" \
            --doc_out="${DOC_FILES_FOLDER}" \
            --doc_opt=markdown,"${DOC_FILE_NAME_MD}" \
            "${PROTO_FILES_FOLDER}"/*.proto
    fi
}

generate_grpc_code "${PROTO_ROOT}"/atlas
generate_grpc_code "${PROTO_ROOT}"/health

delete_string_function "${PROTO_ROOT}"/atlas

generate_docs api atlas "${PROTO_ROOT}"/atlas
generate_docs api health "${PROTO_ROOT}"/health
