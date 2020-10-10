//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//
package model

import "plc4x.apache.org/plc4go-modbus-driver/0.8.0/internal/plc4go/spi"

type MemoryArea uint8

const (
	MemoryArea_COUNTERS                 MemoryArea = 0x1C
	MemoryArea_TIMERS                   MemoryArea = 0x1D
	MemoryArea_DIRECT_PERIPHERAL_ACCESS MemoryArea = 0x80
	MemoryArea_INPUTS                   MemoryArea = 0x81
	MemoryArea_OUTPUTS                  MemoryArea = 0x82
	MemoryArea_FLAGS_MARKERS            MemoryArea = 0x83
	MemoryArea_DATA_BLOCKS              MemoryArea = 0x84
	MemoryArea_INSTANCE_DATA_BLOCKS     MemoryArea = 0x85
	MemoryArea_LOCAL_DATA               MemoryArea = 0x86
)

func (e MemoryArea) GetShortName() string {
	switch e {
	case 0x1C:
		{ /* '0x1C' */
			return "C"
		}
	case 0x1D:
		{ /* '0x1D' */
			return "T"
		}
	case 0x80:
		{ /* '0x80' */
			return "D"
		}
	case 0x81:
		{ /* '0x81' */
			return "I"
		}
	case 0x82:
		{ /* '0x82' */
			return "Q"
		}
	case 0x83:
		{ /* '0x83' */
			return "M"
		}
	case 0x84:
		{ /* '0x84' */
			return "DB"
		}
	case 0x85:
		{ /* '0x85' */
			return "DBI"
		}
	case 0x86:
		{ /* '0x86' */
			return "LD"
		}
	default:
		{
			return ""
		}
	}
}

func MemoryAreaParse(io spi.ReadBuffer) (MemoryArea, error) {
	// TODO: Implement ...
	return 0, nil
}

func (e MemoryArea) Serialize(io spi.WriteBuffer) {
	// TODO: Implement ...
}
