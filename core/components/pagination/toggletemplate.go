package pagination

type ToggleTemplateProps struct {
	// FirstIndex - The first index of the items being paginated. Optional.
	FirstIndex int
	// LastIndex - The last index of the items being paginated. Optional.
	LastIndex int
	// ItemCount - The total number of items being paginated. Optional.
	ItemCount int
	// ItemsTitle - The type or title of the items being paginated. Optional.
	ItemsTitle string
	// OfWord - The word that joins the index and itemCount/itemsTitle. Optional.
	OfWord any //  // React.ReactNode 
}

// contents of react-core/src/components/Pagination/ToggleTemplate.tsx
// import * as React from 'react';
// 
// export interface ToggleTemplateProps {
//   /** The first index of the items being paginated */
//   firstIndex?: number;
//   /** The last index of the items being paginated */
//   lastIndex?: number;
//   /** The total number of items being paginated */
//   itemCount?: number;
//   /** The type or title of the items being paginated */
//   itemsTitle?: string;
//   /** The word that joins the index and itemCount/itemsTitle */
//   ofWord?: React.ReactNode;
// }
// 
// export const ToggleTemplate: React.FunctionComponent<ToggleTemplateProps> = ({
//   firstIndex = 0,
//   lastIndex = 0,
//   itemCount = 0,
//   itemsTitle = 'items',
//   ofWord = 'of'
// }: ToggleTemplateProps) => (
//   <React.Fragment>
//     <b>
//       {firstIndex} - {lastIndex}
//     </b>{' '}
//     {ofWord} <b>{itemCount}</b> {itemsTitle}
//   </React.Fragment>
// );
// ToggleTemplate.displayName = 'ToggleTemplate';
