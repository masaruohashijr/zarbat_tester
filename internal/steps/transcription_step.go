package steps

import (
	"encoding/xml"
	"fmt"
	"zarbat_test/internal/godog/services"
	"zarbat_test/pkg/domains"
)

func IShouldListAtLeastTranscription(from, to string, timeInSeconds int) error {
	calls, err := CallPrimaryPort.FilterCalls(from, to, "in-progress")
	if err != nil {
		return fmt.Errorf("Error %s", err.Error())
	}
	if len(calls) > 0 {
		RecordingPrimaryPort.RecordCall(calls[0].Sid, timeInSeconds)
	} else {
		return fmt.Errorf("There is no in-progress call.")
	}
	return nil
}
func IProvideAnAudioUrl(from, to string) error {
	calls, err := CallPrimaryPort.FilterCalls(from, to, "completed")
	if err != nil {
		return fmt.Errorf("Error %s", err.Error())
	}
	if len(calls) > 0 {
		recordings, err := RecordingPrimaryPort.ListRecordings(calls[0].Sid)
		if err != nil {
			return fmt.Errorf("Error %s", err.Error())
		}
		if len(recordings) == 0 {
			return fmt.Errorf("Error %s", err.Error())
		}
	}
	return nil
}

func IShouldGetTranscriptionTextAs(number string) error {
	testHash := fmt.Sprint(TestHash)
	r := &domains.Record{
		Background:         services.Background,
		MaxLength:          services.MaxLength,
		FileFormat:         services.FileFormat,
		Transcribe:         true,
		TranscribeCallback: services.BaseUrl + "/TranscribeCallback" + "?hash=" + testHash,
	}
	p := &domains.Pause{
		Length: 0,
	}
	ResponseRecord.Pause = *p
	ResponseRecord.Record = *r
	x, _ := xml.MarshalIndent(ResponseRecord, "", "")
	strXML := domains.Header + string(x)
	println(strXML)
	services.WriteActionXML("record", strXML)
	Configuration.VoiceUrl = services.BaseUrl + "/Record"
	Configuration.To, Configuration.ToSid = Configuration.SelectNumber(number)
	NumberPrimaryPort.UpdateNumber()
	return nil
}

func ITranscribeAudioUrl(number string) error {
	testHash := fmt.Sprint(TestHash)
	r := &domains.Record{
		Background: services.Background,
		MaxLength:  services.MaxLength,
		FileFormat: services.FileFormat,
		Transcribe: false,
		Action:     services.BaseUrl + "/RecordAction" + "?hash=" + testHash,
	}
	p := &domains.Pause{
		Length: 0,
	}
	ResponseRecord.Pause = *p
	ResponseRecord.Record = *r
	x, _ := xml.MarshalIndent(ResponseRecord, "", "")
	strXML := domains.Header + string(x)
	println(strXML)
	services.WriteActionXML("record", strXML)
	Configuration.VoiceUrl = services.BaseUrl + "/Record"
	Configuration.To, Configuration.ToSid = Configuration.SelectNumber(number)
	NumberPrimaryPort.UpdateNumber()
	return nil
}

func IShouldGetLastTranscriptionTextAs() error {
	return nil
}

func ITranscribeLastRecodingFromTo() error {
	return nil
}