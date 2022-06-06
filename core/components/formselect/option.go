package formselect

type OptionProps struct {
	// ClassName - additional classes added to the Select Option. Optional.
	ClassName string
	// Value - the value for the option. Optional.
	Value any // 
	// Label - the label for the option.
	Label string
	// IsDisabled - flag indicating if the option is disabled. Optional.
	IsDisabled bool
	// IsPlaceholder - flag indicating if option will have placeholder styling applied when selected *. Optional.
	IsPlaceholder bool
}

// contents of react-core/src/components/FormSelect/FormSelectOption.tsx
// import * as React from 'react';
// 
// export interface FormSelectOptionProps extends Omit<React.HTMLProps<HTMLOptionElement>, 'disabled'> {
//   /** additional classes added to the Select Option */
//   className?: string;
//   /** the value for the option */
//   value?: any;
//   /** the label for the option */
//   label: string;
//   /** flag indicating if the option is disabled */
//   isDisabled?: boolean;
//   /** flag indicating if option will have placeholder styling applied when selected **/
//   isPlaceholder?: boolean;
// }
// 
// export const FormSelectOption: React.FunctionComponent<FormSelectOptionProps> = ({
//   className = '',
//   value = '',
//   isDisabled = false,
//   label,
//   // eslint-disable-next-line @typescript-eslint/no-unused-vars
//   isPlaceholder = false,
//   ...props
// }: FormSelectOptionProps) => (
//   <option {...props} className={className} value={value} disabled={isDisabled}>
//     {label}
//   </option>
// );
// FormSelectOption.displayName = 'FormSelectOption';
