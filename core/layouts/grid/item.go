package grid

// from file "react-core/src/layouts/Grid/GridItem.tsx"

type ItemProps struct {
	// Children - content rendered inside the Grid Layout Item. Optional.
	Children any // React.ReactNode 
	// ClassName - additional classes added to the Grid Layout Item. Optional.
	ClassName string
	// Span - the number of columns the grid item spans. Value should be a number 1-12. Optional.
	Span any // gridSpans 
	// RowSpan - the number of rows the grid item spans. Value should be a number 1-12. Optional.
	RowSpan any // gridSpans 
	// Offset - the number of columns a grid item is offset. Optional.
	Offset any // gridSpans 
	// Sm - the number of columns the grid item spans on small device. Value should be a number 1-12. Optional.
	Sm any // gridSpans 
	// SmRowSpan - the number of rows the grid item spans on medium device. Value should be a number 1-12. Optional.
	SmRowSpan any // gridSpans 
	// SmOffset - the number of columns the grid item is offset on small device. Value should be a number 1-12.
	// Optional.
	SmOffset any // gridSpans 
	// Md - the number of columns the grid item spans on medium device. Value should be a number 1-12. Optional.
	Md any // gridSpans 
	// MdRowSpan - the number of rows the grid item spans on medium device. Value should be a number 1-12. Optional.
	MdRowSpan any // gridSpans 
	// MdOffset - the number of columns the grid item is offset on medium device. Value should be a number 1-12.
	// Optional.
	MdOffset any // gridSpans 
	// Lg - the number of columns the grid item spans on large device. Value should be a number 1-12. Optional.
	Lg any // gridSpans 
	// LgRowSpan - the number of rows the grid item spans on large device. Value should be a number 1-12. Optional.
	LgRowSpan any // gridSpans 
	// LgOffset - the number of columns the grid item is offset on large device. Value should be a number 1-12.
	// Optional.
	LgOffset any // gridSpans 
	// Xl - the number of columns the grid item spans on xLarge device. Value should be a number 1-12. Optional.
	Xl any // gridSpans 
	// XlRowSpan - the number of rows the grid item spans on large device. Value should be a number 1-12. Optional.
	XlRowSpan any // gridSpans 
	// XlOffset - the number of columns the grid item is offset on xLarge device. Value should be a number 1-12.
	// Optional.
	XlOffset any // gridSpans 
	// Xl2 - the number of columns the grid item spans on 2xLarge device. Value should be a number 1-12. Optional.
	Xl2 any // gridSpans 
	// Xl2RowSpan - the number of rows the grid item spans on 2xLarge device. Value should be a number 1-12.
	// Optional.
	Xl2RowSpan any // gridSpans 
	// Xl2Offset - the number of columns the grid item is offset on 2xLarge device. Value should be a number
	// 1-12. Optional.
	Xl2Offset any // gridSpans 
	// Order - Modifies the flex layout element order property. Optional.
	Order map[string]string /* { default:string, md:string, lg:string, xl:string, 2xl:string } */
	// Component - Sets the base component to render. defaults to div. Optional.
	Component any /* any // React.ElementType  | any // React.ComponentType  */
}
