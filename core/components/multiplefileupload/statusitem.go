package multiplefileupload

// from file "react-core/src/components/MultipleFileUpload/MultipleFileUploadStatusItem.tsx"

type StatusItemProps struct {
	// ClassName - Class to add to outer div. Optional.
	ClassName string
	// ButtonAriaLabel - Adds accessibility text to the status item deletion button. Optional.
	ButtonAriaLabel string
	// File - The file object being represented by the status item. Optional.
	File any // File 
	// OnReadStarted - A callback for when a selected file starts loading. Optional.
	OnReadStarted func(fileHandle any // File )
	// OnReadFinished - A callback for when a selected file finishes loading. Optional.
	OnReadFinished func(fileHandle any // File )
	// OnReadSuccess - A callback for when the FileReader successfully reads the file. Optional.
	OnReadSuccess func(data string, file any // File )
	// OnReadFail - A callback for when the FileReader API fails. Optional.
	OnReadFail func(error any // DOMException , onReadFail any // File )
	// OnClearClick - Clear button was clicked. Optional.
	OnClearClick any // React.MouseEventHandler 
	// CustomFileHandler - A callback to process file reading in a custom way. Optional.
	CustomFileHandler func(file any // File )
	// FileIcon - A custom icon to show in place of the generic file icon. Optional.
	FileIcon any // React.ReactNode 
	// FileName - A custom name to display for the file rather than using built in functionality to auto-fill
	// it. Optional.
	FileName string
	// FileSize - A custom file size to display for the file rather than using built in functionality to auto-fill
	// it. Optional.
	FileSize int
	// ProgressValue - A custom value to display for the progress component rather than using built in functionality
	// to auto-fill it. Optional.
	ProgressValue int
	// ProgressVariant - A custom variant to apply to the progress component rather than using built in functionality
	// to auto-fill it. Optional.
	ProgressVariant any /* "danger" | "success" | "warning" */
	// ProgressAriaLabel - Adds accessible text to the progress bar. Required when title not used and there is
	// not any label associated with the progress bar. Optional.
	ProgressAriaLabel string
	// ProgressAriaLabelledBy - Associates the progress bar with it's label for accessibility purposes. Required
	// when title not used. Optional.
	ProgressAriaLabelledBy string
	// ProgressId - Unique identifier for progress. Generated if not specified. Optional.
	ProgressId string
}
