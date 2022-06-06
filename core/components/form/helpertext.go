package form

type HelperTextProps struct {
	// Children - Content rendered inside the Helper Text Item. Optional.
	Children any //  // React.ReactNode 
	// IsError - Adds error styling to the Helper Text  *. Optional.
	IsError bool
	// IsHidden - Hides the helper text *. Optional.
	IsHidden bool
	// ClassName - Additional classes added to the Helper Text Item. Optional.
	ClassName string
	// Icon - Icon displayed to the left of the helper text. Optional.
	Icon any //  // React.ReactNode 
	// Component - Component type of the form helper text. Optional.
	Component any //  /* "p" | "div" */
}

// contents of react-core/src/components/Form/FormHelperText.tsx
// import * as React from 'react';
// import { css } from '@patternfly/react-styles';
// import styles from '@patternfly/react-styles/css/components/Form/form';
// 
// export interface FormHelperTextProps extends React.HTMLProps<HTMLDivElement> {
//   /** Content rendered inside the Helper Text Item */
//   children?: React.ReactNode;
//   /** Adds error styling to the Helper Text  * */
//   isError?: boolean;
//   /** Hides the helper text * */
//   isHidden?: boolean;
//   /** Additional classes added to the Helper Text Item  */
//   className?: string;
//   /** Icon displayed to the left of the helper text. */
//   icon?: React.ReactNode;
//   /** Component type of the form helper text */
//   component?: 'p' | 'div';
// }
// 
// export const FormHelperText: React.FunctionComponent<FormHelperTextProps> = ({
//   children = null,
//   isError = false,
//   isHidden = true,
//   className = '',
//   icon = null,
//   component = 'p',
//   ...props
// }: FormHelperTextProps) => {
//   const Component = component as any;
//   return (
//     <Component
//       className={css(
//         styles.formHelperText,
//         isError && styles.modifiers.error,
//         isHidden && styles.modifiers.hidden,
//         className
//       )}
//       {...props}
//     >
//       {icon && <span className={css(styles.formHelperTextIcon)}>{icon}</span>}
//       {children}
//     </Component>
//   );
// };
// FormHelperText.displayName = 'FormHelperText';
