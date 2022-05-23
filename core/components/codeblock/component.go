package codeblock

// from file "react-core/src/components/CodeBlock/CodeBlock.tsx"

type Props struct {
	// Children - Content rendered inside the code block. Optional.
	Children any // React.ReactNode 
	// ClassName - Additional classes passed to the code block wrapper. Optional.
	ClassName string
	// Actions - Actions in the code block header. Should be wrapped with CodeBlockAction. Optional.
	Actions any // React.ReactNode 
}
