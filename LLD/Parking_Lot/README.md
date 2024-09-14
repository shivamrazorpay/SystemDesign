# Question: Parking Lot System Design

Design a parking lot system that supports the following requirements:

## Functional Requirements:
The car park can accommodate multiple types of vehicles: cars (sedans, SUVs), bikes, and autos.
The car park has a limited number of spots for each vehicle type.
Each vehicle should be assigned a parking spot when it enters the parking lot.
The parking system should record the entry time and calculate the parking fee based on the duration of stay when the vehicle exits.
The system should support different pricing models for each vehicle type.
The system should provide a mechanism for reserving parking spots.
You should be able to query the availability of parking spots at any time.
Non-Functional Requirements:
The design should be scalable to handle thousands of vehicles.
The system should be efficient in terms of time and space complexity.
The system should allow for future extensibility, such as supporting electric vehicle (EV) charging stations.
## Constraints:
**Pricing rules:**
</br>
First 30 minutes are free for all vehicle types.
After that, the fee is calculated based on the type of vehicle and the number of hours stayed.
The system should charge only for full hours beyond the free period.</br>
**Example fee structure:**
</br>Car (SUV): $5 per hour
</br>Car (Sedan): $3 per hour
</br>Bike: $1 per hour
</br>Auto: $2 per hour
## Discussion Points:
How would you design the classes and relationships between them (e.g., Vehicle, ParkingSpot, ParkingLot, Ticket)?
How would you ensure that the parking lot does not allow more vehicles than its capacity?
How would you implement the pricing logic in a scalable way, considering future changes or additions to pricing models?
What data structures would you use to keep track of available parking spots, vehicles, and their entry/exit times?
How would you handle the situation where a vehicle is attempting to enter but no spots are available for its type?