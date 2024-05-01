package constants

const (
	DATABASE           = "eccomerce"
	COLLECTION_USER    = "users"
	COLLECTION_PRODUCT = "products"
	COLLECTION_CARTS   = "carts"
	COLLECTION_ORDERS  = "orders"

	JWT_TOKEN_ISSUEER          = "faizauthar12"
	JWT_TOKEN_LIFESPAN         = 168 // 1 week in hours
	JWT_REFRESH_TOKEN_LIFESPAN = 720 // 30 days in hours

	CONFIG_SMTP_HOST = "smtp.gmail.com"
	CONFIG_SMTP_PORT = 587

	MAIL_SENDER = "faizauthar@gmail.com"
)
