package steps

import (
	"fmt"
	"zarbat_test/internal/logging"
)

func IShouldGetToSeeAsTheFriendlyNameForMyAccount(friendlyName string) error {
	var err error
	if AccountInfo.FriendlyName == "" {
		AccountInfo, err = AccountPrimaryPort.ViewAccount()
		if err != nil {
			return fmt.Errorf("Error %s", err.Error())
		}
	}
	if AccountInfo.FriendlyName != friendlyName {
		return fmt.Errorf("Error %s", "account update did not work")
	}
	return nil
}

func IUpdateTheFriendlyNameForMyAccountTo(friendlyName string) error {
	err := AccountPrimaryPort.UpdateAccount(friendlyName)
	if err != nil {
		return fmt.Errorf("Error %s", err.Error())
	}
	return nil
}

func IViewMyAccountInformation() error {
	var err error
	AccountInfo, err = AccountPrimaryPort.ViewAccount()
	logging.Debug.Printf("My friendly name is %s\n", AccountInfo.FriendlyName)
	if err != nil {
		return fmt.Errorf("Error %s", err.Error())
	}
	return nil
}
