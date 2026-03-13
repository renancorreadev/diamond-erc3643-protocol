export interface SelectorEntry {
  selector: string;
  signature: string;
  facet: string;
}

export interface SelectorMapData {
  entries: SelectorEntry[];
  collisions: { selector: string; facets: string[] }[];
  facetCount: number;
  totalSelectors: number;
}
