package multiplefileupload

// from file "react-core/src/components/MultipleFileUpload/MultipleFileUpload.tsx"

type Props struct {
	// Children - Content rendered inside the multi upload field. Optional.
	Children any // React.ReactNode 
	// ClassName - Class to add to outer div. Optional.
	ClassName string
	// DropzoneProps - Optional extra props to customize react-dropzone. Optional.
	DropzoneProps any // DropzoneProps 
	// IsHorizontal - Flag setting the component to horizontal styling mode. Optional.
	IsHorizontal bool
	// OnFileDrop - When files are dropped or uploaded this callback will be called with all accepted files.
	// Optional.
	OnFileDrop func(data []any // File )
}
