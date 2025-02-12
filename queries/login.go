package queries

var LOGIN_COURIER = `query CheckCourier($phone_number: String, $email: String) @cached {
  couriers(
    where: {
      _or: [
        { phone_number: { _eq: $phone_number } }
        { email: { _eq: $email } }
      ]
    }
  ) {
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
}
`

var LOGIN_ADMIN = `query CheckAdmin($email: String!) @cached
{
  admin(where: {email: {_eq: $email}}) {
    id
    email
    first_name
    last_name
    middle_name
    password
    phone_number
    created_at
    updated_at
  }
}`
