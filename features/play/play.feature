Feature: Play    
  In order to play a tone with frequencies
  As an end user
  I want to set a tone to play after the call is established
  And should be able to record and extract these frequencies
    
  Scenario: Play a tone
    Given my test setup runs 
    And "NumberD" configured to play tone "5000,10,850"
    And "NumberE" configured to record calls for download
    When I make a call from "NumberD" to "NumberE"
    Then "NumberE" should be able to listen to frequencies "850"
    And "NumberD" should be reset
    And "NumberE" should be reset
  