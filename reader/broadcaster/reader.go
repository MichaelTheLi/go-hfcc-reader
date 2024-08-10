package broadcaster

type Reader struct {
	RawBroadcasterData RawBroadcasterData
}

type RawBroadcasterData struct {
	Items    []*RawBroadcasterDataItem
	Metadata RawBroadcasterMetadata
}

type RawBroadcasterMetadata struct {
	Date string `start:"11" end:"22"`
	Name string `start:"23"`
}

// RawBroadcasterDataItem Example:
// ;         23-JAN-2014  Reference Table Broadcaster
// ;
// Aaa Ghotuo
// Aab Alumu-Tesu
type RawBroadcasterDataItem struct {
	Code        string `start:"1" end:"3"`
	EnglishName string `start:"5"` // In A24 pack - some have broken chars like Ag�ncia Goaiana de Comunica��o
}

func NewBroadcasterFileReader() Reader {
	return Reader{
		RawBroadcasterData: RawBroadcasterData{
			Items: []*RawBroadcasterDataItem{},
		},
	}
}

func (source *Reader) ProcessLine(index int, text string) interface{} {
	if index == 0 {
		return &source.RawBroadcasterData.Metadata
	} else if text[0] != ';' {
		Broadcaster := RawBroadcasterDataItem{}
		source.RawBroadcasterData.Items = append(source.RawBroadcasterData.Items, &Broadcaster)
		return &Broadcaster
	}

	return nil
}
