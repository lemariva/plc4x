/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package model

import (
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type FirmataCommandSetPinMode struct {
	Pin    uint8
	Mode   PinMode
	Parent *FirmataCommand
}

// The corresponding interface
type IFirmataCommandSetPinMode interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *FirmataCommandSetPinMode) CommandCode() uint8 {
	return 0x4
}

func (m *FirmataCommandSetPinMode) InitializeParent(parent *FirmataCommand) {
}

func NewFirmataCommandSetPinMode(pin uint8, mode PinMode) *FirmataCommand {
	child := &FirmataCommandSetPinMode{
		Pin:    pin,
		Mode:   mode,
		Parent: NewFirmataCommand(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastFirmataCommandSetPinMode(structType interface{}) *FirmataCommandSetPinMode {
	castFunc := func(typ interface{}) *FirmataCommandSetPinMode {
		if casted, ok := typ.(FirmataCommandSetPinMode); ok {
			return &casted
		}
		if casted, ok := typ.(*FirmataCommandSetPinMode); ok {
			return casted
		}
		if casted, ok := typ.(FirmataCommand); ok {
			return CastFirmataCommandSetPinMode(casted.Child)
		}
		if casted, ok := typ.(*FirmataCommand); ok {
			return CastFirmataCommandSetPinMode(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *FirmataCommandSetPinMode) GetTypeName() string {
	return "FirmataCommandSetPinMode"
}

func (m *FirmataCommandSetPinMode) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *FirmataCommandSetPinMode) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Simple field (pin)
	lengthInBits += 8

	// Enum Field (mode)
	lengthInBits += 8

	return lengthInBits
}

func (m *FirmataCommandSetPinMode) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func FirmataCommandSetPinModeParse(readBuffer utils.ReadBuffer) (*FirmataCommand, error) {
	if pullErr := readBuffer.PullContext("FirmataCommandSetPinMode"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (pin)
	pin, _pinErr := readBuffer.ReadUint8("pin", 8)
	if _pinErr != nil {
		return nil, errors.Wrap(_pinErr, "Error parsing 'pin' field")
	}

	if pullErr := readBuffer.PullContext("mode"); pullErr != nil {
		return nil, pullErr
	}
	// Enum field (mode)
	mode, _modeErr := PinModeParse(readBuffer)
	if _modeErr != nil {
		return nil, errors.Wrap(_modeErr, "Error parsing 'mode' field")
	}
	if closeErr := readBuffer.CloseContext("mode"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("FirmataCommandSetPinMode"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &FirmataCommandSetPinMode{
		Pin:    pin,
		Mode:   mode,
		Parent: &FirmataCommand{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *FirmataCommandSetPinMode) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("FirmataCommandSetPinMode"); pushErr != nil {
			return pushErr
		}

		// Simple Field (pin)
		pin := uint8(m.Pin)
		_pinErr := writeBuffer.WriteUint8("pin", 8, (pin))
		if _pinErr != nil {
			return errors.Wrap(_pinErr, "Error serializing 'pin' field")
		}

		if pushErr := writeBuffer.PushContext("mode"); pushErr != nil {
			return pushErr
		}
		// Enum field (mode)
		mode := CastPinMode(m.Mode)
		_modeErr := mode.Serialize(writeBuffer)
		if _modeErr != nil {
			return errors.Wrap(_modeErr, "Error serializing 'mode' field")
		}
		if popErr := writeBuffer.PopContext("mode"); popErr != nil {
			return popErr
		}

		if popErr := writeBuffer.PopContext("FirmataCommandSetPinMode"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}

func (m *FirmataCommandSetPinMode) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
