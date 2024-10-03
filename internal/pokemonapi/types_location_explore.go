package pokemonapi

type LocationAreaDetail struct {
	Location struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
			Rate int `json:"rate"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
			EncounterDetails []struct {
				Method struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				ConditionValues []any `json:"condition_values"`
				Chance          int   `json:"chance"`
				MaxLevel        int   `json:"max_level"`
				MinLevel        int   `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
}
