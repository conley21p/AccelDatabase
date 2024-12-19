## Users
- `Id` - string
- `Username` - string
- `Password` - string
- `Role` - string

```
Points to `Owner`
Points to `Driver`
```

## Driver
- `Id` - string
- `LicenseNumber` - string
- `LicenseExpireDate` - date
- `InsuranceCompany` - string
- `PolicyNumber` - string

```
Points to `Users`
Points to `Contact Info`
Points to 'List of Haulers`
Points to `List of Trailers`
Points to `List of Transportations`
```

## Owner
- `Id`      - string
- `Userid`  - string
- `Details` - string

```
Points to `User`
Points to `List of Transportations`
```

## User Information
- `Id` - string
- `Username` - string
- `FirstName` - string
- `LastName` - string
- `Email` - string
- `Address` - string
- `PhoneNumber` - string

## Contact Info
- `Id` - string
- `UserId` - User (pointer)
- `ContactUserid` - string

## Transportation
- `Id` - string
- `Description` - string
- `TransportDate` - date
- `PickupAddress` - string
- `DeliveryAddress` - string
- `DeliverByDate` - date
- `PickupByDate` - date
- `PickupAvailableDate` - date
- `AcceptedOfferId` - number
- `VehicleId` - string
- `RequestPrice` - number

```
Points to `Transaction`
Points to `Package`
Points to 'Conversation`
```

## Transaction
- `Id` - string
- `TransportationId` - string
- `PaymentMethod` - string
- `Amount` - number
- `TransactionDate` - datetime

## Package
- `Id` - string
- `TransportationId` - string
- `Weight` - number

```
Points to `Vehicle`
```

## Vehicle
- `Id` - string
- `Auto` - Auto
- `Boat` - Boat

## Auto
- `Id` - string
- `Make` - string
- `Model` - string
- `Year` - number

## Boat
- `Id` - string
- `Make` - string
- `Model` - string
- `Year` - number
- `withTrailer` - boolean

## Hauler
- `Id` - string
- `DriverId` - string
- `Make` - string
- `Model` - string
- `Year` - number
- `Mileage` - number
- `TowingCapacity` - number

## Trailer
- `Id` - string
- `HaulerId` - string
- `Type` - string
- `Length` - number
- `Width` - number
- `Capacity` - number

## Conversation
- `Id` - string
- `SenderId` - string
- `RecipientId` - string
- `Content` - string
- `Timestamp` - datetime

```
Points to `List of Messages`
```

## Message
- `Id` - string
- `Participants` - list
- `Subject` - string







# Potential tables

## Company
- `Id` - string
- `Name` - string

```
This could be for both Customer and the Driver
```