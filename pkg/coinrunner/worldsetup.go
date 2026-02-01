package coinrunner

func InitWorld() WorldData {
	return WorldData{
		Rooms: map[GameState]Room{
			StartPage: {
				ID:          StartPage,
				Name:        StartPage.String(),
				Description: "COINRUNNER",
				Choices:     []Choice{StartAction, QuitAction},
				Creatures: []Creature{
					{
						Name: "",
					},
					{
						Name: "",
					},
				},
				NextRoom:      ProloguePage,
				HideSideViews: true,
			},
			ProloguePage: {
				ID:            ProloguePage,
				Name:          ProloguePage.String(),
				Description:   "Before we begin, please insert below the last item you purchased online or your favorite item in general:",
				Choices:       []Choice{MoveForwardAction},
				Creatures:     []Creature{},
				NextRoom:      MerchantGate,
				HideSideViews: true,
			},
			MerchantGate: {
				ID:          MerchantGate,
				Name:        MerchantGate.String(),
				Description: "You were just created, not by accident, but because someone really really wanted to buy the latest <insert random product>",
				Choices:     []Choice{MoveForwardAction},
				Creatures:   []Creature{},
				NextRoom:    GatewayBridge,
			},
			GatewayBridge: {
				ID:          GatewayBridge,
				Name:        GatewayBridge.String(),
				Description: "You arrived at this fragile bridge crossing, can you go through or will your story end here?",
				Choices:     []Choice{MoveForwardAction},
				Creatures:   []Creature{},
				NextRoom:    RiskEngineWoods,
			},
			RiskEngineWoods: {
				ID:          RiskEngineWoods,
				Name:        RiskEngineWoods.String(),
				Description: "Misterious eyes are watching from afar.. they are analyzing your every move..",
				Choices:     []Choice{MoveForwardAction},
				Creatures:   []Creature{},
				NextRoom:    AcquirerPass,
			},
			AcquirerPass: {
				ID:          AcquirerPass,
				Name:        AcquirerPass.String(),
				Description: "A treacherous pass, will you go through, will you be sent back, or will you be held here for a while?",
				Choices:     []Choice{MoveForwardAction},
				Creatures:   []Creature{},
				NextRoom:    IssuerThrone,
			},
			IssuerThrone: {
				ID:            IssuerThrone,
				Name:          IssuerThrone.String(),
				Description:   "The FINAL destination. What will be the final answer?",
				Choices:       []Choice{MoveForwardAction},
				Creatures:     []Creature{},
				NextRoom:      GameOver,
				HideSideViews: true,
			},
			GameOver: {
				ID:            GameOver,
				Name:          GameOver.String(),
				Description:   "Press q to quit.",
				Choices:       []Choice{MoveForwardAction},
				Creatures:     []Creature{},
				NextRoom:      StartPage,
				HideSideViews: true,
			},
		},
	}
}
