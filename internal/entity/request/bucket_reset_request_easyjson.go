// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package request

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonAf04e3e8DecodeGithubComgmcriptoboxOtusGoFinalProjectInternalEntity(in *jlexer.Lexer, out *BucketResetRequest) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "login":
			out.Login = string(in.String())
		case "ip":
			out.IP = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonAf04e3e8EncodeGithubComgmcriptoboxOtusGoFinalProjectInternalEntity(out *jwriter.Writer, in BucketResetRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"login\":"
		out.RawString(prefix[1:])
		out.String(string(in.Login))
	}
	{
		const prefix string = ",\"ip\":"
		out.RawString(prefix)
		out.String(string(in.IP))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v BucketResetRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonAf04e3e8EncodeGithubComgmcriptoboxOtusGoFinalProjectInternalEntity(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v BucketResetRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonAf04e3e8EncodeGithubComgmcriptoboxOtusGoFinalProjectInternalEntity(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *BucketResetRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonAf04e3e8DecodeGithubComgmcriptoboxOtusGoFinalProjectInternalEntity(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *BucketResetRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonAf04e3e8DecodeGithubComgmcriptoboxOtusGoFinalProjectInternalEntity(l, v)
}
