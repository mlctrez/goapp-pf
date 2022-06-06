package dropdown

type ToggleProps struct {
	// Id - HTML ID of dropdown toggle.
	Id string
	// Type - Type to put on the button. Optional.
	Type any //  /* "button" | "submit" | "reset" */
	// Children - Anything which can be rendered as dropdown toggle. Optional.
	Children any //  // React.ReactNode 
	// ClassName - Classes applied to root element of dropdown toggle. Optional.
	ClassName string
	// IsOpen - Flag to indicate if menu is opened. Optional.
	IsOpen bool
	// OnToggle - Callback called when toggle is clicked. Optional.
	OnToggle func(isOpen bool, event any /* any // MouseEvent  | any // TouchEvent  | any // KeyboardEvent  | any // React.KeyboardEvent  | any // React.MouseEvent  */)
	// OnEnter - Callback called when the Enter key is pressed. Optional.
	OnEnter func()
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
	// IsText - Display the toggle in text only mode. Optional.
	IsText bool
	// IsPrimary - Optional.
	IsPrimary bool
	// IsSplitButton - Style the toggle as a child of a split button. Optional.
	IsSplitButton bool
	// ToggleVariant - Alternate styles for the dropdown toggle button. Optional.
	ToggleVariant any //  /* "primary" | "secondary" | "default" */
	// AriaHaspopup - Flag for aria popup. Optional.
	AriaHaspopup any //  /* bool | "listbox" | "menu" | "dialog" | "grid" | "tree" */
	// BubbleEvent - Allows selecting toggle to select parent. Optional.
	BubbleEvent bool
}

