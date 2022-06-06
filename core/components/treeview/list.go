package treeview

type ListProps struct {
	// IsNested - Flag indicating if the tree view is nested under another tree view. Optional.
	IsNested bool
	// Toolbar - Toolbar to display above the tree view. Optional.
	Toolbar any //  // React.ReactNode 
	// Children - Child nodes of the current tree view.
	Children any //  // React.ReactNode 
}

// contents of react-core/src/components/TreeView/TreeViewList.tsx
// import * as React from 'react';
// import { css } from '@patternfly/react-styles';
// import { Divider } from '../Divider';
// 
// export interface TreeViewListProps extends React.HTMLProps<HTMLUListElement> {
//   /** Flag indicating if the tree view is nested under another tree view */
//   isNested?: boolean;
//   /** Toolbar to display above the tree view */
//   toolbar?: React.ReactNode;
//   /** Child nodes of the current tree view */
//   children: React.ReactNode;
// }
// 
// export const TreeViewList: React.FunctionComponent<TreeViewListProps> = ({
//   isNested = false,
//   toolbar,
//   children,
//   ...props
// }: TreeViewListProps) => (
//   <>
//     {toolbar && (
//       <React.Fragment>
//         {toolbar}
//         <Divider />
//       </React.Fragment>
//     )}
//     <ul className={css('pf-c-tree-view__list')} role={isNested ? 'group' : 'tree'} {...props}>
//       {children}
//     </ul>
//   </>
// );
// TreeViewList.displayName = 'TreeViewList';
