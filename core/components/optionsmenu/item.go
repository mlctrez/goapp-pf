package optionsmenu

type ItemProps struct {
	// Children - Anything which can be rendered as an options menu item. Optional.
	Children any //  // React.ReactNode 
	// ClassName - Classes applied to root element of an options menu item. Optional.
	ClassName string
	// IsSelected - Render options menu item as selected. Optional.
	IsSelected bool
	// IsDisabled - Render options menu item as disabled option. Optional.
	IsDisabled bool
	// OnSelect - Callback for when this options menu item is selected. Optional.
	OnSelect func(event any /* any // React.MouseEvent  | any // React.KeyboardEvent  */)
	// Id - Unique id of this options menu item. Optional.
	Id string
}

// contents of react-core/src/components/OptionsMenu/OptionsMenuItem.tsx
// import * as React from 'react';
// import { css } from '@patternfly/react-styles';
// import styles from '@patternfly/react-styles/css/components/OptionsMenu/options-menu';
// import { DropdownItem } from '../Dropdown';
// import CheckIcon from '@patternfly/react-icons/dist/esm/icons/check-icon';
// 
// export interface OptionsMenuItemProps
//   extends Omit<React.HTMLProps<HTMLAnchorElement>, 'onSelect' | 'onClick' | 'onKeyDown' | 'type'> {
//   /** Anything which can be rendered as an options menu item */
//   children?: React.ReactNode;
//   /** Classes applied to root element of an options menu item */
//   className?: string;
//   /** Render options menu item as selected */
//   isSelected?: boolean;
//   /** Render options menu item as disabled option */
//   isDisabled?: boolean;
//   /** Callback for when this options menu item is selected */
//   onSelect?: (event?: React.MouseEvent<HTMLAnchorElement> | React.KeyboardEvent) => void;
//   /** Unique id of this options menu item */
//   id?: string;
// }
// 
// export const OptionsMenuItem: React.FunctionComponent<OptionsMenuItemProps> = ({
//   children = null as React.ReactNode,
//   isSelected = false,
//   onSelect = () => null as any,
//   id = '',
//   isDisabled,
//   ...props
// }: OptionsMenuItemProps) => (
//   <DropdownItem
//     id={id}
//     component="button"
//     isDisabled={isDisabled}
//     onClick={(event: any) => onSelect(event)}
//     {...(isDisabled && { 'aria-disabled': true })}
//     {...props}
//   >
//     {children}
//     {isSelected && (
//       <span className={css(styles.optionsMenuMenuItemIcon)}>
//         <CheckIcon aria-hidden={isSelected} />
//       </span>
//     )}
//   </DropdownItem>
// );
// OptionsMenuItem.displayName = 'OptionsMenuItem';
