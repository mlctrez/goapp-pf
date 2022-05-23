package fileupload

// from file "react-core/src/components/FileUpload/FileUploadField.tsx"

type FieldProps struct {
	// Id - Unique id for the TextArea, also used to generate ids for accessible labels.
	Id string
	// Type - What type of file. Determines what is is expected by `value` (a string for 'text' and 'dataURL',
	// or a File object otherwise). Optional.
	Type any /* "text" | "dataURL" */
	// Value - Value of the file's contents (string if text file, File object otherwise). Optional.
	Value any /* string | any // File  */
	// Filename - Value to be shown in the read-only filename field. Optional.
	Filename string
	// OnChange - A callback for when the TextArea value changes. Optional.
	OnChange func(value string, filename string, event any /* any // React.ChangeEvent  | any // React.MouseEvent  */)
	// ClassName - Additional classes added to the FileUploadField container element. Optional.
	ClassName string
	// IsDisabled - Flag to show if the field is disabled. Optional.
	IsDisabled bool
	// IsReadOnly - Flag to show if the field is read only. Optional.
	IsReadOnly bool
	// IsLoading - Flag to show if a file is being loaded. Optional.
	IsLoading bool
	// SpinnerAriaValueText - Aria-valuetext for the loading spinner. Optional.
	SpinnerAriaValueText string
	// IsRequired - Flag to show if the field is required. Optional.
	IsRequired bool
	// Validated - Value to indicate if the field is modified to show that validation state. If set to success,
	// field will be modified to indicate valid state. If set to error,  field will be modified to indicate error
	// state. Optional.
	Validated any /* "success" | "error" | "default" */
	// AriaLabel - Aria-label for the TextArea. Optional.
	AriaLabel string
	// FilenamePlaceholder - Placeholder string to display in the empty filename field. Optional.
	FilenamePlaceholder string
	// FilenameAriaLabel - Aria-label for the read-only filename field. Optional.
	FilenameAriaLabel string
	// BrowseButtonText - Text for the Browse button. Optional.
	BrowseButtonText string
	// ClearButtonText - Text for the Clear button. Optional.
	ClearButtonText string
	// IsClearButtonDisabled - Flag to disable the Clear button. Optional.
	IsClearButtonDisabled bool
	// HideDefaultPreview - Flag to hide the built-in preview of the file (where available). If true, you can
	// use children to render an alternate preview. Optional.
	HideDefaultPreview bool
	// AllowEditingUploadedText - Flag to allow editing of a text file's contents after it is selected from disk.
	// Optional.
	AllowEditingUploadedText bool
	// Children - Additional children to render after (or instead of) the file preview. Optional.
	Children any // React.ReactNode 
	// OnBrowseButtonClick - A callback for when the Browse button is clicked. Optional.
	OnBrowseButtonClick func(event any // React.MouseEvent )
	// OnClearButtonClick - A callback for when the Clear button is clicked. Optional.
	OnClearButtonClick func(event any // React.MouseEvent )
	// OnTextAreaClick - A callback from when the text area is clicked. Can also be set via the onClick property
	// of FileUpload. Optional.
	OnTextAreaClick func(event any // React.MouseEvent )
	// IsDragActive - Flag to show if a file is being dragged over the field. Optional.
	IsDragActive bool
	// ContainerRef - A reference object to attach to the FileUploadField container element. Optional.
	ContainerRef any // React.Ref 
	// OnTextChange - Text area text changed. Optional.
	OnTextChange func(text string)
}
