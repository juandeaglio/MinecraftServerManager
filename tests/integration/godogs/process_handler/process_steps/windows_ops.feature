Feature: Windows Operations

  Scenario: We start a process
    Given a process is not running
    When the process starts
    Then the process should be running

  Scenario: We stop a process
    Given a process is running
    When the process stops
    Then the process should not be running
