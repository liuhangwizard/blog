
GEN_TS=./node_modules/.bin/protoc-gen-ts
GEN_PLUGIN=./node_modules/.bin/grpc_tools_node_protoc_plugin
TARGET_DIR=./src/protos

protoc \
--js_out=import_style=commonjs,binary:"${TARGET_DIR}"/ \
--ts_out=import_style=commonjs,binary:"${TARGET_DIR}"/ \
--grpc_out="${TARGET_DIR}"/ \
--plugin=protoc-gen-grpc="${GEN_PLUGIN}" \
--plugin=protoc-gen-ts="${GEN_TS}" \
--proto_path=./protos/ \
-I $TARGET_DIR \
./protos/*.proto

echo 'protos generated'
