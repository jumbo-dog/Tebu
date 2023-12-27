package models

type PlayerSave struct {
	DiscordId    int64             `bson:"discord_id"`
	LastUsername string            `bson:"last_username"`
	Money        int64             `bson:"money"`
	Resources    map[string]uint32 `bson:"resources"`
	Progress     Progress          `bson:"progress"`
}

type Progress struct {
	Planet Planet `bson:"planet"`
	Quest  uint16 `bson:"quest"`
}

type Planet struct {
	Population int64             `bson:"population"`
	Buildings  []bool            `bson:"buildings"`
	Upgrades   map[string]uint16 `bson:"upgrades"`
}
