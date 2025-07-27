Feature: Control the minecraft server

  Scenario: Client asks the server for the status
    Given the Minecraft server is running and ready
    When a client requests the server status
    Then the system returns a status response

  Scenario: Client starts the server
    Given the Minecraft server isn't started
    When a client starts the server
    Then the server starts
    
  Scenario: Client stops the server
    Given a Minecraft server is running
    When a client stops the server
    Then the server stops