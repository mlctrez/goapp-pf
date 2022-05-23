package treeview

// from file "react-core/src/components/TreeView/TreeViewRoot.tsx"

type RootProps struct {
	// Children - Child nodes of the tree view.
	Children any // React.ReactNode 
	// HasChecks - Flag indicating if the tree view has checkboxes. Optional.
	HasChecks bool
	// HasGuides - Flag indicating if tree view has guide lines. Optional.
	HasGuides bool
	// Variant - Variant presentation styles for the tree view. Optional.
	Variant any /* "default" | "compact" | "compactNoBackground" */
	// ClassName - Class to add to add if not passed a parentItem. Optional.
	ClassName string
}
