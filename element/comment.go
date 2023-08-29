package element

type Comment struct {
	TrackChange
	Initials           string
	StartElement       AbstractElement // Mi sa che qui ci va l'interfaccia
	EndElement         AbstractElement // Anche qui
	CollectionRelation bool
}
