package queries

var SIGNUP_COURIERS = `mutation SIGNUP_COURIERS($email: String!, $first_name: String!, $last_name: String!, $location: point!, $middle_name: String!, $phone_number: String!, $rate: Int!, $password:String!) {
	insert_couriers_one(object: {email: $email, first_name: $first_name, last_name: $last_name, location: $location, middle_name: $middle_name, phone_number: $phone_number, rate: $rate, password: $password}) {
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
  }
}
`
