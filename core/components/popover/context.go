package popover

type ContextProps struct {
	// HeaderComponent - Optional.
	HeaderComponent any //  /* "h1" | "h2" | "h3" | "h4" | "h5" | "h6" */
}

// contents of react-core/src/components/Popover/PopoverContext.tsx
// import * as React from 'react';
// 
// interface PopoverContextProps {
//   headerComponent?: 'h1' | 'h2' | 'h3' | 'h4' | 'h5' | 'h6';
// }
// 
// export const PopoverContext = React.createContext<Partial<PopoverContextProps>>({});
