
# DB

<img src="https://team.stacklycode.com/challenges/db.png" />


## Auth

#### `https://challenge.stacklycode.com/api/user/register` - POST
```json
  {
    "fullname":"zEpHiRo",
    "email":"zephirotube@zephiro.dev",
    "password":"123456"
  }
```

**returns** - New User created

#### `https://challenge.stacklycode.com/api/user/login` - POST
```json
  {
    "email":"zephirotube@zephiro.dev",
    "password":"123456"
  }
```

**returns** - JWT token

## Skill

We will need the authentication token for a user and in the above, we generated a token when we logged in user 1. 

#### `https://challenge.stacklycode.com/api/skills/register`- POST
```HEADER Authorization: Bearer Token {JWT token}```
```json
{
  "name":  "C#"
}
```

**returns** - New Skill created


## Add Skill to User

#### `https://challenge.stacklycode.com/api/user/skills` - POST
```HEADER Authorization: Bearer Token {JWT token}```
```json
{
  "ruser_id":  1,
  "rskill_id":  1
}
```

**returns** - New UserSkill created