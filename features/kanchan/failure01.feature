Feature: Gather2  
  In order to read text to the Number B (called) using a text-to-speech engine   
  As an end user
  I want that Number A (caller) listen the speech set to be read on Number B.

  Scenario: Gather something said

    Given my test setup runs 
      And "NumberA" configured to tell "we shall fight on the beaches" 
      And "NumberB" configured to gather speech 
      When I make a call from "NumberA" to "NumberB" 
      Then "NumberB" should get speech "we shall fight on the beaches" 
