package gallery

// from file "react-core/src/layouts/Gallery/Gallery.tsx"

type Props struct {
	// Children - content rendered inside the Gallery layout. Optional.
	Children any // React.ReactNode 
	// ClassName - additional classes added to the Gallery layout. Optional.
	ClassName string
	// HasGutter - Adds space between children. Optional.
	HasGutter bool
	// MinWidths - Minimum widths at various breakpoints. Optional.
	MinWidths map[string]string /* { default:string, sm:string, md:string, lg:string, xl:string, 2xl:string } */
	// MaxWidths - Maximum widths at various breakpoints. Optional.
	MaxWidths map[string]string /* { default:string, sm:string, md:string, lg:string, xl:string, 2xl:string } */
	// Component - Sets the base component to render. defaults to div. Optional.
	Component any /* any // React.ElementType  | any // React.ComponentType  */
}
