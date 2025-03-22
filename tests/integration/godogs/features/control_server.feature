Feature: Minecraft polling

  Scenario: Server should be running
    Given the client asks the server if it's there
    When I receive the client's request
    Then I should greet the client