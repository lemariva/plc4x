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

import (
    "encoding/xml"
    "io"
    "github.com/apache/plc4x/plc4go/internal/plc4go/utils"
)

// The data-structure of this message
type S7MessageUserData struct {
    Parent *S7Message
    IS7MessageUserData
}

// The corresponding interface
type IS7MessageUserData interface {
    LengthInBytes() uint16
    LengthInBits() uint16
    Serialize(io utils.WriteBuffer) error
    xml.Marshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *S7MessageUserData) MessageType() uint8 {
    return 0x07
}


func (m *S7MessageUserData) InitializeParent(parent *S7Message, tpduReference uint16, parameter *S7Parameter, payload *S7Payload) {
    m.Parent.TpduReference = tpduReference
    m.Parent.Parameter = parameter
    m.Parent.Payload = payload
}

func NewS7MessageUserData(tpduReference uint16, parameter *S7Parameter, payload *S7Payload) *S7Message {
    child := &S7MessageUserData{
        Parent: NewS7Message(tpduReference, parameter, payload),
    }
    child.Parent.Child = child
    return child.Parent
}

func CastS7MessageUserData(structType interface{}) *S7MessageUserData {
    castFunc := func(typ interface{}) *S7MessageUserData {
        if casted, ok := typ.(S7MessageUserData); ok {
            return &casted
        }
        if casted, ok := typ.(*S7MessageUserData); ok {
            return casted
        }
        if casted, ok := typ.(S7Message); ok {
            return CastS7MessageUserData(casted.Child)
        }
        if casted, ok := typ.(*S7Message); ok {
            return CastS7MessageUserData(casted.Child)
        }
        return nil
    }
    return castFunc(structType)
}

func (m *S7MessageUserData) GetTypeName() string {
    return "S7MessageUserData"
}

func (m *S7MessageUserData) LengthInBits() uint16 {
    lengthInBits := uint16(0)

    return lengthInBits
}

func (m *S7MessageUserData) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func S7MessageUserDataParse(io *utils.ReadBuffer) (*S7Message, error) {

    // Create a partially initialized instance
    _child := &S7MessageUserData{
        Parent: &S7Message{},
    }
    _child.Parent.Child = _child
    return _child.Parent, nil
}

func (m *S7MessageUserData) Serialize(io utils.WriteBuffer) error {
    ser := func() error {

        return nil
    }
    return m.Parent.SerializeParent(io, m, ser)
}

func (m *S7MessageUserData) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
    var token xml.Token
    var err error
    token = start
    for {
        switch token.(type) {
        case xml.StartElement:
            tok := token.(xml.StartElement)
            switch tok.Name.Local {
            }
        }
        token, err = d.Token()
        if err != nil {
            if err == io.EOF {
                return nil
            }
            return err
        }
    }
}

func (m *S7MessageUserData) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    return nil
}

