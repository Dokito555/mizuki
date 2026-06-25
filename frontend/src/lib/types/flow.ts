export interface Flow {
	id: number;
	src_ip: string;
	dst_ip: string;
	src_port: number;
	dst_port: number;
	protocol: string;
	first_seen: string;
	last_seen: string;
	packet_count: number;
	byte_count: number;
	src_mac?: string;
	dst_mac?: string;
	tls_version?: string;
	tls_sni?: string;
	dns_queries?: string[];
	app_protocol?: string;
	iat_avg_ms: number;
	iat_min_ms: number;
	iat_max_ms: number;
	iat_std_dev_ms: number;
	score: number;
	threats?: string[];
	created_at: string;
}

export interface FlowDetail extends Flow {
	packet_samples?: PacketSample[];
}

export interface PacketSample {
	ts: string;
	size: number;
}
