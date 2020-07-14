// Autogenerated by Thrift Compiler (0.9.2)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package example

import (
	"bytes"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

var GoUnusedProtection__ int

type Data struct {
	Text string `thrift:"text,1" json:"text"`
}

func NewData() *Data {
	return &Data{}
}

func (p *Data) GetText() string {
	return p.Text
}
func (p *Data) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error: %s", p, err)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.ReadField1(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *Data) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return fmt.Errorf("error reading field 1: %s", err)
	} else {
		p.Text = v
	}
	return nil
}

func (p *Data) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("Data"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("write struct stop error: %s", err)
	}
	return nil
}

func (p *Data) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("text", thrift.STRING, 1); err != nil {
		return fmt.Errorf("%T write field begin error 1:text: %s", p, err)
	}
	if err := oprot.WriteString(string(p.Text)); err != nil {
		return fmt.Errorf("%T.text (1) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 1:text: %s", p, err)
	}
	return err
}

func (p *Data) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Data(%+v)", *p)
}
