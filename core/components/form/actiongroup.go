package form

type ActionGroupProps struct {
	// Children - Anything that can be rendered as ActionGroup content. Optional.
	Children any //  // React.ReactNode 
	// ClassName - Additional classes added to the ActionGroup. Optional.
	ClassName string
}

// contents of react-core/src/components/Form/ActionGroup.tsx
// import * as React from 'react';
// import styles from '@patternfly/react-styles/css/components/Form/form';
// import { css } from '@patternfly/react-styles';
// 
// export interface ActionGroupProps extends React.HTMLProps<HTMLDivElement> {
//   /** Anything that can be rendered as ActionGroup content. */
//   children?: React.ReactNode;
//   /** Additional classes added to the ActionGroup. */
//   className?: string;
// }
// 
// export const ActionGroup: React.FunctionComponent<ActionGroupProps> = ({
//   children = null,
//   className = '',
//   ...props
// }: ActionGroupProps) => {
//   const customClassName = css(styles.formGroup, styles.modifiers.action, className);
// 
//   const formActionsComponent = <div className={css(styles.formActions)}>{children}</div>;
//   return (
//     <div {...props} className={customClassName}>
//       <div className={css(styles.formGroupControl)}>{formActionsComponent}</div>
//     </div>
//   );
// };
// ActionGroup.displayName = 'ActionGroup';
