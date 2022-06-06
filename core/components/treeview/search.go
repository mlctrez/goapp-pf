package treeview

type SearchProps struct {
	// OnSearch - Callback for search input. Optional.
	OnSearch func(event any // React.ChangeEvent )
	// Id - Id for the search input. Optional.
	Id string
	// Name - Name for the search input. Optional.
	Name string
	// AriaLabel - Accessible label for the search input. Optional.
	AriaLabel string
	// ClassName - Classes applied to the wrapper for the search input. Optional.
	ClassName string
}

// contents of react-core/src/components/TreeView/TreeViewSearch.tsx
// import * as React from 'react';
// import { css } from '@patternfly/react-styles';
// import styles from '@patternfly/react-styles/css/components/TreeView/tree-view';
// import formStyles from '@patternfly/react-styles/css/components/FormControl/form-control';
// 
// export interface TreeViewSearchProps extends React.HTMLProps<HTMLInputElement> {
//   /** Callback for search input */
//   onSearch?: (event: React.ChangeEvent<HTMLInputElement>) => void;
//   /** Id for the search input */
//   id?: string;
//   /** Name for the search input */
//   name?: string;
//   /** Accessible label for the search input */
//   'aria-label'?: string;
//   /** Classes applied to the wrapper for the search input */
//   className?: string;
// }
// 
// export const TreeViewSearch: React.FunctionComponent<TreeViewSearchProps> = ({
//   className,
//   onSearch,
//   id,
//   name,
//   'aria-label': ariaLabel,
//   ...props
// }: TreeViewSearchProps) => (
//   <div className={css(styles.treeViewSearch, className)}>
//     <input
//       className={css(formStyles.formControl, formStyles.modifiers.search)}
//       onChange={onSearch}
//       id={id}
//       name={name}
//       aria-label={ariaLabel}
//       type="search"
//       {...props}
//     />
//   </div>
// );
// TreeViewSearch.displayName = 'TreeViewSearch';
