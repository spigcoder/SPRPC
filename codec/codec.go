package codec

import "io"

type Header struct {
	ServiceMethod string // format "Service.Method"
	// 用于表明请求的序号，区分不同的请求
	Seq   uint64 // sequence number chosen by client
	Error string
}

type Codec interface {
	io.Closer
	ReadHeader(*Header) error
	ReadBody(any) error
	Write(*Header, any) error
}

type NewCodecFunc func(closer io.ReadWriteCloser) Codec
type Type string

const (
	GobType  Type = "application/gob"
	JsonType Type = "application/json" // not implemented
)

var NewCodecFuncMap map[Type]NewCodecFunc

func init() {
	NewCodecFuncMap = make(map[Type]NewCodecFunc)
	NewCodecFuncMap[GobType] = NewGobCodec
}
