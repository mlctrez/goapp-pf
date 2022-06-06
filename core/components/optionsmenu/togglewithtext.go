package optionsmenu

type ToggleWithTextProps struct {
	// ParentId - Id of the parent options menu component. Optional.
	ParentId string
	// ToggleText - Content to be rendered inside the options menu toggle as text or another non-interactive
	// element.
	ToggleText any //  // React.ReactNode 
	// ToggleTextClassName - classes to be added to the options menu toggle text. Optional.
	ToggleTextClassName string
	// ToggleButtonContents - Content to be rendered inside the options menu toggle button. Optional.
	ToggleButtonContents any //  // React.ReactNode 
	// ToggleButtonContentsClassName - Classes to be added to the options menu toggle button. Optional.
	ToggleButtonContentsClassName string
	// OnToggle - Callback for when this options menu is toggled. Optional.
	OnToggle func(event bool)
	// OnEnter - Inner function to indicate open on Enter. Optional.
	OnEnter func(event any /* any // React.MouseEvent  | any // React.KeyboardEvent  */)
	// IsOpen - Flag to indicate if menu is open. Optional.
	IsOpen bool
	// IsPlain - Flag to indicate if the button is plain. Optional.
	IsPlain bool
	// IsActive - Forces display of the active state of the options menu button. Optional.
	IsActive bool
	// IsDisabled - Disables the options menu toggle. Optional.
	IsDisabled bool
	// ParentRef - Optional.
	ParentRef any //  // React.RefObject 
	// AriaHaspopup - Indicates that the element has a popup context menu or sub-level menu. Optional.
	AriaHaspopup any //  /* bool | "dialog" | "menu" | "listbox" | "tree" | "grid" */
	// AriaLabel - Provides an accessible name for the button when an icon is used instead of text. Optional.
	AriaLabel string
	// IsText - Optional.
	IsText bool
	// GetMenuRef - Optional.
	GetMenuRef func() any // HTMLElement 
}

