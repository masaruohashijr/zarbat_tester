Feature: Number    
  In order to list available numbers from one area code
  As an end user
  I want to list at least one number from this area

  Scenario: List available numbers

    Given my test setup runs    
    When I list all available numbers
    Then I should get to buy 1 from list 

    # error tests - ran out of balance, I get the number
