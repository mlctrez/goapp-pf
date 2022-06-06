package menu

type GroupProps struct {
	// Children - Items within group. Optional.
	Children any //  // React.ReactNode 
	// ClassName - Additional classes added to the MenuGroup. Optional.
	ClassName string
	// Label - Group label. Optional.
	Label string
	// TitleId - ID for title label. Optional.
	TitleId string
	// InnerRef - Forwarded ref. Optional.
	InnerRef any //  // React.Ref 
}

// contents of react-core/src/components/Menu/MenuGroup.tsx
// import * as React from 'react';
// import styles from '@patternfly/react-styles/css/components/Menu/menu';
// import { css } from '@patternfly/react-styles';
// 
// export interface MenuGroupProps extends React.HTMLProps<HTMLElement> {
//   /** Items within group */
//   children?: React.ReactNode;
//   /** Additional classes added to the MenuGroup */
//   className?: string;
//   /** Group label */
//   label?: string;
//   /** ID for title label */
//   titleId?: string;
//   /** Forwarded ref */
//   innerRef?: React.Ref<any>;
// }
// 
// const MenuGroupBase: React.FunctionComponent<MenuGroupProps> = ({
//   children,
//   className = '',
//   label = '',
//   titleId = '',
//   innerRef,
//   ...props
// }: MenuGroupProps) => (
//   <section {...props} className={css('pf-c-menu__group', className)} ref={innerRef}>
//     {label && (
//       <h1 className={css(styles.menuGroupTitle)} id={titleId}>
//         {label}
//       </h1>
//     )}
//     {children}
//   </section>
// );
// 
// export const MenuGroup = React.forwardRef((props: MenuGroupProps, ref: React.Ref<HTMLElement>) => (
//   <MenuGroupBase {...props} innerRef={ref} />
// ));
// MenuGroup.displayName = 'MenuGroup';