// contents of react-core/src/components/OptionsMenu/OptionsMenuToggleWithText.tsx
// import * as React from 'react';
// import { css } from '@patternfly/react-styles';
// import { KEY_CODES } from '../../helpers/constants';
// import styles from '@patternfly/react-styles/css/components/OptionsMenu/options-menu';
// 
// export interface OptionsMenuToggleWithTextProps extends React.HTMLProps<HTMLDivElement> {
//   /** Id of the parent options menu component */
//   parentId?: string;
//   /** Content to be rendered inside the options menu toggle as text or another non-interactive element */
//   toggleText: React.ReactNode;
//   /** classes to be added to the options menu toggle text */
//   toggleTextClassName?: string;
//   /** Content to be rendered inside the options menu toggle button */
//   toggleButtonContents?: React.ReactNode;
//   /** Classes to be added to the options menu toggle button */
//   toggleButtonContentsClassName?: string;
//   /** Callback for when this options menu is toggled */
//   onToggle?: (event: boolean) => void;
//   /** Inner function to indicate open on Enter */
//   onEnter?: (event: React.MouseEvent<HTMLButtonElement> | React.KeyboardEvent<Element>) => void;
//   /** Flag to indicate if menu is open */
//   isOpen?: boolean;
//   /** Flag to indicate if the button is plain */
//   isPlain?: boolean;
//   /** Forces display of the active state of the options menu button */
//   isActive?: boolean;
//   /** Disables the options menu toggle */
//   isDisabled?: boolean;
//   /** @hide Internal parent reference */
//   parentRef?: React.RefObject<HTMLElement>;
//   /** Indicates that the element has a popup context menu or sub-level menu */
//   'aria-haspopup'?: boolean | 'dialog' | 'menu' | 'listbox' | 'tree' | 'grid';
//   /** Provides an accessible name for the button when an icon is used instead of text */
//   'aria-label'?: string;
//   /** @hide Display the toggle in text only mode. */
//   isText?: boolean;
//   /** @hide The menu element */
//   getMenuRef?: () => HTMLElement;
// }
// 
// export const OptionsMenuToggleWithText: React.FunctionComponent<OptionsMenuToggleWithTextProps> = ({
//   parentId = '',
//   toggleText,
//   toggleTextClassName = '',
//   toggleButtonContents,
//   toggleButtonContentsClassName = '',
//   onToggle = () => null as any,
//   isOpen = false,
//   isPlain = false,
//   /* eslint-disable @typescript-eslint/no-unused-vars */
//   isText = true,
//   isDisabled = false,
//   /* eslint-disable @typescript-eslint/no-unused-vars */
//   isActive = false,
//   'aria-haspopup': ariaHasPopup,
//   parentRef,
//   /* eslint-disable @typescript-eslint/no-unused-vars */
//   getMenuRef,
//   onEnter,
//   /* eslint-enable @typescript-eslint/no-unused-vars */
//   'aria-label': ariaLabel = 'Options menu',
//   ...props
// }: OptionsMenuToggleWithTextProps) => {
//   const buttonRef = React.useRef<HTMLButtonElement>();
// 
//   React.useEffect(() => {
//     document.addEventListener('mousedown', onDocClick);
//     document.addEventListener('touchstart', onDocClick);
//     document.addEventListener('keydown', onEscPress);
//     return () => {
//       document.removeEventListener('mousedown', onDocClick);
//       document.removeEventListener('touchstart', onDocClick);
//       document.removeEventListener('keydown', onEscPress);
//     };
//   });
// 
//   const onDocClick = (event: MouseEvent | TouchEvent) => {
//     if (isOpen && parentRef && parentRef.current && !parentRef.current.contains(event.target as Node)) {
//       onToggle(false);
//       buttonRef.current.focus();
//     }
//   };
// 
//   const onKeyDown = (event: React.KeyboardEvent<any>) => {
//     if (event.key === 'Tab' && !isOpen) {
//       return;
//     }
//     event.preventDefault();
//     if ((event.key === 'Enter' || event.key === ' ') && isOpen) {
//       onToggle(!isOpen);
//     } else if ((event.key === 'Enter' || event.key === ' ') && !isOpen) {
//       onToggle(!isOpen);
//       onEnter(event);
//     }
//   };
// 
//   const onEscPress = (event: KeyboardEvent) => {
//     const keyCode = event.keyCode || event.which;
//     if (
//       isOpen &&
//       (keyCode === KEY_CODES.ESCAPE_KEY || event.key === 'Tab') &&
//       parentRef &&
//       parentRef.current &&
//       parentRef.current.contains(event.target as Node)
//     ) {
//       onToggle(false);
//       buttonRef.current.focus();
//     }
//   };
// 
//   return (
//     <div
//       className={css(
//         styles.optionsMenuToggle,
//         styles.modifiers.text,
//         isPlain && styles.modifiers.plain,
//         isDisabled && styles.modifiers.disabled,
//         isActive && styles.modifiers.active
//       )}
//       {...props}
//     >
//       <span className={css(styles.optionsMenuToggleText, toggleTextClassName)}>{toggleText}</span>
//       <button
//         className={css(styles.optionsMenuToggleButton, toggleButtonContentsClassName)}
//         id={`${parentId}-toggle`}
//         aria-haspopup="listbox"
//         aria-label={ariaLabel}
//         aria-expanded={isOpen}
//         ref={buttonRef}
//         disabled={isDisabled}
//         onClick={() => onToggle(!isOpen)}
//         onKeyDown={onKeyDown}
//       >
//         <span className={css(styles.optionsMenuToggleButtonIcon)}>{toggleButtonContents}</span>
//       </button>
//     </div>
//   );
// };
// OptionsMenuToggleWithText.displayName = 'OptionsMenuToggleWithText';
