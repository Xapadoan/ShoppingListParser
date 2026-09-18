package test

import dom "github.com/Xapadoan/shplsprsr/ingredient/domain"

type TestSample struct {
	TestString     string
	ExpectedResult *dom.IngredientQuantity
}

func UnitRecognitionSamples() []TestSample {
	return []TestSample{
		{"[Amount] Courgettes", &dom.IngredientQuantity{Amount: 3, Unit: dom.Unit_Unit, Name: "Courgettes"}},
		{"Potiron [Amount] g", &dom.IngredientQuantity{Amount: 700, Unit: dom.Unit_Gram, Name: "Potirons"}},
		{"Huile de noix ou noisette [Amount] cuillères à soupe", &dom.IngredientQuantity{Amount: 4, Unit: dom.Unit_Cas, Name: "Huile de noix"}},
		{"citron [Amount]", &dom.IngredientQuantity{Amount: 0.5, Unit: dom.Unit_Unit, Name: "Citrons"}},
		{"[Amount] cuillère à café du curcuma", &dom.IngredientQuantity{Amount: 0.5, Unit: dom.Unit_Cac, Name: "Curcuma"}},
		{"[Amount] c. à café vinaigre de cidre", &dom.IngredientQuantity{Amount: 2, Unit: dom.Unit_Cas, Name: "Vinaigre de cidre"}},
		{"huile d'olive", &dom.IngredientQuantity{Amount: 1, Unit: dom.Unit_Unit, Name: "Huile d'olive"}},
		{"[Amount] cas chapelure", &dom.IngredientQuantity{Amount: 2, Unit: dom.Unit_Cac, Name: "Chapelure"}},
		{"[Amount] Oignon rouge", &dom.IngredientQuantity{Amount: 0.5, Unit: dom.Unit_Unit, Name: "Oignons Rouges"}},
		{"[Amount] botte persil", &dom.IngredientQuantity{Amount: 0.5, Unit: dom.Unit_Unit, Name: "Botte de persil"}},
		{"[Amount] cuillère à soupe de cumin moulu", &dom.IngredientQuantity{Amount: 0.25, Unit: dom.Unit_Cac, Name: "Cumin"}},
		{"Gousses d'ail-[Amount]", &dom.IngredientQuantity{Amount: 2, Unit: dom.Unit_Unit, Name: "Gousse d'ail"}},
		{"[Amount] Avocat(s)", &dom.IngredientQuantity{Amount: 1, Unit: dom.Unit_Unit, Name: "Avocats"}},
		{"[Amount] g d'échine de porc", &dom.IngredientQuantity{Amount: 400, Unit: dom.Unit_Gram, Name: "Echine de porc"}},
	}
}

func AmountRecognizedPatternSamples() []TestSample {
	return []TestSample{
		{"3", &dom.IngredientQuantity{Amount: 3, Unit: dom.Unit_Unit, Name: "Courgettes"}},
		{"1/2", &dom.IngredientQuantity{Amount: 0.5, Unit: dom.Unit_Unit, Name: "Citrons"}},
		{"½", &dom.IngredientQuantity{Amount: 0.5, Unit: dom.Unit_Cac, Name: "Curcuma"}},
		{"1 1/2", &dom.IngredientQuantity{Amount: 0.5, Unit: dom.Unit_Unit, Name: "Botte de persil"}},
		{"¼", &dom.IngredientQuantity{Amount: 0.25, Unit: dom.Unit_Cac, Name: "Cumin"}},
	}
}

func RealLifeSamples() []TestSample {
	return []TestSample{
		{"3 Courgettes", &dom.IngredientQuantity{Amount: 3, Unit: dom.Unit_Unit, Name: "Courgettes"}},
		{"Potiron 700 g", &dom.IngredientQuantity{Amount: 700, Unit: dom.Unit_Gram, Name: "Potirons"}},
		{"Huile de noix ou noisette 4 cuillères à soupe", &dom.IngredientQuantity{Amount: 4, Unit: dom.Unit_Cas, Name: "Huile de noix"}},
		{"citron 1/2", &dom.IngredientQuantity{Amount: 0.5, Unit: dom.Unit_Unit, Name: "Citrons"}},
		{"½ c. à café du curcuma", &dom.IngredientQuantity{Amount: 0.5, Unit: dom.Unit_Cac, Name: "Curcuma"}},
		{"2 c. à soupe vinaigre de cidre", &dom.IngredientQuantity{Amount: 2, Unit: dom.Unit_Cas, Name: "Vinaigre de cidre"}},
		{"huile d'olive", &dom.IngredientQuantity{Amount: 1, Unit: dom.Unit_Unit, Name: "Huile d'olive"}},
		{"2 cas chapelure", &dom.IngredientQuantity{Amount: 2, Unit: dom.Unit_Cac, Name: "Chapelure"}},
		{"1/2 Oignon rouge", &dom.IngredientQuantity{Amount: 0.5, Unit: dom.Unit_Unit, Name: "Oignons Rouges"}},
		{"1 1/2 botte persil", &dom.IngredientQuantity{Amount: 0.5, Unit: dom.Unit_Unit, Name: "Botte de persil"}},
		{"¼ cuillère à café de cumin moulu", &dom.IngredientQuantity{Amount: 0.25, Unit: dom.Unit_Cac, Name: "Cumin"}},
		{"Gousses d'ail-2", &dom.IngredientQuantity{Amount: 2, Unit: dom.Unit_Unit, Name: "Gousse d'ail"}},
		{"1 Avocat(s)", &dom.IngredientQuantity{Amount: 1, Unit: dom.Unit_Unit, Name: "Avocats"}},
		{"400 g d'échine de porc", &dom.IngredientQuantity{Amount: 400, Unit: dom.Unit_Gram, Name: "Echine de porc"}},
	}
}
