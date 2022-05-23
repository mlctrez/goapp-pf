package brand

// from file "react-core/src/components/Brand/Brand.tsx"

type Props struct {
	// Children - Transforms the Brand into a <picture> element from an <img> element. Container for <source>
	// child elements. Optional.
	Children any // React.ReactNode 
	// ClassName - Additional classes added to the either type of Brand. Optional.
	ClassName string
	// Src - Attribute that specifies the URL of a <img> Brand. For a <picture> Brand this specifies the fallback
	// <img> URL. Optional.
	Src string
	// Alt - Attribute that specifies the alt text of a <img> Brand. For a <picture> Brand this specifies the
	// fallback <img> alt text.
	Alt string
	// Widths - Widths at various breakpoints for a <picture> Brand. Optional.
	Widths map[string]string /* { default:string, sm:string, md:string, lg:string, xl:string, 2xl:string } */
	// Heights - Heights at various breakpoints for a <picture> Brand. Optional.
	Heights map[string]string /* { default:string, sm:string, md:string, lg:string, xl:string, 2xl:string } */
}
