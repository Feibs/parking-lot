# Simulation of Parking Lot Management

## Description

1. Parking lot manager would like to provide a parking lot that can park and unpark a car.

- Customers can park their cars in the lot and receive a parking ticket.
- Customers can unpark their cars by presenting the correct ticket and will get their car back.
- If a customer attempts to unpark with a wrong ticket or without a ticket or with a used ticket, then no car will be retrieved and an error message will be displayed.
- If the parking lot is already full, then customers can't park there and an error message will be displayed.

2. Parking attendant can assist with parking and unparking cars.

3. Parking attendant can park to multiple parking lots.

- Parking attendant will park cars in one lot until it reaches capacity, then move on to the next available lot.
- If all parking lots are full, then no ticket will be issued, and an error message will be displayed.
- A car cannot be parked more than once, even across different parking lots.

4. Parking attendant would like to be notified exactly when a parking lot is full or back to available, so that the car can be parked to the available lot.

- The subscribers should register/subscribe to get a notification from parking lot. Only the registered subscribers will get notification.
- The subscribers should not poll/always ask a parking lot whether it's full or not. The parking lot will notify/invoke an interface (as a mediator) immediately just after it's full or back to available.
- Given all lot is full, when a car unparked, then attendant should be able to park to a lot again.

5. The customer does not need to worry about the parking style, as the parking attendant will select a lot based on their preferred parking method.

- Parking attendant can switch the parking style at any time.
- Largest Limit Lot: Parking attendant will park cars in the lot with the most capacity.
- Most Vacant Lot: Parking attendant will park cars in the lot with the most available spaces.
- By default, attendant should park cars sequentially.

6. Parking attendant would like to have an interface so that the parking lot can be managed easily.

- The interface should include menus for registering a parking lot, parking, unparking, changing the parking style, and exiting.
- The "Change Parking Style" menu should display all available parking style options.
- The application should keep running until the attendant selects the exit option.

## Tech Stack

Go (Golang)

## How to Run

1. Clone this repository.
2. Make sure Go has been installed.  
   Go version that is used in this app: `go1.19.13`
3. Run the app: `go run .`
4. Run the unit tests: `go test ./...`

## Author

Feibs (2024)
