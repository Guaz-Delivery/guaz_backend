package queries

var SIGNUP_COURIER = `mutation SIGNUP_COURIERS($email: String!, $first_name: String!, $last_name: String!, $location: point!, $middle_name: String!, $phone_number: String!, $rate: Int!, $password: String!, $shipment_size: String!, $shipment_range: String!, $profile_picture: String!) {
  insert_couriers_one(object: {email: $email, first_name: $first_name, last_name: $last_name, location: $location, middle_name: $middle_name, phone_number: $phone_number, rate: $rate, password: $password, profile_picture: $profile_picture, shipment_range: $shipment_range, shipment_size: $shipment_size}) {
    created_at
    email
    first_name
    id
    is_verified
    last_name
    location
    middle_name
    phone_number
    rate
    updated_at
    profile_picture
    shipment_range
    shipment_size
  }
}
`
