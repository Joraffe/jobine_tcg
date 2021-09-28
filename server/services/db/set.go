package db

/*---------------------------------
            Interface
----------------------------------*/

// SetNabager describes all of the methods used
// to interact with the set table in our database
type SetManager interface {
	CreateSet(setCreate *SetCreate) (*Set, error)
}


/*---------------------------------
          Data Structures
----------------------------------*/
// Set describes the required and optional data
// needed to create a new set in our set table
type Set struct {
	SetID    int64   `json:"setId"`
	SetName  string  `json:"setName"`
}



// SetCreate describes the data needed 
// to create a given set in our db
type SetCreate struct {
	SetName  string  `json:"setName"`
}


/*---------------------------------
       Method Implementations
----------------------------------*/
// CreateSet adds a new entry to the card table in our database
func (db *DB) CreateSet(setCreate *SetCreate) (*Set, error) {
	sqlStatement := `
	  INSERT INTO sets
		  (set_name)
		VALUES
		  ($1)
		RETURNING
		  set_id,
			set_name
	`
	row := db.QueryRow(
		sqlStatement,
		setCreate.SetName,
	)

	set := new(Set)
	err := row.Scan(
		&set.SetID,
		&set.SetName,
	)
	if err != nil {
		return nil, err
	}

	return set, nil
}
