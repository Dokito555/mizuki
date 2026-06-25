export interface Meta {
	page: number;
	page_size: number;
	total: number;
	total_pages: number;
}

export interface Stats {
	total_uploads: number;
	total_flows: number;
	total_threats: number;
}

export interface PaginatedResponse<T> {
	data: T[];
	meta: Meta;
}
