Feature: Server process management

  Scenario: The process starts
    Given the server does not have a process
    When the server starts a process
    Then the server process should be running