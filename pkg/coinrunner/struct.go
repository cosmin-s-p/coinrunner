package coinrunner

import (
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
)

type GeneralModel struct {
	WorldData WorldData
	GameData  GameData
	UIData    UIData
}

type WorldData struct {
	Rooms map[GameState]Room
}

type Room struct {
	ID            GameState
	Name          string
	Description   string
	Choices       []Choice
	Creatures     []Creature
	NextRoom      GameState
	HideSideViews bool
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
	IP              string
	CurrentState    GameState
	LatestDialogue  string
	DialogueHistory []string
	IsIdle          bool // more flags?
	CanMoveForward  bool
	FavoriteItem    string
}

type UIData struct {
	Flicker         bool
	Cursor          int
	WindowWidth     int
	WindowHeight    int
	TitleHeight     int
	SidePanelWidth  int
	SidePanelHeight int
	TextInput       textinput.Model
	Viewport        viewport.Model
}

type Token struct {
	IdempotencyKey string
	RiskScore      int
	PaidAmount     float32
	SenderIp       string
	Timestamp      time.Time
}

type GameState int

const (
	StartPage GameState = iota
	ProloguePage
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
	ProloguePage:    "Prologue",
	MerchantGate:    "Merchant Gate",
	GatewayBridge:   "Gateway Bridge",
	RiskEngineWoods: "Risk Engine Woods",
	AcquirerPass:    "Acquirer Pass",
	IssuerThrone:    "Issuer Throne",
	GameOver:        "Game Over",
}
