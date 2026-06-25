export interface Upload {
	id: number;
	filename: string;
	file_size: number;
	file_hash: string;
	file_type: string;
	status: string;
	progress_pct: number;
	packets_processed: number;
	flow_count: number;
	duration_ms: number;
	error?: string;
	created_at: string;
	updated_at: string;
}
