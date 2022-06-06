package simplelist

type ItemProps struct {
	// ItemId - id for the item. Optional.
	ItemId any //  /* int | string */
	// Children - Content rendered inside the SimpleList item. Optional.
	Children any //  // React.ReactNode 
	// ClassName - Additional classes added to the SimpleList <li>. Optional.
	ClassName string
	// Component - Component type of the SimpleList item. Optional.
	Component any //  /* "button" | "a" */
	// ComponentClassName - Additional classes added to the SimpleList <a> or <button>. Optional.
	ComponentClassName string
	// ComponentProps - Additional props added to the SimpleList <a> or <button>. Optional.
	ComponentProps any // 
	// IsActive - Indicates if the link is current/highlighted. Optional.
	IsActive bool
	// IsCurrent - Optional.
	IsCurrent bool
	// OnClick - OnClick callback for the SimpleList item. Optional.
	OnClick func(event any /* any // React.MouseEvent  | any // React.ChangeEvent  */)
	// Type - Type of button SimpleList item. Optional.
	Type any //  /* "button" | "submit" | "reset" */
	// Href - Default hyperlink location. Optional.
	Href string
}

// contents of react-core/src/components/SimpleList/SimpleListItem.tsx
// import * as React from 'react';
// import { css } from '@patternfly/react-styles';
// import styles from '@patternfly/react-styles/css/components/SimpleList/simple-list';
// import { SimpleListContext } from './SimpleList';
// 
// export interface SimpleListItemProps {
//   /** id for the item. */
//   itemId?: number | string;
//   /** Content rendered inside the SimpleList item */
//   children?: React.ReactNode;
//   /** Additional classes added to the SimpleList <li> */
//   className?: string;
//   /** Component type of the SimpleList item */
//   component?: 'button' | 'a';
//   /** Additional classes added to the SimpleList <a> or <button> */
//   componentClassName?: string;
//   /** Additional props added to the SimpleList <a> or <button> */
//   componentProps?: any;
//   /** Indicates if the link is current/highlighted */
//   isActive?: boolean;
//   /** @deprecated please use isActive instead */
//   isCurrent?: boolean;
//   /** OnClick callback for the SimpleList item */
//   onClick?: (event: React.MouseEvent | React.ChangeEvent) => void;
//   /** Type of button SimpleList item */
//   type?: 'button' | 'submit' | 'reset';
//   /** Default hyperlink location */
//   href?: string;
// }
// 
// export class SimpleListItem extends React.Component<SimpleListItemProps> {
//   static displayName = 'SimpleListItem';
//   ref = React.createRef<any>();
//   static defaultProps: SimpleListItemProps = {
//     children: null,
//     className: '',
//     isActive: false,
//     isCurrent: false,
//     component: 'button',
//     componentClassName: '',
//     type: 'button',
//     href: '',
//     onClick: () => {}
//   };
// 
//   render() {
//     const {
//       children,
//       isCurrent,
//       isActive,
//       className,
//       component: Component,
//       componentClassName,
//       componentProps,
//       onClick,
//       type,
//       href,
//       // eslint-disable-next-line @typescript-eslint/no-unused-vars
//       itemId,
//       ...props
//     } = this.props;
//     return (
//       <SimpleListContext.Consumer>
//         {({ currentRef, updateCurrentRef, isControlled }) => {
//           const isButton = Component === 'button';
//           const isCurrentItem =
//             this.ref && currentRef && isControlled ? currentRef.current === this.ref.current : isActive || isCurrent;
// 
//           const additionalComponentProps = isButton
//             ? {
//                 type
//               }
//             : {
//                 tabIndex: 0,
//                 href
//               };
// 
//           return (
//             <li className={css(className)} {...props}>
//               <Component
//                 className={css(
//                   styles.simpleListItemLink,
//                   isCurrentItem && styles.modifiers.current,
//                   componentClassName
//                 )}
//                 onClick={(evt: React.MouseEvent) => {
//                   onClick(evt);
//                   updateCurrentRef(this.ref, this.props);
//                 }}
//                 ref={this.ref}
//                 {...componentProps}
//                 {...additionalComponentProps}
//               >
//                 {children}
//               </Component>
//             </li>
//           );
//         }}
//       </SimpleListContext.Consumer>
//     );
//   }
// }
