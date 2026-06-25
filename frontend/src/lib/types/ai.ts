export interface AIAnalysis {
	flow_id: number;
	status: 'analyzed' | 'not_analyzed';
	model?: string;
	analyzed_at?: string;
	analysis?: {
		model_used: string;
		confidence: number;
		narrative: string;
		mitre_ids: string[];
		attribution: string;
		remediation: string[];
		is_fallback: boolean;
		correlations?: {
			flow_ids: number[];
			pattern: string;
			description: string;
		}[];
	};
}
