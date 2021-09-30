package db

/*---------------------------------
            Interface
----------------------------------*/

// CardManager describes all of the methods used
// to interact with the card table in our database
type CardManager interface {
	GetAllCards() ([]*Card, error)
	GetCardsByCardType(cardType string) ([]*Card, error)
	GetCardByCardName(cardName string) (*Card, error)

	CreateCard(cardCreate *CardCreate) (int64, error)
	UpdateCard(cardUpdate *CardUpdate) (int64, error)
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

// CardUpdate describes the data needed 
// to update a given card in our db
type CardUpdate struct {
	CardID            int64           `json:"cardId,omitempty"`
	CardName          NullStringJSON  `json:"cardName,omitempty"`
	CardType          NullStringJSON  `json:"cardType,omitempty"`
	CardArtist        NullStringJSON  `json:"cardArtist,omitempty"`
	CardEffect        NullStringJSON  `json:"cardEffect,omitempty"`
	CardPower         NullInt64JSON   `json:"cardPower,omitempty"`
	CardIntelligence  NullInt64JSON   `json:"cardIntelligence,omitempty"`
	CardEndurance     NullInt64JSON   `json:"cardEndurance,omitempty"`
}


/*---------------------------------
       Method Implementations
----------------------------------*/

// GetAllCards fetches all cards in our db
func (db *DB) GetAllCards() ([]*Card, error) {
	sqlStatement := `
	  SELECT
		  card_id,
			card_name,
			card_type,
			card_artist,
			card_effect,
			card_power,
			card_intelligence,
			card_endurance
		FROM
		  cards
	`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cards := make([]*Card, 0)
	for rows.Next() {
		card := new(Card)
		err := rows.Scan(
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

		cards = append(cards, card)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return cards, nil
}


// GetCardsByCardType fetches all cards with a specific type
func (db *DB) GetCardsByCardType(cardType string) ([]*Card, error) {
	sqlStatement := `
	  SELECT
		  card_id,
			card_name,
			card_type,
			card_artist,
			card_effect,
			card_power,
			card_intelligence,
			card_endurance
		FROM
		  cards
		WHERE
		  card_type = $1
	`
	rows, err := db.Query(sqlStatement, cardType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cards := make([]*Card, 0)
	for rows.Next() {
		card := new(Card)
		err := rows.Scan(
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
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return cards, nil
}


// GetCardByCardName fetches a specific card by its name from the db
func (db *DB) GetCardByCardName(cardName string) (*Card, error) {
	sqlStatement := `
	  SELECT
			card_id,
			card_name,
			card_type,
			card_artist,
			card_effect,
			card_power,
			card_intelligence,
			card_endurance
		FROM
		  cards
		WHERE
		  card_name = $1
	`
	row := db.QueryRow(sqlStatement, cardName)
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


// CreateCard adds a new entry to the card table in our database
func (db *DB) CreateCard(cardCreate *CardCreate) (int64, error) {
	var cardID int64
	sqlStatement := `
	  INSERT INTO cards
		  (card_name, card_type, card_artist, card_effect, card_power, card_intelligence, card_endurance)
		VALUES
		  ($1, $2, $3, $4, $5, $6, $7)
		RETURNING
		  card_id
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

	err := row.Scan(&cardID)
	if err != nil {
		return 0, err
	}

	return cardID, nil
}


// UpdateCard updates an existing entry in the cards table
func (db *DB) UpdateCard(cardUpdate *CardUpdate) (int64, error) {
	var cardID int64
	sqlStatement := `
	  UPDATE
		  cards
		SET
		  card_name = $1,
			card_type = $2,
			card_artist = $3,
			card_effect = $4,
			card_power = $5,
			card_intelligence = $6,
			card_endurance = $7
		WHERE
		  card_id = $8
		RETURNING
		  card_id
	`

	row := db.QueryRow(
		sqlStatement,
		cardUpdate.CardName,
		cardUpdate.CardType,
		cardUpdate.CardArtist,
		cardUpdate.CardEffect,
		cardUpdate.CardPower,
		cardUpdate.CardIntelligence,
		cardUpdate.CardEndurance,
		cardUpdate.CardID,
	)
	err := row.Scan(&cardID)
	if err != nil {
		return 0, err
	}

	return cardID, nil
}
