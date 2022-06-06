package form

type FieldGroupProps struct {
	// Children - Anything that can be rendered as form field group content. Optional.
	Children any //  // React.ReactNode 
	// ClassName - Additional classes added to the form field group. Optional.
	ClassName string
	// Header - Form field group header. Optional.
	Header any //  // React.ReactNode 
}

// contents of react-core/src/components/Form/FormFieldGroup.tsx
// import * as React from 'react';
// import { InternalFormFieldGroup } from './InternalFormFieldGroup';
// 
// export interface FormFieldGroupProps extends Omit<React.HTMLProps<HTMLDivElement>, 'label'> {
//   /** Anything that can be rendered as form field group content. */
//   children?: React.ReactNode;
//   /** Additional classes added to the form field group. */
//   className?: string;
//   /** Form field group header */
//   header?: React.ReactNode;
// }
// 
// export const FormFieldGroup: React.FunctionComponent<FormFieldGroupProps> = ({
//   children,
//   className,
//   header,
//   ...props
// }: FormFieldGroupProps) => (
//   <InternalFormFieldGroup className={className} header={header} {...props}>
//     {children}
//   </InternalFormFieldGroup>
// );
// FormFieldGroup.displayName = 'FormFieldGroup';