// contents of react-core/src/components/Dropdown/Toggle.tsx
// import * as React from 'react';
// import styles from '@patternfly/react-styles/css/components/Dropdown/dropdown';
// import { DropdownContext } from './dropdownConstants';
// import { css } from '@patternfly/react-styles';
// import { KEY_CODES } from '../../helpers/constants';
// import { PickOptional } from '../../helpers/typeUtils';
// 
// export interface ToggleProps {
//   /** HTML ID of dropdown toggle */
//   id: string;
//   /** Type to put on the button */
//   type?: 'button' | 'submit' | 'reset';
//   /** Anything which can be rendered as dropdown toggle */
//   children?: React.ReactNode;
//   /** Classes applied to root element of dropdown toggle */
//   className?: string;
//   /** Flag to indicate if menu is opened */
//   isOpen?: boolean;
//   /** Callback called when toggle is clicked */
//   onToggle?: (
//     isOpen: boolean,
//     event: MouseEvent | TouchEvent | KeyboardEvent | React.KeyboardEvent<any> | React.MouseEvent<HTMLButtonElement>
//   ) => void;
//   /** Callback called when the Enter key is pressed */
//   onEnter?: () => void;
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
//   /** Display the toggle in text only mode */
//   isText?: boolean;
//   /** @deprecated Use `toggleVariant` instead. Display the toggle with a primary button style */
//   isPrimary?: boolean;
//   /** Style the toggle as a child of a split button */
//   isSplitButton?: boolean;
//   /** Alternate styles for the dropdown toggle button */
//   toggleVariant?: 'primary' | 'secondary' | 'default';
//   /** Flag for aria popup */
//   'aria-haspopup'?: boolean | 'listbox' | 'menu' | 'dialog' | 'grid' | 'tree';
//   /** Allows selecting toggle to select parent */
//   bubbleEvent?: boolean;
// }
// 
// const buttonVariantStyles = {
//   default: '',
//   primary: styles.modifiers.primary,
//   secondary: styles.modifiers.secondary
// };
// 
// export class Toggle extends React.Component<ToggleProps> {
//   static displayName = 'Toggle';
//   private buttonRef = React.createRef<HTMLButtonElement>();
// 
//   static defaultProps: PickOptional<ToggleProps> = {
//     className: '',
//     isOpen: false,
//     isActive: false,
//     isDisabled: false,
//     isPlain: false,
//     isText: false,
//     isPrimary: false,
//     isSplitButton: false,
//     onToggle: () => {},
//     onEnter: () => {},
//     bubbleEvent: false
//   };
// 
//   componentDidMount = () => {
//     document.addEventListener('click', this.onDocClick);
//     document.addEventListener('touchstart', this.onDocClick);
//     document.addEventListener('keydown', this.onEscPress);
//   };
// 
//   componentWillUnmount = () => {
//     document.removeEventListener('click', this.onDocClick);
//     document.removeEventListener('touchstart', this.onDocClick);
//     document.removeEventListener('keydown', this.onEscPress);
//   };
// 
//   onDocClick = (event: MouseEvent | TouchEvent) => {
//     const { isOpen, parentRef, onToggle, getMenuRef } = this.props;
//     const menuRef = getMenuRef && getMenuRef();
//     const clickedOnToggle = parentRef && parentRef.current && parentRef.current.contains(event.target as Node);
//     const clickedWithinMenu = menuRef && menuRef.contains && menuRef.contains(event.target as Node);
//     if (isOpen && !(clickedOnToggle || clickedWithinMenu)) {
//       onToggle(false, event);
//     }
//   };
// 
//   onEscPress = (event: KeyboardEvent) => {
//     const { parentRef, getMenuRef } = this.props;
//     const keyCode = event.keyCode || event.which;
//     const menuRef = getMenuRef && getMenuRef();
//     const escFromToggle = parentRef && parentRef.current && parentRef.current.contains(event.target as Node);
//     const escFromWithinMenu = menuRef && menuRef.contains && menuRef.contains(event.target as Node);
//     if (
//       this.props.isOpen &&
//       (keyCode === KEY_CODES.ESCAPE_KEY || event.key === 'Tab') &&
//       (escFromToggle || escFromWithinMenu)
//     ) {
//       this.props.onToggle(false, event);
//       this.buttonRef.current.focus();
//     }
//   };
// 
//   onKeyDown = (event: React.KeyboardEvent<any>) => {
//     if (event.key === 'Tab' && !this.props.isOpen) {
//       return;
//     }
//     if ((event.key === 'Tab' || event.key === 'Enter' || event.key === ' ') && this.props.isOpen) {
//       if (!this.props.bubbleEvent) {
//         event.stopPropagation();
//       }
//       event.preventDefault();
// 
//       this.props.onToggle(!this.props.isOpen, event);
//     } else if ((event.key === 'Enter' || event.key === ' ') && !this.props.isOpen) {
//       if (!this.props.bubbleEvent) {
//         event.stopPropagation();
//       }
//       event.preventDefault();
// 
//       this.props.onToggle(!this.props.isOpen, event);
//       this.props.onEnter();
//     }
//   };
// 
//   render() {
//     const {
//       className,
//       children,
//       isOpen,
//       isDisabled,
//       isPlain,
//       isText,
//       isPrimary,
//       isSplitButton,
//       toggleVariant,
//       onToggle,
//       'aria-haspopup': ariaHasPopup,
//       /* eslint-disable @typescript-eslint/no-unused-vars */
//       isActive,
//       bubbleEvent,
//       onEnter,
//       parentRef,
//       getMenuRef,
//       /* eslint-enable @typescript-eslint/no-unused-vars */
//       id,
//       type,
//       ...props
//     } = this.props;
//     return (
//       <DropdownContext.Consumer>
//         {({ toggleClass }) => (
//           <button
//             {...props}
//             id={id}
//             ref={this.buttonRef}
//             className={css(
//               isSplitButton ? styles.dropdownToggleButton : toggleClass || styles.dropdownToggle,
//               isActive && styles.modifiers.active,
//               isPlain && styles.modifiers.plain,
//               isText && styles.modifiers.text,
//               isPrimary && styles.modifiers.primary,
//               buttonVariantStyles[toggleVariant],
//               className
//             )}
//             type={type || 'button'}
//             onClick={event => onToggle(!isOpen, event)}
//             aria-expanded={isOpen}
//             aria-haspopup={ariaHasPopup}
//             onKeyDown={event => this.onKeyDown(event)}
//             disabled={isDisabled}
//           >
//             {children}
//           </button>
//         )}
//       </DropdownContext.Consumer>
//     );
//   }
// }
