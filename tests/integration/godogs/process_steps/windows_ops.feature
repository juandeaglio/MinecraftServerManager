Feature: Windows Operations

  Scenario: We start a process
    Given a process is not running
    When the process starts
    Then the process should be running
