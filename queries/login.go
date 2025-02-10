package queries

var LOGIN_COURIER = `query CheckCourier($phone_number: String!, $email: String!) @cached {
  couriers(where: {phone_number: {_eq: $phone_number}, _or: {email: {_eq: $email}}}) {
    created_at
    email
    first_name
    id
    is_verified
    last_name
    location
    middle_name
    password
    phone_number
    rate
    updated_at
  }
}`
