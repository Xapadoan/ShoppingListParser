package test

import dom "github.com/Xapadoan/shplsprsr/ingredient/domain"

type testSample struct {
	testString     string
	expectedResult *dom.IngredientQuantity
}

func Samples() []testSample {
	return []testSample{
		{"3 Courgettes", &dom.IngredientQuantity{Amount: 3, Unit: dom.Unit, Name: "Courgettes"}},
		{"Potiron 700 g", &dom.IngredientQuantity{Amount: 700, Unit: dom.Gram, Name: "Potirons"}},
		{"Huile de noix ou noisette 4 cuillères à soupe", &dom.IngredientQuantity{Amount: 4, Unit: dom.Cas, Name: "Huile de noix"}},
		{"citron 1/2", &dom.IngredientQuantity{Amount: 0.5, Unit: dom.Unit, Name: "Citrons"}},
		{"½ c. à café du curcuma", &dom.IngredientQuantity{Amount: 0.5, Unit: dom.Cac, Name: "Curcuma"}},
		{"2 c. à soupe vinaigre de cidre", &dom.IngredientQuantity{Amount: 2, Unit: dom.Cas, Name: "Vinaigre de cidre"}},
		{"huile d'olive", &dom.IngredientQuantity{Amount: 1, Unit: dom.Unit, Name: "Huile d'olive"}},
		{"2 cas chapelure", &dom.IngredientQuantity{Amount: 2, Unit: dom.Cac, Name: "Chapelure"}},
		{"1/2 Oignon rouge", &dom.IngredientQuantity{Amount: 0.5, Unit: dom.Unit, Name: "Oignons Rouges"}},
		{"1 1/2 botte persil", &dom.IngredientQuantity{Amount: 0.5, Unit: dom.Unit, Name: "Botte de persil"}},
		{"¼ cuillère à café de cumin moulu", &dom.IngredientQuantity{Amount: 0.25, Unit: dom.Cac, Name: "Cumin"}},
		{"Gousses d'ail-2", &dom.IngredientQuantity{Amount: 2, Unit: dom.Unit, Name: "Gousse d'ail"}},
		{"1 Avocat(s)", &dom.IngredientQuantity{Amount: 1, Unit: dom.Unit, Name: "Avocats"}},
		{"400 g d'échine de porc", &dom.IngredientQuantity{Amount: 400, Unit: dom.Gram, Name: "Echine de porc"}},
	}
}
