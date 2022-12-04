package ValorantAPI

type ContentDto struct {
	Version      string            `json:"version"`
	Characters   []*ContentItemDto `json:"characters"`
	Maps         []*ContentItemDto `json:"maps"`
	Chromas      []*ContentItemDto `json:"chromas"`
	Skins        []*ContentItemDto `json:"skins"`
	SkinLevels   []*ContentItemDto `json:"skinLevels"`
	Equips       []*ContentItemDto `json:"equips"`
	GameModes    []*ContentItemDto `json:"gameModes"`
	Sprays       []*ContentItemDto `json:"sprays"`
	SprayLevels  []*ContentItemDto `json:"sprayLevels"`
	Charms       []*ContentItemDto `json:"charms"`
	CharmLevels  []*ContentItemDto `json:"charmLevels"`
	PlayerCards  []*ContentItemDto `json:"playerCards"`
	PlayerTitles []*ContentItemDto `json:"playerTitles"`
	Acts         []*ActDto         `json:"acts"`
}
type ContentItemDto struct {
	Name           string             `json:"name"`
	LocalizedNames *LocalizedNamesDto `json:"localizedNames"`
	ID             string             `json:"id"`
	AssetName      string             `json:"assetName"`
	AssetPath      string             `json:"assetPath"`
}
type LocalizedNamesDto struct {
	AR_AE string `json:"ar-AE"`
	DE_DE string `json:"de-DE"`
	EN_GB string `json:"en-GB"`
	EN_US string `json:"en-US"`
	ES_ES string `json:"es-ES"`
	ES_MX string `json:"es-MX"`
	FR_FR string `json:"fr-FR"`
	ID_ID string `json:"id-ID"`
	IT_IT string `json:"it-IT"`
	JA_JP string `json:"ja-JP"`
	KO_KR string `json:"ko-KR"`
	PL_PL string `json:"pl-PL"`
	PT_BR string `json:"pt-BR"`
	RU_RU string `json:"ru-RU"`
	TH_TH string `json:"th-TH"`
	TR_TR string `json:"tr-TR"`
	VI_VN string `json:"vi-VN"`
	ZH_CN string `json:"zh-CN"`
	ZH_TW string `json:"zh-TW"`
}
type ActDto struct {
	Name           string             `json:"name"`
	LocalizedNames *LocalizedNamesDto `json:"localizedNames"`
	ID             string             `json:"id"`
	IsActive       bool               `json:"isActive"`
}
