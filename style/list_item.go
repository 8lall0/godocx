package style

type ListType int

const (
	TYPE_SQUARE_FILLED ListType = 1
	TYPE_BULLET_FILLED ListType = 3 // default
	TYPE_BULLET_EMPTY  ListType = 5
	TYPE_NUMBER        ListType = 7
	TYPE_NUMBER_NESTED ListType = 8
	TYPE_ALPHANUM      ListType = 9
)

type ListItem struct {
	Type     ListType
	NumStyle string
	NumId    int
}

// TODO capire numstyle
