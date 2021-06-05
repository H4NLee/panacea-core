package types

// this line is used by starport scaffolding # ibc/genesistype/import

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		// this line is used by starport scaffolding # ibc/genesistype/default
		// this line is used by starport scaffolding # genesis/types/default
		Owners:  map[string]*Owner{},
		Topics:  map[string]*Topic{},
		Writers: map[string]*Writer{},
		Records: map[string]*Record{},
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// this line is used by starport scaffolding # ibc/genesistype/validate

	// this line is used by starport scaffolding # genesis/types/validate

	for keyStr := range gs.Owners {
		var key OwnerCompositeKey
		if err := key.Unmarshal(keyStr); err != nil {
			return err
		}
	}
	for keyStr := range gs.Topics {
		var key TopicCompositeKey
		if err := key.Unmarshal(keyStr); err != nil {
			return err
		}
	}
	for keyStr := range gs.Writers {
		var key WriterCompositeKey
		if err := key.Unmarshal(keyStr); err != nil {
			return err
		}
	}
	for keyStr := range gs.Records {
		var key RecordCompositeKey
		if err := key.Unmarshal(keyStr); err != nil {
			return err
		}
	}
	return nil
}
