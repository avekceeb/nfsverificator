package v4

type AuthsysParms struct {
	Stamp       uint32
	Machinename string
	Uid         uint32
	Gid         uint32
	GidLen      uint32
	// this was producing extra 4-byte field
	//Gids        uint32
}

