package masking

var DefaultMaskingKey = []string{
	// Personal Identification
	"ktp",
	"idcard", "passport",

	// Finance
	"cardnumber", "creditcardnumber", "ccnumber", "cardno",
	"cvv", "cvc", "securitycode", "cardsecuritycode",
	"expirydate", "expdate",
	"iban", "bic", "swift",
	"bankaccount", "accountnumber",
	"balance", "amount", "salary",
	"paymenttoken", "transactiontoken",

	// Authentication
	"privatekey", "publickey",
	"password", "passcode", "pin",
	"otp", "onetimepassword",
	"secret", "secretkey",
	"apikey", "apitoken", "jwttoken", "credential",
	"authtoken", "accesstoken", "refreshtoken", "sessiontoken",
	"jwt", "bearer",
	"signature",
}

const DefaultMaskingPercentage int = 75
const DefaultMaskingPosition = MaskLeft
const DefaultMaskingCharacter = "*"

type MaskPosition int

const (
	MaskLeft MaskPosition = iota
	MaskRight
	MaskCenter
)

type Config struct {
	MaskingKeys       []string
	MaskingPercentage int
	MaskingPosition   MaskPosition
	MaskingCharacter  string
}
