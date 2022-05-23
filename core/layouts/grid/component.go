package grid

// from file "react-core/src/layouts/Grid/Grid.tsx"

type Props struct {
	// Children - content rendered inside the Grid layout. Optional.
	Children any // React.ReactNode 
	// ClassName - additional classes added to the Grid layout. Optional.
	ClassName string
	// HasGutter - Adds space between children. Optional.
	HasGutter bool
	// Span - The number of rows a column in the grid should span.  Value should be a number 1-12. Optional.
	Span any // gridItemSpanValueShape 
	// Sm - the number of columns all grid items should span on a small device. Optional.
	Sm any // gridItemSpanValueShape 
	// Md - the number of columns all grid items should span on a medium device. Optional.
	Md any // gridItemSpanValueShape 
	// Lg - the number of columns all grid items should span on a large device. Optional.
	Lg any // gridItemSpanValueShape 
	// Xl - the number of columns all grid items should span on a xLarge device. Optional.
	Xl any // gridItemSpanValueShape 
	// Xl2 - the number of columns all grid items should span on a 2xLarge device. Optional.
	Xl2 any // gridItemSpanValueShape 
	// Order - Modifies the flex layout element order property. Optional.
	Order map[string]string /* { default:string, md:string, lg:string, xl:string, 2xl:string } */
	// Component - Sets the base component to render. defaults to div. Optional.
	Component any /* any // React.ElementType  | any // React.ComponentType  */
}
