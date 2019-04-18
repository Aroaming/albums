// Autogenerated by Thrift Compiler (0.9.2)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package imagehandle

import (
	"bytes"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

var GoUnusedProtection__ int

type Result_ struct {
	Desc string `thrift:"desc,1" json:"desc"`
	Ret  []byte `thrift:"ret,2" json:"ret"`
}

func NewResult_() *Result_ {
	return &Result_{}
}

func (p *Result_) GetDesc() string {
	return p.Desc
}

func (p *Result_) GetRet() []byte {
	return p.Ret
}
func (p *Result_) Read(iprot thrift.TProtocol) error {
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
		case 2:
			if err := p.ReadField2(iprot); err != nil {
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

func (p *Result_) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return fmt.Errorf("error reading field 1: %s", err)
	} else {
		p.Desc = v
	}
	return nil
}

func (p *Result_) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(); err != nil {
		return fmt.Errorf("error reading field 2: %s", err)
	} else {
		p.Ret = v
	}
	return nil
}

func (p *Result_) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("Result"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
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

func (p *Result_) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("desc", thrift.STRING, 1); err != nil {
		return fmt.Errorf("%T write field begin error 1:desc: %s", p, err)
	}
	if err := oprot.WriteString(string(p.Desc)); err != nil {
		return fmt.Errorf("%T.desc (1) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 1:desc: %s", p, err)
	}
	return err
}

func (p *Result_) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("ret", thrift.STRING, 2); err != nil {
		return fmt.Errorf("%T write field begin error 2:ret: %s", p, err)
	}
	if err := oprot.WriteBinary(p.Ret); err != nil {
		return fmt.Errorf("%T.ret (2) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 2:ret: %s", p, err)
	}
	return err
}

func (p *Result_) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Result_(%+v)", *p)
}