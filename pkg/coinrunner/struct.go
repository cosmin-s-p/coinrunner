package coinrunner

type GeneralModel struct {
	WorldData WorldData
	GameData  GameData
	UIData    UIData
}

type WorldData struct {
	Rooms []Room
}

type Room struct {
	ID          GameState
	Name        string
	Description string
	Choices     []Choice
	Creatures   []Creature
	NextRoom    GameState
}

type Choice int

const (
	StartAction Choice = iota
	QuitAction
	MoveForwardAction
	MoveBackwardAction
	ScanAction
	ShieldAction
	IdentifyAction
	StatusAction
)

func (c Choice) String() string {
	return Choices[c]
}

var Choices = map[Choice]string{
	StartAction:        "Start",
	QuitAction:         "Quit",
	MoveForwardAction:  "Move Forward",
	MoveBackwardAction: "Move Backward",
	ScanAction:         "Scan",
	ShieldAction:       "Shield",
	IdentifyAction:     "Identify",
	StatusAction:       "Status",
}

type Creature struct {
	Name string
}

type GameData struct {
	Token           *Token
	CurrentState    GameState
	DialogueHistory map[GameState][]string
	IsIdle          bool // more flags?
}

type UIData struct {
	Flicker      bool
	Cursor       int
	WindowWidth  int
	WindowHeight int
}

type Token struct {
	IdempotencyKey string
	RiskScore      int
	PaidAmount     float32
	SenderIp       string
}

type GameState int

const (
	StartPage GameState = iota
	MerchantGate
	GatewayBridge
	RiskEngineWoods
	AcquirerPass
	IssuerThrone
	GameOver
)

func (s GameState) String() string {
	return GameStates[s]
}

var GameStates = map[GameState]string{
	StartPage:       "Start page",
	MerchantGate:    "Merchant Gate",
	GatewayBridge:   "Gateway Bridge",
	RiskEngineWoods: "Risk Engine Woods",
	AcquirerPass:    "Acquirer Pass",
	IssuerThrone:    "Issuer Throne",
	GameOver:        "Game Over",
}
