package models

type Game struct {
	Title     string
	Highlight string
	Discount  string
	Net_price string
	Platform  string
}

func GetAllGames() []Game {
	games := []Game{
		{
			Title:     "Disney Dreamlight Valley",
			Highlight: "",
			Discount:  "",
			Net_price: "29.99",
			Platform:  "Nintendo Switch",
		},
		{
			Title:     "Ori and the Will of the Wisps",
			Highlight: "Now on sale",
			Discount:  "-66%",
			Net_price: "10.99",
			Platform:  "Nintendo Switch",
		},
		{
			Title:     "Splatoon™ 3",
			Highlight: "",
			Discount:  "",
			Net_price: "59.99",
			Platform:  "Nintendo Switch",
		},
		{
			Title:     "No Man's Sky",
			Highlight: "New!",
			Discount:  "",
			Net_price: "59.99",
			Platform:  "Nintendo Switch",
		},
		{
			Title:     "Boomerang Fu",
			Highlight: "Now on sale",
			Discount:  "-86%",
			Net_price: "1.99",
			Platform:  "Nintendo Switch",
		},
		{
			Title:     "Minecraft",
			Highlight: "",
			Discount:  "",
			Net_price: "29.99",
			Platform:  "Nintendo Switch",
		},
		{
			Title:     "Mortal Kombat 11",
			Highlight: "",
			Discount:  "",
			Net_price: "49.99",
			Platform:  "Nintendo Switch",
		},
		{
			Title:     "NieR:Automata The End of YoRHa EditionNieR:Automata The End of YoRHa Edition",
			Highlight: "New!",
			Discount:  "",
			Net_price: "39.99",
			Platform:  "Nintendo Switch",
		},
		{
			Title:     "Mario Kart™ 8 Deluxe",
			Highlight: "",
			Discount:  "",
			Net_price: "59.99",
			Platform:  "Nintendo Switch",
		},
		{
			Title:     "DRAGON BALL Xenoverse 2 for Nintendo Switch",
			Highlight: "Now on sale",
			Discount:  "-85%",
			Net_price: "7.49",
			Platform:  "Nintendo Switch",
		},
		{
			Title:     "Stardew Valley",
			Highlight: "",
			Discount:  "",
			Net_price: "14.99",
			Platform:  "Nintendo Switch",
		},
		{
			Title:     "SpongeBob SquarePants: Battle for Bikini Bottom - Rehydrated",
			Highlight: "Now on sale",
			Discount:  "-50%",
			Net_price: "14.99",
			Platform:  "Nintendo Switch",
		},
		{
			Title:     "Kirby™ and the Forgotten Land",
			Highlight: "Free demo",
			Discount:  "",
			Net_price: "59.99",
			Platform:  "Nintendo Switch",
		},
		{
			Title:     "New Super Mario Bros.™ U Deluxe",
			Highlight: "",
			Discount:  "",
			Net_price: "59.99",
			Platform:  "Nintendo Switch",
		},
		{
			Title:     "Real Boxing 2",
			Highlight: "Now on sale",
			Discount:  "-86%",
			Net_price: "1.99",
			Platform:  "Nintendo Switch",
		},
	}

	return games
}
