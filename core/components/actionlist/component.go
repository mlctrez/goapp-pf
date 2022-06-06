package actionlist

type Props struct {
	// Children - Children of the action list. Optional.
	Children any //  // React.ReactNode 
	// IsIconList - Flag indicating the action list contains multiple icons and item padding should be removed.
	// Optional.
	IsIconList bool
	// ClassName - Additional classes added to the action list. Optional.
	ClassName string
}

// contents of react-core/src/components/ActionList/ActionList.tsx
// import * as React from 'react';
// import { css } from '@patternfly/react-styles';
// import styles from '@patternfly/react-styles/css/components/ActionList/action-list';
// 
// export interface ActionListProps extends React.HTMLProps<HTMLDivElement> {
//   /** Children of the action list */
//   children?: React.ReactNode;
//   /** Flag indicating the action list contains multiple icons and item padding should be removed */
//   isIconList?: boolean;
//   /** Additional classes added to the action list */
//   className?: string;
// }
// 
// export const ActionList: React.FunctionComponent<ActionListProps> = ({
//   children,
//   isIconList,
//   className = '',
//   ...props
// }: ActionListProps) => (
//   <div className={css(styles.actionList, isIconList && styles.modifiers.icons, className)} {...props}>
//     {children}
//   </div>
// );
// ActionList.displayName = 'ActionList';
