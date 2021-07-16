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

#
#
#
function clean_grpc_code_go() {
    CODE_FILES_FOLDER="${1}"

    echo "Go code generator. Clean previously generated .pb.go files in ${CODE_FILES_FOLDER}"
    rm -f "${PROTO_FILES_FOLDER}"/*.pb.go
}

#
# Generate Go code from .proto files
#
function generate_grpc_code_go() {
    PROTO_FILES_FOLDER="${1}"
    RESULT_FILES_FOLDER="${1}"

    echo "Go code generator. Generate code from .proto files in ${PROTO_FILES_FOLDER} into ${RESULT_FILES_FOLDER}"

    if [[ -z "${PROTO_FILES_FOLDER}" ]]; then
        echo "Go code generator. Need to specify folder where to look for .proto files to generate code from "
        exit 1
    fi

    clean_grpc_code_go "${PROTO_FILES_FOLDER}"

    echo "Go code generator. Compile .proto files in ${PROTO_FILES_FOLDER}"
    # Specify the directory in which to search for imports. May be specified multiple times
    IMPORTS_FOLDER="${PROTO_FILES_FOLDER}"
    # --go_out requires list of plugins to be used
    "${PROTOC}" \
        -I "${IMPORTS_FOLDER}" \
        --go_out=plugins=grpc:"${RESULT_FILES_FOLDER}" \
        "${PROTO_FILES_FOLDER}"/*.proto

    #protoc -I "${PROTO_FILES_FOLDER}" --go_out=plugins=grpc:"${PROTO_FILES_FOLDER}" "${PROTO_FILES_FOLDER}"/*.proto
}

#
#
#
function clean_grpc_code_js() {
    CODE_FILES_FOLDER="${1}"

    echo "JS code generator. Clean previously generated *_pb.js files in ${CODE_FILES_FOLDER}"
    rm -f "${CODE_FILES_FOLDER}"/*_pb.js
}


#
# Generate JS code from .proto files
#
function generate_grpc_code_js() {
    PROTO_FILES_FOLDER="${1}"
    RESULT_FILES_FOLDER="${1}"

    echo "JS code generator. Generate code from .proto files in ${PROTO_FILES_FOLDER} into ${RESULT_FILES_FOLDER}"

    if [[ -z "${PROTO_FILES_FOLDER}" ]]; then
        echo "JS code generator. Need to specify folder where to look for .proto files to generate code from "
        exit 1
    fi

    clean_grpc_code_js "${PROTO_FILES_FOLDER}"

    echo "JS code generator. Compile .proto files in ${PROTO_FILES_FOLDER}"
    # Specify the directory in which to search for imports. May be specified multiple times
    IMPORTS_FOLDER="${PROTO_FILES_FOLDER}"

    # Generate Protobuf Messages and Service Client Stub
    # with the help of protoc

    # To generate the protobuf message classes from our echo.proto, run the following command:
    "${PROTOC}" \
        -I "${IMPORTS_FOLDER}" \
        --js_out=import_style=commonjs:"${RESULT_FILES_FOLDER}" \
        "${PROTO_FILES_FOLDER}"/*.proto

    # Generate the service client stub
    # In the --grpc-web_out param:
    # import_style can be closure (default) or commonjs
    # mode can be grpcwebtext (default) or grpcweb
    "${PROTOC}" \
        -I "${IMPORTS_FOLDER}" \
        --grpc-web_out=import_style=commonjs,mode=grpcwebtext:"${RESULT_FILES_FOLDER}" \
        "${PROTO_FILES_FOLDER}"/*.proto
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

#
#
#
function rename_uuid_function() {
    PROTO_FILES_FOLDER="${1}"

    if [[ -z "${PROTO_FILES_FOLDER}" ]]; then
        echo "need to specify folder where to look for .pb.go files to process"
        exit 1
    fi

    for FILE in "${PROTO_FILES_FOLDER}"/*.pb.go; do
        sed -i "s/GetUuid/GetUUID/g" "${FILE}"
    done
}

#
#
#
BUILD_DOCS_HTML="yes"
BUILD_DOCS_MD="yes"
function generate_docs() {
    AREA="${1}"
    PACKAGE_NAME="${2}"
    PROTO_FILES_FOLDER="${3}"

    echo "AutoDoc generator. Generate docs start."

    DOC_FILES_FOLDER="${DOCS_ROOT}/${AREA}/${PACKAGE_NAME}"
    echo "AutoDoc generator. Prepare folder for docs ${DOC_FILES_FOLDER}"
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

generate_grpc_code_go "${PROTO_ROOT}"/atlas
generate_grpc_code_go "${PROTO_ROOT}"/health
generate_grpc_code_js "${PROTO_ROOT}"/health
generate_grpc_code_js "${PROTO_ROOT}"/atlas

delete_string_function "${PROTO_ROOT}"/atlas
rename_uuid_function "${PROTO_ROOT}"/atlas

generate_docs api atlas "${PROTO_ROOT}"/atlas
generate_docs api health "${PROTO_ROOT}"/health
