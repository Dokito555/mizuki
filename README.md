# Mizuki

Network traffic analysis platform. Upload PCAP/PCAPNG files, extract flows, detect anomalies with heuristic scoring, and enrich findings with AI-powered analysis via Ollama.

## Features

- **PCAP Ingestion** — Upload PCAP/PCAPNG files via drag-and-drop or file picker; async pipeline parses packets, extracts flows, and stores them with packet samples
- **Flow Extraction** — Identifies bidirectional flows with MAC, IPs, ports, protocol (TCP/UDP/ICMP), TLS handshake info (SNI, version), DNS queries, application protocol detection (HTTP, SSH, SMTP, FTP, DNS), and inter-arrival time (IAT) statistics
- **Heuristic Detection** — 5 built-in anomaly detectors score every flow on a 0–100 scale:
  - **C2 Beaconing** — detects periodic communication via IAT coefficient of variation
  - **Multi-Port Scanning** — flags hosts connecting to many unique ports
  - **Protocol Anomaly** — detects protocols on non-standard ports (e.g., TLS on port 25)
  - **Payload Download** — flags large data transfers
  - **TLS/DNS Anomaly** — detects TLS without SNI, DNS floods
- **AI Enrichment** — Optional Ollama integration generates threat narratives, MITRE ATT&CK mapping, attribution, remediation steps, and cross-flow correlations using LLM analysis
- **Threat Dashboard** — Aggregated stats (total uploads, flows, threats), recent uploads, and per-flow threat scores
- **Interactive Flow Browser** — Sortable, filterable table with pagination; filter by source/dest IP, protocol, minimum score
- **Flow Detail** — Metadata cards, IAT statistics, TLS/DNS info, packet sample timeline, detected threats
- **Async Upload Pipeline** — File processing (hashing → parsing → inserting → detection) runs asynchronously with live progress polling

## Architecture

```
mizuki/
├── cmd/                     # Go entry point
│   └── main.go
├── internal/                # Go backend
│   ├── configs/             # Dependency injection, DB, Gin, logging
│   ├── constants/           # Shared errors and status constants
│   ├── entities/            # GORM models (Upload, Flow, FlowAI, etc.)
│   ├── models/              # Request/response DTOs
│   ├── repositories/        # Data access layer
│   ├── services/            # Business logic
│   │   ├── upload/          # Upload pipeline (parse, hash, insert)
│   │   ├── flow/            # Flow listing and detail
│   │   ├── pcap/            # PCAP/PCAPNG parsing engine
│   │   ├── detection/       # Heuristic scoring engine (5 detectors)
│   │   └── ai/              # AI enrichment (Ollama + fallback + queue)
│   └── deliveries/http/     # Gin HTTP handlers + routes + middleware
├── frontend/                # SvelteKit frontend
│   └── src/
│       ├── lib/api/         # Axios API clients
│       ├── lib/hooks/       # TanStack Query hooks
│       ├── lib/stores/      # Svelte 5 runes stores (toast)
│       ├── lib/types/       # TypeScript interfaces
│       ├── lib/components/  # UI components
│       └── routes/          # SvelteKit pages
└── .env.example             # Environment config template
```

**Backend:** Go + Gin (HTTP), GORM + PostgreSQL (data), gopacket (packet parsing), Ollama (AI)
**Frontend:** Svelte 5 (runes), SvelteKit 2, Tailwind CSS v4, TanStack Query 5, Axios, Lucide icons

## Prerequisites

- Go 1.26+
- Node.js 20+
- PostgreSQL 15+
- Ollama (optional, for AI enrichment)

## Setup

### 1. Clone and configure

```bash
git clone <repo-url> && cd mizuki
cp .env.example .env
```

### 2. Configure environment (`.env`)

| Variable | Description | Default |
|----------|-------------|---------|
| `APP_PORT` | Backend server port | `8080` |
| `DB_HOST` / `DB_PORT` / `DB_USER` / `DB_PASSWORD` / `DB_NAME` | PostgreSQL connection | — |
| `MAX_FILE_SIZE_MB` | Max upload size in MB | `500` |
| `MIZUKI_TEMP_DIR` | Temp directory for upload processing | system temp |
| `AI_ENABLED` | Enable AI enrichment | `false` |
| `OLLAMA_URL` | Ollama API URL | `http://localhost:11434` |
| `OLLAMA_MODEL` | LLM model for analysis | `qwen2.5:3b` |

See [`.env.example`](.env.example) for the full list.

### 3. Backend

```bash
go mod download
go run ./cmd/main.go
```

The backend starts on the port specified in `APP_PORT` (default `8080`). It auto-migrates all database tables on startup.

### 4. Frontend

```bash
cd frontend
npm install
npm run dev
```

The frontend starts on `http://localhost:5173` and proxies `/api` requests to the backend at `http://localhost:8080`.

## API Reference

All endpoints are under `/api` and return JSON envelopes:

- `{ "data": ... }` — single resource
- `{ "data": [...], "meta": { page, page_size, total, total_pages } }` — paginated list
- `{ "error": "..." }` — error response

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/api/healthcheck` | Health check |
| `GET` | `/api/stats` | Aggregate stats (uploads, flows, threats) |
| `POST` | `/api/pcap/upload` | Upload PCAP file (multipart, field: `file`) |
| `GET` | `/api/uploads` | List uploads (paginated) |
| `GET` | `/api/uploads/:id` | Get upload status/progress |
| `POST` | `/api/uploads/:id/analyze` | Re-run analysis |
| `POST` | `/api/uploads/:id/ai-analyze` | Batch AI analysis for all flows |
| `GET` | `/api/flows` | List flows (paginated, filtered, sortable) |
| `GET` | `/api/flows/:id` | Get flow detail with packet samples |
| `GET` | `/api/flows/:id/ai` | Get AI analysis for a flow |
| `POST` | `/api/flows/:id/ai-analyze` | Trigger single-flow AI analysis |

### Query parameters for `GET /api/flows`

```
src_ip     — source IP
dst_ip     — destination IP
protocol   — TCP, UDP, ICMP
min_score  — minimum threat score (0–100)
upload_id  — filter by upload
since      — RFC3339 start time
until      — RFC3339 end time
page       — page number (default: 1)
page_size  — items per page (default: 20, max: 500)
sort_by    — first_seen, last_seen, score, src_ip, dst_ip, protocol, packet_count, byte_count
sort_desc  — sort descending (default: true for first_seen)
```

If no `since`/`until`/`upload_id` is provided, defaults to the last 24 hours.

## Tech Stack

| Layer | Technology |
|-------|-----------|
| Backend | Go 1.26, Gin 1.12, GORM 1.31, PostgreSQL |
| Packet parsing | gopacket 1.19 (supports pcap + pcapng) |
| AI | Ollama with Qwen 2.5 (or any chat model) |
| Frontend | Svelte 5, SvelteKit 2, TypeScript 5 |
| Styling | Tailwind CSS v4 |
| State | @tanstack/svelte-query 5 |
| Icons | Lucide Svelte |
| HTTP | Axios 1.9 |
| Auth | None (CORS open) |

## License

MIT
