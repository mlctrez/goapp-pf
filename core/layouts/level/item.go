package level

type ItemProps struct {
	// Children - content rendered inside the Level Layout Item. Optional.
	Children any //  // React.ReactNode 
}

// contents of react-core/src/layouts/Level/LevelItem.tsx
// import * as React from 'react';
// 
// export interface LevelItemProps extends React.HTMLProps<HTMLDivElement> {
//   /** content rendered inside the Level Layout Item */
//   children?: React.ReactNode;
// }
// 
// export const LevelItem: React.FunctionComponent<LevelItemProps> = ({ children = null, ...props }: LevelItemProps) => (
//   <div {...props}>{children}</div>
// );
// LevelItem.displayName = 'LevelItem';
