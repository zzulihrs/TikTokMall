package utils

const ServiceName = "frontend"

type SessionUserIdKey string

const UserIdKey = SessionUserIdKey("user_id")
const UsernameKey = SessionUserIdKey("username")
const EmailKey = SessionUserIdKey("email")
