Feature: Number    
  In order to make a seamless call transfer to another number
  As an end user
  I want to make an outgoing dial from an already made and current call
  And send digits to the dialed number

Background: Setup
    Given my test setup runs

Scenario: Number
    Given "NumberB" configured to dial and send digits "wwww1234#" to "NumberC"
    And "NumberC" configured to gather digits until "#"
    When I make a call from "NumberA" to "NumberB"
    Then "NumberC" should get digits "1234" from "NumberB"