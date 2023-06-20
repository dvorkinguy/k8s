// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package module

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

func easyjsonF0077844DecodeGithubComDataDogDatadogAgentPkgSecurityModule(in *jlexer.Lexer, out *SelfTestEvent) {
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
		case "succeeded_tests":
			if in.IsNull() {
				in.Skip()
				out.Success = nil
			} else {
				in.Delim('[')
				if out.Success == nil {
					if !in.IsDelim(']') {
						out.Success = make([]string, 0, 4)
					} else {
						out.Success = []string{}
					}
				} else {
					out.Success = (out.Success)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.Success = append(out.Success, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "failed_tests":
			if in.IsNull() {
				in.Skip()
				out.Fails = nil
			} else {
				in.Delim('[')
				if out.Fails == nil {
					if !in.IsDelim(']') {
						out.Fails = make([]string, 0, 4)
					} else {
						out.Fails = []string{}
					}
				} else {
					out.Fails = (out.Fails)[:0]
				}
				for !in.IsDelim(']') {
					var v2 string
					v2 = string(in.String())
					out.Fails = append(out.Fails, v2)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "date":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Timestamp).UnmarshalJSON(data))
			}
		case "service":
			out.Service = string(in.String())
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
func easyjsonF0077844EncodeGithubComDataDogDatadogAgentPkgSecurityModule(out *jwriter.Writer, in SelfTestEvent) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"succeeded_tests\":"
		out.RawString(prefix[1:])
		if in.Success == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v3, v4 := range in.Success {
				if v3 > 0 {
					out.RawByte(',')
				}
				out.String(string(v4))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"failed_tests\":"
		out.RawString(prefix)
		if in.Fails == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Fails {
				if v5 > 0 {
					out.RawByte(',')
				}
				out.String(string(v6))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"date\":"
		out.RawString(prefix)
		out.Raw((in.Timestamp).MarshalJSON())
	}
	{
		const prefix string = ",\"service\":"
		out.RawString(prefix)
		out.String(string(in.Service))
	}
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SelfTestEvent) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF0077844EncodeGithubComDataDogDatadogAgentPkgSecurityModule(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SelfTestEvent) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF0077844DecodeGithubComDataDogDatadogAgentPkgSecurityModule(l, v)
}
