package jumplinks

type ItemProps struct {
	// IsActive - Whether this item is active. Parent JumpLinks component sets this when passed a `scrollableSelector`.
	// Optional.
	IsActive bool
	// Href - Href for this link. Optional.
	Href string
	// Node - Selector or HTMLElement to spy on. Optional.
	Node any //  /* string | any // HTMLElement  */
	// Children - Text to be rendered inside span. Optional.
	Children any //  // React.ReactNode 
	// OnClick - Click handler for anchor tag. Parent JumpLinks components tap into this. Optional.
	OnClick func(ev any // React.MouseEvent )
	// ClassName - Class to add to li. Optional.
	ClassName string
}

// contents of react-core/src/components/JumpLinks/JumpLinksItem.tsx
// import * as React from 'react';
// import { css } from '@patternfly/react-styles';
// import styles from '@patternfly/react-styles/css/components/JumpLinks/jump-links';
// import { JumpLinksList } from './JumpLinksList';
// 
// export interface JumpLinksItemProps extends Omit<React.HTMLProps<HTMLLIElement>, 'onClick'> {
//   /** Whether this item is active. Parent JumpLinks component sets this when passed a `scrollableSelector`. */
//   isActive?: boolean;
//   /** Href for this link */
//   href?: string;
//   /** Selector or HTMLElement to spy on */
//   node?: string | HTMLElement;
//   /** Text to be rendered inside span */
//   children?: React.ReactNode;
//   /** Click handler for anchor tag. Parent JumpLinks components tap into this. */
//   onClick?: (ev: React.MouseEvent<HTMLAnchorElement>) => void;
//   /** Class to add to li */
//   className?: string;
// }
// 
// export const JumpLinksItem: React.FunctionComponent<JumpLinksItemProps> = ({
//   isActive,
//   href,
//   // eslint-disable-next-line
//   node,
//   children,
//   onClick,
//   className,
//   ...props
// }: JumpLinksItemProps) => {
//   const childrenArr = React.Children.toArray(children) as any[];
//   const sublists = childrenArr.filter(child => child.type === JumpLinksList);
//   children = childrenArr.filter(child => child.type !== JumpLinksList);
// 
//   return (
//     <li
//       className={css(styles.jumpLinksItem, isActive && styles.modifiers.current, className)}
//       {...(isActive && { 'aria-current': 'location' })}
//       {...props}
//     >
//       <a className={styles.jumpLinksLink} href={href} onClick={onClick}>
//         <span className={styles.jumpLinksLinkText}>{children}</span>
//       </a>
//       {sublists}
//     </li>
//   );
// };
// JumpLinksItem.displayName = 'JumpLinksItem';
