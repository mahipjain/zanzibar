// Code generated by zanzibar
// @generated
// Checksum : cLHkVccw2xPIXtyZLkj7oQ==
// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package contacts

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

func easyjson472eaa72DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsContactsContactsContactsTestUrlUrl(in *jlexer.Lexer, out *Contacts_TestUrlUrl_Result) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "success":
			if in.IsNull() {
				in.Skip()
				out.Success = nil
			} else {
				if out.Success == nil {
					out.Success = new(string)
				}
				*out.Success = string(in.String())
			}
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
func easyjson472eaa72EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsContactsContactsContactsTestUrlUrl(out *jwriter.Writer, in Contacts_TestUrlUrl_Result) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Success != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"success\":")
		if in.Success == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Success))
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Contacts_TestUrlUrl_Result) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson472eaa72EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsContactsContactsContactsTestUrlUrl(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Contacts_TestUrlUrl_Result) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson472eaa72EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsContactsContactsContactsTestUrlUrl(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Contacts_TestUrlUrl_Result) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson472eaa72DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsContactsContactsContactsTestUrlUrl(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Contacts_TestUrlUrl_Result) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson472eaa72DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsContactsContactsContactsTestUrlUrl(l, v)
}
func easyjson472eaa72DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsContactsContactsContactsTestUrlUrl1(in *jlexer.Lexer, out *Contacts_TestUrlUrl_Args) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
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
func easyjson472eaa72EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsContactsContactsContactsTestUrlUrl1(out *jwriter.Writer, in Contacts_TestUrlUrl_Args) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Contacts_TestUrlUrl_Args) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson472eaa72EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsContactsContactsContactsTestUrlUrl1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Contacts_TestUrlUrl_Args) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson472eaa72EncodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsContactsContactsContactsTestUrlUrl1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Contacts_TestUrlUrl_Args) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson472eaa72DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsContactsContactsContactsTestUrlUrl1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Contacts_TestUrlUrl_Args) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson472eaa72DecodeGithubComUberZanzibarExamplesExampleGatewayBuildGenCodeEndpointsContactsContactsContactsTestUrlUrl1(l, v)
}
