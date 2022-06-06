package formselect

type OptionGroupProps struct {
	// Children - content rendered inside the Select Option Group. Optional.
	Children any //  // React.ReactNode 
	// ClassName - additional classes added to the Select Option. Optional.
	ClassName string
	// Label - the label for the option.
	Label string
	// IsDisabled - flag indicating if the Option Group is disabled. Optional.
	IsDisabled bool
}

// contents of react-core/src/components/FormSelect/FormSelectOptionGroup.tsx
// import * as React from 'react';
// 
// export interface FormSelectOptionGroupProps extends Omit<React.HTMLProps<HTMLOptGroupElement>, 'disabled'> {
//   /** content rendered inside the Select Option Group */
//   children?: React.ReactNode;
//   /** additional classes added to the Select Option */
//   className?: string;
//   /** the label for the option */
//   label: string;
//   /** flag indicating if the Option Group is disabled */
//   isDisabled?: boolean;
// }
// 
// export const FormSelectOptionGroup: React.FunctionComponent<FormSelectOptionGroupProps> = ({
//   children = null,
//   className = '',
//   isDisabled = false,
//   label,
//   ...props
// }: FormSelectOptionGroupProps) => (
//   <optgroup {...props} disabled={!!isDisabled} className={className} label={label}>
//     {children}
//   </optgroup>
// );
// FormSelectOptionGroup.displayName = 'FormSelectOptionGroup';
