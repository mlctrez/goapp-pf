package helpertext

type ItemProps struct {
	// Children - Content rendered inside the helper text item. Optional.
	Children any //  // React.ReactNode 
	// ClassName - Additional classes applied to the helper text item. Optional.
	ClassName string
	// Component - Sets the component type of the helper text item. Optional.
	Component any //  /* "div" | "li" */
	// Variant - Variant styling of the helper text item. Optional.
	Variant any //  /* "default" | "indeterminate" | "warning" | "success" | "error" */
	// Icon - Custom icon prefixing the helper text. This property will override the default icon paired with
	// each helper text variant. Optional.
	Icon any //  // React.ReactNode 
	// IsDynamic - Flag indicating the helper text item is dynamic. This prop should be used when the text content
	// of the helper text item will never change, but the icon and styling will be dynamically updated via the
	// `variant` prop. Optional.
	IsDynamic bool
	// HasIcon - Flag indicating the helper text should have an icon. Dynamic helper texts include icons by default
	// while static helper texts do not. Optional.
	HasIcon bool
	// Id - ID for the helper text item. The value of this prop can be passed into a form component's aria-describedby
	// prop when you intend for only specific helper text items to be announced to assistive technologies. Optional.
	Id string
}

// contents of react-core/src/components/HelperText/HelperTextItem.tsx
// import * as React from 'react';
// import styles from '@patternfly/react-styles/css/components/HelperText/helper-text';
// import { css } from '@patternfly/react-styles';
// import MinusIcon from '@patternfly/react-icons/dist/esm/icons/minus-icon';
// import ExclamationTriangleIcon from '@patternfly/react-icons/dist/esm/icons/exclamation-triangle-icon';
// import CheckCircleIcon from '@patternfly/react-icons/dist/esm/icons/check-circle-icon';
// import ExclamationCircleIcon from '@patternfly/react-icons/dist/esm/icons/exclamation-circle-icon';
// 
// export interface HelperTextItemProps extends React.HTMLProps<HTMLDivElement | HTMLLIElement> {
//   /** Content rendered inside the helper text item. */
//   children?: React.ReactNode;
//   /** Additional classes applied to the helper text item. */
//   className?: string;
//   /** Sets the component type of the helper text item. */
//   component?: 'div' | 'li';
//   /** Variant styling of the helper text item. */
//   variant?: 'default' | 'indeterminate' | 'warning' | 'success' | 'error';
//   /** Custom icon prefixing the helper text. This property will override the default icon paired with each helper text variant. */
//   icon?: React.ReactNode;
//   /** Flag indicating the helper text item is dynamic. This prop should be used when the
//    * text content of the helper text item will never change, but the icon and styling will
//    * be dynamically updated via the `variant` prop.
//    */
//   isDynamic?: boolean;
//   /** Flag indicating the helper text should have an icon. Dynamic helper texts include icons by default while static helper texts do not. */
//   hasIcon?: boolean;
//   /** ID for the helper text item. The value of this prop can be passed into a form component's
//    * aria-describedby prop when you intend for only specific helper text items to be announced to
//    * assistive technologies.
//    */
//   id?: string;
// }
// 
// const variantStyle = {
//   default: '',
//   indeterminate: styles.modifiers.indeterminate,
//   warning: styles.modifiers.warning,
//   success: styles.modifiers.success,
//   error: styles.modifiers.error
// };
// 
// export const HelperTextItem: React.FunctionComponent<HelperTextItemProps> = ({
//   children,
//   className,
//   component = 'div',
//   variant = 'default',
//   icon,
//   isDynamic = false,
//   hasIcon = isDynamic,
//   id,
//   ...props
// }: HelperTextItemProps) => {
//   const Component = component as any;
//   return (
//     <Component
//       className={css(styles.helperTextItem, variantStyle[variant], isDynamic && styles.modifiers.dynamic, className)}
//       id={id}
//       {...props}
//     >
//       {icon && (
//         <span className={css(styles.helperTextItemIcon)} aria-hidden>
//           {icon}
//         </span>
//       )}
//       {hasIcon && !icon && (
//         <span className={css(styles.helperTextItemIcon)} aria-hidden>
//           {(variant === 'default' || variant === 'indeterminate') && <MinusIcon />}
//           {variant === 'warning' && <ExclamationTriangleIcon />}
//           {variant === 'success' && <CheckCircleIcon />}
//           {variant === 'error' && <ExclamationCircleIcon />}
//         </span>
//       )}
//       <span className={css(styles.helperTextItemText)}>{children}</span>
//     </Component>
//   );
// };
// HelperTextItem.displayName = 'HelperTextItem';
