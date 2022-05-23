package fileupload

// from file "react-core/src/components/FileUpload/FileUpload.tsx"

type DropzoneInputPropsWithRef struct {
	// Ref - 
	Ref any // React.RefCallback 
}

type Props struct {
	// Id - Unique id for the TextArea, also used to generate ids for accessible labels.
	Id string
	// Type - What type of file. Determines what is is passed to `onChange` and expected by `value` (a string
	// for 'text' and 'dataURL', or a File object otherwise. Optional.
	Type any /* "text" | "dataURL" */
	// Value - Value of the file's contents (string if text file, File object otherwise). Optional.
	Value any /* string | any // File  */
	// Filename - Value to be shown in the read-only filename field. Optional.
	Filename string
	// OnChange - Optional.
	OnChange func(value any /* string | any // File  */, filename string, event any /* any // React.MouseEvent  | any // React.DragEvent  | any // React.ChangeEvent  */)
	// OnFileInputChange - Change event emitted from the hidden \<input type="file" \> field associated with
	// the component. Optional.
	OnFileInputChange func(event any /* any // React.ChangeEvent  | any // React.DragEvent  */, file any // File )
	// OnClick - Callback for clicking on the FileUploadField text area. By default, prevents a click in the
	// text area from opening file dialog. Optional.
	OnClick func(event any // React.MouseEvent )
	// ClassName - Additional classes added to the FileUpload container element. Optional.
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
	// HideDefaultPreview - Flag to hide the built-in preview of the file (where available). If true, you can
	// use children to render an alternate preview. Optional.
	HideDefaultPreview bool
	// AllowEditingUploadedText - Flag to allow editing of a text file's contents after it is selected from disk.
	// Optional.
	AllowEditingUploadedText bool
	// Children - Additional children to render after (or instead of) the file preview. Optional.
	Children any // React.ReactNode 
	// OnReadStarted - A callback for when a selected file starts loading. Optional.
	OnReadStarted func(fileHandle any // File )
	// OnReadFinished - A callback for when a selected file finishes loading. Optional.
	OnReadFinished func(fileHandle any // File )
	// OnReadFailed - A callback for when the FileReader API fails. Optional.
	OnReadFailed func(error any // DOMException , fileHandle any // File )
	// DropzoneProps - Optional extra props to customize react-dropzone. Optional.
	DropzoneProps any // DropzoneProps 
	// OnClearClick - Clear button was clicked. Optional.
	OnClearClick any // React.MouseEventHandler 
	// OnTextChange - Text area text changed. Optional.
	OnTextChange func(text string)
	// OnDataChange - On data changed - if type='text' or type='dataURL' and file was loaded it will call this
	// method. Optional.
	OnDataChange func(data string)
}
