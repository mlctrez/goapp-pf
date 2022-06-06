package dropdown

type InternalDropdownItemProps struct {
	// Children - Anything which can be rendered as dropdown item. Optional.
	Children any //  // React.ReactNode 
	// StyleChildren - Whether to set className on component when React.isValidElement(component). Optional.
	StyleChildren bool
	// ClassName - Classes applied to root element of dropdown item. Optional.
	ClassName string
	// ListItemClassName - Class applied to list element. Optional.
	ListItemClassName string
	// Component - Indicates which component will be used as dropdown item. Will have className injected if React.isValidElement(component).
	// Optional.
	Component any //  // React.ReactNode 
	// Role - Role for the item. Optional.
	Role string
	// IsDisabled - Render dropdown item as disabled option. Optional.
	IsDisabled bool
	// IsAriaDisabled - Render dropdown item as aria disabled option. Optional.
	IsAriaDisabled bool
	// IsPlainText - Render dropdown item as a non-interactive item. Optional.
	IsPlainText bool
	// IsHovered - Forces display of the hover state of the element. Optional.
	IsHovered bool
	// Href - Default hyperlink location. Optional.
	Href string
	// Tooltip - Tooltip to display when hovered over the item. Optional.
	Tooltip any //  // React.ReactNode 
	// TooltipProps - Additional tooltip props forwarded to the Tooltip component. Optional.
	TooltipProps any // 
	// Index - Optional.
	Index int
	// Context - Optional.
	Context map[string]string /* { "keyHandler":func(index int, innerIndex int, direction string), "sendRef":func(index int, ref any, isDisabled bool, isSeparator bool) } */
	// OnClick - Callback for click event. Optional.
	OnClick func(event any /* any // React.MouseEvent  | any // React.KeyboardEvent  | any // MouseEvent  */)
	// Id - ID for the list element. Optional.
	Id string
	// ComponentID - ID for the component element. Optional.
	ComponentID string
	// AdditionalChild - Additional content to include alongside item within the <li>. Optional.
	AdditionalChild any //  // React.ReactNode 
	// CustomChild - Custom item rendering that receives the DropdownContext. Optional.
	CustomChild any //  // React.ReactNode 
	// EnterTriggersArrowDown - Flag indicating if hitting enter on an item also triggers an arrow down key press.
	// Optional.
	EnterTriggersArrowDown bool
	// Icon - An image to display within the InternalDropdownItem, appearing before any component children. Optional.
	Icon any //  // React.ReactNode 
	// AutoFocus - Initial focus on the item when the menu is opened (Note: Only applicable to one of the items).
	// Optional.
	AutoFocus bool
	// Description - A short description of the dropdown item, displayed under the dropdown item content. Optional.
	Description any //  // React.ReactNode 
	// InoperableEvents - Events to prevent when the item is disabled. Optional.
	InoperableEvents []string
}

