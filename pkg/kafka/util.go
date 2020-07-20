// Copyright 2020 The Atlas Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
)

// MsgAddressPrintable
func MsgAddressPrintable(_msg interface{}) string {
	switch _msg.(type) {
	case *sarama.ConsumerMessage:
		msg := _msg.(*sarama.ConsumerMessage)
		return fmt.Sprintf("topic:%s partition:%d offset:%d", msg.Topic, msg.Partition, msg.Offset)
	case *sarama.ProducerMessage:
		msg := _msg.(*sarama.ProducerMessage)
		return fmt.Sprintf("topic:%s partition:%d offset:%d", msg.Topic, msg.Partition, msg.Offset)
	default:
		return fmt.Sprintf("unknown type to print msg address")
	}
}
