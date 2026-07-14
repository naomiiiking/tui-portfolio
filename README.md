# Terminal Website Portfolio

`ssh naomiking.xyz`

Made a terminal based website as my portfolio. Felt more on brand for an AI Engineer. Written in GO with bubbletea for the TUI framework and lipgloss for the styling.

## Deployment

The app takes port 22 so visitors connect with a bare `ssh naomiking.xyz`, and admin SSH is moved to a high port (2200). See [`deploy/DEPLOY-OCI.md`](deploy/DEPLOY-OCI.md) for the full walkthrough and [`deploy/terminal-website.service`](deploy/terminal-website.service) for the systemd unit.

The service runs as a dedicated non-root `terminal` user and binds the privileged port 22 via `CAP_NET_BIND_SERVICE` rather than running as root. The listen port is configurable with the `PORT` env var (defaults to `2222` locally).
