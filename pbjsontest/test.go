//go:generate ../../../scripts/protoc -I../../../ -I../../../vendor --grpc_out=../../../ --go_out=paths=source_relative,plugins=grpc:../../.. pkg/pbjson/pbjsontest/test.proto
package pbjsontest
