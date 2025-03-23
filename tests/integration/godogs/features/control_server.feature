Feature: Control the minecraft server

  Scenario: Client asks the server for the status
    Given the server is started
    When a client asks the status
    Then the server should tell the client the status

  Scenario: Client asks populated server for status
    Given the server is started with players
    When a client asks the status with players
    Then the server should tell the client the status with players