// contents of react-core/src/components/Dropdown/InternalDropdownItem.tsx
// import * as React from 'react';
// import { css } from '@patternfly/react-styles';
// import { DropdownContext } from './dropdownConstants';
// import { KEYHANDLER_DIRECTION } from '../../helpers/constants';
// import { preventedEvents } from '../../helpers/util';
// import { Tooltip } from '../Tooltip';
// import styles from '@patternfly/react-styles/css/components/Dropdown/dropdown';
// 
// export interface InternalDropdownItemProps extends React.HTMLProps<HTMLAnchorElement> {
//   /** Anything which can be rendered as dropdown item */
//   children?: React.ReactNode;
//   /** Whether to set className on component when React.isValidElement(component) */
//   styleChildren?: boolean;
//   /** Classes applied to root element of dropdown item */
//   className?: string;
//   /** Class applied to list element */
//   listItemClassName?: string;
//   /** Indicates which component will be used as dropdown item. Will have className injected if React.isValidElement(component) */
//   component?: React.ReactNode;
//   /** Role for the item */
//   role?: string;
//   /** Render dropdown item as disabled option */
//   isDisabled?: boolean;
//   /** Render dropdown item as aria disabled option */
//   isAriaDisabled?: boolean;
//   /** Render dropdown item as a non-interactive item */
//   isPlainText?: boolean;
//   /** Forces display of the hover state of the element */
//   isHovered?: boolean;
//   /** Default hyperlink location */
//   href?: string;
//   /** Tooltip to display when hovered over the item */
//   tooltip?: React.ReactNode;
//   /** Additional tooltip props forwarded to the Tooltip component */
//   tooltipProps?: any;
//   index?: number;
//   context?: {
//     keyHandler?: (index: number, innerIndex: number, direction: string) => void;
//     sendRef?: (index: number, ref: any, isDisabled: boolean, isSeparator: boolean) => void;
//   };
//   /** Callback for click event */
//   onClick?: (event: React.MouseEvent<any> | React.KeyboardEvent | MouseEvent) => void;
//   /** ID for the list element */
//   id?: string;
//   /** ID for the component element */
//   componentID?: string;
//   /** Additional content to include alongside item within the <li> */
//   additionalChild?: React.ReactNode;
//   /** Custom item rendering that receives the DropdownContext */
//   customChild?: React.ReactNode;
//   /** Flag indicating if hitting enter on an item also triggers an arrow down key press */
//   enterTriggersArrowDown?: boolean;
//   /** An image to display within the InternalDropdownItem, appearing before any component children */
//   icon?: React.ReactNode;
//   /** Initial focus on the item when the menu is opened (Note: Only applicable to one of the items) */
//   autoFocus?: boolean;
//   /** A short description of the dropdown item, displayed under the dropdown item content */
//   description?: React.ReactNode;
//   /** Events to prevent when the item is disabled */
//   inoperableEvents?: string[];
// }
// 
// export class InternalDropdownItem extends React.Component<InternalDropdownItemProps> {
//   static displayName = 'InternalDropdownItem';
//   ref = React.createRef<HTMLLIElement>();
//   additionalRef = React.createRef<any>();
// 
//   static defaultProps: InternalDropdownItemProps = {
//     className: '',
//     isHovered: false,
//     component: 'a',
//     role: 'none',
//     isDisabled: false,
//     isPlainText: false,
//     tooltipProps: {},
//     // eslint-disable-next-line @typescript-eslint/no-unused-vars
//     onClick: (event: React.MouseEvent<any> | React.KeyboardEvent | MouseEvent) => undefined as any,
//     index: -1,
//     context: {
//       keyHandler: () => {},
//       sendRef: () => {}
//     },
//     enterTriggersArrowDown: false,
//     icon: null,
//     styleChildren: true,
//     description: null,
//     inoperableEvents: ['onClick', 'onKeyPress']
//   };
// 
//   componentDidMount() {
//     const { context, index, isDisabled, role, customChild, autoFocus } = this.props;
//     const customRef = customChild ? this.getInnerNode(this.ref.current) : this.ref.current;
//     context.sendRef(
//       index,
//       [customRef, customChild ? customRef : this.additionalRef.current],
//       isDisabled,
//       role === 'separator'
//     );
//     autoFocus && setTimeout(() => customRef.focus());
//   }
// 
//   componentDidUpdate() {
//     const { context, index, isDisabled, role, customChild } = this.props;
//     const customRef = customChild ? this.getInnerNode(this.ref.current) : this.ref.current;
//     context.sendRef(
//       index,
//       [customRef, customChild ? customRef : this.additionalRef.current],
//       isDisabled,
//       role === 'separator'
//     );
//   }
// 
//   getInnerNode = (node: any) => (node && node.childNodes && node.childNodes.length ? node.childNodes[0] : node);
// 
//   onKeyDown = (event: any) => {
//     // Detected key press on this item, notify the menu parent so that the appropriate item can be focused
//     const innerIndex = event.target === this.ref.current ? 0 : 1;
//     if (!this.props.customChild) {
//       event.preventDefault();
//     }
//     if (event.key === 'ArrowUp') {
//       this.props.context.keyHandler(this.props.index, innerIndex, KEYHANDLER_DIRECTION.UP);
//       event.stopPropagation();
//     } else if (event.key === 'ArrowDown') {
//       this.props.context.keyHandler(this.props.index, innerIndex, KEYHANDLER_DIRECTION.DOWN);
//       event.stopPropagation();
//     } else if (event.key === 'ArrowRight') {
//       this.props.context.keyHandler(this.props.index, innerIndex, KEYHANDLER_DIRECTION.RIGHT);
//       event.stopPropagation();
//     } else if (event.key === 'ArrowLeft') {
//       this.props.context.keyHandler(this.props.index, innerIndex, KEYHANDLER_DIRECTION.LEFT);
//       event.stopPropagation();
//     } else if (event.key === 'Enter' || event.key === ' ') {
//       event.target.click();
//       this.props.enterTriggersArrowDown &&
//         this.props.context.keyHandler(this.props.index, innerIndex, KEYHANDLER_DIRECTION.DOWN);
//     }
//   };
// 
//   extendAdditionalChildRef() {
//     const { additionalChild } = this.props;
// 
//     return React.cloneElement(additionalChild as React.ReactElement<any>, {
//       ref: this.additionalRef
//     });
//   }
// 
//   componentRef = (element: HTMLLIElement) => {
//     (this.ref as React.MutableRefObject<any>).current = element;
//     const { component } = this.props;
//     const ref = (component as any).ref;
//     if (ref) {
//       if (typeof ref === 'function') {
//         ref(element);
//       } else {
//         (ref as React.MutableRefObject<any>).current = element;
//       }
//     }
//   };
// 
//   render() {
//     /* eslint-disable @typescript-eslint/no-unused-vars */
//     const {
//       className,
//       children,
//       isHovered,
//       context,
//       onClick,
//       component,
//       role,
//       isDisabled,
//       isAriaDisabled,
//       isPlainText,
//       index,
//       href,
//       tooltip,
//       tooltipProps,
//       id,
//       componentID,
//       listItemClassName,
//       additionalChild,
//       customChild,
//       enterTriggersArrowDown,
//       icon,
//       autoFocus,
//       styleChildren,
//       description,
//       inoperableEvents,
//       ...additionalProps
//     } = this.props;
//     /* eslint-enable @typescript-eslint/no-unused-vars */
//     let classes = css(icon && styles.modifiers.icon, isAriaDisabled && styles.modifiers.ariaDisabled, className);
// 
//     if (component === 'a') {
//       additionalProps['aria-disabled'] = isDisabled || isAriaDisabled;
//     } else if (component === 'button') {
//       additionalProps['aria-disabled'] = isDisabled || isAriaDisabled;
//       additionalProps.type = additionalProps.type || 'button';
//     }
//     const renderWithTooltip = (childNode: React.ReactNode) =>
//       tooltip ? (
//         <Tooltip content={tooltip} {...tooltipProps}>
//           {childNode}
//         </Tooltip>
//       ) : (
//         childNode
//       );
// 
//     const renderClonedComponent = (element: React.ReactElement<any>) =>
//       React.cloneElement(element, {
//         ...(styleChildren && {
//           className: css(element.props.className, classes)
//         }),
//         ...(this.props.role !== 'separator' && { ref: this.componentRef })
//       });
// 
//     const renderDefaultComponent = (tag: string) => {
//       const Component = tag as any;
// 
//       const componentContent = description ? (
//         <>
//           <div className={styles.dropdownMenuItemMain}>
//             {icon && <span className={css(styles.dropdownMenuItemIcon)}>{icon}</span>}
//             {children}
//           </div>
//           <div className={styles.dropdownMenuItemDescription}>{description}</div>
//         </>
//       ) : (
//         <>
//           {icon && <span className={css(styles.dropdownMenuItemIcon)}>{icon}</span>}
//           {children}
//         </>
//       );
// 
//       return (
//         <Component
//           {...additionalProps}
//           {...(isDisabled || isAriaDisabled ? preventedEvents(inoperableEvents) : null)}
//           href={href}
//           ref={this.ref}
//           className={classes}
//           id={componentID}
//         >
//           {componentContent}
//         </Component>
//       );
//     };
// 
//     return (
//       <DropdownContext.Consumer>
//         {({ onSelect, itemClass, disabledClass, plainTextClass }) => {
//           if (this.props.role !== 'separator') {
//             classes = css(
//               classes,
//               isDisabled && disabledClass,
//               isPlainText && plainTextClass,
//               itemClass,
//               description && styles.modifiers.description
//             );
//           }
//           if (customChild) {
//             return React.cloneElement(customChild as React.ReactElement<any>, {
//               ref: this.ref,
//               onKeyDown: this.onKeyDown
//             });
//           }
// 
//           return (
//             <li
//               className={listItemClassName || null}
//               role={role}
//               onKeyDown={this.onKeyDown}
//               onClick={(event: any) => {
//                 if (!isDisabled && !isAriaDisabled) {
//                   onClick(event);
//                   onSelect(event);
//                 }
//               }}
//               id={id}
//             >
//               {renderWithTooltip(
//                 React.isValidElement(component)
//                   ? renderClonedComponent(component)
//                   : renderDefaultComponent(component as string)
//               )}
//               {additionalChild && this.extendAdditionalChildRef()}
//             </li>
//           );
//         }}
//       </DropdownContext.Consumer>
//     );
//   }
// }
