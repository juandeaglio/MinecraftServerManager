Feature: Minecraft polling

  Scenario: Server should be running
    Given the server is running RCON on port 25565
    When I query the port
    Then I should see a response from the server