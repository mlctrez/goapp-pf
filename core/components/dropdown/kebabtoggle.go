package dropdown

type KebabToggleProps struct {
	// Id - HTML ID of dropdown toggle. Optional.
	Id string
	// Children - Anything which can be rendered as dropdown toggle. Optional.
	Children any //  // React.ReactNode 
	// ClassName - Classess applied to root element of dropdown toggle. Optional.
	ClassName string
	// IsOpen - Flag to indicate if menu is opened. Optional.
	IsOpen bool
	// AriaLabel - Label Toggle button. Optional.
	AriaLabel string
	// OnToggle - Callback called when toggle is clicked. Optional.
	OnToggle func(value bool, event any)
	// ParentRef - Element which wraps toggle. Optional.
	ParentRef any // 
	// GetMenuRef - The menu element. Optional.
	GetMenuRef func() any // HTMLElement 
	// IsActive - Forces active state. Optional.
	IsActive bool
	// IsDisabled - Disables the dropdown toggle. Optional.
	IsDisabled bool
	// IsPlain - Display the toggle with no border or background. Optional.
	IsPlain bool
	// Type - Type to put on the button. Optional.
	Type any //  /* "button" | "submit" | "reset" */
	// BubbleEvent - Allows selecting toggle to select parent. Optional.
	BubbleEvent bool
}

// contents of react-core/src/components/Dropdown/KebabToggle.tsx
// import * as React from 'react';
// import { DropdownToggleProps } from './DropdownToggle';
// import EllipsisVIcon from '@patternfly/react-icons/dist/esm/icons/ellipsis-v-icon';
// import { Toggle } from './Toggle';
// 
// export interface KebabToggleProps extends DropdownToggleProps {
//   /** HTML ID of dropdown toggle */
//   id?: string;
//   /** Anything which can be rendered as dropdown toggle */
//   children?: React.ReactNode;
//   /** Classess applied to root element of dropdown toggle */
//   className?: string;
//   /** Flag to indicate if menu is opened */
//   isOpen?: boolean;
//   /** Label Toggle button */
//   'aria-label'?: string;
//   /** Callback called when toggle is clicked */
//   onToggle?: (value: boolean, event: any) => void;
//   /** Element which wraps toggle */
//   parentRef?: any;
//   /** The menu element */
//   getMenuRef?: () => HTMLElement;
//   /** Forces active state */
//   isActive?: boolean;
//   /** Disables the dropdown toggle */
//   isDisabled?: boolean;
//   /** Display the toggle with no border or background */
//   isPlain?: boolean;
//   /** Type to put on the button */
//   type?: 'button' | 'submit' | 'reset';
//   /** Allows selecting toggle to select parent */
//   bubbleEvent?: boolean;
// }
// 
// export const KebabToggle: React.FunctionComponent<KebabToggleProps> = ({
//   id = '',
//   // eslint-disable-next-line @typescript-eslint/no-unused-vars
//   children = null,
//   className = '',
//   isOpen = false,
//   'aria-label': ariaLabel = 'Actions',
//   parentRef = null,
//   getMenuRef = null,
//   isActive = false,
//   isPlain = false,
//   isDisabled = false,
//   bubbleEvent = false,
//   onToggle = () => undefined as void,
//   // eslint-disable-next-line @typescript-eslint/no-unused-vars
//   ref, // Types of Ref are different for React.FunctionComponent vs React.Component
//   ...props
// }: KebabToggleProps) => (
//   <Toggle
//     id={id}
//     className={className}
//     isOpen={isOpen}
//     aria-label={ariaLabel}
//     parentRef={parentRef}
//     getMenuRef={getMenuRef}
//     isActive={isActive}
//     isPlain={isPlain}
//     isDisabled={isDisabled}
//     onToggle={onToggle}
//     bubbleEvent={bubbleEvent}
//     {...props}
//   >
//     <EllipsisVIcon />
//   </Toggle>
// );
// KebabToggle.displayName = 'KebabToggle';
