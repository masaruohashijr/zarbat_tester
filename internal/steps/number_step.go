package steps

import (
	"encoding/xml"
	"fmt"
	"zarbat_test/internal/godog/services"
	"zarbat_test/internal/logging"
	"zarbat_test/pkg/domains"
)

func ConfiguredToDialAndSendDigitsTo(dialerNumber, digits, dialedNumber string) error {
	services.CloseChannel = true
	dialed, _ := Configuration.SelectNumber(dialedNumber)
	n := &domains.Number{
		Value:      dialed,
		SendDigits: digits,
	}
	d := &domains.DialNumber{
		Number: *n,
	}
	ResponseDialNumber.DialNumber = *d
	p := &domains.Hangup{}
	ResponseDialNumber.Hangup = *p
	x, _ := xml.MarshalIndent(ResponseDialNumber, "", "")
	strXML := domains.Header + string(x)
	services.WriteActionXML("number", strXML)
	Configuration.To, Configuration.ToSid = Configuration.SelectNumber(dialerNumber)
	Configuration.VoiceUrl = services.BaseUrl + "/Number"
	NumberPrimaryPort.UpdateNumber()
	println(string(x))
	return nil
}

func ShouldBeReset(number string) error {
	Configuration.To, Configuration.ToSid = Configuration.SelectNumber(number)
	Configuration.VoiceUrl = ""
	NumberPrimaryPort.UpdateNumber()
	return nil
}

func IListAllAvailableNumbers() error {
	anumbers, err := NumberSecondaryPort.ListAvailableNumbers()
	IncomingNumbers = anumbers
	if err != nil {
		return fmt.Errorf("Error %s", "Not able to list my numbers.")
	}
	for _, inumbers := range IncomingNumbers {
		logging.Debug.Println(inumbers)
		println(inumbers)
	}
	return nil
}

func IShouldGetToBuyFromList(amount int) error {
	ok := false
	for i := 0; i < amount; i++ {
		logging.Debug.Println("Buying number is: ", AvailableNumbers[i])
		NumberSecondaryPort.AddNumber(AvailableNumbers[i])
		purchased, _ := NumberSecondaryPort.ListNumbers()
		for _, n := range purchased {
			if AvailableNumbers[i] == n {
				logging.Debug.Println("Purchased number is: ", AvailableNumbers[i])
				ok = true
				break
			}
		}
		if !ok {
			return fmt.Errorf("Error %s", "Not able to list available numbers.")
		}
	}
	return nil
}

func IListMyNumbers() error {
	mynumbers, err := NumberSecondaryPort.ListNumbers()
	IncomingNumbers = mynumbers
	if err != nil {
		return fmt.Errorf("Error %s", "Not able to list available numbers.")
	}
	for _, in := range IncomingNumbers {
		logging.Debug.Println(in)
		println(in)
	}
	return nil
}

func IReleaseAllMyNumbersExcept(exceptionList string) error {
	anumbers, err := NumberSecondaryPort.ListAvailableNumbers()
	AvailableNumbers = anumbers
	if err != nil {
		return fmt.Errorf("Error %s", "Not able to list available numbers.")
	}
	for _, a := range AvailableNumbers {
		logging.Debug.Println(a)
		println(a)
	}
	return nil
}

func IShouldGetNumbersFromMyList(amount int) error {
	ok := false
	for i := 0; i < amount; i++ {
		logging.Debug.Println("Buying number is: ", AvailableNumbers[i])
		NumberSecondaryPort.AddNumber(AvailableNumbers[i])
		purchased, _ := NumberSecondaryPort.ListNumbers()
		for _, n := range purchased {
			if AvailableNumbers[i] == n {
				logging.Debug.Println("Purchased number is: ", AvailableNumbers[i])
				ok = true
				break
			}
		}
		if !ok {
			return fmt.Errorf("Error %s", "Not able to list available numbers.")
		}
	}
	return nil
}

func ConfiguredWithFriendlyNameAs(number, friendlyName string) error {
	services.CloseChannel = true
	Configuration.To, Configuration.ToSid = Configuration.SelectNumber(number)
	Configuration.FriendlyName = friendlyName
	Configuration.VoiceUrl = ""
	NumberPrimaryPort.UpdateNumber()
	return nil
	// return fmt.Errorf("Error %s", "Not able to configure friendly name on number.")
}
func IShouldGetFriendlyNameOn(friendlyName, number string) error {
	return fmt.Errorf("Error %s", "Not able to get friendly name on number.")
}
func IViewInfo(number string) error {
	return fmt.Errorf("Error %s", "Not able to view number info.")
}