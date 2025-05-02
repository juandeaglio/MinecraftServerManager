Feature: Server process management

  Scenario: The process starts
    Given a process is not running
    When the process starts
    Then the process should be running