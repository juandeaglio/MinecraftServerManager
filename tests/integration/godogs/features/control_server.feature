Feature: Control the minecraft server

  Scenario: Client asks the server for the status
    Given the server is started
    When a client asks the status
    Then I should tell the client the status