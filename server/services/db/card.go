package db

/*---------------------------------
            Interface
----------------------------------*/

// CardManager describes all of the methods used
// to interact with the card table in our database
type CardManager interface {
	CreateCard(cardCreate *CardCreate) (*Card, error)
}


/*---------------------------------
          Data Structures
----------------------------------*/
// Card describes the required and optional data
// needed to create a new card in our card table
type Card struct {
	CardID            int64           `json:"cardId"`
	CardName          string          `json:"cardName"`
	CardType          string          `json:"cardType"`
	CardArtist        string          `json:"cardArtist"`
	CardEffect        NullStringJSON  `json:"cardEffect,omitempty"`
	CardPower         NullInt64JSON   `json:"cardPower,omitempty"`
	CardIntelligence  NullInt64JSON   `json:"cardIntelligence,omitempty"`
	CardEndurance     NullInt64JSON   `json:"cardEndurance,omitempty"`
}




// CardCreate describes the data needed 
// to create a given card in our db
type CardCreate struct {
	CardName          string          `json:"cardName"`
	CardType          string          `json:"cardType"`
	CardArtist        string          `json:"cardArtist"`
	CardEffect        NullStringJSON  `json:"cardEffect,omitempty"`
	CardPower         NullInt64JSON   `json:"cardPower,omitempty"`
	CardIntelligence  NullInt64JSON   `json:"cardIntelligence,omitempty"`
	CardEndurance     NullInt64JSON   `json:"cardEndurance,omitempty"`
}


/*---------------------------------
       Method Implementations
----------------------------------*/
// CreateCard adds a new entry to the card table in our database
func (db *DB) CreateCard(cardCreate *CardCreate) (*Card, error) {
	sqlStatement := `
	  INSERT INTO cards
		  (card_name, card_type, card_artist, card_effect, card_power, card_intelligence, card_endurance)
		VALUES
		  ($1, $2, $3, $4, $5, $6, $7)
		RETURNING
		  card_id,
			card_name,
			card_type,
			card_artist,
			card_effect,
			card_power,
			card_intelligence,
			card_endurance
	`
	row := db.QueryRow(
		sqlStatement,
		cardCreate.CardName,
		cardCreate.CardType,
		cardCreate.CardArtist,
		cardCreate.CardEffect,
		cardCreate.CardPower,
		cardCreate.CardIntelligence,
		cardCreate.CardEndurance,
	)

	card := new(Card)
	err := row.Scan(
		&card.CardID,
		&card.CardName,
		&card.CardType,
		&card.CardArtist,
		&card.CardEffect,
		&card.CardPower,
		&card.CardIntelligence,
		&card.CardEndurance,
	)
	if err != nil {
		return nil, err
	}

	return card, nil
}
