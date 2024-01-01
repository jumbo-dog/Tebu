package models

type PlayerSave struct {
	DiscordId    string            `bson:"discord_id"`
	LastUsername string            `bson:"last_username"`
	Money        int64             `bson:"money"`
	Resources    map[string]uint32 `bson:"resources"`
	Items        map[string]uint16 `bson:"items"`
	Progress     Progress          `bson:"progress"`
}

type Progress struct {
	Planet Planet `bson:"planet"`
	Quest  Quest  `bson:"quest"`
}

type Quest struct {
	QuestNumber   uint16 `bson:"quest_number"`
	QuestProgress uint8  `bson:"quest_progress"`
}

type Planet struct {
	Population uint64            `bson:"population"`
	Buildings  []bool            `bson:"buildings"`
	Upgrades   map[string]uint16 `bson:"upgrades"`
}
