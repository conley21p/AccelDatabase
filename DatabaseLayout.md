## Users
- `id` - text (primary key, default nanoid())
- `username` - varchar(255)
- `password` - varchar(255)
- `driver_id` - text (references drivers)
- `owner_id` - text (references owners)
- `created_at` - timestamp with time zone
- `updated_at` - timestamp with time zone

## Driver
- `id` - text (primary key, default nanoid())
- `user_id` - text (references users)
- `first_name` - varchar(100)
- `last_name` - varchar(100)
- `contact_info_id` - text (references contact_info)
- `insurance_id` - text (references insurance)
- `license_id` - text (references license)
- `rating_id` - text (references ratings)
- `created_at` - timestamp with time zone
- `updated_at` - timestamp with time zone
- `haulers` - []Hauler (one-to-many relationship)

## Contact Info
- `id` - text (primary key, default nanoid())
- `phone_number` - varchar(20)
- `street_address` - varchar(255)
- `city` - varchar(100)
- `state` - varchar(50)
- `zip_code` - varchar(20)
- `country` - varchar(50)
- `created_at` - timestamp with time zone
- `updated_at` - timestamp with time zone

## Insurance
- `id` - text (primary key, default nanoid())
- `driver_id` - text (references drivers)
- `policy_number` - varchar(50)
- `ins_provider` - varchar(100)
- `policy_start_date` - date
- `policy_end_date` - date
- `created_at` - timestamp with time zone
- `updated_at` - timestamp with time zone

## License
- `id` - text (primary key, default nanoid())
- `driver_id` - text (references drivers)
- `license_number` - varchar(50)
- `license_expire_date` - date
- `created_at` - timestamp with time zone
- `updated_at` - timestamp with time zone

## Hauler
- `id` - text (primary key, default nanoid())
- `driver_id` - text (references drivers)
- `make` - varchar(50)
- `model` - varchar(50)
- `year` - integer
- `mileage` - double precision
- `towing_capacity` - double precision
- `created_at` - timestamp with time zone
- `updated_at` - timestamp with time zone

## Trailer
- `id` - text (primary key, default nanoid())
- `type` - varchar(50)
- `length` - double precision
- `width` - double precision
- `capacity` - double precision
- `created_at` - timestamp with time zone
- `updated_at` - timestamp with time zone

## Hauler_Trailers (Join Table)
- `hauler_id` - text (references haulers)
- `trailer_id` - text (references trailers)
- `created_at` - timestamp with time zone
- Primary Key (hauler_id, trailer_id)

## Owner
- `id` - text (primary key, default nanoid())
- `user_id` - text (references users)
- `first_name` - varchar(100)
- `last_name` - varchar(100)
- `contact_info_id` - text (references contact_info)
- `created_at` - timestamp with time zone
- `updated_at` - timestamp with time zone

## Transportation
- `id` - text (primary key, default nanoid())
- `driver_id` - text (references drivers)
- `description` - text
- `transport_date` - timestamp with time zone
- `pickup_address` - text
- `delivery_address` - text
- `deliver_by_date` - timestamp with time zone
- `pickup_by_date` - timestamp with time zone
- `pickup_available_date` - timestamp with time zone
- `accepted_offer_id` - text (references offers)
- `vehicle_id` - text (references vehicles)
- `request_price` - double precision
- `created_at` - timestamp with time zone
- `updated_at` - timestamp with time zone

## Rating
- `id` - text (primary key, default nanoid())
- `driver_id` - text (references drivers)
- `past_deliveries` - text
- `average_rating` - varchar(50)
- `created_at` - timestamp with time zone
- `updated_at` - timestamp with time zone

## Transaction
- `id` - text (primary key, default nanoid())
- `payment_method` - varchar(50)
- `amount` - double precision
- `transaction_date` - timestamp with time zone
- `created_at` - timestamp with time zone
- `updated_at` - timestamp with time zone

## Vehicle
- `id` - text (primary key, default nanoid())
- `transportation_id` - text (references transportation)
- `length` - integer
- `width` - integer
- `height` - integer
- `auto_id` - text (references autos)
- `boat_id` - text (references boats)
- `created_at` - timestamp with time zone
- `updated_at` - timestamp with time zone

## Auto
- `id` - text (primary key, default nanoid())
- `make` - varchar(50)
- `model` - varchar(50)
- `year` - integer
- `created_at` - timestamp with time zone
- `updated_at` - timestamp with time zone

## Boat
- `id` - text (primary key, default nanoid())
- `make` - varchar(50)
- `model` - varchar(50)
- `year` - integer
- `with_trailer` - boolean
- `created_at` - timestamp with time zone
- `updated_at` - timestamp with time zone

## Conversation
- `id` - text (primary key, default nanoid())
- `sender_id` - varchar(255)
- `recipient_id` - varchar(255)
- `content` - text
- `timestamp` - timestamp with time zone

## Message
- `id` - text (primary key, default nanoid())
- `subject` - varchar(255)

## Offers
- `id` - text (primary key, default nanoid())
- `driver_id` - text (references drivers)
- `amount` - double precision
- `deadline_date` - timestamp with time zone
- `created_at` - timestamp with time zone
- `updated_at` - timestamp with time zone
