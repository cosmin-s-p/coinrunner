package coinrunner

func InitWorld() WorldData {
	return WorldData{
		Rooms: []Room{
			{
				ID:          StartPage,
				Name:        StartPage.String(),
				Description: "",
				Choices:     []Choice{StartAction, QuitAction},
				Creatures: []Creature{
					{
						Name: "",
					},
					{
						Name: "",
					},
				},
			},
			{
				ID:          MerchantGate,
				Name:        MerchantGate.String(),
				Description: "You were just created, not my accident, but because someone \nreally really wanted to buy the latest <insert random product>",
				Choices:     []Choice{MoveForwardAction},
				Creatures:   []Creature{},
			},
			{
				ID:          GatewayBridge,
				Name:        GatewayBridge.String(),
				Description: "You arrived at this fragile bridge crossing,\n can you go through or will your story end here?",
				Choices:     []Choice{MoveForwardAction},
				Creatures:   []Creature{},
			},
			{
				ID:          RiskEngineWoods,
				Name:        RiskEngineWoods.String(),
				Description: "Misterious eyes are watching from afar.. they are analyzing your every move..",
				Choices:     []Choice{MoveForwardAction},
				Creatures:   []Creature{},
			},
			{
				ID:          AcquirerPass,
				Name:        AcquirerPass.String(),
				Description: "A treacherous pass, will you go through, will you \nbe sent back, or will you be held here for a while?",
				Choices:     []Choice{MoveForwardAction},
				Creatures:   []Creature{},
			},
			{
				ID:          IssuerThrone,
				Name:        IssuerThrone.String(),
				Description: "The FINAL destination. What will be the final answer?",
				Choices:     []Choice{MoveForwardAction},
				Creatures:   []Creature{},
			},
			{
				ID:          GameOver,
				Name:        GameOver.String(),
				Description: "Press q to quit.",
				Choices:     []Choice{MoveForwardAction},
				Creatures:   []Creature{},
			},
		},
	}
}
