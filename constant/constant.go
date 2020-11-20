package constant

const (
	StatusError            = "Error"
	MessageErrorEmptyField = "Data Input Tidak Boleh Kosong"
	MessageExpiredJWT      = "Token Anda Telah Expired"
	MessageFailedJWT       = "Unauthorized"
	MessageFailedServer    = "Gagal Menghubungkan Server"
)

// LOAD ENV
const (
	MessageEnvironment = "Gagal Load Environment"
	JWTKey             = "JWT_KEY"
	Environment        = "ENVIRONMENT"
	Local              = "local"
	APPPort            = "APP_PORT"
	APPHost            = "APP_HOST"
	CertificateFile    = "CERTIFICATE_FILE"
	KeyFile            = "KEY_FILE"
)

// DB ENV
const (
	DBHost       = "DB_HOST"
	DBName       = "DB_NAME"
	DBUser       = "DB_USER"
	DBPassword   = "DB_PASSWORD"
	DBSchema     = "DB_SCHEMA"
	TokenDropBox = "TOKEN_DROPBOX"
)

const Authorization = "Authorization"
const Bearer = "Bearer"
