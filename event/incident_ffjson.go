// DO NOT EDIT!
// Code generated by ffjson <https://github.com/pquerna/ffjson>
// source: incident.go
// DO NOT EDIT!

package event

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	fflib "github.com/pquerna/ffjson/fflib/v1"
	"reflect"
)

func (mj *Incident) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	err := mj.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func (mj *Incident) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{ "event":`)
	if mj.EventName != nil {
		buf.WriteString(`"`)
		{
			enc := base64.NewEncoder(base64.StdEncoding, buf)
			enc.Write(reflect.Indirect(reflect.ValueOf(mj.EventName)).Bytes())
			enc.Close()
		}
		buf.WriteString(`"`)
	} else {
		buf.WriteString(`null`)
	}
	buf.WriteString(`,"time":`)
	fflib.FormatBits2(buf, uint64(mj.Time), 10, mj.Time < 0)
	buf.WriteString(`,"id":`)
	fflib.FormatBits2(buf, uint64(mj.Id), 10, mj.Id < 0)
	if mj.Active {
		buf.WriteString(`,"active":true`)
	} else {
		buf.WriteString(`,"active":false`)
	}
	buf.WriteString(`,"escalation":`)
	fflib.WriteJsonString(buf, string(mj.Escalation))
	buf.WriteString(`,"description":`)
	fflib.WriteJsonString(buf, string(mj.Description))
	buf.WriteString(`,"policy":`)
	fflib.WriteJsonString(buf, string(mj.Policy))
	buf.WriteString(`,"host":`)
	fflib.WriteJsonString(buf, string(mj.Host))
	buf.WriteString(`,"service":`)
	fflib.WriteJsonString(buf, string(mj.Service))
	buf.WriteString(`,"sub_type":`)
	fflib.WriteJsonString(buf, string(mj.SubService))
	buf.WriteString(`,"metric":`)
	fflib.AppendFloat(buf, float64(mj.Metric), 'g', -1, 64)
	buf.WriteString(`,"occurences":`)
	fflib.FormatBits2(buf, uint64(mj.Occurences), 10, mj.Occurences < 0)
	if mj.Tags == nil {
		buf.WriteString(`,"tags":null`)
	} else {
		buf.WriteString(`,"tags":{ `)
		for key, value := range mj.Tags {
			fflib.WriteJsonString(buf, key)
			buf.WriteString(`:`)
			fflib.WriteJsonString(buf, string(value))
			buf.WriteByte(',')
		}
		buf.Rewind(1)
		buf.WriteByte('}')
	}
	buf.WriteString(`,"status":`)
	fflib.FormatBits2(buf, uint64(mj.Status), 10, mj.Status < 0)
	buf.WriteByte(',')
	if mj.IncidentId != nil {
		if true {
			buf.WriteString(`"incident":`)
			fflib.FormatBits2(buf, uint64(*mj.IncidentId), 10, *mj.IncidentId < 0)
			buf.WriteByte(',')
		}
	}
	buf.Rewind(1)
	buf.WriteByte('}')
	return nil
}

const (
	ffj_t_Incidentbase = iota
	ffj_t_Incidentno_such_key

	ffj_t_Incident_EventName

	ffj_t_Incident_Time

	ffj_t_Incident_Id

	ffj_t_Incident_Active

	ffj_t_Incident_Escalation

	ffj_t_Incident_Description

	ffj_t_Incident_Policy

	ffj_t_Incident_Host

	ffj_t_Incident_Service

	ffj_t_Incident_SubService

	ffj_t_Incident_Metric

	ffj_t_Incident_Occurences

	ffj_t_Incident_Tags

	ffj_t_Incident_Status

	ffj_t_Incident_IncidentId
)

var ffj_key_Incident_EventName = []byte("event")

var ffj_key_Incident_Time = []byte("time")

var ffj_key_Incident_Id = []byte("id")

var ffj_key_Incident_Active = []byte("active")

var ffj_key_Incident_Escalation = []byte("escalation")

var ffj_key_Incident_Description = []byte("description")

var ffj_key_Incident_Policy = []byte("policy")

var ffj_key_Incident_Host = []byte("host")

var ffj_key_Incident_Service = []byte("service")

var ffj_key_Incident_SubService = []byte("sub_type")

var ffj_key_Incident_Metric = []byte("metric")

var ffj_key_Incident_Occurences = []byte("occurences")

var ffj_key_Incident_Tags = []byte("tags")

var ffj_key_Incident_Status = []byte("status")

var ffj_key_Incident_IncidentId = []byte("incident")

func (uj *Incident) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return uj.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

func (uj *Incident) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error = nil
	currentKey := ffj_t_Incidentbase
	_ = currentKey
	tok := fflib.FFTok_init
	wantedTok := fflib.FFTok_init

