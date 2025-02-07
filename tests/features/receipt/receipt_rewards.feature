Feature: Receipt Rewards System
  As a customer,
  I want to upload my receipts and earn rewards based on various conditions
  So that I can accumulate points for discounts and offers

  Background:
    Given I have a valid receipt with the required information

  Scenario: Earning points based on retailer name
    Given I have a receipt with a retailer name "SuperMarket123"
    When I submit the receipt
    Then the total points should be 14

  Scenario: Earning points for a round dollar total
    Given I have a receipt with a total of 50.00
    When I submit the receipt
    Then the total points should be 50

  Scenario: Earning points for a total that is a multiple of 0.25
    Given I have a receipt with a total of 30.25
    When I submit the receipt
    Then the total points should be 25

  Scenario: Earning points for the number of items in the receipt
    Given I have a receipt with 4 items "Shampoo, Conditioner, Soap, Toothpaste" and a final total of 100
    When I submit the receipt
    Then the total points should be 10

  Scenario: Earning points based on item description length being a multiple of 3
    Given I have a receipt with an item "butter" priced at 12.99
    When I submit the receipt
    Then the total points should be 3

    Given I have a receipt with an item "Conditioner" priced at 8.99
    When I submit the receipt
    Then the total points should be 2

  Scenario: Earning points based on the day of the purchase being odd
    Given I have a receipt with a purchase date of "2025-02-05"
    When I submit the receipt
    Then the total points should be 6

  Scenario: Earning points based on the time of purchase being between 2:00 PM and 4:00 PM
    Given I have a receipt with a purchase time of "15:30"
    When I submit the receipt
    Then the total points should be 10

  Scenario: Earning total points from multiple conditions
    Given I have a receipt with a retailer name "SuperStoreXYZ"
    Given I have a receipt with 6 items "Butter, Conditioner, Toothpaste, Soap, Towel, Toothbrush" and a final total of 100
    Given I have a receipt with a purchase date of "2025-02-05"
    Given I have a receipt with a purchase time of "15:30"
    When I submit the receipt
    Then the total points should be 121