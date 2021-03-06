package steps

import (
	"encoding/xml"
	"fmt"
	"strings"
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
	logging.Debug.Println(string(x))
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
	AvailableNumbers = anumbers
	if err != nil {
		logging.Debug.Printf("Error %s", "Not able to list all available numbers.")
		return fmt.Errorf("Error %s", "Not able to list all available numbers.")
	}
	for _, a := range AvailableNumbers {
		logging.Debug.Println(a)
		logging.Debug.Println(a)
	}
	return nil
}

func IShouldGetToBuyFromList(amount int) error {
	ok := false
	for i := 0; i < amount; i++ {
		logging.Debug.Println("Buying number is: ", AvailableNumbers[i])
		NumberSecondaryPort.AddNumber(AvailableNumbers[i])
		purchased, _ := NumberSecondaryPort.ListNumbers()
		for _, n := range *purchased {
			if AvailableNumbers[i] == n.PhoneNumber {
				logging.Debug.Println("Purchased number is: ", AvailableNumbers[i])
				ok = true
				break
			}
		}
		if !ok {
			logging.Debug.Printf("Error %s", "Not able to list available numbers.")
			return fmt.Errorf("Error %s", "Not able to list available numbers.")
		}
	}
	return nil
}

func IListMyNumbers() error {
	myNumbers, err := NumberSecondaryPort.ListNumbers()
	if err != nil {
		logging.Debug.Printf("Error %s", "Not able to list available numbers.")
		return fmt.Errorf("Error %s", "Not able to list available numbers.")
	}
	for _, in := range *myNumbers {
		logging.Debug.Println(in.PhoneNumber)
		logging.Debug.Println(in.PhoneNumber)
	}
	return nil
}

func IShouldListMyNumbers(amount int) error {
	myNumbers, _ := NumberSecondaryPort.ListNumbers()
	if len(*myNumbers) != amount {
		logging.Debug.Printf("Error %s", "The list has more numbers than expected")
		return fmt.Errorf("Error %s", "The list has more numbers than expected")
	}
	return nil
}

func IReleaseAllMyNumbersExcept(exceptionList string) error {
	myNumbers, err := NumberPrimaryPort.ListNumbers()
	if err != nil {
		logging.Debug.Printf("Error %s", "Not able to list my numbers.")
		return fmt.Errorf("Error %s", "Not able to list my numbers.")
	}
	exList := strings.Split(exceptionList, ",")
	for _, a := range *myNumbers {
		exceptionNumberFound := false
		for _, e := range exList {
			pn, _ := Configuration.SelectNumber(e)
			if pn == a.PhoneNumber {
				exceptionNumberFound = true
				break
			}
		}
		if !exceptionNumberFound {
			logging.Debug.Println("Releasing " + a.PhoneNumber)
			NumberPrimaryPort.DeleteNumber(a.Sid)
		}
	}
	return nil
}

func IShouldGetNumbersFromMyList(expectedAmount int) error {
	phoneNumbers, err := NumberSecondaryPort.ListNumbers()
	if err != nil {
		return fmt.Errorf("Could not perform List Numbers.")
	}
	if len(*phoneNumbers) != expectedAmount {
		return fmt.Errorf("Expected %d phone numbers, but got %d.", expectedAmount, len(*phoneNumbers))
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
}

func ConfiguredWithVoiceUrlAs(number, voiceUrl string) error {
	services.CloseChannel = true
	Configuration.To, Configuration.ToSid = Configuration.SelectNumber(number)
	Configuration.VoiceUrl = voiceUrl
	NumberPrimaryPort.UpdateNumber()
	return nil
}
func IShouldGetFriendlyNameOn(friendlyName, number string) error {
	selectedNumber, sid := Configuration.SelectNumber(number)
	ipn, err := NumberPrimaryPort.ViewNumber(sid)
	IncomingPhoneNumber = ipn
	if err != nil {
		logging.Debug.Printf("Error %s", "Not able to view number info.")
		return fmt.Errorf("Error %s", "Not able to view number info.")
	}
	logging.Debug.Println(IncomingPhoneNumber.FriendlyName)
	if IncomingPhoneNumber != nil {
		if IncomingPhoneNumber.PhoneNumber == selectedNumber {
			if IncomingPhoneNumber.FriendlyName == friendlyName {
				return nil
			}
		}
	}
	logging.Debug.Printf("Error %s", "Not able to get friendly name on number.")
	return fmt.Errorf("Error %s", "Not able to get friendly name on number.")

}
func IViewInfo(number string) error {
	_, sid := Configuration.SelectNumber(number)
	ipn, err := NumberPrimaryPort.ViewNumber(sid)
	IncomingPhoneNumber = ipn
	if err != nil {
		logging.Debug.Printf("Error %s", "Not able to view number info.")
		return fmt.Errorf("Error %s", "Not able to view number info.")
	}
	logging.Debug.Println(IncomingPhoneNumber.FriendlyName)
	logging.Debug.Println(IncomingPhoneNumber.FriendlyName)
	return nil

}

func IShouldListMyNumbersAs(list string) error {
	myNumbers, err := NumberSecondaryPort.ListNumbers()
	if err != nil {
		logging.Debug.Printf("Error %s", "Not able to list my numbers.")
		return fmt.Errorf("Error %s", "Not able to list my numbers.")
	}
	arr := strings.Split(list, ",")
	for _, n := range *myNumbers {
		found := false

		for j := 0; j < len(arr); j++ {
			pn, _ := Configuration.SelectNumber(arr[j])
			if pn == n.PhoneNumber {
				found = true
				break
			}
		}
		if !found {
			logging.Debug.Printf("Error %s", "List is different from expected.")
			return fmt.Errorf("Error %s", "List is different from expected.")
		}
	}
	return nil
}

func IShouldGetVoiceUrlOn(voiceUrl, number string) error {
	selectedNumber, sid := Configuration.SelectNumber(number)
	ipn, err := NumberPrimaryPort.ViewNumber(sid)
	IncomingPhoneNumber = ipn
	if err != nil {
		logging.Debug.Printf("Error %s", "Not able to view number info.")
		return fmt.Errorf("Error %s", "Not able to view number info.")
	}
	logging.Debug.Println(IncomingPhoneNumber.FriendlyName)
	if IncomingPhoneNumber != nil {
		if IncomingPhoneNumber.PhoneNumber == selectedNumber {
			if IncomingPhoneNumber.VoiceURL == voiceUrl {
				logging.Debug.Printf("Successs %s", voiceUrl)
				return nil
			}
		}
	}
	logging.Debug.Printf("Error %s", "Not able to get voice url on number.")
	return fmt.Errorf("Error %s", "Not able to get voice url on number.")

}
