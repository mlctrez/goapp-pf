package fileupload

type FieldProps struct {
	// Id - Unique id for the TextArea, also used to generate ids for accessible labels.
	Id string
	// Type - What type of file. Determines what is is expected by `value` (a string for 'text' and 'dataURL',
	// or a File object otherwise). Optional.
	Type any //  /* "text" | "dataURL" */
	// Value - Value of the file's contents (string if text file, File object otherwise). Optional.
	Value any //  /* string | any // File  */
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
	Validated any //  /* "success" | "error" | "default" */
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
	Children any //  // React.ReactNode 
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
	ContainerRef any //  // React.Ref 
	// OnTextChange - Text area text changed. Optional.
	OnTextChange func(text string)
}

// contents of react-core/src/components/FileUpload/FileUploadField.tsx
// import * as React from 'react';
// import styles from '@patternfly/react-styles/css/components/FileUpload/file-upload';
// import { css } from '@patternfly/react-styles';
// import { InputGroup } from '../InputGroup';
// import { TextInput } from '../TextInput';
// import { Button, ButtonVariant } from '../Button';
// import { TextArea, TextAreResizeOrientation } from '../TextArea';
// import { Spinner, spinnerSize } from '../Spinner';
// import { fileReaderType } from '../../helpers/fileUtils';
// 
// export interface FileUploadFieldProps extends Omit<React.HTMLProps<HTMLDivElement>, 'value' | 'onChange'> {
//   /** Unique id for the TextArea, also used to generate ids for accessible labels */
//   id: string;
//   /** What type of file. Determines what is is expected by `value`
//    * (a string for 'text' and 'dataURL', or a File object otherwise). */
//   type?: 'text' | 'dataURL';
//   /** Value of the file's contents
//    * (string if text file, File object otherwise) */
//   value?: string | File;
//   /** Value to be shown in the read-only filename field. */
//   filename?: string;
//   /** A callback for when the TextArea value changes. */
//   onChange?: (
//     value: string,
//     filename: string,
//     event:
//       | React.ChangeEvent<HTMLTextAreaElement> // User typed in the TextArea
//       | React.MouseEvent<HTMLButtonElement, MouseEvent> // User clicked Clear button
//   ) => void;
//   /** Additional classes added to the FileUploadField container element. */
//   className?: string;
//   /** Flag to show if the field is disabled. */
//   isDisabled?: boolean;
//   /** Flag to show if the field is read only. */
//   isReadOnly?: boolean;
//   /** Flag to show if a file is being loaded. */
//   isLoading?: boolean;
//   /** Aria-valuetext for the loading spinner */
//   spinnerAriaValueText?: string;
//   /** Flag to show if the field is required. */
//   isRequired?: boolean;
//   /** Value to indicate if the field is modified to show that validation state.
//    * If set to success, field will be modified to indicate valid state.
//    * If set to error,  field will be modified to indicate error state.
//    */
//   validated?: 'success' | 'error' | 'default';
//   /** Aria-label for the TextArea. */
//   'aria-label'?: string;
//   /** Placeholder string to display in the empty filename field */
//   filenamePlaceholder?: string;
//   /** Aria-label for the read-only filename field */
//   filenameAriaLabel?: string;
//   /** Text for the Browse button */
//   browseButtonText?: string;
//   /** Text for the Clear button */
//   clearButtonText?: string;
//   /** Flag to disable the Clear button */
//   isClearButtonDisabled?: boolean;
//   /** Flag to hide the built-in preview of the file (where available).
//    * If true, you can use children to render an alternate preview. */
//   hideDefaultPreview?: boolean;
//   /** Flag to allow editing of a text file's contents after it is selected from disk */
//   allowEditingUploadedText?: boolean;
//   /** Additional children to render after (or instead of) the file preview. */
//   children?: React.ReactNode;
// 
//   // Props available in FileUploadField but not FileUpload:
// 
//   /** A callback for when the Browse button is clicked. */
//   onBrowseButtonClick?: (event: React.MouseEvent<HTMLButtonElement, MouseEvent>) => void;
//   /** A callback for when the Clear button is clicked. */
//   onClearButtonClick?: (event: React.MouseEvent<HTMLButtonElement, MouseEvent>) => void;
//   /** A callback from when the text area is clicked. Can also be set via the onClick property of FileUpload. */
//   onTextAreaClick?: (event: React.MouseEvent<HTMLTextAreaElement, MouseEvent>) => void;
//   /** Flag to show if a file is being dragged over the field */
//   isDragActive?: boolean;
//   /** A reference object to attach to the FileUploadField container element. */
//   containerRef?: React.Ref<HTMLDivElement>;
//   /** Text area text changed */
//   onTextChange?: (text: string) => void;
// }
// 
// export const FileUploadField: React.FunctionComponent<FileUploadFieldProps> = ({
//   id,
//   type,
//   value = '',
//   filename = '',
//   onChange = () => {},
//   onBrowseButtonClick = () => {},
//   onClearButtonClick = () => {},
//   onTextAreaClick,
//   onTextChange,
//   className = '',
//   isDisabled = false,
//   isReadOnly = false,
//   isLoading = false,
//   spinnerAriaValueText,
//   isRequired = false,
//   isDragActive = false,
//   validated = 'default' as 'success' | 'error' | 'default',
//   'aria-label': ariaLabel = 'File upload',
//   filenamePlaceholder = 'Drag a file here or browse to upload',
//   filenameAriaLabel = filename ? 'Read only filename' : filenamePlaceholder,
//   browseButtonText = 'Browse...',
//   clearButtonText = 'Clear',
//   isClearButtonDisabled = !filename && !value,
//   containerRef = null as React.Ref<HTMLDivElement>,
//   allowEditingUploadedText = false,
//   hideDefaultPreview = false,
//   children = null,
// 
//   ...props
// }: FileUploadFieldProps) => {
//   const onTextAreaChange = (newValue: string, event: React.ChangeEvent<HTMLTextAreaElement>) => {
//     onChange(newValue, filename, event);
//     onTextChange?.(newValue);
//   };
//   return (
//     <div
//       className={css(
//         styles.fileUpload,
//         isDragActive && styles.modifiers.dragHover,
//         isLoading && styles.modifiers.loading,
//         className
//       )}
//       ref={containerRef}
//       {...props}
//     >
//       <div className={styles.fileUploadFileSelect}>
//         <InputGroup>
//           <TextInput
//             isReadOnly // Always read-only regardless of isReadOnly prop (which is just for the TextArea)
//             isDisabled={isDisabled}
//             id={`${id}-filename`}
//             name={`${id}-filename`}
//             aria-label={filenameAriaLabel}
//             placeholder={filenamePlaceholder}
//             aria-describedby={`${id}-browse-button`}
//             value={filename}
//           />
//           <Button
//             id={`${id}-browse-button`}
//             variant={ButtonVariant.control}
//             onClick={onBrowseButtonClick}
//             isDisabled={isDisabled}
//           >
//             {browseButtonText}
//           </Button>
//           <Button
//             variant={ButtonVariant.control}
//             isDisabled={isDisabled || isClearButtonDisabled}
//             onClick={onClearButtonClick}
//           >
//             {clearButtonText}
//           </Button>
//         </InputGroup>
//       </div>
//       <div className={styles.fileUploadFileDetails}>
//         {!hideDefaultPreview && type === fileReaderType.text && (
//           <TextArea
//             readOnly={isReadOnly || (!!filename && !allowEditingUploadedText)}
//             disabled={isDisabled}
//             isRequired={isRequired}
//             resizeOrientation={TextAreResizeOrientation.vertical}
//             validated={validated}
//             id={id}
//             name={id}
//             aria-label={ariaLabel}
//             value={value as string}
//             onChange={onTextAreaChange}
//             onClick={onTextAreaClick}
//           />
//         )}
//         {isLoading && (
//           <div className={styles.fileUploadFileDetailsSpinner}>
//             <Spinner size={spinnerSize.lg} aria-valuetext={spinnerAriaValueText} />
//           </div>
//         )}
//       </div>
//       {children}
//     </div>
//   );
// };
// FileUploadField.displayName = 'FileUploadField';
