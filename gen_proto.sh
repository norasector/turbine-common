#!/bin/bash

protoc -I=./protobuf --go_out=./types/ --go_opt=module=github.com/norasector/turbine-common/types ./protobuf/*.proto