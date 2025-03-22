Feature: Minecraft polling

  Scenario: Server should be running
    Given the client asks the server if it's there
    When a client asks the status
    Then I should greet the client