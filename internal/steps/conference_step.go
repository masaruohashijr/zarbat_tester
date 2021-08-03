package steps

import (
	"encoding/xml"
	"fmt"
	"net/url"
	"time"
	"zarbat_test/internal/godog/services"
	"zarbat_test/pkg/domains"
)

func ConfiguredAsConferenceWithSize(conferenceNumber string, conferenceName string, size int) error {
	Configuration.To, Configuration.ToSid = Configuration.SelectNumber(conferenceNumber)
	c := &domains.Conference{
		Value:                  conferenceName,
		StartConferenceOnEnter: true,
		CallbackUrl:            services.BaseUrl + "/ConferenceCallback",
		HangupOnStar:           true,
		MaxParticipants:        size,
	}
	dc := &domains.DialConference{
		Conference: *c,
	}
	ResponseConference.DialConference = *dc
	p := &domains.Hangup{}
	ResponseConference.Pause = domains.Pause{
		Length: 5,
	}
	ResponseConference.Hangup = *p
	x, _ := xml.MarshalIndent(ResponseConference, "", "")
	strXML := domains.Header + string(x)
	services.WriteActionXML("conference", strXML)
	Configuration.VoiceUrl = services.BaseUrl + "/Conference"
	NumberPrimaryPort.UpdateNumber()
	println(string(x))
	return nil
}

func ShouldBeEnterConference(number, conferenceName string) error {
	bodyContent := ""
	select {
	case bodyContent = <-Ch:
		fmt.Printf("Result: %s\n", bodyContent)
	case <-time.After(time.Duration(services.TestTimeout) * time.Second):
		fmt.Println("timeout")
		Ch = nil
		return fmt.Errorf("timeout")
	}

	url_parameters, _ := url.ParseQuery(bodyContent)
	cname := url_parameters["ConferenceName"][0]
	if cname != conferenceName {
		return fmt.Errorf("Expected message %s different from %s.", conferenceName, cname)
	}
	Configuration.VoiceUrl = ""
	Configuration.To, Configuration.ToSid = Configuration.SelectNumber(number)
	NumberPrimaryPort.UpdateNumber()
	return nil
}
