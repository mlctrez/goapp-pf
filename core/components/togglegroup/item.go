package togglegroup

type ItemProps struct {
	// Text - Text rendered inside the toggle group item. Optional.
	Text any //  // React.ReactNode 
	// Icon - Icon rendered inside the toggle group item. Optional.
	Icon any //  // React.ReactNode 
	// ClassName - Additional classes added to the toggle group item. Optional.
	ClassName string
	// IsDisabled - Flag indicating if the toggle group item is disabled. Optional.
	IsDisabled bool
	// IsSelected - Flag indicating if the toggle group item is selected. Optional.
	IsSelected bool
	// AriaLabel - required when icon is used with no supporting text. Optional.
	AriaLabel string
	// ButtonId - Optional id for the button within the toggle group item. Optional.
	ButtonId string
	// OnChange - A callback for when the toggle group item selection changes. Optional.
	OnChange func(selected bool, event any /* any // React.MouseEvent  | any // React.KeyboardEvent  | any // MouseEvent  */)
}

// contents of react-core/src/components/ToggleGroup/ToggleGroupItem.tsx
// import * as React from 'react';
// import { css } from '@patternfly/react-styles';
// import styles from '@patternfly/react-styles/css/components/ToggleGroup/toggle-group';
// import { ToggleGroupItemVariant, ToggleGroupItemElement } from './ToggleGroupItemElement';
// 
// export interface ToggleGroupItemProps extends Omit<React.HTMLProps<HTMLDivElement>, 'onChange'> {
//   /** Text rendered inside the toggle group item */
//   text?: React.ReactNode;
//   /** Icon rendered inside the toggle group item */
//   icon?: React.ReactNode;
//   /** Additional classes added to the toggle group item */
//   className?: string;
//   /** Flag indicating if the toggle group item is disabled */
//   isDisabled?: boolean;
//   /** Flag indicating if the toggle group item is selected */
//   isSelected?: boolean;
//   /** required when icon is used with no supporting text */
//   'aria-label'?: string;
//   /** Optional id for the button within the toggle group item */
//   buttonId?: string;
//   /** A callback for when the toggle group item selection changes. */
//   onChange?: (selected: boolean, event: React.MouseEvent<any> | React.KeyboardEvent | MouseEvent) => void;
// }
// 
// export const ToggleGroupItem: React.FunctionComponent<ToggleGroupItemProps> = ({
//   text,
//   icon,
//   className,
//   isDisabled = false,
//   isSelected = false,
//   'aria-label': ariaLabel = '',
//   onChange = () => {},
//   buttonId = '',
//   ...props
// }: ToggleGroupItemProps) => {
//   const handleChange = (event: any): void => {
//     onChange(!isSelected, event);
//   };
// 
//   if (!ariaLabel && icon && !text) {
//     /* eslint-disable no-console */
//     console.warn('An accessible aria-label is required when using the toggle group item icon variant.');
//   }
// 
//   return (
//     <div className={css(styles.toggleGroupItem, className)} {...props}>
//       <button
//         type="button"
//         className={css(styles.toggleGroupButton, isSelected && styles.modifiers.selected)}
//         aria-pressed={isSelected}
//         onClick={handleChange}
//         {...(ariaLabel && { 'aria-label': ariaLabel })}
//         {...(isDisabled && { disabled: true })}
//         {...(buttonId && { id: buttonId })}
//       >
//         {icon ? <ToggleGroupItemElement variant={ToggleGroupItemVariant.icon}>{icon}</ToggleGroupItemElement> : null}
//         {text ? <ToggleGroupItemElement variant={ToggleGroupItemVariant.text}>{text}</ToggleGroupItemElement> : null}
//       </button>
//     </div>
//   );
// };
// ToggleGroupItem.displayName = 'ToggleGroupItem';
