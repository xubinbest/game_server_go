package msg

import (
    "github.com/name5566/leaf/network"
    "github.com/name5566/leaf/network/protobuf"
)

var Processor network.Processor

func init() {
    Processor = new(protobuf.Processor)
}
