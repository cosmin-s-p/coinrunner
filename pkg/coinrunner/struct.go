package coinrunner

import "github.com/charmbracelet/lipgloss"

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
	Choices     []string
	Creatures   []Creature
}

type Creature struct {
	Name string
}

type GameData struct {
	Token           *Token
	CurrentState    GameState
	DialogueHistory map[GameState][]string
}

type UIData struct {
	Flicker      bool
	Cursor       int
	WindowWidth  int
	WindowHeight int
	Style        lipgloss.Style
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
}