mainparse:
	for {
		tok = fs.Scan()
		//	println(fmt.Sprintf("debug: tok: %v  state: %v", tok, state))
		if tok == fflib.FFTok_error {
			goto tokerror
		}

		switch state {

		case fflib.FFParse_map_start:
			if tok != fflib.FFTok_left_bracket {
				wantedTok = fflib.FFTok_left_bracket
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_key
			continue

		case fflib.FFParse_after_value:
			if tok == fflib.FFTok_comma {
				state = fflib.FFParse_want_key
			} else if tok == fflib.FFTok_right_bracket {
				goto done
			} else {
				wantedTok = fflib.FFTok_comma
				goto wrongtokenerror
			}

		case fflib.FFParse_want_key:
			// json {} ended. goto exit. woo.
			if tok == fflib.FFTok_right_bracket {
				goto done
			}
			if tok != fflib.FFTok_string {
				wantedTok = fflib.FFTok_string
				goto wrongtokenerror
			}

			kn := fs.Output.Bytes()
			if len(kn) <= 0 {
				// "" case. hrm.
				currentKey = ffj_t_Incidentno_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'a':

					if bytes.Equal(ffj_key_Incident_Active, kn) {
						currentKey = ffj_t_Incident_Active
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'd':

					if bytes.Equal(ffj_key_Incident_Description, kn) {
						currentKey = ffj_t_Incident_Description
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'e':

					if bytes.Equal(ffj_key_Incident_EventName, kn) {
						currentKey = ffj_t_Incident_EventName
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_Incident_Escalation, kn) {
						currentKey = ffj_t_Incident_Escalation
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'h':

					if bytes.Equal(ffj_key_Incident_Host, kn) {
						currentKey = ffj_t_Incident_Host
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'i':

					if bytes.Equal(ffj_key_Incident_Id, kn) {
						currentKey = ffj_t_Incident_Id
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_Incident_IncidentId, kn) {
						currentKey = ffj_t_Incident_IncidentId
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'm':

					if bytes.Equal(ffj_key_Incident_Metric, kn) {
						currentKey = ffj_t_Incident_Metric
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'o':

					if bytes.Equal(ffj_key_Incident_Occurences, kn) {
						currentKey = ffj_t_Incident_Occurences
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'p':

					if bytes.Equal(ffj_key_Incident_Policy, kn) {
						currentKey = ffj_t_Incident_Policy
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 's':

					if bytes.Equal(ffj_key_Incident_Service, kn) {
						currentKey = ffj_t_Incident_Service
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_Incident_SubService, kn) {
						currentKey = ffj_t_Incident_SubService
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_Incident_Status, kn) {
						currentKey = ffj_t_Incident_Status
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 't':

					if bytes.Equal(ffj_key_Incident_Time, kn) {
						currentKey = ffj_t_Incident_Time
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_Incident_Tags, kn) {
						currentKey = ffj_t_Incident_Tags
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.SimpleLetterEqualFold(ffj_key_Incident_IncidentId, kn) {
					currentKey = ffj_t_Incident_IncidentId
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_Incident_Status, kn) {
					currentKey = ffj_t_Incident_Status
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_Incident_Tags, kn) {
					currentKey = ffj_t_Incident_Tags
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_Incident_Occurences, kn) {
					currentKey = ffj_t_Incident_Occurences
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Incident_Metric, kn) {
					currentKey = ffj_t_Incident_Metric
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_Incident_SubService, kn) {
					currentKey = ffj_t_Incident_SubService
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_Incident_Service, kn) {
					currentKey = ffj_t_Incident_Service
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_Incident_Host, kn) {
					currentKey = ffj_t_Incident_Host
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Incident_Policy, kn) {
					currentKey = ffj_t_Incident_Policy
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_Incident_Description, kn) {
					currentKey = ffj_t_Incident_Description
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_Incident_Escalation, kn) {
					currentKey = ffj_t_Incident_Escalation
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Incident_Active, kn) {
					currentKey = ffj_t_Incident_Active
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Incident_Id, kn) {
					currentKey = ffj_t_Incident_Id
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Incident_Time, kn) {
					currentKey = ffj_t_Incident_Time
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Incident_EventName, kn) {
					currentKey = ffj_t_Incident_EventName
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffj_t_Incidentno_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			}

		case fflib.FFParse_want_colon:
			if tok != fflib.FFTok_colon {
				wantedTok = fflib.FFTok_colon
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_value
			continue
		case fflib.FFParse_want_value:

			if tok == fflib.FFTok_left_brace || tok == fflib.FFTok_left_bracket || tok == fflib.FFTok_integer || tok == fflib.FFTok_double || tok == fflib.FFTok_string || tok == fflib.FFTok_bool || tok == fflib.FFTok_null {
				switch currentKey {

				case ffj_t_Incident_EventName:
					goto handle_EventName

				case ffj_t_Incident_Time:
					goto handle_Time

				case ffj_t_Incident_Id:
					goto handle_Id

				case ffj_t_Incident_Active:
					goto handle_Active

				case ffj_t_Incident_Escalation:
					goto handle_Escalation

				case ffj_t_Incident_Description:
					goto handle_Description

				case ffj_t_Incident_Policy:
					goto handle_Policy

				case ffj_t_Incident_Host:
					goto handle_Host

				case ffj_t_Incident_Service:
					goto handle_Service

				case ffj_t_Incident_SubService:
					goto handle_SubService

				case ffj_t_Incident_Metric:
					goto handle_Metric

				case ffj_t_Incident_Occurences:
					goto handle_Occurences

				case ffj_t_Incident_Tags:
					goto handle_Tags

				case ffj_t_Incident_Status:
					goto handle_Status

				case ffj_t_Incident_IncidentId:
					goto handle_IncidentId

				case ffj_t_Incidentno_such_key:
					err = fs.SkipField(tok)
					if err != nil {
						return fs.WrapErr(err)
					}
					state = fflib.FFParse_after_value
					goto mainparse
				}
			} else {
				goto wantedvalue
			}
		}
	}

handle_EventName:

	/* handler: uj.EventName type=[]uint8 kind=slice */

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			uj.EventName = nil
		} else {
			b := make([]byte, base64.StdEncoding.DecodedLen(fs.Output.Len()))
			n, err := base64.StdEncoding.Decode(b, fs.Output.Bytes())
			if err != nil {
				return fs.WrapErr(err)
			}

			v := reflect.ValueOf(&uj.EventName).Elem()
			v.SetBytes(b[0:n])

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Time:

	/* handler: uj.Time type=int64 kind=int64 */

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int64", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			uj.Time = int64(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Id:

	/* handler: uj.Id type=int64 kind=int64 */

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int64", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			uj.Id = int64(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Active:

	/* handler: uj.Active type=bool kind=bool */

	{

		{
			if tok != fflib.FFTok_bool && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for bool", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {
			tmpb := fs.Output.Bytes()

			if bytes.Compare([]byte{'t', 'r', 'u', 'e'}, tmpb) == 0 {

				uj.Active = true

			} else if bytes.Compare([]byte{'f', 'a', 'l', 's', 'e'}, tmpb) == 0 {

				uj.Active = false

			} else {
				err = errors.New("unexpected bytes for true/false value")
				return fs.WrapErr(err)
			}

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Escalation:

	/* handler: uj.Escalation type=string kind=string */

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			uj.Escalation = string(fs.Output.String())

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Description:

	/* handler: uj.Description type=string kind=string */

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			uj.Description = string(fs.Output.String())

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Policy:

	/* handler: uj.Policy type=string kind=string */

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			uj.Policy = string(fs.Output.String())

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Host:

	/* handler: uj.Host type=string kind=string */

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			uj.Host = string(fs.Output.String())

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Service:

	/* handler: uj.Service type=string kind=string */

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			uj.Service = string(fs.Output.String())

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_SubService:

	/* handler: uj.SubService type=string kind=string */

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			uj.SubService = string(fs.Output.String())

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Metric:

	/* handler: uj.Metric type=float64 kind=float64 */

	{
		if tok != fflib.FFTok_double && tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for float64", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseFloat(fs.Output.Bytes(), 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			uj.Metric = float64(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Occurences:

	/* handler: uj.Occurences type=int kind=int */

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			uj.Occurences = int(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Tags:

	/* handler: uj.Tags type=map[string]string kind=map */

	{
		/* Falling back. type=map[string]string kind=map */
		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = json.Unmarshal(tbuf, &uj.Tags)
		if err != nil {
			return fs.WrapErr(err)
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Status:

	/* handler: uj.Status type=int kind=int */

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			uj.Status = int(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_IncidentId:

	/* handler: uj.IncidentId type=int64 kind=int64 */

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int64", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

			uj.IncidentId = nil

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			ttypval := int64(tval)
			uj.IncidentId = &ttypval

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

wantedvalue:
	return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
wrongtokenerror:
	return fs.WrapErr(fmt.Errorf("ffjson: wanted token: %v, but got token: %v output=%s", wantedTok, tok, fs.Output.String()))
tokerror:
	if fs.BigError != nil {
		return fs.WrapErr(fs.BigError)
	}
	err = fs.Error.ToError()
	if err != nil {
		return fs.WrapErr(err)
	}
	panic("ffjson-generated: unreachable, please report bug.")
done:
	return nil
}
