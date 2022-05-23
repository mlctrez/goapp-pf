package pfselect

// from file "react-core/src/components/Select/selectConstants.tsx"

type SelectContextInterface struct {
	// OnSelect - 
	OnSelect func(event any /* any // React.MouseEvent  | any // React.ChangeEvent  */, value any /* string | any // SelectOptionObject  */, isPlaceholder bool)
	// OnClose - 
	OnClose func()
	// OnFavorite - 
	OnFavorite func(itemId string, isFavorite bool)
	// Variant - 
	Variant string
	// InputIdPrefix - 
	InputIdPrefix string
	// ShouldResetOnSelect - 
	ShouldResetOnSelect bool
}
