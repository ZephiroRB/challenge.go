
# DB

<img src="https://team.stacklycode.com/challenges/db.png" />


## Auth

#### `/api/user/register` - POST
```json
  {
    "fullname":"zEpHiRo",
    "email":"zephirotube@zephiro.dev",
    "password":"123456"
  }
```

**returns** - New User created

#### `/api/user/login` - POST
```json
  {
    "email":"zephirotube@zephiro.dev",
    "password":"123456"
  }
```

**returns** - JWT token

## Skill

We will need the authentication token for a user and in the above, we generated a token when we logged in user 1. 

#### `/api/skills/register`- POST
```HEADER Authorization: Bearer Token {JWT token}```
```json
{
  "name":  "C#"
}
```

**returns** - New Skill created


## Add Skill to User

#### `/api/user/skills` - POST
```HEADER Authorization: Bearer Token {JWT token}```
```json
{
  "ruser_id":  1,
  "rskill_id":  1
}
```

**returns** - New UserSkill created